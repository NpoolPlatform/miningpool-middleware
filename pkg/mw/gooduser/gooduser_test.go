package gooduser

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	npool "github.com/NpoolPlatform/message/npool/miningpool/mw/v1/gooduser"
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

var ret = &npool.GoodUser{
	Name:           "sssss",
	EntID:          uuid.NewString(),
	RootUserID:     uuid.NewString(),
	GoodID:         uuid.NewString(),
	MiningpoolType: basetypes.MiningpoolType_AntPool,
	CoinType:       basetypes.CoinType_BitCoin,
	HashRate:       5.0,
	ReadPageLink:   "sssss",
	RevenueType:    basetypes.RevenueType_FPPS,
}

var req = &npool.GoodUserReq{
	Name:           &ret.Name,
	EntID:          &ret.EntID,
	RootUserID:     &ret.RootUserID,
	GoodID:         &ret.GoodID,
	MiningpoolType: &ret.MiningpoolType,
	CoinType:       &ret.CoinType,
	HashRate:       &ret.HashRate,
	ReadPageLink:   &ret.ReadPageLink,
	RevenueType:    &ret.RevenueType,
}

func create(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithName(req.Name, true),
		WithEntID(req.EntID, true),
		WithRootUserID(req.RootUserID, true),
		WithGoodID(req.GoodID, true),
		WithMiningpoolType(req.MiningpoolType, true),
		WithCoinType(req.CoinType, true),
		WithHashRate(req.HashRate, true),
		WithReadPageLink(req.ReadPageLink, true),
		WithRevenueType(req.RevenueType, true),
	)
	assert.Nil(t, err)

	info, err := handler.CreateGoodUser(context.Background())
	if assert.Nil(t, err) {
		ret.UpdatedAt = info.UpdatedAt
		ret.CreatedAt = info.CreatedAt
		ret.MiningpoolTypeStr = info.MiningpoolTypeStr
		ret.CoinTypeStr = info.CoinTypeStr
		ret.RevenueTypeStr = info.RevenueTypeStr
		ret.ID = info.ID
		ret.EntID = info.EntID
		assert.Equal(t, info, ret)
	}
}

func update(t *testing.T) {
	ret.MiningpoolType = basetypes.MiningpoolType_F2Pool
	ret.CoinType = basetypes.CoinType_BitCoin
	ret.RevenueType = basetypes.RevenueType_PPLNS
	ret.HashRate = 666

	handler, err := NewHandler(
		context.Background(),
		WithID(&ret.ID, true),
		WithMiningpoolType(&ret.MiningpoolType, false),
		WithCoinType(&ret.CoinType, false),
		WithRevenueType(&ret.RevenueType, false),
		WithHashRate(&ret.HashRate, false),
		WithReadPageLink(nil, false),
	)
	assert.Nil(t, err)

	info, err := handler.UpdateGoodUser(context.Background())
	if assert.Nil(t, err) {
		ret.MiningpoolTypeStr = info.MiningpoolTypeStr
		ret.CoinTypeStr = info.CoinTypeStr
		ret.RevenueTypeStr = info.RevenueTypeStr
		ret.UpdatedAt = info.UpdatedAt
		assert.Equal(t, info, ret)
	}

	ret.MiningpoolType = basetypes.MiningpoolType_AntPool
	ret.GoodID = uuid.NewString()

	handler, err = NewHandler(
		context.Background(),
		WithID(&ret.ID, true),
		WithMiningpoolType(&ret.MiningpoolType, false),
		WithCoinType(&ret.CoinType, false),
		WithRevenueType(&ret.RevenueType, false),
		WithGoodID(&ret.GoodID, false),
		WithReadPageLink(nil, false),
	)
	assert.Nil(t, err)

	_, err = handler.UpdateGoodUser(context.Background())
	assert.Nil(t, err)
}

func delete(t *testing.T) {
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

	infos, total, err := handler.GetGoodUsers(context.Background())
	if assert.Nil(t, err) {
		assert.Equal(t, uint32(1), total)
		ret.MiningpoolTypeStr = infos[0].MiningpoolTypeStr
		ret.CoinTypeStr = infos[0].CoinTypeStr
		ret.RevenueTypeStr = infos[0].RevenueTypeStr
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

	exist, err := handler.ExistGoodUser(context.Background())
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
	deletedItem, err := handler.DeleteGoodUser(context.Background())
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

	exist, err = handler.ExistGoodUserConds(context.Background())
	if assert.Nil(t, err) {
		assert.Equal(t, false, exist)
	}
}

func TestGoodUser(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	t.Run("create", create)
	t.Run("update", update)
	t.Run("delete", delete)
}
