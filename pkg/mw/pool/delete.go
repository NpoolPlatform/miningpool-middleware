//nolint:dupl
package pool

import (
	"context"
	"fmt"
	"time"

	npool "github.com/NpoolPlatform/message/npool/miningpool/mw/v1/pool"

	crud "github.com/NpoolPlatform/miningpool-middleware/pkg/crud/pool"

	"github.com/NpoolPlatform/miningpool-middleware/pkg/db"
	"github.com/NpoolPlatform/miningpool-middleware/pkg/db/ent"

	"github.com/google/uuid"
)

type deleteHandler struct {
	*Handler
}

func (h *deleteHandler) deletePoolBase(ctx context.Context, tx *ent.Tx) error {
	now := uint32(time.Now().Unix())
	updateOne, err := crud.UpdateSet(tx.Pool.UpdateOneID(*h.ID), &crud.Req{DeletedAt: &now})
	if err != nil {
		return err
	}
	_, err = updateOne.Save(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (h *Handler) DeletePool(ctx context.Context) (*npool.Pool, error) {
	if h.ID == nil {
		return nil, fmt.Errorf("invalid id")
	}

	info, err := h.GetPool(ctx)
	if err != nil {
		return nil, err
	}

	entID, err := uuid.Parse(info.EntID)
	if err != nil {
		return nil, err
	}
	h.EntID = &entID
	handler := &deleteHandler{
		Handler: h,
	}

	err = db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		if err := handler.deletePoolBase(_ctx, tx); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return info, nil
}