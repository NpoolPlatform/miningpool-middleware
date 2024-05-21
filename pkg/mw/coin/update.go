package coin

import (
	"context"
	"fmt"

	"github.com/NpoolPlatform/miningpool-middleware/pkg/db"
	"github.com/NpoolPlatform/miningpool-middleware/pkg/db/ent"
	"github.com/google/uuid"
)

func (h *Handler) UpdateCoin(ctx context.Context) error {
	info, err := h.GetCoin(ctx)
	if err != nil {
		return err
	}

	if info == nil {
		return fmt.Errorf("invalid id or ent_id")
	}

	poolID, err := uuid.Parse(info.PoolID)
	if err != nil {
		return fmt.Errorf("invalid poolid")
	}

	sqlH := h.newSQLHandler()
	sqlH.BondPoolID = &poolID
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
			return fmt.Errorf("failed to update pool: %v", err)
		}
		return nil
	})
}
