package coin

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/wlog"
	coincrud "github.com/NpoolPlatform/miningpool-middleware/pkg/crud/coin"
	"github.com/NpoolPlatform/miningpool-middleware/pkg/db"
	"github.com/NpoolPlatform/miningpool-middleware/pkg/db/ent"
	coinent "github.com/NpoolPlatform/miningpool-middleware/pkg/db/ent/coin"
)

func (h *Handler) ExistCoin(ctx context.Context) (bool, error) {
	exist := false
	var err error

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		exist, err = cli.
			Coin.
			Query().
			Where(
				coinent.EntID(*h.EntID),
				coinent.DeletedAt(0),
			).
			Exist(_ctx)
		if err != nil {
			return wlog.WrapError(err)
		}
		return nil
	})
	if err != nil {
		return false, wlog.WrapError(err)
	}
	return exist, nil
}

func (h *Handler) ExistCoinConds(ctx context.Context) (bool, error) {
	exist := false
	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		stm, err := coincrud.SetQueryConds(cli.Coin.Query(), h.Conds)
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
		return false, wlog.WrapError(err)
	}
	return exist, nil
}
