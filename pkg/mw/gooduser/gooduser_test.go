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
	"github.com/NpoolPlatform/miningpool-middleware/pkg/pools/registetestinfo"
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

var gooduserRet = &npool.GoodUser{
	EntID:          uuid.NewString(),
	RootUserID:     rootuserRet.EntID,
	MiningpoolType: mpbasetypes.MiningpoolType_F2Pool,
}

var gooduserReq = &npool.GoodUserReq{
	EntID:      &gooduserRet.EntID,
	RootUserID: &gooduserRet.RootUserID,
}

func create(t *testing.T) {
	coinH, err := coin.NewHandler(context.Background(),
		coin.WithConds(&coinmw.Conds{
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

	for _, coinInfo := range coinInfos {
		gooduserReq.CoinTypeIDs = append(gooduserReq.CoinTypeIDs, coinInfo.CoinTypeID)
	}

	handler, err := NewHandler(
		context.Background(),
		WithEntID(gooduserReq.EntID, true),
		WithRootUserID(gooduserReq.RootUserID, true),
		WithCoinTypeIDs(gooduserReq.CoinTypeIDs, true),
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
		gooduserRet.UpdatedAt = info.UpdatedAt
		gooduserRet.CreatedAt = info.CreatedAt
		gooduserRet.PoolID = info.PoolID
		gooduserRet.MiningpoolTypeStr = info.MiningpoolTypeStr
		gooduserRet.FeeRatio = info.FeeRatio
		gooduserRet.ID = info.ID
		gooduserRet.EntID = info.EntID
		gooduserRet.Name = info.Name
		gooduserRet.ReadPageLink = info.ReadPageLink
		assert.Equal(t, info, gooduserRet)
	}
}

func update(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithID(&gooduserRet.ID, true),
	)
	assert.Nil(t, err)

	err = handler.UpdateGoodUser(context.Background())
	assert.NotNil(t, err)

	info, err := handler.GetGoodUser(context.Background())
	if assert.Nil(t, err) {
		gooduserRet.MiningpoolTypeStr = info.MiningpoolTypeStr
		gooduserRet.UpdatedAt = info.UpdatedAt
		assert.Equal(t, info, gooduserRet)
	}
}

func deleteRow(t *testing.T) {
	conds := &npool.Conds{
		EntID: &basetypes.StringVal{
			Op:    cruder.EQ,
			Value: gooduserRet.EntID,
		},
	}
	handler, err := NewHandler(
		context.Background(),
		WithConds(conds),
		WithID(&gooduserRet.ID, true),
		WithOffset(0),
		WithLimit(2),
	)
	assert.Nil(t, err)

	infos, total, err := handler.GetGoodUsers(context.Background())
	if assert.Nil(t, err) {
		assert.Equal(t, uint32(1), total)
		gooduserRet.MiningpoolTypeStr = infos[0].MiningpoolTypeStr
		gooduserRet.UpdatedAt = infos[0].UpdatedAt
		assert.Equal(t, infos[0], gooduserRet)
	}

	gooduserRet.EntID = infos[0].EntID
	handler, err = NewHandler(
		context.Background(),
		WithEntID(&gooduserRet.EntID, true),
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
		WithID(&gooduserRet.ID, true),
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
				Value: gooduserRet.EntID,
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
	registetestinfo.InitTestInfo(context.Background())
	t.Run("create", createRootUser)
	t.Run("create", create)
	t.Run("update", update)
	t.Run("deleteRow", deleteRow)
	t.Run("deleteRow", deleteRootUser)
	registetestinfo.CleanTestInfo(context.Background())
}
