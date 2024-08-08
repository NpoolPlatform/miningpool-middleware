package gooduser

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"github.com/NpoolPlatform/go-service-framework/pkg/wlog"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	mpbasetypes "github.com/NpoolPlatform/message/npool/basetypes/miningpool/v1"
	v1 "github.com/NpoolPlatform/message/npool/basetypes/v1"
	coinpb "github.com/NpoolPlatform/message/npool/miningpool/mw/v1/coin"
	npool "github.com/NpoolPlatform/message/npool/miningpool/mw/v1/gooduser"

	"github.com/NpoolPlatform/miningpool-middleware/pkg/db"
	"github.com/NpoolPlatform/miningpool-middleware/pkg/db/ent"
	gooduserent "github.com/NpoolPlatform/miningpool-middleware/pkg/db/ent/gooduser"
	rootusermw "github.com/NpoolPlatform/miningpool-middleware/pkg/mw/rootuser"

	"github.com/NpoolPlatform/miningpool-middleware/pkg/db/ent/pool"
	"github.com/NpoolPlatform/miningpool-middleware/pkg/db/ent/rootuser"
	"github.com/NpoolPlatform/miningpool-middleware/pkg/mw/coin"
	"github.com/NpoolPlatform/miningpool-middleware/pkg/pools"

	goodusercrud "github.com/NpoolPlatform/miningpool-middleware/pkg/crud/gooduser"
)

type queryHandler struct {
	*Handler
	stm   *ent.GoodUserSelect
	infos []*npool.GoodUser
	total uint32
}

func (h *queryHandler) selectGoodUser(stm *ent.GoodUserQuery) {
	h.stm = stm.Select(
		gooduserent.FieldID,
		gooduserent.FieldCreatedAt,
		gooduserent.FieldUpdatedAt,
		gooduserent.FieldEntID,
		gooduserent.FieldRootUserID,
		gooduserent.FieldName,
		gooduserent.FieldReadPageLink,
	)
}

func (h *queryHandler) queryGoodUser(cli *ent.Client) error {
	if h.ID == nil && h.EntID == nil {
		return wlog.Errorf("invalid id")
	}
	stm := cli.GoodUser.Query().Where(gooduserent.DeletedAt(0))
	if h.ID != nil {
		stm.Where(gooduserent.ID(*h.ID))
	}
	if h.EntID != nil {
		stm.Where(gooduserent.EntID(*h.EntID))
	}
	h.selectGoodUser(stm)
	return nil
}

func (h *queryHandler) queryGoodUsers(ctx context.Context, cli *ent.Client) error {
	stm, err := goodusercrud.SetQueryConds(cli.GoodUser.Query(), h.Conds)
	if err != nil {
		return wlog.WrapError(err)
	}

	stmCount, err := goodusercrud.SetQueryConds(cli.GoodUser.Query(), h.Conds)
	if err != nil {
		return wlog.WrapError(err)
	}
	stmCount.Modify(h.queryJoinRootUserAndPool)
	total, err := stmCount.Count(ctx)
	if err != nil {
		return wlog.WrapError(err)
	}
	h.total = uint32(total)

	h.selectGoodUser(stm)
	return nil
}

func (h *queryHandler) queryJoin() {
	h.stm.Modify(h.queryJoinRootUserAndPool)
}

func (h *queryHandler) queryJoinRootUserAndPool(s *sql.Selector) {
	ruT := sql.Table(rootuser.Table)
	poolT := sql.Table(pool.Table)
	s.Join(ruT).On(
		s.C(gooduserent.FieldRootUserID),
		ruT.C(rootuser.FieldEntID),
	).OnP(
		sql.EQ(ruT.C(rootuser.FieldDeletedAt), 0),
	).Join(poolT).On(
		ruT.C(rootuser.FieldPoolID),
		poolT.C(pool.FieldEntID),
	).OnP(
		sql.EQ(poolT.C(pool.FieldDeletedAt), 0),
	).AppendSelect(
		poolT.C(pool.FieldMiningpoolType),
		ruT.C(rootuser.FieldPoolID),
	).Distinct()
}

func (h *queryHandler) scan(ctx context.Context) error {
	return h.stm.Scan(ctx, &h.infos)
}

func (h *queryHandler) formalize() {
	for _, info := range h.infos {
		info.MiningpoolType = mpbasetypes.MiningpoolType(mpbasetypes.MiningpoolType_value[info.MiningpoolTypeStr])
	}
}

func (h *Handler) GetGoodUser(ctx context.Context) (*npool.GoodUser, error) {
	handler := &queryHandler{
		Handler: h,
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.queryGoodUser(cli); err != nil {
			return wlog.WrapError(err)
		}
		handler.queryJoin()
		const singleRowLimit = 2
		handler.stm.Offset(0).Limit(singleRowLimit)
		return handler.scan(_ctx)
	})
	if err != nil {
		return nil, wlog.WrapError(err)
	}
	if len(handler.infos) == 0 {
		return nil, nil
	}
	if len(handler.infos) > 1 {
		return nil, wlog.Errorf("too many record")
	}

	handler.formalize()
	return handler.infos[0], nil
}

func (h *Handler) GetGoodUsers(ctx context.Context) ([]*npool.GoodUser, uint32, error) {
	handler := &queryHandler{
		Handler: h,
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.queryGoodUsers(ctx, cli); err != nil {
			return wlog.WrapError(err)
		}
		handler.queryJoin()
		handler.stm.
			Offset(int(h.Offset)).
			Limit(int(h.Limit)).
			Order(ent.Desc(gooduserent.FieldUpdatedAt))
		return handler.scan(_ctx)
	})
	if err != nil {
		return nil, 0, wlog.WrapError(err)
	}

	handler.formalize()
	return handler.infos, handler.total, nil
}

func (h *Handler) getGoodUserHashRate(ctx context.Context) (float64, error) {
	guInfo, err := h.GetGoodUser(ctx)
	if err != nil {
		return 0, wlog.WrapError(err)
	}

	if guInfo == nil {
		return 0, wlog.Errorf("invalid id or ent_id")
	}

	rootuserID := guInfo.RootUserID
	rootuserH, err := rootusermw.NewHandler(ctx, rootusermw.WithEntID(&rootuserID, true))
	if err != nil {
		return 0, wlog.WrapError(err)
	}
	rootuserToken, err := rootuserH.GetAuthToken(ctx)
	if err != nil {
		return 0, wlog.WrapError(err)
	}
	if rootuserToken == nil {
		return 0, wlog.Errorf("have no rootuser,entid: %v", rootuserID)
	}

	coinH, err := coin.NewHandler(ctx,
		coin.WithConds(&coinpb.Conds{
			CoinTypeIDs: &v1.StringSliceVal{
				Op:    cruder.IN,
				Value: h.CoinTypeIDs,
			},
			PoolID: &v1.StringVal{
				Op:    cruder.EQ,
				Value: guInfo.PoolID,
			},
		}),
		coin.WithLimit(int32(len(h.CoinTypeIDs))),
		coin.WithOffset(0))
	if err != nil {
		return 0, wlog.WrapError(err)
	}

	// check if cointypes is suppored by the miningpool
	coinInfos, _, err := coinH.GetCoins(ctx)
	if err != nil {
		return 0, wlog.WrapError(err)
	}

	coinTypes := []v1.CoinType{}
	for _, coinTypeID := range h.CoinTypeIDs {
		existID := false
		for _, coinInfo := range coinInfos {
			if coinInfo.CoinTypeID == coinTypeID {
				coinTypes = append(coinTypes, coinInfo.CoinType)
				existID = true
				break
			}
		}

		if !existID {
			return 0, wlog.Errorf("cannot support all cointype in cointypeids")
		}
	}

	mgr, err := pools.NewPoolManager(guInfo.MiningpoolType, nil, rootuserToken.AuthTokenPlain)
	if err != nil {
		return 0, wlog.WrapError(err)
	}
	hashRate, err := mgr.GetHashRate(ctx, guInfo.Name, coinTypes)
	if err != nil {
		return 0, wlog.WrapError(err)
	}
	return hashRate, nil
}

func (h *Handler) GetGoodUserHashRate(ctx context.Context) (float64, error) {
	return h.getGoodUserHashRate(ctx)
}
