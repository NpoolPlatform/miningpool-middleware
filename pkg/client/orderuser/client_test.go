package orderuser

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"

	"github.com/NpoolPlatform/miningpool-middleware/pkg/pools"
	"github.com/NpoolPlatform/miningpool-middleware/pkg/testinit"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"

	apppoolclient "github.com/NpoolPlatform/miningpool-middleware/pkg/client/app/pool"
	"github.com/stretchr/testify/assert"

	cruder "github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	mpbasetypes "github.com/NpoolPlatform/message/npool/basetypes/miningpool/v1"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	"github.com/NpoolPlatform/message/npool/miningpool/mw/v1/app/pool"
	npool "github.com/NpoolPlatform/message/npool/miningpool/mw/v1/orderuser"
)

func init() {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}
	if err := testinit.Init(); err != nil {
		fmt.Printf("cannot init test stub: %v\n", err)
	}
}

var ret = &npool.OrderUser{
	EntID:          uuid.NewString(),
	GoodUserID:     goodUserRet.EntID,
	AppID:          uuid.NewString(),
	UserID:         uuid.NewString(),
	MiningpoolType: mpbasetypes.MiningpoolType_F2Pool,
	CoinType:       basetypes.CoinType_CoinTypeBitCoin,
}

var req = &npool.OrderUserReq{
	EntID:      &ret.EntID,
	GoodUserID: &ret.GoodUserID,
	AppID:      &ret.AppID,
	UserID:     &ret.UserID,
}

func createOrderUser(t *testing.T) {
	err := apppoolclient.CreatePool(context.Background(), &pool.PoolReq{
		AppID:  &ret.AppID,
		PoolID: &goodUserRet.PoolID,
	})
	assert.Nil(t, err)

	err = CreateOrderUser(context.Background(), req)
	assert.Nil(t, err)

	info, err := GetOrderUser(context.Background(), *req.EntID)
	if assert.Nil(t, err) {
		ret.CreatedAt = info.CreatedAt
		ret.Name = info.Name
		ret.ReadPageLink = info.ReadPageLink
		ret.AutoPay = info.AutoPay
		ret.MiningpoolTypeStr = info.MiningpoolTypeStr
		ret.CoinTypeStr = info.CoinTypeStr
		ret.RevenueTypeStr = info.RevenueTypeStr
		ret.RevenueType = info.RevenueType
		ret.RootUserID = info.RootUserID
		ret.UpdatedAt = info.UpdatedAt
		ret.ID = info.ID
		ret.Proportion = info.Proportion
		ret.EntID = info.EntID
		assert.Equal(t, ret, info)
	}
}

func updateOrderUser(t *testing.T) {
	req.ID = &ret.ID
	dec, err := decimal.NewFromString("50.1")
	assert.Nil(t, err)

	ret.Proportion = dec.String()
	req.Proportion = &ret.Proportion
	err = UpdateOrderUser(context.Background(), req)
	assert.Nil(t, err)

	info, err := GetOrderUser(context.Background(), *req.EntID)
	if assert.Nil(t, err) {
		ret.UpdatedAt = info.UpdatedAt
		assert.Equal(t, info, ret)
	}

	dec, err = decimal.NewFromString("99")
	assert.Nil(t, err)

	ret.Proportion = dec.String()
	req.Proportion = &ret.Proportion
	err = UpdateOrderUser(context.Background(), &npool.OrderUserReq{
		ID:         req.ID,
		EntID:      &ret.EntID,
		Proportion: &ret.Proportion,
	})
	assert.Nil(t, err)
	info, err = GetOrderUser(context.Background(), *req.EntID)
	if assert.Nil(t, err) {
		ret.UpdatedAt = info.UpdatedAt
		assert.Equal(t, info, ret)
	}

	ret.RevenueAddress = "1PWMfNSb3auXwZ1qhu96WRJL7BCgG4mGB4"
	req.RevenueAddress = &ret.RevenueAddress
	err = UpdateOrderUser(context.Background(), &npool.OrderUserReq{
		ID:             req.ID,
		EntID:          &ret.EntID,
		RevenueAddress: &ret.RevenueAddress,
	})
	assert.Nil(t, err)

	info, err = GetOrderUser(context.Background(), *req.EntID)
	if assert.Nil(t, err) {
		ret.UpdatedAt = info.UpdatedAt
		ret.AutoPay = info.AutoPay
		assert.Equal(t, info, ret)
	}

	ret.AutoPay = true
	req.AutoPay = &ret.AutoPay
	err = UpdateOrderUser(context.Background(), &npool.OrderUserReq{
		ID:      req.ID,
		EntID:   &ret.EntID,
		AutoPay: &ret.AutoPay,
	})
	assert.Nil(t, err)

	info, err = GetOrderUser(context.Background(), *req.EntID)
	if assert.Nil(t, err) {
		ret.UpdatedAt = info.UpdatedAt
		assert.Equal(t, info, ret)
	}

	ret.AutoPay = false
	req.AutoPay = &ret.AutoPay
	err = UpdateOrderUser(context.Background(), &npool.OrderUserReq{
		ID:      req.ID,
		EntID:   &ret.EntID,
		AutoPay: &ret.AutoPay,
	})
	assert.Nil(t, err)

	info, err = GetOrderUser(context.Background(), *req.EntID)
	if assert.Nil(t, err) {
		ret.UpdatedAt = info.UpdatedAt
		assert.Equal(t, info, ret)
	}
}

func getOrderUser(t *testing.T) {
	info, err := GetOrderUser(context.Background(), ret.EntID)
	if assert.Nil(t, err) {
		assert.Equal(t, info, ret)
	}
}

func getOrderUsers(t *testing.T) {
	infos, total, err := GetOrderUsers(context.Background(), &npool.Conds{
		EntID: &basetypes.StringVal{Op: cruder.EQ, Value: ret.EntID},
	}, 0, 1)
	if assert.Nil(t, err) {
		assert.Equal(t, len(infos), 1)
		assert.Equal(t, total, uint32(1))
		assert.Equal(t, infos[0], ret)
	}
}

func deleteOrderUser(t *testing.T) {
	exist, err := ExistOrderUserConds(context.Background(), &npool.Conds{
		EntID: &basetypes.StringVal{
			Op:    cruder.EQ,
			Value: ret.EntID,
		},
	})
	if assert.Nil(t, err) {
		assert.Equal(t, true, exist)
	}

	err = DeleteOrderUser(context.Background(), ret.ID, ret.EntID)
	assert.Nil(t, err)

	exist, err = ExistOrderUserConds(context.Background(), &npool.Conds{
		EntID: &basetypes.StringVal{
			Op:    cruder.EQ,
			Value: ret.EntID,
		},
	})
	if assert.Nil(t, err) {
		assert.Equal(t, false, exist)
	}
}

func TestClient(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	pools.InitTestInfo(context.Background())
	t.Run("createRootUser", createRootUser)
	t.Run("createGoodUser", createGoodUser)
	t.Run("createOrderUser", createOrderUser)
	t.Run("updateOrderUser", updateOrderUser)
	t.Run("getOrderUser", getOrderUser)
	t.Run("getOrderUsers", getOrderUsers)
	t.Run("deleteOrderUser", deleteOrderUser)
	t.Run("deleteGoodUser", deleteGoodUser)
	t.Run("deleteRootUser", deleteRootUser)
	pools.CleanTestInfo(context.Background())
}
