package gooduser

import (
	"context"
	"fmt"

	"github.com/NpoolPlatform/miningpool-middleware/pkg/db"
	"github.com/NpoolPlatform/miningpool-middleware/pkg/db/ent"
	"github.com/google/uuid"
)

func (h *Handler) UpdateGoodUser(ctx context.Context) error {
	info, err := h.GetGoodUser(ctx)
	if err != nil {
		return err
	}

	if info == nil {
		return fmt.Errorf("invalid id or ent_id")
	}

	coinID, err := uuid.Parse(info.PoolCoinTypeID)
	if err != nil {
		return err
	}
	sqlH := h.newSQLHandler()
	sqlH.BondName = &info.Name
	sqlH.BondPoolCoinTypeID = &coinID

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
			return fmt.Errorf("failed to update gooduser: %v", err)
		}
		return nil
	})
}
