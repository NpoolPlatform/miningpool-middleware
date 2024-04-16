//nolint:dupl
package rootuser

import (
	"context"
	"time"

	npool "github.com/NpoolPlatform/message/npool/miningpool/mw/v1/rootuser"

	crud "github.com/NpoolPlatform/miningpool-middleware/pkg/crud/rootuser"

	"github.com/NpoolPlatform/miningpool-middleware/pkg/db"
	"github.com/NpoolPlatform/miningpool-middleware/pkg/db/ent"
)

type deleteHandler struct {
	*Handler
}

func (h *deleteHandler) deleteRootUserBase(ctx context.Context, tx *ent.Tx) error {
	now := uint32(time.Now().Unix())
	updateOne, err := crud.UpdateSet(tx.RootUser.UpdateOneID(*h.ID), &crud.Req{DeletedAt: &now})
	if err != nil {
		return err
	}
	_, err = updateOne.Save(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (h *Handler) DeleteRootUser(ctx context.Context) (*npool.RootUser, error) {
	info, err := h.GetRootUser(ctx)
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
		if err := handler.deleteRootUserBase(_ctx, tx); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return info, nil
}
