package gooduser

import (
	"context"
	"fmt"

	basetypes "github.com/NpoolPlatform/message/npool/basetypes/miningpool/v1"
	npool "github.com/NpoolPlatform/message/npool/miningpool/mw/v1/gooduser"

	"github.com/NpoolPlatform/miningpool-middleware/pkg/db"
	"github.com/NpoolPlatform/miningpool-middleware/pkg/db/ent"
	gooduserent "github.com/NpoolPlatform/miningpool-middleware/pkg/db/ent/gooduser"

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
		gooduserent.FieldCoinID,
		gooduserent.FieldRevenueID,
		gooduserent.FieldHashRate,
		gooduserent.FieldReadPageLink,
	)
}

func (h *queryHandler) queryGoodUser(cli *ent.Client) error {
	if h.ID == nil && h.EntID == nil {
		return fmt.Errorf("invalid id")
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
		return err
	}

	total, err := stm.Count(ctx)
	if err != nil {
		return err
	}
	h.total = uint32(total)

	h.selectGoodUser(stm)
	return nil
}

func (h *queryHandler) scan(ctx context.Context) error {
	return h.stm.Scan(ctx, &h.infos)
}

func (h *queryHandler) formalize() {
	for _, info := range h.infos {
		info.MiningpoolType = basetypes.MiningpoolType(basetypes.MiningpoolType_value[info.MiningpoolTypeStr])
		info.CoinType = basetypes.CoinType(basetypes.CoinType_value[info.CoinTypeStr])
		info.RevenueType = basetypes.RevenueType(basetypes.RevenueType_value[info.RevenueTypeStr])
	}
}

func (h *Handler) GetGoodUser(ctx context.Context) (*npool.GoodUser, error) {
	handler := &queryHandler{
		Handler: h,
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.queryGoodUser(cli); err != nil {
			return err
		}
		const singleRowLimit = 2
		handler.stm.Offset(0).Limit(singleRowLimit)
		return handler.scan(_ctx)
	})
	if err != nil {
		return nil, err
	}
	if len(handler.infos) == 0 {
		return nil, nil
	}
	if len(handler.infos) > 1 {
		return nil, fmt.Errorf("too many record")
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
			return err
		}
		handler.stm.
			Offset(int(h.Offset)).
			Limit(int(h.Limit)).
			Order(ent.Desc(gooduserent.FieldUpdatedAt))
		return handler.scan(_ctx)
	})
	if err != nil {
		return nil, 0, err
	}

	handler.formalize()
	return handler.infos, handler.total, nil
}
