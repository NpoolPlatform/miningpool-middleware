//nolint:dupl
package fraction

import (
	"context"
	"time"

	npool "github.com/NpoolPlatform/message/npool/miningpool/mw/v1/fraction"

	crud "github.com/NpoolPlatform/miningpool-middleware/pkg/crud/fraction"

	"github.com/NpoolPlatform/miningpool-middleware/pkg/db"
	"github.com/NpoolPlatform/miningpool-middleware/pkg/db/ent"
)

type deleteHandler struct {
	*Handler
}

func (h *deleteHandler) deleteFractionBase(ctx context.Context, tx *ent.Tx) error {
	now := uint32(time.Now().Unix())
	updateOne, err := crud.UpdateSet(tx.Fraction.UpdateOneID(*h.ID), &crud.Req{DeletedAt: &now})
	if err != nil {
		return err
	}
	_, err = updateOne.Save(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (h *Handler) DeleteFraction(ctx context.Context) (*npool.Fraction, error) {
	info, err := h.GetFraction(ctx)
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
		if err := handler.deleteFractionBase(_ctx, tx); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return info, nil
}
