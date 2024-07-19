package apppool

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/wlog"
	apppoolcrud "github.com/NpoolPlatform/miningpool-middleware/pkg/crud/app/pool"
	"github.com/NpoolPlatform/miningpool-middleware/pkg/db"
	"github.com/NpoolPlatform/miningpool-middleware/pkg/db/ent"
	apppoolent "github.com/NpoolPlatform/miningpool-middleware/pkg/db/ent/apppool"
)

func (h *Handler) ExistPool(ctx context.Context) (bool, error) {
	exist := false
	var err error

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		exist, err = cli.
			AppPool.
			Query().
			Where(
				apppoolent.EntID(*h.EntID),
				apppoolent.DeletedAt(0),
			).
			Exist(_ctx)
		if err != nil {
			return wlog.WrapError(err)
		}
		return nil
	})
	if err != nil {
		return false, err
	}
	return exist, nil
}

func (h *Handler) ExistPoolConds(ctx context.Context) (bool, error) {
	exist := false
	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		stm, err := apppoolcrud.SetQueryConds(cli.AppPool.Query(), h.Conds)
		if err != nil {
			return wlog.WrapError(err)
		}
		exist, err = stm.Exist(_ctx)
		if err != nil {
			return wlog.WrapError(err)
		}
		return nil
	})
	if err != nil {
		return false, err
	}
	return exist, nil
}
