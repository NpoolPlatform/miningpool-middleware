package gooduser

import (
	"context"

	goodusercrud "github.com/NpoolPlatform/miningpool-middleware/pkg/crud/gooduser"
	"github.com/NpoolPlatform/miningpool-middleware/pkg/db"
	"github.com/NpoolPlatform/miningpool-middleware/pkg/db/ent"
	gooduserent "github.com/NpoolPlatform/miningpool-middleware/pkg/db/ent/gooduser"
)

func (h *Handler) ExistGoodUser(ctx context.Context) (bool, error) {
	exist := false
	var err error

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		exist, err = cli.
			GoodUser.
			Query().
			Where(
				gooduserent.EntID(*h.EntID),
				gooduserent.DeletedAt(0),
			).
			Exist(_ctx)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return false, err
	}
	return exist, nil
}

func (h *Handler) ExistGoodUserConds(ctx context.Context) (bool, error) {
	exist := false
	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		stm, err := goodusercrud.SetQueryConds(cli.GoodUser.Query(), h.Conds)
		if err != nil {
			return err
		}
		exist, err = stm.Exist(_ctx)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return false, err
	}
	return exist, nil
}
