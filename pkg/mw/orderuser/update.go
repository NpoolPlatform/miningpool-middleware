package orderuser

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/wlog"
	mpbasetypes "github.com/NpoolPlatform/message/npool/basetypes/miningpool/v1"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	"github.com/NpoolPlatform/miningpool-middleware/pkg/db"
	"github.com/NpoolPlatform/miningpool-middleware/pkg/db/ent"
	"github.com/NpoolPlatform/miningpool-middleware/pkg/mw/gooduser"
	"github.com/NpoolPlatform/miningpool-middleware/pkg/mw/rootuser"
	"github.com/NpoolPlatform/miningpool-middleware/pkg/pools"
	"github.com/google/uuid"
)

func (h *Handler) UpdateOrderUser(ctx context.Context) error {
	info, err := h.GetOrderUser(ctx)
	if err != nil {
		return wlog.WrapError(err)
	}

	if info == nil {
		return wlog.Errorf("invalid id or ent_id")
	}

	updateH := updateInPoolHandle{
		Handler: h,
	}

	err = updateH.handleUpdateReq(ctx)
	if err != nil {
		return wlog.WrapError(err)
	}

	gooduserID, err := uuid.Parse(info.GoodUserID)
	if err != nil {
		return wlog.WrapError(err)
	}

	sqlH := h.newSQLHandler()
	sqlH.BondGoodUserID = &gooduserID
	sqlH.BondName = &info.Name

	return db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		sql, err := sqlH.genUpdateSQL()
		if err != nil {
			return wlog.WrapError(err)
		}

		rc, err := tx.ExecContext(ctx, sql)
		if err != nil {
			return wlog.WrapError(err)
		}

		if n, err := rc.RowsAffected(); err != nil || n != 1 {
			return wlog.Errorf("failed to update orderuser: %v", err)
		}
		return nil
	})
}

type baseInfo struct {
	OrderUserID    uint32
	MiningpoolType mpbasetypes.MiningpoolType
	CoinType       basetypes.CoinType
	AuthToken      string
	Recipient      string
	Distributor    string
	RevenueAddress string
}

type updateInPoolHandle struct {
	*Handler
	baseInfo *baseInfo
}

//nolint:gocognit
func (h *updateInPoolHandle) handleUpdateReq(ctx context.Context) error {
	err := h.getBaseInfo(ctx)
	if err != nil {
		return wlog.WrapError(err)
	}
	mgr, err := pools.NewPoolManager(h.baseInfo.MiningpoolType, h.baseInfo.CoinType, h.baseInfo.AuthToken)
	if err != nil {
		return wlog.WrapError(err)
	}
	if h.Proportion != nil {
		err = mgr.SetRevenueProportion(ctx, h.baseInfo.Distributor, h.baseInfo.Recipient, h.Proportion.String())
		if err != nil {
			return wlog.WrapError(err)
		}
	}

	if h.RevenueAddress != nil {
		err = mgr.SetRevenueAddress(ctx, h.baseInfo.Recipient, *h.RevenueAddress)
		if err != nil {
			return wlog.WrapError(err)
		}
		if h.AutoPay == nil {
			autopay := true
			h.AutoPay = &autopay
		}
	}

	if h.AutoPay != nil &&
		*h.AutoPay &&
		h.RevenueAddress == nil &&
		h.baseInfo.RevenueAddress == "" {
		return wlog.Errorf("cannot set autopay to true without an revenue address")
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
			return wlog.WrapError(err)
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
		return wlog.WrapError(err)
	}
	if orderUser == nil {
		err = wlog.Errorf("have no record of orderuser")
		return wlog.WrapError(err)
	}

	gooduserH, err := gooduser.NewHandler(ctx, gooduser.WithEntID(&orderUser.GoodUserID, true))
	if err != nil {
		return wlog.WrapError(err)
	}
	goodUser, err := gooduserH.GetGoodUser(ctx)
	if err != nil {
		return wlog.WrapError(err)
	}
	if goodUser == nil {
		err = wlog.Errorf("have no record of gooduser with entid %v", orderUser.GoodUserID)
		return wlog.WrapError(err)
	}

	rootuserH, err := rootuser.NewHandler(ctx, rootuser.WithEntID(&goodUser.RootUserID, true))
	if err != nil {
		return wlog.WrapError(err)
	}
	rootUser, err := rootuserH.GetAuthToken(ctx)
	if err != nil {
		return wlog.WrapError(err)
	}
	if rootUser == nil {
		err = wlog.Errorf("have no record of rootuser with entid %v", goodUser.RootUserID)
		return wlog.WrapError(err)
	}

	h.baseInfo = &baseInfo{
		OrderUserID:    orderUser.ID,
		MiningpoolType: orderUser.MiningpoolType,
		Distributor:    goodUser.Name,
		Recipient:      orderUser.Name,
		AuthToken:      rootUser.AuthTokenPlain,
	}
	return nil
}
