package rootuser

import (
	"context"
	"fmt"

	v1 "github.com/NpoolPlatform/message/npool/basetypes/miningpool/v1"
	npool "github.com/NpoolPlatform/message/npool/miningpool/mw/v1/rootuser"

	"github.com/NpoolPlatform/miningpool-middleware/pkg/db"
	"github.com/NpoolPlatform/miningpool-middleware/pkg/db/ent"
	"github.com/NpoolPlatform/miningpool-middleware/pkg/mw/pool"
	"github.com/NpoolPlatform/miningpool-middleware/pkg/pools"
	"github.com/google/uuid"
)

func (h *Handler) UpdateRootUser(ctx context.Context) error {
	info, err := h.GetRootUser(ctx)
	if err != nil {
		return err
	}
	if info == nil {
		return fmt.Errorf("invalid id or ent_id")
	}

	poolID, err := uuid.Parse(info.PoolID)
	if err != nil {
		return err
	}

	err = h.checkUpdateAuthed(ctx, info)
	if err != nil {
		return err
	}

	sqlH := h.newSQLHandler()
	sqlH.BondPoolID = &poolID
	sqlH.BondEmail = &info.Email
	sqlH.BondName = &info.Name

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

func (h *Handler) checkUpdateAuthed(ctx context.Context, oldInfo *npool.RootUser) error {
	if h.AuthToken == nil && h.PoolID == nil {
		return nil
	}

	poolID := oldInfo.PoolID
	if h.PoolID != nil {
		poolID = h.PoolID.String()
	}

	poolH, err := pool.NewHandler(ctx, pool.WithEntID(&poolID, true))
	if err != nil {
		return err
	}
	info, err := poolH.GetPool(ctx)
	if err != nil {
		return err
	}
	if info == nil {
		return fmt.Errorf("invalid poolid")
	}

	defaultCoinType := v1.CoinType_BitCoin
	miningtype := info.MiningpoolType
	authtoken := h.AuthTokenPlain
	authed := false
	h.Authed = &authed

	if h.AuthTokenPlain == nil || h.PoolID == nil {
		info, err := h.GetAuthToken(ctx)
		if err != nil {
			return err
		}
		if h.AuthToken == nil {
			authtoken = &info.AuthTokenPlain
		}
	}

	mgr, err := pools.NewPoolManager(miningtype, defaultCoinType, *authtoken)
	if err != nil {
		return err
	}
	err = mgr.CheckAuth(ctx)
	if err != nil {
		return fmt.Errorf("have no permission to opreate pool,err: %v", err)
	}
	authed = true
	h.Authed = &authed
	return nil
}
