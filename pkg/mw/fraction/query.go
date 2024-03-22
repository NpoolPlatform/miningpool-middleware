package fraction

import (
	"context"
	"fmt"

	basetypes "github.com/NpoolPlatform/message/npool/basetypes/miningpool/v1"
	npool "github.com/NpoolPlatform/message/npool/miningpool/mw/v1/fraction"

	"github.com/NpoolPlatform/miningpool-middleware/pkg/db"
	"github.com/NpoolPlatform/miningpool-middleware/pkg/db/ent"
	fractionent "github.com/NpoolPlatform/miningpool-middleware/pkg/db/ent/fraction"

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
		fractionent.FieldWithdrawTime,
		fractionent.FieldPayTime,
		fractionent.FieldMsg,
		fractionent.FieldCreatedAt,
		fractionent.FieldUpdatedAt,
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
