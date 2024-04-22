package rootuser

import (
	"context"
	"fmt"

	v1 "github.com/NpoolPlatform/message/npool/basetypes/miningpool/v1"
	"github.com/NpoolPlatform/miningpool-middleware/pkg/db"
	"github.com/NpoolPlatform/miningpool-middleware/pkg/db/ent"
	"github.com/NpoolPlatform/miningpool-middleware/pkg/pools"
)

func (h *Handler) UpdateRootUser(ctx context.Context) error {
	info, err := h.GetRootUser(ctx)
	if err != nil {
		return err
	}

	if info == nil {
		return fmt.Errorf("invalid id or ent_id")
	}

	err = h.checkUpdateAuthed(ctx)
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
			return fmt.Errorf("failed to update rootuser: %v", err)
		}
		return nil
	})
}

func (h *Handler) checkUpdateAuthed(ctx context.Context) error {
	if h.AuthToken == nil && h.MiningpoolType == nil {
		return nil
	}

	defaultCoinType := v1.CoinType_BitCoin
	miningtype := h.MiningpoolType
	authtoken := h.AuthTokenPlain
	authed := false
	h.Authed = &authed

	if h.AuthTokenPlain == nil || h.MiningpoolType == nil {
		info, err := h.GetAuthToken(ctx)
		if err != nil {
			return err
		}
		if h.AuthToken == nil {
			authtoken = &info.AuthTokenPlain
		}
		if h.MiningpoolType == nil {
			miningtype = &info.MiningpoolType
		}
	}

	mgr, err := pools.NewPoolManager(*miningtype, defaultCoinType, *authtoken)
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
