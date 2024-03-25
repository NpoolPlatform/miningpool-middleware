package coin

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	redis2 "github.com/NpoolPlatform/go-service-framework/pkg/redis"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	npool "github.com/NpoolPlatform/message/npool/miningpool/mw/v1/coin"
	coincrud "github.com/NpoolPlatform/miningpool-middleware/pkg/crud/coin"
	"github.com/NpoolPlatform/miningpool-middleware/pkg/db"
	"github.com/NpoolPlatform/miningpool-middleware/pkg/db/ent"
	"github.com/google/uuid"
)

func (h *Handler) CreateCoin(ctx context.Context) (*npool.Coin, error) {
	lockKey := basetypes.Prefix_PrefixCreateMiningpoolCoin.String()
	if err := redis2.TryLock(lockKey, 0); err != nil {
		return nil, err
	}
	defer func() {
		err := redis2.Unlock(lockKey)
		if err != nil {
			logger.Sugar().Error(err)
		}
	}()
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
	lockKey := basetypes.Prefix_PrefixCreateMiningpoolCoin.String()
	if err := redis2.TryLock(lockKey, 0); err != nil {
		return nil, err
	}
	defer func() {
		err := redis2.Unlock(lockKey)
		if err != nil {
			logger.Sugar().Error(err)
		}
	}()
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
