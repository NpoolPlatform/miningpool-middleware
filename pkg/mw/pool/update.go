package pool

import (
	"context"
	"fmt"

	basetypes "github.com/NpoolPlatform/message/npool/basetypes/miningpool/v1"
	npool "github.com/NpoolPlatform/message/npool/miningpool/mw/v1/pool"
	poolcrud "github.com/NpoolPlatform/miningpool-middleware/pkg/crud/pool"

	"github.com/NpoolPlatform/miningpool-middleware/pkg/db"
	"github.com/NpoolPlatform/miningpool-middleware/pkg/db/ent"
	poolent "github.com/NpoolPlatform/miningpool-middleware/pkg/db/ent/pool"
)

type updateHandler struct {
	*Handler
}

//nolint:gocyclo
func (h *updateHandler) validateState(info *ent.Pool) error {
	if info.MiningpoolType == basetypes.MiningpoolType_DefaultMiningpoolType.String() {
		return fmt.Errorf("invalid miningpooltype")
	}
	return nil
}

func (h *Handler) UpdatePool(ctx context.Context) (*npool.Pool, error) {
	if h.ID == nil {
		return nil, fmt.Errorf("invalid id")
	}

	handler := &updateHandler{
		Handler: h,
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		info, err := cli.
			Pool.
			Query().
			Where(
				poolent.ID(*h.ID),
			).
			Only(_ctx)
		if err != nil {
			return err
		}

		if err := handler.validateState(info); err != nil {
			return err
		}

		stm, err := poolcrud.UpdateSet(
			info.Update(),
			&poolcrud.Req{
				MiningpoolType: h.MiningpoolType,
				Name:           h.Name,
				Site:           h.Site,
				Description:    h.Description,
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

	return h.GetPool(ctx)
}
