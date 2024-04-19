package orderuser

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	npool "github.com/NpoolPlatform/message/npool/miningpool/mw/v1/orderuser"
	testinit "github.com/NpoolPlatform/miningpool-middleware/pkg/testinit"
	"github.com/google/uuid"

	basetypes "github.com/NpoolPlatform/message/npool/basetypes/miningpool/v1"
	v1 "github.com/NpoolPlatform/message/npool/basetypes/v1"

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
	Name:           "sssss",
	EntID:          uuid.NewString(),
	RootUserID:     uuid.NewString(),
	GoodUserID:     uuid.NewString(),
	AppID:          uuid.NewString(),
	UserID:         uuid.NewString(),
	MiningpoolType: basetypes.MiningpoolType_AntPool,
	CoinType:       basetypes.CoinType_BitCoin,
	Proportion:     "5.0",
	RevenueAddress: "sssss",
	ReadPageLink:   "sssss",
	AutoPay:        false,
}

var req = &npool.OrderUserReq{
	Name:           &ret.Name,
	EntID:          &ret.EntID,
	RootUserID:     &ret.RootUserID,
	GoodUserID:     &ret.GoodUserID,
	AppID:          &ret.AppID,
	UserID:         &ret.UserID,
	MiningpoolType: &ret.MiningpoolType,
	CoinType:       &ret.CoinType,
	Proportion:     &ret.Proportion,
	RevenueAddress: &ret.RevenueAddress,
	ReadPageLink:   &ret.ReadPageLink,
	AutoPay:        &ret.AutoPay,
}

func create(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithName(req.Name, true),
		WithEntID(req.EntID, true),
		WithRootUserID(req.RootUserID, true),
		WithGoodUserID(req.GoodUserID, true),
		WithAppID(req.AppID, true),
		WithUserID(req.UserID, true),
		WithMiningpoolType(req.MiningpoolType, true),
		WithCoinType(req.CoinType, true),
		WithProportion(req.Proportion, true),
		WithRevenueAddress(req.RevenueAddress, true),
		WithReadPageLink(req.ReadPageLink, true),
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
		ret.ID = info.ID
		ret.EntID = info.EntID
		assert.Equal(t, info, ret)
	}
}

func update(t *testing.T) {
	ret.MiningpoolType = basetypes.MiningpoolType_F2Pool
	ret.CoinType = basetypes.CoinType_BitCoin
	ret.Proportion = "66"

	handler, err := NewHandler(
		context.Background(),
		WithID(&ret.ID, true),
		WithMiningpoolType(&ret.MiningpoolType, false),
		WithCoinType(&ret.CoinType, false),
		WithProportion(&ret.Proportion, false),
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

	ret.MiningpoolType = basetypes.MiningpoolType_AntPool
	ret.AppID = uuid.NewString()
	ret.UserID = uuid.NewString()

	handler, err = NewHandler(
		context.Background(),
		WithID(&ret.ID, true),
		WithMiningpoolType(&ret.MiningpoolType, false),
		WithCoinType(&ret.CoinType, false),
		WithAppID(&ret.AppID, false),
		WithUserID(&ret.UserID, false),
		WithRevenueAddress(nil, false),
	)
	assert.Nil(t, err)

	err = handler.UpdateOrderUser(context.Background())
	assert.Nil(t, err)
}

func deleteRow(t *testing.T) {
	conds := &npool.Conds{
		EntID: &v1.StringVal{
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
	deletedItem, err := handler.DeleteOrderUser(context.Background())
	if assert.Nil(t, err) {
		assert.Equal(t, true, exist)
		assert.Equal(t, deletedItem, ret)
	}

	handler, err = NewHandler(
		context.Background(),
		WithConds(&npool.Conds{
			EntID: &v1.StringVal{
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

	t.Run("create", create)
	t.Run("update", update)
	t.Run("deleteRow", deleteRow)
}
