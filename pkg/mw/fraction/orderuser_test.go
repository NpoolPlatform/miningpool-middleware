package fraction

import (
	"context"
	"testing"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	npool "github.com/NpoolPlatform/message/npool/miningpool/mw/v1/orderuser"
	apppool "github.com/NpoolPlatform/miningpool-middleware/pkg/mw/app/pool"
	"github.com/NpoolPlatform/miningpool-middleware/pkg/mw/orderuser"
	"github.com/NpoolPlatform/miningpool-middleware/pkg/pools/f2pool"
	"github.com/google/uuid"

	basetypes "github.com/NpoolPlatform/message/npool/basetypes/miningpool/v1"
	v1 "github.com/NpoolPlatform/message/npool/basetypes/v1"

	"github.com/stretchr/testify/assert"
)

var orderuserRet = &npool.OrderUser{
	EntID:          uuid.NewString(),
	GoodUserID:     gooduserRet.EntID,
	AppID:          uuid.NewString(),
	UserID:         uuid.NewString(),
	MiningpoolType: basetypes.MiningpoolType_F2Pool,
	CoinType:       v1.CoinType_CoinTypeBitCoin,
	RevenueAddress: "sssss",
	AutoPay:        false,
}

var orderuserReq = &npool.OrderUserReq{
	EntID:          &orderuserRet.EntID,
	GoodUserID:     &orderuserRet.GoodUserID,
	AppID:          &orderuserRet.AppID,
	UserID:         &orderuserRet.UserID,
	RevenueAddress: &orderuserRet.RevenueAddress,
	AutoPay:        &orderuserRet.AutoPay,
}

func createOrderUser(t *testing.T) {
	apppoolH, err := apppool.NewHandler(
		context.Background(),
		apppool.WithAppID(&ret.AppID, true),
		apppool.WithPoolID(&gooduserRet.PoolID, true),
	)
	assert.Nil(t, err)
	err = apppoolH.CreatePool(context.Background())
	assert.Nil(t, err)

	name, err := f2pool.RandomF2PoolUser(8)
	if !assert.Nil(t, err) {
		return
	}
	orderuserRet.Name = name
	orderuserReq.Name = &name

	handler, err := orderuser.NewHandler(
		context.Background(),
		orderuser.WithName(orderuserReq.Name, true),
		orderuser.WithEntID(orderuserReq.EntID, true),
		orderuser.WithGoodUserID(orderuserReq.GoodUserID, true),
		orderuser.WithAppID(orderuserReq.AppID, true),
		orderuser.WithUserID(orderuserReq.UserID, true),
		orderuser.WithRevenueAddress(orderuserReq.RevenueAddress, true),
		orderuser.WithAutoPay(orderuserReq.AutoPay, true),
	)
	assert.Nil(t, err)

	err = handler.CreateOrderUser(context.Background())
	assert.Nil(t, err)

	info, err := handler.GetOrderUser(context.Background())
	if assert.Nil(t, err) {
		orderuserRet.UpdatedAt = info.UpdatedAt
		orderuserRet.CreatedAt = info.CreatedAt
		orderuserRet.MiningpoolTypeStr = info.MiningpoolTypeStr
		orderuserRet.CoinTypeStr = info.CoinTypeStr
		orderuserRet.RevenueTypeStr = info.RevenueTypeStr
		orderuserRet.RevenueType = info.RevenueType
		orderuserRet.Proportion = info.Proportion
		orderuserRet.RootUserID = info.RootUserID
		orderuserRet.ID = info.ID
		orderuserRet.EntID = info.EntID
		orderuserRet.Name = info.Name
		orderuserRet.ReadPageLink = info.ReadPageLink
		assert.Equal(t, info, orderuserRet)
	}
}

func deleteOrderUser(t *testing.T) {
	conds := &npool.Conds{
		EntID: &v1.StringVal{
			Op:    cruder.EQ,
			Value: orderuserRet.EntID,
		},
	}
	handler, err := orderuser.NewHandler(
		context.Background(),
		orderuser.WithConds(conds),
		orderuser.WithID(&orderuserRet.ID, true),
		orderuser.WithOffset(0),
		orderuser.WithLimit(2),
	)
	assert.Nil(t, err)

	infos, total, err := handler.GetOrderUsers(context.Background())
	if assert.Nil(t, err) {
		assert.Equal(t, uint32(1), total)
		orderuserRet.MiningpoolTypeStr = infos[0].MiningpoolTypeStr
		orderuserRet.CoinTypeStr = infos[0].CoinTypeStr
		orderuserRet.UpdatedAt = infos[0].UpdatedAt
		assert.Equal(t, infos[0], orderuserRet)
	}

	orderuserRet.EntID = infos[0].EntID
	handler, err = orderuser.NewHandler(
		context.Background(),
		orderuser.WithEntID(&orderuserRet.EntID, true),
		orderuser.WithOffset(0),
		orderuser.WithLimit(2),
	)
	assert.Nil(t, err)

	exist, err := handler.ExistOrderUser(context.Background())
	if assert.Nil(t, err) {
		assert.Equal(t, true, exist)
	}

	handler, err = orderuser.NewHandler(
		context.Background(),
		orderuser.WithID(&orderuserRet.ID, true),
		orderuser.WithOffset(0),
		orderuser.WithLimit(2),
	)
	assert.Nil(t, err)
	err = handler.DeleteOrderUser(context.Background())
	assert.Nil(t, err)

	handler, err = orderuser.NewHandler(
		context.Background(),
		orderuser.WithConds(&npool.Conds{
			EntID: &v1.StringVal{
				Op:    cruder.EQ,
				Value: orderuserRet.EntID,
			},
		}),
	)
	assert.Nil(t, err)

	exist, err = handler.ExistOrderUserConds(context.Background())
	if assert.Nil(t, err) {
		assert.Equal(t, false, exist)
	}
}
