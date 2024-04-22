package orderuser

import (
	"context"
	"fmt"

	basetypes "github.com/NpoolPlatform/message/npool/basetypes/miningpool/v1"
	npool "github.com/NpoolPlatform/message/npool/miningpool/mw/v1/orderuser"
	"github.com/shopspring/decimal"

	"github.com/NpoolPlatform/miningpool-middleware/pkg/db"
	"github.com/NpoolPlatform/miningpool-middleware/pkg/db/ent"
	orderuserent "github.com/NpoolPlatform/miningpool-middleware/pkg/db/ent/orderuser"

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
		orderuserent.FieldRootUserID,
		orderuserent.FieldGoodUserID,
		orderuserent.FieldAppID,
		orderuserent.FieldUserID,
		orderuserent.FieldMiningpoolType,
		orderuserent.FieldCoinType,
		orderuserent.FieldProportion,
		orderuserent.FieldRevenueAddress,
		orderuserent.FieldReadPageLink,
		orderuserent.FieldAutoPay,
		orderuserent.FieldCreatedAt,
		orderuserent.FieldUpdatedAt,
	)
}

func (h *queryHandler) queryOrderUser(cli *ent.Client) error {
	if h.ID == nil && h.EntID == nil {
		return fmt.Errorf("invalid id or ent_id")
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
		return err
	}

	total, err := stm.Count(ctx)
	if err != nil {
		return err
	}
	h.total = uint32(total)

	h.selectOrderUser(stm)
	return nil
}

func (h *queryHandler) scan(ctx context.Context) error {
	return h.stm.Scan(ctx, &h.infos)
}

func (h *queryHandler) formalize() {
	for _, info := range h.infos {
		info.Proportion = func() string { amount, _ := decimal.NewFromString(info.Proportion); return amount.String() }()
		info.MiningpoolType = basetypes.MiningpoolType(basetypes.MiningpoolType_value[info.MiningpoolTypeStr])
		info.CoinType = basetypes.CoinType(basetypes.CoinType_value[info.CoinTypeStr])
	}
}

func (h *Handler) GetOrderUser(ctx context.Context) (*npool.OrderUser, error) {
	handler := &queryHandler{
		Handler: h,
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.queryOrderUser(cli); err != nil {
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

func (h *Handler) GetOrderUsers(ctx context.Context) ([]*npool.OrderUser, uint32, error) {
	handler := &queryHandler{
		Handler: h,
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.queryOrderUsers(ctx, cli); err != nil {
			return err
		}
		handler.stm.
			Offset(int(h.Offset)).
			Limit(int(h.Limit)).
			Order(ent.Desc(orderuserent.FieldUpdatedAt))
		return handler.scan(_ctx)
	})
	if err != nil {
		return nil, 0, err
	}

	handler.formalize()
	return handler.infos, handler.total, nil
}
