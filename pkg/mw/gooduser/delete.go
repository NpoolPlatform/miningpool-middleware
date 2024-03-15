//nolint:dupl
package gooduser

import (
	"context"
	"fmt"
	"time"

	npool "github.com/NpoolPlatform/message/npool/miningpool/mw/v1/gooduser"

	crud "github.com/NpoolPlatform/miningpool-middleware/pkg/crud/gooduser"

	"github.com/NpoolPlatform/miningpool-middleware/pkg/db"
	"github.com/NpoolPlatform/miningpool-middleware/pkg/db/ent"

	"github.com/google/uuid"
)

type deleteHandler struct {
	*Handler
}

func (h *deleteHandler) deleteGoodUserBase(ctx context.Context, tx *ent.Tx) error {
	now := uint32(time.Now().Unix())
	updateOne, err := crud.UpdateSet(tx.GoodUser.UpdateOneID(*h.ID), &crud.Req{DeletedAt: &now})
	if err != nil {
		return err
	}
	_, err = updateOne.Save(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (h *Handler) DeleteGoodUser(ctx context.Context) (*npool.GoodUser, error) {
	if h.ID == nil {
		return nil, fmt.Errorf("invalid id")
	}

	info, err := h.GetGoodUser(ctx)
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
		if err := handler.deleteGoodUserBase(_ctx, tx); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return info, nil
}
