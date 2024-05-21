package rootuser

import (
	"context"
	"fmt"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
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
	if h.AuthTokenPlain == nil && h.PoolID == nil {
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

	defaultCoinType := basetypes.CoinType_CoinTypeBitCoin
	miningtype := info.MiningpoolType
	authed := false
	h.Authed = &authed

	authInfo, err := h.GetAuthToken(ctx)
	if err != nil {
		return err
	}
	if h.AuthTokenPlain == nil {
		h.AuthTokenPlain = &authInfo.AuthTokenPlain
	}
	logger.Sugar().Error("ssssss1", *h.AuthToken)
	logger.Sugar().Error("ssssss1", *h.AuthTokenPlain)
	logger.Sugar().Error("ssssss1", *h.AuthTokenSalt)
	logger.Sugar().Error("ssssss1", authInfo.AuthToken)
	logger.Sugar().Error("ssssss1", authInfo.AuthTokenPlain)
	logger.Sugar().Error("ssssss1", authInfo.AuthTokenSalt)
	if *h.AuthTokenPlain == authInfo.AuthToken && (h.PoolID == nil || h.PoolID.String() == poolID) {
		return nil
	}

	mgr, err := pools.NewPoolManager(miningtype, defaultCoinType, *h.AuthTokenPlain)
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
