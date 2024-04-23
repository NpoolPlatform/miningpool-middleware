package apppool

import (
	"context"
	"fmt"

	"github.com/NpoolPlatform/miningpool-middleware/pkg/db"
	"github.com/NpoolPlatform/miningpool-middleware/pkg/db/ent"
	"github.com/google/uuid"
)

//nolint:gocognit
func (h *Handler) UpdatePool(ctx context.Context) error {
	info, err := h.GetPool(ctx)
	if err != nil {
		return err
	}

	if info == nil {
		return fmt.Errorf("invalid id or ent_id")
	}

	sqlH := h.newSQLHandler()
	if appID, err := uuid.Parse(info.AppID); err == nil {
		sqlH.BondAppID = &appID
	} else {
		return fmt.Errorf("invalid appid")
	}

	if poolID, err := uuid.Parse(info.PoolID); err == nil {
		sqlH.BondPoolID = &poolID
	} else {
		return fmt.Errorf("invalid poolid")
	}

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
			return fmt.Errorf("failed to update rootuser: %v", err)
		}
		return nil
	})
}
