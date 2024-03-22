package coin

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	npool "github.com/NpoolPlatform/message/npool/miningpool/mw/v1/coin"
	testinit "github.com/NpoolPlatform/miningpool-middleware/pkg/testinit"

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

var ret = &npool.Coin{
	CoinType:         basetypes.CoinType_BitCoin,
	MiningpoolType:   basetypes.MiningpoolType_AntPool,
	RevenueTypes:     []basetypes.RevenueType{basetypes.RevenueType_FPPS, basetypes.RevenueType_PPLNS},
	FeeRate:          5.4,
	FixedRevenueAble: false,
	Threshold:        5.4,
	Remark:           "asdfaf",
}

var req = &npool.CoinReq{
	CoinType:         &ret.CoinType,
	MiningpoolType:   &ret.MiningpoolType,
	RevenueTypes:     ret.RevenueTypes,
	FeeRate:          &ret.FeeRate,
	FixedRevenueAble: &ret.FixedRevenueAble,
	Threshold:        &ret.Threshold,
	Remark:           &ret.Remark,
}

func create(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithCoinType(req.CoinType, true),
		WithMiningpoolType(req.MiningpoolType, true),
		WithRevenueTypes(&req.RevenueTypes, true),
		WithFeeRate(req.FeeRate, true),
		WithFixedRevenueAble(req.FixedRevenueAble, true),
		WithThreshold(req.Threshold, true),
		WithRemark(req.Remark, true),
	)
	assert.Nil(t, err)

	info, err := handler.CreateCoin(context.Background())
	if assert.Nil(t, err) {
		ret.UpdatedAt = info.UpdatedAt
		ret.CreatedAt = info.CreatedAt
		ret.MiningpoolTypeStr = info.MiningpoolTypeStr
		ret.CoinTypeStr = info.CoinTypeStr
		ret.RevenueTypesStr = info.RevenueTypesStr
		ret.ID = info.ID
		ret.EntID = info.EntID
		assert.Equal(t, ret, info)
	}
}

func update(t *testing.T) {
	ret.MiningpoolType = basetypes.MiningpoolType_F2Pool
	ret.FeeRate = 6.0

	handler, err := NewHandler(
		context.Background(),
		WithID(&ret.ID, true),
		WithMiningpoolType(&ret.MiningpoolType, false),
		WithFeeRate(&ret.FeeRate, false),
		WithRevenueTypes(nil, false),
	)
	assert.Nil(t, err)

	info, err := handler.UpdateCoin(context.Background())
	if assert.Nil(t, err) {
		ret.MiningpoolTypeStr = info.MiningpoolTypeStr
		ret.UpdatedAt = info.UpdatedAt
		assert.Equal(t, info, ret)
	}

	ret.MiningpoolType = basetypes.MiningpoolType_AntPool
	handler, err = NewHandler(
		context.Background(),
		WithID(&ret.ID, true),
		WithMiningpoolType(&ret.MiningpoolType, false),
		WithFeeRate(&ret.FeeRate, false),
		WithRevenueTypes(nil, false),
	)
	assert.Nil(t, err)

	_, err = handler.UpdateCoin(context.Background())
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

	infos, total, err := handler.GetCoins(context.Background())
	if assert.Nil(t, err) {
		assert.Equal(t, uint32(1), total)
		ret.MiningpoolTypeStr = infos[0].MiningpoolTypeStr
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

	exist, err := handler.ExistCoin(context.Background())
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
	deletedItem, err := handler.DeleteCoin(context.Background())
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

	exist, err = handler.ExistCoinConds(context.Background())
	if assert.Nil(t, err) {
		assert.Equal(t, false, exist)
	}
}

func TestCoin(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	t.Run("create", create)
	t.Run("update", update)
	t.Run("delete", delete)
}
