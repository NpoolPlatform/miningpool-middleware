package pool

import (
	"context"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	npool "github.com/NpoolPlatform/message/npool/miningpool/mw/v1/pool"
	poolcrud "github.com/NpoolPlatform/miningpool-middleware/pkg/crud/pool"
	"github.com/NpoolPlatform/miningpool-middleware/pkg/db"
	"github.com/NpoolPlatform/miningpool-middleware/pkg/db/ent"
	"github.com/google/uuid"
)

func (h *Handler) CreatePool(ctx context.Context) (*npool.Pool, error) {
	err := db.WithClient(ctx, func(ctx context.Context, cli *ent.Client) error {
		info, err := poolcrud.CreateSet(
			cli.Pool.Create(),
			&poolcrud.Req{
				EntID:          h.EntID,
				MiningpoolType: h.MiningpoolType,
				Name:           h.Name,
				Site:           h.Site,
				Description:    h.Description,
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
	return h.GetPool(ctx)
}

func (h *Handler) CreatePools(ctx context.Context) ([]*npool.Pool, error) {
	ids := []uuid.UUID{}
	err := db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		for _, req := range h.Reqs {
			info, err := poolcrud.CreateSet(tx.Pool.Create(), req).Save(_ctx)
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

	h.Conds = &poolcrud.Conds{
		EntIDs: &cruder.Cond{Op: cruder.IN, Val: ids},
	}
	h.Offset = 0
	h.Limit = int32(len(ids))

	infos, _, err := h.GetPools(ctx)
	return infos, err
}
