package gooduser

import (
	"context"
	"fmt"

	"github.com/NpoolPlatform/miningpool-middleware/pkg/db"
	"github.com/NpoolPlatform/miningpool-middleware/pkg/db/ent"

	"github.com/google/uuid"
)

func (h *Handler) CreateGoodUser(ctx context.Context) error {
	if h.EntID == nil {
		h.EntID = func() *uuid.UUID { uid := uuid.New(); return &uid }()
	}
	return db.WithTx(ctx, func(ctx context.Context, tx *ent.Tx) error {
		sql, err := h.genCreateSQL()
		if err != nil {
			return err
		}
		rc, err := tx.ExecContext(ctx, sql)
		if err != nil {
			return err
		}
		if n, err := rc.RowsAffected(); err != nil || n != 1 {
			return fmt.Errorf("fail create pool: %v", err)
		}
		return nil
	})
}
