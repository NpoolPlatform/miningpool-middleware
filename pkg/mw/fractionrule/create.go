package fractionrule

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	redis2 "github.com/NpoolPlatform/go-service-framework/pkg/redis"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	npool "github.com/NpoolPlatform/message/npool/miningpool/mw/v1/fractionrule"
	fractionrulecrud "github.com/NpoolPlatform/miningpool-middleware/pkg/crud/fractionrule"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	"github.com/NpoolPlatform/miningpool-middleware/pkg/db"
	"github.com/NpoolPlatform/miningpool-middleware/pkg/db/ent"

	"github.com/google/uuid"
)

func (h *Handler) CreateFractionRule(ctx context.Context) (*npool.FractionRule, error) {
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
		info, err := fractionrulecrud.CreateSet(
			cli.FractionRule.Create(),
			&fractionrulecrud.Req{
				EntID:            h.EntID,
				MiningpoolType:   h.MiningpoolType,
				CoinType:         h.CoinType,
				WithdrawInterval: h.WithdrawInterval,
				MinAmount:        h.MinAmount,
				WithdrawRate:     h.WithdrawRate,
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

	return h.GetFractionRule(ctx)
}

func (h *Handler) CreateFractionRules(ctx context.Context) ([]*npool.FractionRule, error) {
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
			info, err := fractionrulecrud.CreateSet(tx.FractionRule.Create(), req).Save(_ctx)
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

	h.Conds = &fractionrulecrud.Conds{
		EntIDs: &cruder.Cond{Op: cruder.IN, Val: ids},
	}
	h.Offset = 0
	h.Limit = int32(len(ids))

	infos, _, err := h.GetFractionRules(ctx)
	return infos, err
}
