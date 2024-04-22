package gooduser

import (
	"context"
	"fmt"

	"github.com/NpoolPlatform/miningpool-middleware/pkg/db"
	"github.com/NpoolPlatform/miningpool-middleware/pkg/db/ent"
	"github.com/NpoolPlatform/miningpool-middleware/pkg/mw/rootuser"

	"github.com/google/uuid"
)

func (h *Handler) CreateGoodUser(ctx context.Context) error {
	rootuserID := h.RootUserID.String()
	rootuserH, err := rootuser.NewHandler(ctx, rootuser.WithEntID(&rootuserID, true))
	if err != nil {
		return err
	}

	info, err := rootuserH.GetRootUser(ctx)
	if info == nil || err != nil {
		return fmt.Errorf("invalid rootuserid")
	}

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
			return fmt.Errorf("fail create gooduser: %v", err)
		}
		return nil
	})
}
