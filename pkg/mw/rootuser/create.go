package rootuser

import (
	"context"

	npool "github.com/NpoolPlatform/message/npool/miningpool/mw/v1/rootuser"
	rootusercrud "github.com/NpoolPlatform/miningpool-middleware/pkg/crud/rootuser"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	"github.com/NpoolPlatform/miningpool-middleware/pkg/db"
	"github.com/NpoolPlatform/miningpool-middleware/pkg/db/ent"

	"github.com/google/uuid"
)

func (h *Handler) CreateRootUser(ctx context.Context) (*npool.RootUser, error) {
	id := uuid.New()
	if h.EntID == nil {
		h.EntID = &id
	}

	err := db.WithClient(ctx, func(ctx context.Context, cli *ent.Client) error {
		info, err := rootusercrud.CreateSet(
			cli.RootUser.Create(),
			&rootusercrud.Req{
				EntID:          h.EntID,
				Name:           h.Name,
				MiningpoolType: h.MiningpoolType,
				Email:          h.Email,
				AuthToken:      h.AuthToken,
				Authed:         h.Authed,
				Remark:         h.Remark,
			},
		).Save(ctx)
		if err != nil {
			return err
		}
		h.ID = &info.ID
		return nil
	})
	if err != nil {
		return nil, err
	}

	return h.GetRootUser(ctx)
}

func (h *Handler) CreateRootUsers(ctx context.Context) ([]*npool.RootUser, error) {
	ids := []uuid.UUID{}

	err := db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		for _, req := range h.Reqs {
			info, err := rootusercrud.CreateSet(tx.RootUser.Create(), req).Save(_ctx)
			if err != nil {
				return err
			}
			ids = append(ids, info.EntID)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	h.Conds = &rootusercrud.Conds{
		EntIDs: &cruder.Cond{Op: cruder.IN, Val: ids},
	}
	h.Offset = 0
	h.Limit = int32(len(ids))

	infos, _, err := h.GetRootUsers(ctx)
	return infos, err
}
