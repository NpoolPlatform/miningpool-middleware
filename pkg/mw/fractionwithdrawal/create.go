package fractionwithdrawal

import (
	"context"
	"time"

	"github.com/NpoolPlatform/go-service-framework/pkg/wlog"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	v1 "github.com/NpoolPlatform/message/npool/basetypes/miningpool/v1"
	basetypesv1 "github.com/NpoolPlatform/message/npool/basetypes/v1"
	orderusermwpb "github.com/NpoolPlatform/message/npool/miningpool/mw/v1/orderuser"

	fractionwithdrawalcrud "github.com/NpoolPlatform/miningpool-middleware/pkg/crud/fractionwithdrawal"
	"github.com/NpoolPlatform/miningpool-middleware/pkg/mw/coin"
	"github.com/NpoolPlatform/miningpool-middleware/pkg/mw/orderuser"
	"github.com/NpoolPlatform/miningpool-middleware/pkg/mw/rootuser"
	"github.com/NpoolPlatform/miningpool-middleware/pkg/pools"

	coinpb "github.com/NpoolPlatform/message/npool/miningpool/mw/v1/coin"

	"github.com/NpoolPlatform/miningpool-middleware/pkg/db"
	"github.com/NpoolPlatform/miningpool-middleware/pkg/db/ent"

	"github.com/google/uuid"
)

type createHandler struct {
	*Handler
	orderUser    *orderusermwpb.OrderUser
	rootUserAuth *rootuser.TokenInfo
}

func (h *createHandler) _getOrderUser(ctx context.Context, orderUserID string) error {
	orderuserH, err := orderuser.NewHandler(ctx, orderuser.WithEntID(&orderUserID, true))
	if err != nil {
		return wlog.WrapError(err)
	}
	orderUser, err := orderuserH.GetOrderUser(ctx)
	if err != nil {
		return wlog.WrapError(err)
	}
	if orderUser == nil {
		return wlog.Errorf("have no orderuser,entid: %v", orderUserID)
	}

	if h.AppID == nil || h.AppID.String() != orderUser.AppID {
		return wlog.Errorf("invalid appid")
	}

	if h.UserID == nil || h.UserID.String() != orderUser.UserID {
		return wlog.Errorf("invalid userid")
	}

	h.orderUser = orderUser
	return nil
}

func (h *createHandler) _getRootUserAuth(ctx context.Context, rootUserID string) error {
	rootuserH, err := rootuser.NewHandler(ctx, rootuser.WithEntID(&rootUserID, true))
	if err != nil {
		return wlog.WrapError(err)
	}
	rootUser, err := rootuserH.GetAuthToken(ctx)
	if err != nil {
		return wlog.WrapError(err)
	}
	if rootUser == nil {
		return wlog.Errorf("have no rootuser,entid: %v", h.orderUser.RootUserID)
	}

	h.rootUserAuth = rootUser
	return nil
}

func (h *createHandler) fractionwithdrawalInPool(ctx context.Context) error {
	if h.OrderUserID == nil {
		return wlog.Errorf("invalid orderuserid")
	}

	if h.CoinTypeID == nil {
		return wlog.Errorf("invalid cointypeid")
	}

	if err := h._getOrderUser(ctx, h.OrderUserID.String()); err != nil {
		return wlog.WrapError(err)
	}

	if err := h._getRootUserAuth(ctx, h.orderUser.RootUserID); err != nil {
		return wlog.WrapError(err)
	}

	coinH, err := coin.NewHandler(ctx, coin.WithConds(&coinpb.Conds{
		MiningpoolType: &basetypesv1.Uint32Val{
			Op:    cruder.EQ,
			Value: uint32(*h.orderUser.MiningpoolType.Enum()),
		},
		CoinTypeIDs: &basetypesv1.StringSliceVal{
			Op:    cruder.EQ,
			Value: []string{h.CoinTypeID.String()},
		},
	}), coin.WithOffset(0), coin.WithLimit(1))
	if err != nil {
		return wlog.WrapError(err)
	}

	coinInfos, _, err := coinH.GetCoins(ctx)
	if err != nil {
		return wlog.WrapError(err)
	}

	if len(coinInfos) == 0 {
		return wlog.Errorf("cannot support cointypeid: %v", h.CoinTypeID.String())
	}

	mgr, err := pools.NewPoolManager(h.orderUser.MiningpoolType, &coinInfos[0].CoinType, h.rootUserAuth.AuthTokenPlain)
	if err != nil {
		return wlog.WrapError(err)
	}
	withdrawTime := uint32(time.Now().Unix())
	h.WithdrawAt = &withdrawTime
	_PromisePayAt, err := mgr.WithdrawFractionWithdrawal(ctx, h.orderUser.Name)
	PromisePayAt := uint32(_PromisePayAt)
	if err != nil {
		errMsg := err.Error()
		h.Msg = &errMsg
		h.FractionWithdrawState = v1.FractionWithdrawState_FractionWithdrawStateFailed.Enum()
	} else if PromisePayAt == 0 {
		h.PromisePayAt = &PromisePayAt
		h.FractionWithdrawState = v1.FractionWithdrawState_FractionWithdrawStateFailed.Enum()
	} else {
		h.PromisePayAt = &PromisePayAt
		h.FractionWithdrawState = v1.FractionWithdrawState_FractionWithdrawStateSuccess.Enum()
	}
	return nil
}

func (h *Handler) CreateFractionWithdrawal(ctx context.Context) error {
	handle := createHandler{Handler: h}

	if handle.EntID == nil {
		handle.EntID = func() *uuid.UUID { uid := uuid.New(); return &uid }()
	}

	err := handle.fractionwithdrawalInPool(ctx)
	if err != nil {
		return wlog.WrapError(err)
	}

	return db.WithClient(ctx, func(ctx context.Context, cli *ent.Client) error {
		info, err := fractionwithdrawalcrud.CreateSet(
			cli.FractionWithdrawal.Create(),
			&fractionwithdrawalcrud.Req{
				EntID:                 h.EntID,
				AppID:                 h.AppID,
				UserID:                h.UserID,
				OrderUserID:           h.OrderUserID,
				CoinTypeID:            h.CoinTypeID,
				FractionWithdrawState: h.FractionWithdrawState,
				WithdrawAt:            h.WithdrawAt,
				PromisePayAt:          h.PromisePayAt,
				Msg:                   h.Msg,
			},
		).Save(ctx)
		if err != nil {
			return wlog.WrapError(err)
		}
		h.ID = &info.ID
		return nil
	})
}
