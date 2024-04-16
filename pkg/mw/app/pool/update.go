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

	if appID, err := uuid.Parse(info.AppID); h.AppID == nil && err != nil {
		h.AppID = &appID
	} else if err != nil {
		return fmt.Errorf("invalid appid")
	}

	if poolID, err := uuid.Parse(info.PoolID); h.PoolID == nil && err != nil {
		h.PoolID = &poolID
	} else if err != nil {
		return fmt.Errorf("invalid poolid")
	}

	return db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		sql, err := h.genUpdateSQL()
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
