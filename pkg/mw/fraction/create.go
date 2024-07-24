package fraction

import (
	"context"
	"time"

	"github.com/NpoolPlatform/go-service-framework/pkg/wlog"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	v1 "github.com/NpoolPlatform/message/npool/basetypes/miningpool/v1"
	basetypesv1 "github.com/NpoolPlatform/message/npool/basetypes/v1"

	fractioncrud "github.com/NpoolPlatform/miningpool-middleware/pkg/crud/fraction"
	"github.com/NpoolPlatform/miningpool-middleware/pkg/mw/coin"
	"github.com/NpoolPlatform/miningpool-middleware/pkg/mw/orderuser"
	"github.com/NpoolPlatform/miningpool-middleware/pkg/mw/rootuser"
	"github.com/NpoolPlatform/miningpool-middleware/pkg/pools"

	coinpb "github.com/NpoolPlatform/message/npool/miningpool/mw/v1/coin"

	"github.com/NpoolPlatform/miningpool-middleware/pkg/db"
	"github.com/NpoolPlatform/miningpool-middleware/pkg/db/ent"

	"github.com/google/uuid"
)

func (h *Handler) fractionInPool(ctx context.Context) error {
	orderUserID := h.OrderUserID.String()
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

	rootuserH, err := rootuser.NewHandler(ctx, rootuser.WithEntID(&orderUser.RootUserID, true))
	if err != nil {
		return wlog.WrapError(err)
	}
	rootUser, err := rootuserH.GetAuthToken(ctx)
	if err != nil {
		return wlog.WrapError(err)
	}
	if rootUser == nil {
		return wlog.Errorf("have no rootuser,entid: %v", orderUser.RootUserID)
	}

	coinH, err := coin.NewHandler(ctx, coin.WithConds(&coinpb.Conds{
		MiningpoolType: &basetypesv1.Uint32Val{
			Op:    cruder.EQ,
			Value: uint32(*orderUser.MiningpoolType.Enum()),
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

	mgr, err := pools.NewPoolManager(orderUser.MiningpoolType, &coinInfos[0].CoinType, rootUser.AuthTokenPlain)
	if err != nil {
		return wlog.WrapError(err)
	}
	withdrawTime := uint32(time.Now().Unix())
	h.WithdrawAt = &withdrawTime
	_PromisePayAt, err := mgr.WithdrawFraction(ctx, orderUser.Name)
	PromisePayAt := uint32(_PromisePayAt)
	if err != nil {
		errMsg := err.Error()
		h.Msg = &errMsg
		h.WithdrawState = v1.WithdrawState_WithdrawStateFailed.Enum()
	} else if PromisePayAt == 0 {
		h.PromisePayAt = &PromisePayAt
		h.WithdrawState = v1.WithdrawState_WithdrawStateFailed.Enum()
	} else {
		h.PromisePayAt = &PromisePayAt
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
		return wlog.WrapError(err)
	}

	return db.WithClient(ctx, func(ctx context.Context, cli *ent.Client) error {
		info, err := fractioncrud.CreateSet(
			cli.Fraction.Create(),
			&fractioncrud.Req{
				EntID:         h.EntID,
				AppID:         h.AppID,
				UserID:        h.UserID,
				OrderUserID:   h.OrderUserID,
				CoinTypeID:    h.CoinTypeID,
				WithdrawState: h.WithdrawState,
				WithdrawAt:    h.WithdrawAt,
				PromisePayAt:  h.PromisePayAt,
				Msg:           h.Msg,
			},
		).Save(ctx)
		if err != nil {
			return wlog.WrapError(err)
		}
		h.ID = &info.ID
		return nil
	})
}
