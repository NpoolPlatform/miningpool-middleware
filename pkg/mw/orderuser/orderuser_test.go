package orderuser

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	npool "github.com/NpoolPlatform/message/npool/miningpool/mw/v1/orderuser"
	apppool "github.com/NpoolPlatform/miningpool-middleware/pkg/mw/app/pool"
	"github.com/NpoolPlatform/miningpool-middleware/pkg/pools"
	"github.com/NpoolPlatform/miningpool-middleware/pkg/pools/f2pool"
	testinit "github.com/NpoolPlatform/miningpool-middleware/pkg/testinit"
	"github.com/google/uuid"

	mpbasetypes "github.com/NpoolPlatform/message/npool/basetypes/miningpool/v1"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"

	"github.com/stretchr/testify/assert"
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
	GoodUserID:     gooduserRet.EntID,
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

func create(t *testing.T) {
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
	ret.Name = name
	req.Name = &name

	handler, err := NewHandler(
		context.Background(),
		WithName(req.Name, true),
		WithEntID(req.EntID, true),
		WithGoodUserID(req.GoodUserID, true),
		WithAppID(req.AppID, true),
		WithUserID(req.UserID, true),
		WithRevenueAddress(req.RevenueAddress, true),
		WithAutoPay(req.AutoPay, true),
	)
	assert.Nil(t, err)

	err = handler.CreateOrderUser(context.Background())
	assert.Nil(t, err)

	info, err := handler.GetOrderUser(context.Background())
	if assert.Nil(t, err) {
		ret.UpdatedAt = info.UpdatedAt
		ret.CreatedAt = info.CreatedAt
		ret.MiningpoolTypeStr = info.MiningpoolTypeStr
		ret.CoinTypeStr = info.CoinTypeStr
		ret.RootUserID = info.RootUserID
		ret.ID = info.ID
		ret.EntID = info.EntID
		ret.Name = info.Name
		ret.ReadPageLink = info.ReadPageLink
		assert.Equal(t, info, ret)
	}
}

func update(t *testing.T) {
	ret.MiningpoolType = mpbasetypes.MiningpoolType_F2Pool
	ret.CoinType = basetypes.CoinType_CoinTypeBitCoin

	handler, err := NewHandler(
		context.Background(),
		WithID(&ret.ID, true),
		WithRevenueAddress(nil, false),
	)
	assert.Nil(t, err)

	err = handler.UpdateOrderUser(context.Background())
	assert.Nil(t, err)

	info, err := handler.GetOrderUser(context.Background())
	if assert.Nil(t, err) {
		ret.MiningpoolTypeStr = info.MiningpoolTypeStr
		ret.CoinTypeStr = info.CoinTypeStr
		ret.UpdatedAt = info.UpdatedAt
		assert.Equal(t, info, ret)
	}
}

func deleteRow(t *testing.T) {
	conds := &npool.Conds{
		EntID: &basetypes.StringVal{
			Op:    cruder.EQ,
			Value: ret.EntID,
		},
	}
	handler, err := NewHandler(
		context.Background(),
		WithConds(conds),
		WithID(&ret.ID, true),
		WithOffset(0),
		WithLimit(2),
	)
	assert.Nil(t, err)

	infos, total, err := handler.GetOrderUsers(context.Background())
	if assert.Nil(t, err) {
		assert.Equal(t, uint32(1), total)
		ret.MiningpoolTypeStr = infos[0].MiningpoolTypeStr
		ret.CoinTypeStr = infos[0].CoinTypeStr
		ret.UpdatedAt = infos[0].UpdatedAt
		assert.Equal(t, infos[0], ret)
	}

	ret.EntID = infos[0].EntID
	handler, err = NewHandler(
		context.Background(),
		WithEntID(&ret.EntID, true),
		WithOffset(0),
		WithLimit(2),
	)
	assert.Nil(t, err)

	exist, err := handler.ExistOrderUser(context.Background())
	if assert.Nil(t, err) {
		assert.Equal(t, true, exist)
	}

	handler, err = NewHandler(
		context.Background(),
		WithID(&ret.ID, true),
		WithOffset(0),
		WithLimit(2),
	)
	assert.Nil(t, err)
	err = handler.DeleteOrderUser(context.Background())
	assert.Nil(t, err)

	handler, err = NewHandler(
		context.Background(),
		WithConds(&npool.Conds{
			EntID: &basetypes.StringVal{
				Op:    cruder.EQ,
				Value: ret.EntID,
			},
		}),
	)
	assert.Nil(t, err)

	exist, err = handler.ExistOrderUserConds(context.Background())
	if assert.Nil(t, err) {
		assert.Equal(t, false, exist)
	}
}

func TestOrderUser(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	pools.InitTestInfo(context.Background())
	t.Run("create", createRootUser)
	t.Run("create", createGoodUser)
	t.Run("create", create)
	t.Run("update", update)
	t.Run("deleteRow", deleteRow)
	t.Run("deleteRow", deleteGoodUser)
	t.Run("deleteRow", deleteRootUser)
	pools.CleanTestInfo(context.Background())
}
