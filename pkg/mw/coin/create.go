package coin

import (
	"context"

	npool "github.com/NpoolPlatform/message/npool/miningpool/mw/v1/coin"
	coincrud "github.com/NpoolPlatform/miningpool-middleware/pkg/crud/coin"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	"github.com/NpoolPlatform/miningpool-middleware/pkg/db"
	"github.com/NpoolPlatform/miningpool-middleware/pkg/db/ent"

	"github.com/google/uuid"
)

func (h *Handler) CreateCoin(ctx context.Context) (*npool.Coin, error) {
	id := uuid.New()
	if h.EntID == nil {
		h.EntID = &id
	}

	err := db.WithClient(ctx, func(ctx context.Context, cli *ent.Client) error {
		info, err := coincrud.CreateSet(
			cli.Coin.Create(),
			&coincrud.Req{
				EntID:            h.EntID,
				CoinType:         h.CoinType,
				MiningpoolType:   h.MiningpoolType,
				RevenueTypes:     h.RevenueTypes,
				Site:             h.Site,
				FeeRate:          h.FeeRate,
				FixedRevenueAble: h.FixedRevenueAble,
				Threshold:        h.Threshold,
				Remark:           h.Remark,
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
	return h.GetCoin(ctx)
}

func (h *Handler) CreateCoins(ctx context.Context) ([]*npool.Coin, error) {
	ids := []uuid.UUID{}

	err := db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		for _, req := range h.Reqs {
			info, err := coincrud.CreateSet(tx.Coin.Create(), req).Save(_ctx)
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

	h.Conds = &coincrud.Conds{
		EntIDs: &cruder.Cond{Op: cruder.IN, Val: ids},
	}
	h.Offset = 0
	h.Limit = int32(len(ids))

	infos, _, err := h.GetCoins(ctx)
	return infos, err
}
