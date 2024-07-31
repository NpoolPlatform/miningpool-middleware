package orderuser

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"github.com/NpoolPlatform/go-service-framework/pkg/wlog"
	mpbasetypes "github.com/NpoolPlatform/message/npool/basetypes/miningpool/v1"
	npool "github.com/NpoolPlatform/message/npool/miningpool/mw/v1/orderuser"

	"github.com/NpoolPlatform/miningpool-middleware/pkg/db"
	"github.com/NpoolPlatform/miningpool-middleware/pkg/db/ent"
	"github.com/NpoolPlatform/miningpool-middleware/pkg/db/ent/gooduser"
	orderuserent "github.com/NpoolPlatform/miningpool-middleware/pkg/db/ent/orderuser"
	"github.com/NpoolPlatform/miningpool-middleware/pkg/db/ent/pool"
	"github.com/NpoolPlatform/miningpool-middleware/pkg/db/ent/rootuser"

	orderusercrud "github.com/NpoolPlatform/miningpool-middleware/pkg/crud/orderuser"
)

type queryHandler struct {
	*Handler
	stm   *ent.OrderUserSelect
	infos []*npool.OrderUser
	total uint32
}

func (h *queryHandler) selectOrderUser(stm *ent.OrderUserQuery) {
	h.stm = stm.Select(
		orderuserent.FieldID,
		orderuserent.FieldEntID,
		orderuserent.FieldName,
		orderuserent.FieldGoodUserID,
		orderuserent.FieldAppID,
		orderuserent.FieldUserID,
		orderuserent.FieldReadPageLink,
		orderuserent.FieldCreatedAt,
		orderuserent.FieldUpdatedAt,
	)
}

func (h *queryHandler) queryOrderUser(cli *ent.Client) error {
	if h.ID == nil && h.EntID == nil {
		return wlog.Errorf("invalid id or ent_id")
	}
	stm := cli.OrderUser.Query().Where(orderuserent.DeletedAt(0))
	if h.ID != nil {
		stm.Where(orderuserent.ID(*h.ID))
	}
	if h.EntID != nil {
		stm.Where(orderuserent.EntID(*h.EntID))
	}
	h.selectOrderUser(stm)
	return nil
}

func (h *queryHandler) queryOrderUsers(ctx context.Context, cli *ent.Client) error {
	stm, err := orderusercrud.SetQueryConds(cli.OrderUser.Query(), h.Conds)
	if err != nil {
		return wlog.WrapError(err)
	}

	stmCount, err := orderusercrud.SetQueryConds(cli.OrderUser.Query(), h.Conds)
	if err != nil {
		return wlog.WrapError(err)
	}
	stmCount.Modify(h.queryJoinCoinAndPool)
	total, err := stmCount.Count(ctx)
	if err != nil {
		return wlog.WrapError(err)
	}
	h.total = uint32(total)

	h.selectOrderUser(stm)
	return nil
}

func (h *queryHandler) queryJoin() {
	h.stm.Modify(h.queryJoinCoinAndPool)
}

func (h *queryHandler) queryJoinCoinAndPool(s *sql.Selector) {
	guT := sql.Table(gooduser.Table)
	ruT := sql.Table(rootuser.Table)
	poolT := sql.Table(pool.Table)

	s.Join(guT).On(
		s.C(orderuserent.FieldGoodUserID),
		guT.C(gooduser.FieldEntID),
	).OnP(
		sql.EQ(guT.C(gooduser.FieldDeletedAt), 0),
	).Join(ruT).On(
		guT.C(gooduser.FieldRootUserID),
		ruT.C(rootuser.FieldEntID),
	).OnP(
		sql.EQ(ruT.C(rootuser.FieldDeletedAt), 0),
	).Join(poolT).On(
		ruT.C(rootuser.FieldPoolID),
		poolT.C(pool.FieldEntID),
	).OnP(
		sql.EQ(poolT.C(pool.FieldDeletedAt), 0),
	).AppendSelect(
		pool.FieldMiningpoolType,
		sql.As(poolT.C(pool.FieldName), "miningpool_name"),
		sql.As(poolT.C(pool.FieldLogo), "miningpool_logo"),
		sql.As(poolT.C(pool.FieldSite), "miningpool_site"),
		gooduser.FieldRootUserID,
	)
}

func (h *queryHandler) scan(ctx context.Context) error {
	return h.stm.Scan(ctx, &h.infos)
}

func (h *queryHandler) formalize() {
	for _, info := range h.infos {
		info.MiningpoolType = mpbasetypes.MiningpoolType(mpbasetypes.MiningpoolType_value[info.MiningpoolTypeStr])
	}
}

func (h *Handler) GetOrderUser(ctx context.Context) (*npool.OrderUser, error) {
	handler := &queryHandler{
		Handler: h,
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.queryOrderUser(cli); err != nil {
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

func (h *Handler) GetOrderUsers(ctx context.Context) ([]*npool.OrderUser, uint32, error) {
	handler := &queryHandler{
		Handler: h,
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.queryOrderUsers(ctx, cli); err != nil {
			return wlog.WrapError(err)
		}
		handler.queryJoin()
		handler.stm.
			Offset(int(h.Offset)).
			Limit(int(h.Limit)).
			Order(ent.Desc(orderuserent.FieldUpdatedAt))
		return handler.scan(_ctx)
	})
	if err != nil {
		return nil, 0, wlog.WrapError(err)
	}

	handler.formalize()
	return handler.infos, handler.total, nil
}
