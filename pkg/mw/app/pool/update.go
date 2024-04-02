package apppool

import (
	"context"
	"fmt"

	npool "github.com/NpoolPlatform/message/npool/miningpool/mw/v1/app/pool"
	apppoolcrud "github.com/NpoolPlatform/miningpool-middleware/pkg/crud/app/pool"

	"github.com/NpoolPlatform/miningpool-middleware/pkg/db"
	"github.com/NpoolPlatform/miningpool-middleware/pkg/db/ent"
	apppoolent "github.com/NpoolPlatform/miningpool-middleware/pkg/db/ent/apppool"
)

func (h *Handler) UpdatePool(ctx context.Context) (*npool.Pool, error) {
	if h.ID == nil {
		return nil, fmt.Errorf("invalid id")
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		info, err := cli.
			AppPool.
			Query().
			Where(
				apppoolent.ID(*h.ID),
			).
			Only(_ctx)
		if err != nil {
			return err
		}

		stm, err := apppoolcrud.UpdateSet(
			info.Update(),
			&apppoolcrud.Req{
				AppID:  h.AppID,
				PoolID: h.PoolID,
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
