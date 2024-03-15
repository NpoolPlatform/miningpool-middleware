package gooduser

import (
	"context"
	"fmt"

	basetypes "github.com/NpoolPlatform/message/npool/basetypes/miningpool/v1"
	npool "github.com/NpoolPlatform/message/npool/miningpool/mw/v1/gooduser"
	goodusercrud "github.com/NpoolPlatform/miningpool-middleware/pkg/crud/gooduser"

	"github.com/NpoolPlatform/miningpool-middleware/pkg/db"
	"github.com/NpoolPlatform/miningpool-middleware/pkg/db/ent"
	gooduserent "github.com/NpoolPlatform/miningpool-middleware/pkg/db/ent/gooduser"
)

type updateHandler struct {
	*Handler
}

//nolint:gocyclo
func (h *updateHandler) validateState(info *ent.GoodUser) error {
	if info.MiningpoolType == basetypes.MiningpoolType_DefaultMiningpoolType.String() {
		return fmt.Errorf("invalid miningpooltype")
	}
	if info.CoinType == basetypes.CoinType_DefaultCoinType.String() {
		return fmt.Errorf("invalid cointype")
	}
	return nil
}

func (h *Handler) UpdateGoodUser(ctx context.Context) (*npool.GoodUser, error) {
	if h.ID == nil {
		return nil, fmt.Errorf("invalid id")
	}

	handler := &updateHandler{
		Handler: h,
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		info, err := cli.
			GoodUser.
			Query().
			Where(
				gooduserent.ID(*h.ID),
			).
			Only(_ctx)
		if err != nil {
			return err
		}

		if err := handler.validateState(info); err != nil {
			return err
		}

		stm, err := goodusercrud.UpdateSet(
			info.Update(),
			&goodusercrud.Req{
				Name:           h.Name,
				RootUserID:     h.RootUserID,
				GoodID:         h.GoodID,
				MiningpoolType: h.MiningpoolType,
				CoinType:       h.CoinType,
				HashRate:       h.HashRate,
				ReadPageLink:   h.ReadPageLink,
				RevenueType:    h.RevenueType,
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

	return h.GetGoodUser(ctx)
}
