package orderuser

import (
	"context"
	"fmt"

	"github.com/NpoolPlatform/miningpool-middleware/pkg/db"
	"github.com/NpoolPlatform/miningpool-middleware/pkg/db/ent"
	"github.com/NpoolPlatform/miningpool-middleware/pkg/mw/gooduser"
	"github.com/NpoolPlatform/miningpool-middleware/pkg/mw/rootuser"
	"github.com/NpoolPlatform/miningpool-middleware/pkg/pools"

	"github.com/google/uuid"
)

func (h *Handler) CreateOrderUser(ctx context.Context) error {
	if h.EntID == nil {
		h.EntID = func() *uuid.UUID { uid := uuid.New(); return &uid }()
	}

	err := h.newOrderUserInPool(ctx)
	if err != nil {
		return err
	}

	sqlH := h.newSQLHandler()
	sql, err := sqlH.genCreateSQL()
	if err != nil {
		return err
	}
	return db.WithTx(ctx, func(ctx context.Context, tx *ent.Tx) error {
		rc, err := tx.ExecContext(ctx, sql)
		if err != nil {
			return err
		}
		if n, err := rc.RowsAffected(); err != nil || n != 1 {
			return fmt.Errorf("fail create orderuser: %v", err)
		}
		return nil
	})
}

func (h *Handler) newOrderUserInPool(ctx context.Context) error {
	if h.GoodUserID == nil {
		return fmt.Errorf("invalid gooduserid")
	}
	gooduserID := h.GoodUserID.String()
	gooduserH, err := gooduser.NewHandler(ctx, gooduser.WithEntID(&gooduserID, true))
	if err != nil {
		return err
	}
	guInfo, err := gooduserH.GetGoodUser(ctx)
	if err != nil {
		return err
	}
	if guInfo == nil {
		return fmt.Errorf("have no gooduser,entid: %v", gooduserID)
	}

	rootuserH, err := rootuser.NewHandler(ctx, rootuser.WithEntID(&guInfo.RootUserID, true))
	if err != nil {
		return err
	}
	ruInfo, err := rootuserH.GetAuthToken(ctx)
	if err != nil {
		return err
	}
	if ruInfo == nil {
		return fmt.Errorf("have no rootuser,entid: %v", guInfo.RootUserID)
	}

	mgr, err := pools.NewPoolManager(guInfo.MiningpoolType, guInfo.CoinType, ruInfo.AuthTokenPlain)
	if err != nil {
		return err
	}

	name, pagelink, err := mgr.AddMiningUser(ctx)
	if err != nil {
		return err
	}
	h.Name = &name
	h.ReadPageLink = &pagelink

	paused, err := mgr.PausePayment(ctx, name)
	if err != nil {
		return err
	}
	autopay := !paused
	h.AutoPay = &autopay

	return nil
}
