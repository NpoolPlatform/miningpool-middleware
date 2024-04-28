package fraction

import (
	"context"
	"testing"

	npool "github.com/NpoolPlatform/message/npool/miningpool/mw/v1/orderuser"
	"github.com/NpoolPlatform/miningpool-middleware/pkg/mw/orderuser"
	"github.com/NpoolPlatform/miningpool-middleware/pkg/pools/f2pool"
	"github.com/google/uuid"

	basetypes "github.com/NpoolPlatform/message/npool/basetypes/miningpool/v1"

	"github.com/stretchr/testify/assert"
)

var orderuserRet = &npool.OrderUser{
	EntID:          uuid.NewString(),
	GoodUserID:     gooduserRet.EntID,
	AppID:          uuid.NewString(),
	UserID:         uuid.NewString(),
	MiningpoolType: basetypes.MiningpoolType_F2Pool,
	CoinType:       basetypes.CoinType_BitCoin,
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
		orderuserRet.Proportion = info.Proportion
		orderuserRet.ID = info.ID
		orderuserRet.EntID = info.EntID
		orderuserRet.Name = info.Name
		orderuserRet.ReadPageLink = info.ReadPageLink
		assert.Equal(t, info, orderuserRet)
	}
}

func deleteOrderUser(t *testing.T) {
	handler, err := orderuser.NewHandler(
		context.Background(),
		orderuser.WithID(&orderuserRet.ID, true),
		orderuser.WithEntID(&orderuserRet.EntID, true),
	)
	assert.Nil(t, err)
	err = handler.DeleteOrderUser(context.Background())
	assert.Nil(t, err)
}
