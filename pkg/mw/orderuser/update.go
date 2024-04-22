package orderuser

import (
	"context"
	"fmt"

	"github.com/NpoolPlatform/miningpool-middleware/pkg/db"
	"github.com/NpoolPlatform/miningpool-middleware/pkg/db/ent"
)

func (h *Handler) UpdateOrderUser(ctx context.Context) error {
	info, err := h.GetOrderUser(ctx)
	if err != nil {
		return err
	}

	if info == nil {
		return fmt.Errorf("invalid id or ent_id")
	}

	if h.MiningpoolType == nil {
		h.MiningpoolType = &info.MiningpoolType
	}

	if h.Name == nil {
		h.Name = &info.Name
	}

	sqlH := newSQLHandler(h)
	sql, err := sqlH.genUpdateSQL()
	if err != nil {
		return err
	}

	return db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		rc, err := tx.ExecContext(ctx, sql)
		if err != nil {
			return err
		}

		if n, err := rc.RowsAffected(); err != nil || n != 1 {
			return fmt.Errorf("failed to update orderuser: %v", err)
		}
		return nil
	})
}
