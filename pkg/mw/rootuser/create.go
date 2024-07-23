package rootuser

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/wlog"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	"github.com/NpoolPlatform/miningpool-middleware/pkg/db"
	"github.com/NpoolPlatform/miningpool-middleware/pkg/db/ent"
	"github.com/NpoolPlatform/miningpool-middleware/pkg/mw/pool"
	"github.com/NpoolPlatform/miningpool-middleware/pkg/pools"

	"github.com/google/uuid"
)

func (h *Handler) CreateRootUser(ctx context.Context) error {
	if h.EntID == nil {
		h.EntID = func() *uuid.UUID { uid := uuid.New(); return &uid }()
	}
	err := h.checkCreateAuthed(ctx)
	if err != nil {
		return wlog.WrapError(err)
	}
	sqlH := h.newSQLHandler()
	return db.WithTx(ctx, func(ctx context.Context, tx *ent.Tx) error {
		sql, err := sqlH.genCreateSQL()
		if err != nil {
			return wlog.WrapError(err)
		}
		rc, err := tx.ExecContext(ctx, sql)
		if err != nil {
			return wlog.WrapError(err)
		}
		if n, err := rc.RowsAffected(); err != nil || n != 1 {
			return wlog.Errorf("fail create rootuser: %v", err)
		}
		return nil
	})
}

func (h *Handler) checkCreateAuthed(ctx context.Context) error {
	if h.PoolID == nil || h.Name == nil {
		return wlog.Errorf("have no poolid or name")
	}
	poolID := h.PoolID.String()
	poolH, err := pool.NewHandler(ctx, pool.WithEntID(&poolID, true))
	if err != nil {
		return wlog.WrapError(err)
	}
	info, err := poolH.GetPool(ctx)
	if err != nil {
		return wlog.WrapError(err)
	}
	if info == nil {
		return wlog.Errorf("invalid poolid")
	}
	defaultCoinType := basetypes.CoinType_CoinTypeBitCoin
	mgr, err := pools.NewPoolManager(info.MiningpoolType, defaultCoinType, *h.AuthTokenPlain)
	if err != nil {
		return wlog.WrapError(err)
	}

	err = mgr.CheckAuth(ctx)
	if err != nil {
		err = wlog.Errorf("have no permission to opreate pool, miningpool: %v, username: %v , err: %v", h.PoolID, *h.Name, err)
		return wlog.WrapError(err)
	}

	exist, err := mgr.ExistMiningUser(ctx, *h.Name)
	if err != nil {
		err = wlog.Errorf("failed to queary in %v,which called %v, err: %v", info.MiningpoolType, *h.Name, err)
		return wlog.WrapError(err)
	}

	if !exist {
		return wlog.Errorf("have no username in %v,which called %v", info.MiningpoolType, *h.Name)
	}

	authed := true
	h.Authed = &authed
	return nil
}
