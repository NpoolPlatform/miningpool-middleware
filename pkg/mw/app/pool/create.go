package apppool

import (
	"context"
	"fmt"

	"github.com/NpoolPlatform/miningpool-middleware/pkg/db"
	"github.com/NpoolPlatform/miningpool-middleware/pkg/db/ent"
	"github.com/NpoolPlatform/miningpool-middleware/pkg/mw/pool"
	"github.com/google/uuid"
)

func (h *Handler) CreatePool(ctx context.Context) error {
	poolID := h.PoolID.String()
	poolH, err := pool.NewHandler(ctx, pool.WithEntID(&poolID, true))
	if err != nil {
		return err
	}

	exist, err := poolH.ExistPool(ctx)
	if err != nil {
		return err
	}
	if !exist {
		return fmt.Errorf("invalid pool id")
	}

	if h.EntID == nil {
		h.EntID = func() *uuid.UUID { uid := uuid.New(); return &uid }()
	}

	sqlH := h.newSQLHandler()

	return db.WithTx(ctx, func(ctx context.Context, tx *ent.Tx) error {
		sql, err := sqlH.genCreateSQL()
		if err != nil {
			return err
		}

		rc, err := tx.ExecContext(ctx, sql)
		if err != nil {
			return err
		}
		if n, err := rc.RowsAffected(); err != nil || n != 1 {
			return fmt.Errorf("fail create app pool: %v", err)
		}
		return nil
	})
}
