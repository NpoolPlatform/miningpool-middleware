package fractionwithdrawal

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/wlog"
	fractionwithdrawalcrud "github.com/NpoolPlatform/miningpool-middleware/pkg/crud/fractionwithdrawal"
	"github.com/NpoolPlatform/miningpool-middleware/pkg/db"
	"github.com/NpoolPlatform/miningpool-middleware/pkg/db/ent"
	fractionwithdrawalent "github.com/NpoolPlatform/miningpool-middleware/pkg/db/ent/fractionwithdrawal"
)

func (h *Handler) ExistFractionWithdrawal(ctx context.Context) (bool, error) {
	exist := false
	var err error

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		exist, err = cli.
			FractionWithdrawal.
			Query().
			Where(
				fractionwithdrawalent.EntID(*h.EntID),
				fractionwithdrawalent.DeletedAt(0),
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

func (h *Handler) ExistFractionWithdrawalConds(ctx context.Context) (bool, error) {
	exist := false
	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		stm, err := fractionwithdrawalcrud.SetQueryConds(cli.FractionWithdrawal.Query(), h.Conds)
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
