package gooduser

import (
	"context"

	npool "github.com/NpoolPlatform/message/npool/miningpool/mw/v1/gooduser"
	goodusercrud "github.com/NpoolPlatform/miningpool-middleware/pkg/crud/gooduser"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	"github.com/NpoolPlatform/miningpool-middleware/pkg/db"
	"github.com/NpoolPlatform/miningpool-middleware/pkg/db/ent"

	"github.com/google/uuid"
)

func (h *Handler) CreateGoodUser(ctx context.Context) (*npool.GoodUser, error) {
	id := uuid.New()
	if h.EntID == nil {
		h.EntID = &id
	}

	err := db.WithClient(ctx, func(ctx context.Context, cli *ent.Client) error {
		info, err := goodusercrud.CreateSet(
			cli.GoodUser.Create(),
			&goodusercrud.Req{
				EntID:          h.EntID,
				Name:           h.Name,
				RootUserID:     h.RootUserID,
				MiningpoolType: h.MiningpoolType,
				CoinType:       h.CoinType,
				HashRate:       h.HashRate,
				ReadPageLink:   h.ReadPageLink,
				RevenueType:    h.RevenueType,
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

	return h.GetGoodUser(ctx)
}

func (h *Handler) CreateGoodUsers(ctx context.Context) ([]*npool.GoodUser, error) {
	ids := []uuid.UUID{}

	err := db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		for _, req := range h.Reqs {
			info, err := goodusercrud.CreateSet(tx.GoodUser.Create(), req).Save(_ctx)
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

	h.Conds = &goodusercrud.Conds{
		EntIDs: &cruder.Cond{Op: cruder.IN, Val: ids},
	}
	h.Offset = 0
	h.Limit = int32(len(ids))

	infos, _, err := h.GetGoodUsers(ctx)
	return infos, err
}
