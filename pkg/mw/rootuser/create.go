package rootuser

import (
	"context"
	"fmt"

	v1 "github.com/NpoolPlatform/message/npool/basetypes/miningpool/v1"
	"github.com/NpoolPlatform/miningpool-middleware/pkg/db"
	"github.com/NpoolPlatform/miningpool-middleware/pkg/db/ent"
	"github.com/NpoolPlatform/miningpool-middleware/pkg/pools"

	"github.com/google/uuid"
)

func (h *Handler) CreateRootUser(ctx context.Context) error {
	if h.EntID == nil {
		h.EntID = func() *uuid.UUID { uid := uuid.New(); return &uid }()
	}
	err := h.checkCreateAuthed(ctx)
	if err != nil {
		return err
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
			return fmt.Errorf("fail create rootuser: %v", err)
		}
		return nil
	})
}

func (h *Handler) checkCreateAuthed(ctx context.Context) error {
	defaultCoinType := v1.CoinType_BitCoin
	mgr, err := pools.NewPoolManager(*h.MiningpoolType, defaultCoinType, *h.AuthTokenPlain)
	if err != nil {
		return err
	}

	err = mgr.CheckAuth(ctx)
	if err != nil {
		err = fmt.Errorf("have no permission to opreate pool, miningpool: %v, username: %v , err: %v", h.MiningpoolType, *h.Name, err)
		return err
	}
	authed := true
	h.Authed = &authed
	return nil
}
