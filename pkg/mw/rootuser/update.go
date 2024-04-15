package rootuser

import (
	"context"
	"fmt"

	"github.com/NpoolPlatform/miningpool-middleware/pkg/db"
	"github.com/NpoolPlatform/miningpool-middleware/pkg/db/ent"
)

func (h *Handler) UpdateRootUser(ctx context.Context) error {
	info, err := h.GetRootUser(ctx)
	if err != nil {
		return err
	}

	if h.MiningpoolType == nil {
		h.MiningpoolType = &info.MiningpoolType
	}
	if h.Email == nil {
		h.Email = &info.Email
	}
	if h.Name == nil {
		h.Name = &info.Name
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
			return fmt.Errorf("fail update rootuser: %v", err)
		}
		return nil
	})
}
