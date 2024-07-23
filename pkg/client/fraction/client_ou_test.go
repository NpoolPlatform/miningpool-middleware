package fraction

import (
	"context"
	"testing"

	orderuserclient "github.com/NpoolPlatform/miningpool-middleware/pkg/client/orderuser"
	"github.com/google/uuid"

	cruder "github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	mpbasetypes "github.com/NpoolPlatform/message/npool/basetypes/miningpool/v1"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	apppool "github.com/NpoolPlatform/message/npool/miningpool/mw/v1/app/pool"
	npool "github.com/NpoolPlatform/message/npool/miningpool/mw/v1/orderuser"
	apppoolclient "github.com/NpoolPlatform/miningpool-middleware/pkg/client/app/pool"
	"github.com/stretchr/testify/assert"
)

var orderserRet = &npool.OrderUser{
	EntID:          uuid.NewString(),
	GoodUserID:     goodUserRet.EntID,
	AppID:          uuid.NewString(),
	UserID:         uuid.NewString(),
	MiningpoolType: mpbasetypes.MiningpoolType_F2Pool,
}

var orderuserReq = &npool.OrderUserReq{
	EntID:      &orderserRet.EntID,
	GoodUserID: &orderserRet.GoodUserID,
	AppID:      &orderserRet.AppID,
	UserID:     &orderserRet.UserID,
}

func createOrderUser(t *testing.T) {
	err := apppoolclient.CreatePool(context.Background(), &apppool.PoolReq{
		AppID:  &orderserRet.AppID,
		PoolID: &goodUserRet.PoolID,
	})
	assert.Nil(t, err)

	err = orderuserclient.CreateOrderUser(context.Background(), orderuserReq)
	assert.Nil(t, err)

	info, err := orderuserclient.GetOrderUser(context.Background(), *orderuserReq.EntID)
	if assert.Nil(t, err) {
		orderserRet.CreatedAt = info.CreatedAt
		orderserRet.Name = info.Name
		orderserRet.ReadPageLink = info.ReadPageLink
		orderserRet.MiningpoolTypeStr = info.MiningpoolTypeStr
		orderserRet.RootUserID = info.RootUserID
		orderserRet.UpdatedAt = info.UpdatedAt
		orderserRet.ID = info.ID
		orderserRet.EntID = info.EntID
		assert.Equal(t, orderserRet, info)
	}
}

func deleteOrderUser(t *testing.T) {
	exist, err := orderuserclient.ExistOrderUserConds(context.Background(), &npool.Conds{
		EntID: &basetypes.StringVal{
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
		EntID: &basetypes.StringVal{
			Op:    cruder.EQ,
			Value: orderserRet.EntID,
		},
	})
	if assert.Nil(t, err) {
		assert.Equal(t, false, exist)
	}
}
