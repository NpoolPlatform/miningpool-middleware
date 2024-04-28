package gooduser

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	coinmw "github.com/NpoolPlatform/message/npool/miningpool/mw/v1/coin"
	npool "github.com/NpoolPlatform/message/npool/miningpool/mw/v1/gooduser"
	"github.com/NpoolPlatform/miningpool-middleware/pkg/mw/coin"
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
	EntID:       uuid.NewString(),
	RootUserID:  rootuserRet.EntID,
	HashRate:    5.0,
	CoinType:    basetypes.CoinType_BitCoin,
	RevenueType: basetypes.RevenueType_FPPS,
}

var req = &npool.GoodUserReq{
	EntID:      &ret.EntID,
	RootUserID: &ret.RootUserID,
	HashRate:   &ret.HashRate,
}

func create(t *testing.T) {
	coinH, err := coin.NewHandler(context.Background(),
		coin.WithConds(&coinmw.Conds{
			CoinType: &v1.Uint32Val{
				Op:    cruder.EQ,
				Value: uint32(ret.CoinType),
			},
		}),
		coin.WithOffset(0),
		coin.WithLimit(2),
	)
	assert.Nil(t, err)

	coinInfos, _, err := coinH.GetCoins(context.Background())
	assert.Nil(t, err)

	if !assert.NotEqual(t, 0, len(coinInfos)) {
		return
	}

	ret.CoinID = coinInfos[0].EntID
	req.CoinID = &coinInfos[0].EntID

	handler, err := NewHandler(
		context.Background(),
		WithEntID(req.EntID, true),
		WithRootUserID(req.RootUserID, true),
		WithCoinID(req.CoinID, true),
		WithHashRate(req.HashRate, true),
	)
	if !assert.Nil(t, err) {
		return
	}

	err = handler.CreateGoodUser(context.Background())
	if !assert.Nil(t, err) {
		return
	}

	info, err := handler.GetGoodUser(context.Background())
	if assert.Nil(t, err) {
		ret.UpdatedAt = info.UpdatedAt
		ret.CreatedAt = info.CreatedAt
		ret.MiningpoolTypeStr = info.MiningpoolTypeStr
		ret.RevenueTypeStr = info.RevenueTypeStr
		ret.ID = info.ID
		ret.EntID = info.EntID
		ret.Name = info.Name
		ret.ReadPageLink = info.ReadPageLink
		assert.Equal(t, info, ret)
	}
}

func update(t *testing.T) {
	ret.MiningpoolType = basetypes.MiningpoolType_F2Pool
	ret.RevenueType = basetypes.RevenueType_PPLNS
	ret.HashRate = 666

	handler, err := NewHandler(
		context.Background(),
		WithID(&ret.ID, true),
		WithHashRate(&ret.HashRate, false),
	)
	assert.Nil(t, err)

	err = handler.UpdateGoodUser(context.Background())
	assert.Nil(t, err)

	info, err := handler.GetGoodUser(context.Background())
	if assert.Nil(t, err) {
		ret.MiningpoolTypeStr = info.MiningpoolTypeStr
		ret.RevenueTypeStr = info.RevenueTypeStr
		ret.UpdatedAt = info.UpdatedAt
		assert.Equal(t, info, ret)
	}

	ret.HashRate = 6.18

	handler, err = NewHandler(
		context.Background(),
		WithID(&ret.ID, true),
		WithHashRate(&ret.HashRate, false),
		WithReadPageLink(nil, false),
	)
	assert.Nil(t, err)

	err = handler.UpdateGoodUser(context.Background())
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

	infos, total, err := handler.GetGoodUsers(context.Background())
	if assert.Nil(t, err) {
		assert.Equal(t, uint32(1), total)
		ret.MiningpoolTypeStr = infos[0].MiningpoolTypeStr
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
	err = handler.DeleteGoodUser(context.Background())
	assert.Nil(t, err)

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

	t.Run("create", createRootUser)
	t.Run("create", create)
	t.Run("update", update)
	t.Run("deleteRow", deleteRow)
	t.Run("deleteRow", deleteRootUser)
}
