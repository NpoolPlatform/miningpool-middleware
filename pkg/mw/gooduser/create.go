package gooduser

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/wlog"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	v1 "github.com/NpoolPlatform/message/npool/basetypes/v1"
	coinpb "github.com/NpoolPlatform/message/npool/miningpool/mw/v1/coin"
	"github.com/NpoolPlatform/miningpool-middleware/pkg/db"
	"github.com/NpoolPlatform/miningpool-middleware/pkg/db/ent"
	"github.com/NpoolPlatform/miningpool-middleware/pkg/mw/coin"
	"github.com/NpoolPlatform/miningpool-middleware/pkg/mw/rootuser"
	"github.com/NpoolPlatform/miningpool-middleware/pkg/pools"

	"github.com/google/uuid"
)

func (h *Handler) CreateGoodUser(ctx context.Context) error {
	rootuserID := h.RootUserID.String()
	rootuserH, err := rootuser.NewHandler(ctx, rootuser.WithEntID(&rootuserID, true))
	if err != nil {
		return wlog.WrapError(err)
	}

	info, err := rootuserH.GetRootUser(ctx)
	if info == nil || err != nil {
		return wlog.Errorf("invalid rootuserid")
	}

	if h.EntID == nil {
		h.EntID = func() *uuid.UUID { uid := uuid.New(); return &uid }()
	}

	err = h.newGoodUserInPool(ctx)
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
			return wlog.Errorf("fail create gooduser: %v", err)
		}
		return nil
	})
}

func (h *Handler) newGoodUserInPool(ctx context.Context) error {
	rootuserID := h.RootUserID.String()
	rootuserH, err := rootuser.NewHandler(ctx, rootuser.WithEntID(&rootuserID, true))
	if err != nil {
		return wlog.WrapError(err)
	}
	rootuserToken, err := rootuserH.GetAuthToken(ctx)
	if err != nil {
		return wlog.WrapError(err)
	}
	if rootuserToken == nil {
		return wlog.Errorf("have no rootuser,entid: %v", rootuserID)
	}

	ruInfo, err := rootuserH.GetRootUser(ctx)
	if err != nil {
		return wlog.WrapError(err)
	}
	if ruInfo == nil {
		return wlog.Errorf("have no rootuser,entid: %v", rootuserID)
	}

	coinH, err := coin.NewHandler(ctx,
		coin.WithConds(&coinpb.Conds{
			CoinTypeIDs: &v1.StringSliceVal{
				Op:    cruder.EQ,
				Value: h.CoinTypeIDs,
			},
			PoolID: &v1.StringVal{
				Op:    cruder.EQ,
				Value: ruInfo.PoolID,
			},
		}),
		coin.WithLimit(1),
		coin.WithOffset(0))
	if err != nil {
		return wlog.WrapError(err)
	}

	// check if cointypes is suppored by the miningpool
	_, total, err := coinH.GetCoins(ctx)
	if err != nil {
		return wlog.WrapError(err)
	}

	if total == 0 {
		return wlog.Errorf("cannot support all cointype in cointypeids")
	}

	mgr, err := pools.NewPoolManager(ruInfo.MiningpoolType, nil, rootuserToken.AuthTokenPlain)
	if err != nil {
		return wlog.WrapError(err)
	}
	name, pageLink, err := mgr.AddMiningUser(ctx)
	if err != nil {
		return wlog.WrapError(err)
	}
	h.Name = &name
	h.ReadPageLink = &pageLink
	return nil
}
