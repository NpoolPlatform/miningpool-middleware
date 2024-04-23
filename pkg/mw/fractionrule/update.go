package fractionrule

import (
	"context"
	"fmt"

	"github.com/NpoolPlatform/miningpool-middleware/pkg/db"
	"github.com/NpoolPlatform/miningpool-middleware/pkg/db/ent"
)

func (h *Handler) UpdateFractionRule(ctx context.Context) error {
	info, err := h.GetFractionRule(ctx)
	if err != nil {
		return err
	}

	if info == nil {
		return fmt.Errorf("invalid id or ent_id")
	}

	sqlH := h.newSQLHandler()
	sqlH.BondMiningpoolType = &info.MiningpoolType
	sqlH.BondCoinType = &info.CoinType

	return db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		sql, err := sqlH.genUpdateSQL()
		if err != nil {
			return err
		}

		rc, err := tx.ExecContext(ctx, sql)
		if err != nil {
			return err
		}

		if n, err := rc.RowsAffected(); err != nil || n != 1 {
			return fmt.Errorf("failed to update fractionrule: %v", err)
		}
		return nil
	})
}
