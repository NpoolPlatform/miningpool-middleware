package fraction

import (
	"context"
	"fmt"
	"time"

	v1 "github.com/NpoolPlatform/message/npool/basetypes/miningpool/v1"
	fractioncrud "github.com/NpoolPlatform/miningpool-middleware/pkg/crud/fraction"
	"github.com/NpoolPlatform/miningpool-middleware/pkg/mw/orderuser"
	"github.com/NpoolPlatform/miningpool-middleware/pkg/mw/rootuser"
	"github.com/NpoolPlatform/miningpool-middleware/pkg/pools"

	"github.com/NpoolPlatform/miningpool-middleware/pkg/db"
	"github.com/NpoolPlatform/miningpool-middleware/pkg/db/ent"

	"github.com/google/uuid"
)

func (h *Handler) fractionInPool(ctx context.Context) error {
	orderUserID := h.OrderUserID.String()
	orderuserH, err := orderuser.NewHandler(ctx, orderuser.WithEntID(&orderUserID, true))
	if err != nil {
		return err
	}
	orderUser, err := orderuserH.GetOrderUser(ctx)
	if err != nil {
		return err
	}
	if orderUser == nil {
		return fmt.Errorf("have no orderuser,entid: %v", orderUserID)
	}

	if h.AppID == nil || h.AppID.String() != orderUser.AppID {
		return fmt.Errorf("invalid appid")
	}

	if h.UserID == nil || h.UserID.String() != orderUser.UserID {
		return fmt.Errorf("invalid userid")
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
		return fmt.Errorf("have no rootuser,entid: %v", orderUser.RootUserID)
	}

	mgr, err := pools.NewPoolManager(orderUser.MiningpoolType, orderUser.CoinType, rootUser.AuthTokenPlain)
	if err != nil {
		return err
	}
	withdrawTime := uint32(time.Now().Unix())
	h.WithdrawTime = &withdrawTime
	_paytime, err := mgr.WithdrawPraction(ctx, orderUser.Name)
	paytime := uint32(_paytime)
	if err != nil {
		errMsg := err.Error()
		h.Msg = &errMsg
		h.WithdrawState = v1.WithdrawState_WithdrawStateFailed.Enum()
	} else if paytime == 0 {
		h.PayTime = &paytime
		h.WithdrawState = v1.WithdrawState_WithdrawStateFailed.Enum()
	} else {
		h.PayTime = &paytime
		h.WithdrawState = v1.WithdrawState_WithdrawStateSuccess.Enum()
	}
	return nil
}

func (h *Handler) CreateFraction(ctx context.Context) error {
	if h.EntID == nil {
		h.EntID = func() *uuid.UUID { uid := uuid.New(); return &uid }()
	}

	err := h.fractionInPool(ctx)
	if err != nil {
		return err
	}

	return db.WithClient(ctx, func(ctx context.Context, cli *ent.Client) error {
		info, err := fractioncrud.CreateSet(
			cli.Fraction.Create(),
			&fractioncrud.Req{
				EntID:         h.EntID,
				AppID:         h.AppID,
				UserID:        h.UserID,
				OrderUserID:   h.OrderUserID,
				WithdrawState: h.WithdrawState,
				WithdrawTime:  h.WithdrawTime,
				PayTime:       h.PayTime,
				Msg:           h.Msg,
			},
		).Save(ctx)
		if err != nil {
			return err
		}
		h.ID = &info.ID
		return nil
	})
}
