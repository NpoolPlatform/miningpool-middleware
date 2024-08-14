package orderuser

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/wlog"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	mpbasetypes "github.com/NpoolPlatform/message/npool/basetypes/miningpool/v1"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	coinpb "github.com/NpoolPlatform/message/npool/miningpool/mw/v1/coin"

	"github.com/NpoolPlatform/miningpool-middleware/pkg/mw/coin"
	"github.com/NpoolPlatform/miningpool-middleware/pkg/mw/gooduser"
	"github.com/NpoolPlatform/miningpool-middleware/pkg/mw/rootuser"
)

type baseInfo struct {
	OrderUserID    uint32
	MiningpoolType mpbasetypes.MiningpoolType
	CoinType       basetypes.CoinType
	AuthToken      string
	Recipient      string
	Distributor    string
}

type baseInfoHandle struct {
	*Handler
	baseInfo *baseInfo
}

func (h *baseInfoHandle) getBaseInfo(ctx context.Context) error {
	if h.CoinTypeID == nil {
		return wlog.Errorf("have invalid cointypeid")
	}
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

	coinH, err := coin.NewHandler(ctx, coin.WithConds(&coinpb.Conds{
		MiningpoolType: &basetypes.Uint32Val{
			Op:    cruder.EQ,
			Value: uint32(*orderUser.MiningpoolType.Enum()),
		},
		CoinTypeID: &basetypes.StringVal{
			Op:    cruder.EQ,
			Value: *h.CoinTypeID,
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
		return wlog.Errorf("cannot support cointypeid: %v", h.CoinTypeID)
	}

	h.baseInfo = &baseInfo{
		OrderUserID:    orderUser.ID,
		MiningpoolType: orderUser.MiningpoolType,
		CoinType:       coinInfos[0].CoinType,
		Distributor:    goodUser.Name,
		Recipient:      orderUser.Name,
		AuthToken:      rootUser.AuthTokenPlain,
	}
	return nil
}
