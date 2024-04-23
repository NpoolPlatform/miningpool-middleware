package orderuser

import (
	"context"
	"fmt"

	v1 "github.com/NpoolPlatform/message/npool/basetypes/miningpool/v1"
	"github.com/NpoolPlatform/miningpool-middleware/pkg/db"
	"github.com/NpoolPlatform/miningpool-middleware/pkg/db/ent"
	"github.com/NpoolPlatform/miningpool-middleware/pkg/mw/gooduser"
	"github.com/NpoolPlatform/miningpool-middleware/pkg/mw/rootuser"
	"github.com/NpoolPlatform/miningpool-middleware/pkg/pools"
)

func (h *Handler) UpdateOrderUser(ctx context.Context) error {
	info, err := h.GetOrderUser(ctx)
	if err != nil {
		return err
	}

	if info == nil {
		return fmt.Errorf("invalid id or ent_id")
	}

	updateH := updateInPoolHandle{
		Handler: h,
	}

	err = updateH.handleUpdateReq(ctx)
	if err != nil {
		return err
	}

	sqlH := h.newSQLHandler()
	sqlH.BondMiningpoolType = &info.MiningpoolType
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
			return fmt.Errorf("failed to update orderuser: %v", err)
		}
		return nil
	})
}

type baseInfo struct {
	OrderUserID    uint32
	MiningpoolType v1.MiningpoolType
	CoinType       v1.CoinType
	AuthToken      string
	Recipient      string
	Distributor    string
}

type updateInPoolHandle struct {
	*Handler
	baseInfo *baseInfo
}

func (h *updateInPoolHandle) handleUpdateReq(ctx context.Context) error {
	err := h.getBaseInfo(ctx)
	if err != nil {
		return err
	}
	mgr, err := pools.NewPoolManager(h.baseInfo.MiningpoolType, h.baseInfo.CoinType, h.baseInfo.AuthToken)
	if err != nil {
		return err
	}
	if h.Proportion != nil {
		err = mgr.SetRevenueProportion(ctx, h.baseInfo.Distributor, h.baseInfo.Recipient, h.Proportion.String())
		if err != nil {
			return err
		}
	}

	if h.RevenueAddress != nil {
		err = mgr.SetRevenueAddress(ctx, h.baseInfo.Recipient, *h.RevenueAddress)
		if err != nil {
			return err
		}
		if h.AutoPay == nil {
			autopay := true
			h.AutoPay = &autopay
		}
	}

	if h.AutoPay != nil {
		autoPay := *h.AutoPay
		paused := true

		if autoPay {
			autoPay, err = mgr.ResumePayment(ctx, h.baseInfo.Recipient)
		} else {
			paused, err = mgr.PausePayment(ctx, h.baseInfo.Recipient)
		}
		if err != nil {
			return err
		}
		if !paused {
			autoPay = false
		}

		h.AutoPay = &autoPay
	}

	return nil
}

func (h *updateInPoolHandle) getBaseInfo(ctx context.Context) error {
	orderUser, err := h.GetOrderUser(ctx)
	if err != nil {
		return err
	}
	if orderUser == nil {
		err = fmt.Errorf("have no record of orderuser")
		return err
	}

	if h.MiningpoolType == nil {
		h.MiningpoolType = &orderUser.MiningpoolType
	}

	if h.Name == nil {
		h.Name = &orderUser.Name
	}

	gooduserH, err := gooduser.NewHandler(ctx, gooduser.WithEntID(&orderUser.GoodUserID, true))
	if err != nil {
		return err
	}
	goodUser, err := gooduserH.GetGoodUser(ctx)
	if err != nil {
		return err
	}
	if goodUser == nil {
		err = fmt.Errorf("have no record of gooduser with entid %v", orderUser.GoodUserID)
		return err
	}

	rootuserH, err := rootuser.NewHandler(ctx, rootuser.WithEntID(&orderUser.RootUserID, true))
	if err != nil {
		return err
	}
	rootUser, err := rootuserH.GetAuthToken(ctx)
	if err != nil {
		return err
	}
	if rootUser == nil {
		err = fmt.Errorf("have no record of rootuser with entid %v", orderUser.RootUserID)
		return err
	}
	h.baseInfo = &baseInfo{
		OrderUserID:    orderUser.ID,
		MiningpoolType: orderUser.MiningpoolType,
		CoinType:       orderUser.CoinType,
		Distributor:    goodUser.Name,
		Recipient:      orderUser.Name,
		AuthToken:      rootUser.AuthTokenPlain,
	}
	return nil
}
