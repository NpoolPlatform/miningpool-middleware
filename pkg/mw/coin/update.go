package coin

import (
	"context"
	"fmt"

	basetypes "github.com/NpoolPlatform/message/npool/basetypes/miningpool/v1"
	npool "github.com/NpoolPlatform/message/npool/miningpool/mw/v1/coin"
	coincrud "github.com/NpoolPlatform/miningpool-middleware/pkg/crud/coin"

	"github.com/NpoolPlatform/miningpool-middleware/pkg/db"
	"github.com/NpoolPlatform/miningpool-middleware/pkg/db/ent"
	coinent "github.com/NpoolPlatform/miningpool-middleware/pkg/db/ent/coin"
)

type updateHandler struct {
	*Handler
}

//nolint:gocyclo
func (h *updateHandler) validateState(info *ent.Coin) error {
	if info.MiningpoolType == basetypes.MiningpoolType_DefaultMiningpoolType.String() {
		return fmt.Errorf("invalid miningpooltype")
	}
	return nil
}

func (h *Handler) UpdateCoin(ctx context.Context) (*npool.Coin, error) {
	if h.ID == nil {
		return nil, fmt.Errorf("invalid id")
	}

	handler := &updateHandler{
		Handler: h,
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		info, err := cli.
			Coin.
			Query().
			Where(
				coinent.ID(*h.ID),
			).
			Only(_ctx)
		if err != nil {
			return err
		}

		if err := handler.validateState(info); err != nil {
			return err
		}

		stm, err := coincrud.UpdateSet(
			info.Update(),
			&coincrud.Req{
				CoinType:         h.CoinType,
				MiningpoolType:   h.MiningpoolType,
				RevenueTypes:     h.RevenueTypes,
				FeeRate:          h.FeeRate,
				FixedRevenueAble: h.FixedRevenueAble,
				Threshold:        h.Threshold,
				Remark:           h.Remark,
			},
		)
		if err != nil {
			return err
		}
		if _, err := stm.Save(_ctx); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return h.GetCoin(ctx)
}
