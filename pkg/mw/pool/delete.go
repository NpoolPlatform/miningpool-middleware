//nolint:dupl
package pool

import (
	"context"
	"time"

	"github.com/NpoolPlatform/go-service-framework/pkg/wlog"
	crud "github.com/NpoolPlatform/miningpool-middleware/pkg/crud/pool"

	"github.com/NpoolPlatform/miningpool-middleware/pkg/db"
	"github.com/NpoolPlatform/miningpool-middleware/pkg/db/ent"
)

type deleteHandler struct {
	*Handler
}

func (h *deleteHandler) deletePoolBase(ctx context.Context, tx *ent.Tx) error {
	now := uint32(time.Now().Unix())
	updateOne, err := crud.UpdateSet(tx.Pool.UpdateOneID(*h.ID), &crud.Req{DeletedAt: &now})
	if err != nil {
		return wlog.WrapError(err)
	}
	_, err = updateOne.Save(ctx)
	if err != nil {
		return wlog.WrapError(err)
	}
	return nil
}

func (h *Handler) DeletePool(ctx context.Context) error {
	info, err := h.GetPool(ctx)
	if err != nil {
		return wlog.WrapError(err)
	}
	if info == nil {
		return nil
	}

	h.ID = &info.ID
	handler := &deleteHandler{
		Handler: h,
	}

	return db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		if err := handler.deletePoolBase(_ctx, tx); err != nil {
			return wlog.WrapError(err)
		}
		return nil
	})
}
