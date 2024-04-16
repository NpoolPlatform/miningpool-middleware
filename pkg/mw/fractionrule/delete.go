//nolint:dupl
package fractionrule

import (
	"context"
	"time"

	npool "github.com/NpoolPlatform/message/npool/miningpool/mw/v1/fractionrule"

	crud "github.com/NpoolPlatform/miningpool-middleware/pkg/crud/fractionrule"

	"github.com/NpoolPlatform/miningpool-middleware/pkg/db"
	"github.com/NpoolPlatform/miningpool-middleware/pkg/db/ent"
)

type deleteHandler struct {
	*Handler
}

func (h *deleteHandler) deleteFractionRuleBase(ctx context.Context, tx *ent.Tx) error {
	now := uint32(time.Now().Unix())
	updateOne, err := crud.UpdateSet(tx.FractionRule.UpdateOneID(*h.ID), &crud.Req{DeletedAt: &now})
	if err != nil {
		return err
	}
	_, err = updateOne.Save(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (h *Handler) DeleteFractionRule(ctx context.Context) (*npool.FractionRule, error) {
	info, err := h.GetFractionRule(ctx)
	if err != nil {
		return nil, err
	}

	if info == nil {
		return nil, nil
	}

	h.ID = &info.ID
	handler := &deleteHandler{
		Handler: h,
	}

	err = db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		if err := handler.deleteFractionRuleBase(_ctx, tx); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return info, nil
}
