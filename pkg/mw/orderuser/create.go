package orderuser

import (
	"context"
	"fmt"

	"github.com/NpoolPlatform/miningpool-middleware/pkg/db"
	"github.com/NpoolPlatform/miningpool-middleware/pkg/db/ent"
	"github.com/NpoolPlatform/miningpool-middleware/pkg/mw/gooduser"

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
	gooduserID := h.GoodUserID.String()
	gooduserH, err := gooduser.NewHandler(ctx, gooduser.WithEntID(&gooduserID, true))
	if err != nil {
		return err
	}
	goodUser, err := gooduserH.GetGoodUser(ctx)
	if err != nil {
		return err
	}
	if goodUser == nil {
		return fmt.Errorf("have no gooduser,entid: %v", gooduserID)
	}

	// if h.MiningpoolType == nil {
	// 	h.MiningpoolType = &goodUser.MiningpoolType
	// }
	// if h.CoinType == nil {
	// 	h.CoinType = &goodUser.CoinType
	// }

	// if h.MiningpoolType.String() != goodUser.MiningpoolType.String() {
	// 	return fmt.Errorf("the miningpool type is different from type of gooduser")
	// }

	// if h.CoinType.String() != goodUser.CoinType.String() {
	// 	return fmt.Errorf("the coin type is different from type of gooduser")
	// }

	// if h.RootUserID == nil || h.RootUserID.String() != goodUser.RootUserID {
	// 	return fmt.Errorf("invalid rootuserid")
	// }

	// rootuserID := h.RootUserID.String()
	// rootuserH, err := rootuser.NewHandler(ctx, rootuser.WithEntID(&rootuserID, true))
	// if err != nil {
	// 	return err
	// }
	// rootUser, err := rootuserH.GetAuthToken(ctx)
	// if err != nil {
	// 	return err
	// }
	// if rootUser == nil {
	// 	return fmt.Errorf("have no rootuser,entid: %v", rootuserID)
	// }

	// mgr, err := pools.NewPoolManager(*h.MiningpoolType, *h.CoinType, rootUser.AuthTokenPlain)
	// if err != nil {
	// 	return err
	// }

	// name, pagelink, err := mgr.AddMiningUser(ctx)
	// if err != nil {
	// 	return err
	// }
	// h.Name = &name
	// h.ReadPageLink = &pagelink

	// paused, err := mgr.PausePayment(ctx, name)
	// if err != nil {
	// 	return err
	// }
	// autopay := !paused
	// h.AutoPay = &autopay

	return nil
}
