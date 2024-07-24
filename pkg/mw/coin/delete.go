//nolint:dupl
package coin

import (
	"context"
	"time"

	"github.com/NpoolPlatform/go-service-framework/pkg/wlog"
	crud "github.com/NpoolPlatform/miningpool-middleware/pkg/crud/coin"

	"github.com/NpoolPlatform/miningpool-middleware/pkg/db"
	"github.com/NpoolPlatform/miningpool-middleware/pkg/db/ent"
)

type deleteHandler struct {
	*Handler
}

func (h *deleteHandler) deleteCoinBase(ctx context.Context, tx *ent.Tx) error {
	now := uint32(time.Now().Unix())
	updateOne, err := crud.UpdateSet(tx.Coin.UpdateOneID(*h.ID), &crud.Req{DeletedAt: &now})
	if err != nil {
		return wlog.WrapError(err)
	}

	_, err = updateOne.Save(ctx)
	if err != nil {
		return wlog.WrapError(err)
	}
	return nil
}

func (h *Handler) DeleteCoin(ctx context.Context) error {
	info, err := h.GetCoin(ctx)
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
		if err := handler.deleteCoinBase(_ctx, tx); err != nil {
			return wlog.WrapError(err)
		}
		return nil
	})
}
