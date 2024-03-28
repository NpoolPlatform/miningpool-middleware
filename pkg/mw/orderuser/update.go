package orderuser

import (
	"context"
	"fmt"

	basetypes "github.com/NpoolPlatform/message/npool/basetypes/miningpool/v1"
	npool "github.com/NpoolPlatform/message/npool/miningpool/mw/v1/orderuser"
	orderusercrud "github.com/NpoolPlatform/miningpool-middleware/pkg/crud/orderuser"

	"github.com/NpoolPlatform/miningpool-middleware/pkg/db"
	"github.com/NpoolPlatform/miningpool-middleware/pkg/db/ent"
	orderuserent "github.com/NpoolPlatform/miningpool-middleware/pkg/db/ent/orderuser"
)

type updateHandler struct {
	*Handler
}

//nolint:gocyclo
func (h *updateHandler) validateState(info *ent.OrderUser) error {
	if info.MiningpoolType == basetypes.MiningpoolType_DefaultMiningpoolType.String() {
		return fmt.Errorf("invalid miningpooltype")
	}
	if info.CoinType == basetypes.CoinType_DefaultCoinType.String() {
		return fmt.Errorf("invalid cointype")
	}
	return nil
}

func (h *Handler) UpdateOrderUser(ctx context.Context) (*npool.OrderUser, error) {
	if h.ID == nil {
		return nil, fmt.Errorf("invalid id")
	}

	handler := &updateHandler{
		Handler: h,
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		info, err := cli.
			OrderUser.
			Query().
			Where(
				orderuserent.ID(*h.ID),
			).
			Only(_ctx)
		if err != nil {
			return err
		}

		if err := handler.validateState(info); err != nil {
			return err
		}

		stm, err := orderusercrud.UpdateSet(
			info.Update(),
			&orderusercrud.Req{
				Name:           h.Name,
				RootUserID:     h.RootUserID,
				GoodUserID:     h.GoodUserID,
				AppID:          h.AppID,
				UserID:         h.UserID,
				MiningpoolType: h.MiningpoolType,
				CoinType:       h.CoinType,
				Proportion:     h.Proportion,
				RevenueAddress: h.RevenueAddress,
				ReadPageLink:   h.ReadPageLink,
				AutoPay:        h.AutoPay,
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

	return h.GetOrderUser(ctx)
}
