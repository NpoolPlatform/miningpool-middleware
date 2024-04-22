package fraction

import (
	"context"
	"testing"

	orderuserclient "github.com/NpoolPlatform/miningpool-middleware/pkg/client/orderuser"
	"github.com/google/uuid"

	"github.com/stretchr/testify/assert"

	cruder "github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/miningpool/v1"
	v1 "github.com/NpoolPlatform/message/npool/basetypes/v1"
	npool "github.com/NpoolPlatform/message/npool/miningpool/mw/v1/orderuser"
)

var orderserRet = &npool.OrderUser{
	EntID:          uuid.NewString(),
	RootUserID:     rootUserRet.EntID,
	GoodUserID:     goodUserRet.EntID,
	AppID:          uuid.NewString(),
	UserID:         uuid.NewString(),
	MiningpoolType: basetypes.MiningpoolType_F2Pool,
	CoinType:       basetypes.CoinType_BitCoin,
}

var orderuserReq = &npool.OrderUserReq{
	EntID:          &orderserRet.EntID,
	RootUserID:     &orderserRet.RootUserID,
	GoodUserID:     &orderserRet.GoodUserID,
	AppID:          &orderserRet.AppID,
	UserID:         &orderserRet.UserID,
	MiningpoolType: &orderserRet.MiningpoolType,
	CoinType:       &orderserRet.CoinType,
}

func createOrderUser(t *testing.T) {
	err := orderuserclient.CreateOrderUser(context.Background(), orderuserReq)
	assert.Nil(t, err)

	info, err := orderuserclient.GetOrderUser(context.Background(), *orderuserReq.EntID)
	if assert.Nil(t, err) {
		orderserRet.CreatedAt = info.CreatedAt
		orderserRet.Name = info.Name
		orderserRet.ReadPageLink = info.ReadPageLink
		orderserRet.AutoPay = info.AutoPay
		orderserRet.MiningpoolTypeStr = info.MiningpoolTypeStr
		orderserRet.CoinTypeStr = info.CoinTypeStr
		orderserRet.UpdatedAt = info.UpdatedAt
		orderserRet.Proportion = info.Proportion
		orderserRet.ID = info.ID
		orderserRet.EntID = info.EntID
		assert.Equal(t, orderserRet, info)
	}
}

func deleteOrderUser(t *testing.T) {
	exist, err := orderuserclient.ExistOrderUserConds(context.Background(), &npool.Conds{
		EntID: &v1.StringVal{
			Op:    cruder.EQ,
			Value: orderserRet.EntID,
		},
	})
	if assert.Nil(t, err) {
		assert.Equal(t, true, exist)
	}

	err = orderuserclient.DeleteOrderUser(context.Background(), orderserRet.ID, orderserRet.EntID)
	assert.Nil(t, err)

	exist, err = orderuserclient.ExistOrderUserConds(context.Background(), &npool.Conds{
		EntID: &v1.StringVal{
			Op:    cruder.EQ,
			Value: orderserRet.EntID,
		},
	})
	if assert.Nil(t, err) {
		assert.Equal(t, false, exist)
	}
}
