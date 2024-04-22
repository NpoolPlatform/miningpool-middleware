//nolint:dupl
package orderuser

import (
	"context"
	"time"

	crud "github.com/NpoolPlatform/miningpool-middleware/pkg/crud/orderuser"

	"github.com/NpoolPlatform/miningpool-middleware/pkg/db"
	"github.com/NpoolPlatform/miningpool-middleware/pkg/db/ent"
)

type deleteHandler struct {
	*Handler
}

func (h *deleteHandler) deleteOrderUserBase(ctx context.Context, tx *ent.Tx) error {
	now := uint32(time.Now().Unix())
	updateOne, err := crud.UpdateSet(tx.OrderUser.UpdateOneID(*h.ID), &crud.Req{DeletedAt: &now})
	if err != nil {
		return err
	}
	_, err = updateOne.Save(ctx)
	if err != nil {
		return err
	}

	return nil
}

func (h *Handler) DeleteOrderUser(ctx context.Context) error {
	info, err := h.GetOrderUser(ctx)
	if err != nil {
		return err
	}

	if info == nil {
		return nil
	}
	h.ID = &info.ID

	zeroProportion := "0"
	updateH, err := NewHandler(ctx,
		WithID(h.ID, true),
		WithEntID(&info.EntID, true),
		WithProportion(&zeroProportion, true),
	)
	if err != nil {
		return err
	}

	err = updateH.UpdateOrderUser(ctx)
	if err != nil {
		return err
	}

	handler := &deleteHandler{
		Handler: h,
	}

	return db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		if err := handler.deleteOrderUserBase(_ctx, tx); err != nil {
			return err
		}
		return nil
	})
}
