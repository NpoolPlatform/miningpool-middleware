//nolint:dupl
package orderuser

import (
	"context"
	"fmt"
	"time"

	npool "github.com/NpoolPlatform/message/npool/miningpool/mw/v1/orderuser"

	crud "github.com/NpoolPlatform/miningpool-middleware/pkg/crud/orderuser"

	"github.com/NpoolPlatform/miningpool-middleware/pkg/db"
	"github.com/NpoolPlatform/miningpool-middleware/pkg/db/ent"

	"github.com/google/uuid"
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

func (h *Handler) DeleteOrderUser(ctx context.Context) (*npool.OrderUser, error) {
	if h.ID == nil {
		return nil, fmt.Errorf("invalid id")
	}

	info, err := h.GetOrderUser(ctx)
	if err != nil {
		return nil, err
	}

	if info == nil {
		return nil, fmt.Errorf("have no orderuser %v", h.ID)
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
		if err := handler.deleteOrderUserBase(_ctx, tx); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return info, nil
}
