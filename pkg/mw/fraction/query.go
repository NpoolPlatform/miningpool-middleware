package fraction

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/miningpool/v1"
	npool "github.com/NpoolPlatform/message/npool/miningpool/mw/v1/fraction"

	"github.com/NpoolPlatform/miningpool-middleware/pkg/db"
	"github.com/NpoolPlatform/miningpool-middleware/pkg/db/ent"
	"github.com/NpoolPlatform/miningpool-middleware/pkg/db/ent/coin"
	fractionent "github.com/NpoolPlatform/miningpool-middleware/pkg/db/ent/fraction"
	"github.com/NpoolPlatform/miningpool-middleware/pkg/db/ent/gooduser"
	"github.com/NpoolPlatform/miningpool-middleware/pkg/db/ent/orderuser"
	"github.com/NpoolPlatform/miningpool-middleware/pkg/db/ent/pool"

	fractioncrud "github.com/NpoolPlatform/miningpool-middleware/pkg/crud/fraction"
)

type queryHandler struct {
	*Handler
	stm   *ent.FractionSelect
	infos []*npool.Fraction
	total uint32
}

func (h *queryHandler) selectFraction(stm *ent.FractionQuery) {
	h.stm = stm.Select(
		fractionent.FieldID,
		fractionent.FieldEntID,
		fractionent.FieldAppID,
		fractionent.FieldUserID,
		fractionent.FieldOrderUserID,
		fractionent.FieldWithdrawState,
		fractionent.FieldWithdrawAt,
		fractionent.FieldPromisePayAt,
		fractionent.FieldMsg,
		fractionent.FieldCreatedAt,
		fractionent.FieldUpdatedAt,
	)
}

func (h *queryHandler) queryJoin() {
	h.stm.Modify(h.queryJoinCoinAndPool)
}

func (h *queryHandler) queryJoinCoinAndPool(s *sql.Selector) {
	ouT := sql.Table(orderuser.Table)
	guT := sql.Table(gooduser.Table)
	coinT := sql.Table(coin.Table)
	poolT := sql.Table(pool.Table)

	s.LeftJoin(ouT).On(
		s.C(fractionent.FieldOrderUserID),
		ouT.C(orderuser.FieldEntID),
	).Where(
		sql.EQ(ouT.C(orderuser.FieldDeletedAt), 0),
	).LeftJoin(guT).On(
		ouT.C(orderuser.FieldGoodUserID),
		guT.C(gooduser.FieldEntID),
	).Where(
		sql.EQ(guT.C(gooduser.FieldDeletedAt), 0),
	).LeftJoin(coinT).On(
		guT.C(gooduser.FieldPoolCoinTypeID),
		coinT.C(coin.FieldEntID),
	).Where(
		sql.EQ(coinT.C(coin.FieldDeletedAt), 0),
	).LeftJoin(poolT).On(
		coinT.C(coin.FieldPoolID),
		poolT.C(pool.FieldEntID),
	).Where(
		sql.EQ(poolT.C(pool.FieldDeletedAt), 0),
	)
}

func (h *queryHandler) queryFraction(cli *ent.Client) error {
	if h.ID == nil && h.EntID == nil {
		return fmt.Errorf("invalid id")
	}
	stm := cli.Fraction.Query().Where(fractionent.DeletedAt(0))
	if h.ID != nil {
		stm.Where(fractionent.ID(*h.ID))
	}
	if h.EntID != nil {
		stm.Where(fractionent.EntID(*h.EntID))
	}
	h.selectFraction(stm)
	return nil
}

func (h *queryHandler) queryFractions(ctx context.Context, cli *ent.Client) error {
	stm, err := fractioncrud.SetQueryConds(cli.Fraction.Query(), h.Conds)
	if err != nil {
		return err
	}

	total, err := stm.Count(ctx)
	if err != nil {
		return err
	}
	h.total = uint32(total)

	h.selectFraction(stm)
	return nil
}

func (h *queryHandler) scan(ctx context.Context) error {
	return h.stm.Scan(ctx, &h.infos)
}

func (h *queryHandler) formalize() {
	for _, info := range h.infos {
		info.WithdrawState = basetypes.WithdrawState(basetypes.WithdrawState_value[info.WithdrawStateStr])
	}
}

func (h *Handler) GetFraction(ctx context.Context) (*npool.Fraction, error) {
	handler := &queryHandler{
		Handler: h,
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.queryFraction(cli); err != nil {
			return err
		}
		handler.queryJoin()
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

func (h *Handler) GetFractions(ctx context.Context) ([]*npool.Fraction, uint32, error) {
	handler := &queryHandler{
		Handler: h,
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.queryFractions(ctx, cli); err != nil {
			return err
		}
		handler.queryJoin()
		handler.stm.
			Offset(int(h.Offset)).
			Limit(int(h.Limit)).
			Order(ent.Desc(fractionent.FieldUpdatedAt))
		return handler.scan(_ctx)
	})
	if err != nil {
		return nil, 0, err
	}

	handler.formalize()
	return handler.infos, handler.total, nil
}
