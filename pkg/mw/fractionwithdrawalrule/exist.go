package fractionwithdrawalrule

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/wlog"
	fractionwithdrawalrulecrud "github.com/NpoolPlatform/miningpool-middleware/pkg/crud/fractionwithdrawalrule"
	"github.com/NpoolPlatform/miningpool-middleware/pkg/db"
	"github.com/NpoolPlatform/miningpool-middleware/pkg/db/ent"
	fractionwithdrawalruleent "github.com/NpoolPlatform/miningpool-middleware/pkg/db/ent/fractionwithdrawalrule"
)

func (h *Handler) ExistFractionWithdrawalRule(ctx context.Context) (bool, error) {
	exist := false
	var err error

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		exist, err = cli.
			FractionWithdrawalRule.
			Query().
			Where(
				fractionwithdrawalruleent.EntID(*h.EntID),
				fractionwithdrawalruleent.DeletedAt(0),
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

func (h *Handler) ExistFractionWithdrawalRuleConds(ctx context.Context) (bool, error) {
	exist := false
	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		stm, err := fractionwithdrawalrulecrud.SetQueryConds(cli.FractionWithdrawalRule.Query(), h.Conds)
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
