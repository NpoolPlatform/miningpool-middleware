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
	"github.com/NpoolPlatform/miningpool-middleware/pkg/pools"
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

var ret = &npool.GoodUser{
	EntID:          uuid.NewString(),
	RootUserID:     rootuserRet.EntID,
	CoinType:       basetypes.CoinType_CoinTypeBitCoin,
	MiningpoolType: mpbasetypes.MiningpoolType_F2Pool,
}

var req = &npool.GoodUserReq{
	EntID:      &ret.EntID,
	RootUserID: &ret.RootUserID,
}

func create(t *testing.T) {
	coinH, err := coin.NewHandler(context.Background(),
		coin.WithConds(&coinmw.Conds{
			CoinType: &basetypes.Uint32Val{
				Op:    cruder.EQ,
				Value: uint32(ret.CoinType),
			},
			PoolID: &basetypes.StringVal{
				Op:    cruder.EQ,
				Value: rootuserRet.PoolID,
			},
		}),
		coin.WithOffset(0),
		coin.WithLimit(2),
	)
	assert.Nil(t, err)

	coinInfos, _, err := coinH.GetCoins(context.Background())
	assert.Nil(t, err)

	if len(coinInfos) == 0 {
		return
	}

	ret.PoolCoinTypeID = coinInfos[0].EntID
	req.PoolCoinTypeID = &coinInfos[0].EntID

	handler, err := NewHandler(
		context.Background(),
		WithEntID(req.EntID, true),
		WithRootUserID(req.RootUserID, true),
		WithPoolCoinTypeID(req.PoolCoinTypeID, true),
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
		ret.PoolID = info.PoolID
		ret.MiningpoolTypeStr = info.MiningpoolTypeStr
		ret.CoinTypeStr = info.CoinTypeStr
		ret.FeeRatio = info.FeeRatio
		ret.ID = info.ID
		ret.EntID = info.EntID
		ret.Name = info.Name
		ret.ReadPageLink = info.ReadPageLink
		assert.Equal(t, info, ret)
	}
}

func update(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithID(&ret.ID, true),
	)
	assert.Nil(t, err)

	err = handler.UpdateGoodUser(context.Background())
	assert.NotNil(t, err)

	info, err := handler.GetGoodUser(context.Background())
	if assert.Nil(t, err) {
		ret.MiningpoolTypeStr = info.MiningpoolTypeStr
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

	infos, total, err := handler.GetGoodUsers(context.Background())
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
			EntID: &basetypes.StringVal{
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
	pools.InitTestInfo(context.Background())
	t.Run("create", createRootUser)
	t.Run("create", create)
	t.Run("update", update)
	t.Run("deleteRow", deleteRow)
	t.Run("deleteRow", deleteRootUser)
	pools.CleanTestInfo(context.Background())
}
