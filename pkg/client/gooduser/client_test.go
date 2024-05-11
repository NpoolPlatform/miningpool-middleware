package gooduser

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"

	"github.com/NpoolPlatform/miningpool-middleware/pkg/client/coin"
	"github.com/NpoolPlatform/miningpool-middleware/pkg/testinit"
	"github.com/google/uuid"

	coinmw "github.com/NpoolPlatform/message/npool/miningpool/mw/v1/coin"
	"github.com/stretchr/testify/assert"

	cruder "github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	mpbasetypes "github.com/NpoolPlatform/message/npool/basetypes/miningpool/v1"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	npool "github.com/NpoolPlatform/message/npool/miningpool/mw/v1/gooduser"
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
	Name:           "fffff",
	RootUserID:     rootUserRet.EntID,
	PoolCoinTypeID: uuid.NewString(),
	MiningpoolType: mpbasetypes.MiningpoolType_F2Pool,
	CoinType:       basetypes.CoinType_CoinTypeBitCoin,
	ReadPageLink:   "fffff",
	RevenueType:    mpbasetypes.RevenueType_FPPS,
}

var req = &npool.GoodUserReq{
	EntID:          &ret.EntID,
	PoolCoinTypeID: &ret.PoolCoinTypeID,
	Name:           &ret.Name,
	RootUserID:     &ret.RootUserID,
	ReadPageLink:   &ret.ReadPageLink,
}

func createGoodUser(t *testing.T) {
	coinInfos, _, err := coin.GetCoins(context.Background(), &coinmw.Conds{
		CoinType: &basetypes.Uint32Val{
			Op:    cruder.EQ,
			Value: uint32(ret.CoinType),
		},
		PoolID: &basetypes.StringVal{
			Op:    cruder.EQ,
			Value: rootUserRet.PoolID,
		},
	}, 0, 2)
	assert.Nil(t, err)
	assert.NotEqual(t, 0, len(coinInfos))

	ret.PoolCoinTypeID = coinInfos[0].EntID
	req.PoolCoinTypeID = &coinInfos[0].EntID

	err = CreateGoodUser(context.Background(), req)
	assert.Nil(t, err)

	info, err := GetGoodUser(context.Background(), *req.EntID)
	if assert.Nil(t, err) {
		ret.Name = info.Name
		ret.ReadPageLink = info.ReadPageLink
		ret.CreatedAt = info.CreatedAt
		ret.PoolID = info.PoolID
		ret.MiningpoolTypeStr = info.MiningpoolTypeStr
		ret.CoinTypeStr = info.CoinTypeStr
		ret.RevenueTypeStr = info.RevenueTypeStr
		ret.FeeRatio = info.FeeRatio
		ret.UpdatedAt = info.UpdatedAt
		ret.ID = info.ID
		ret.EntID = info.EntID
		assert.Equal(t, ret, info)
	}
}

func updateGoodUser(t *testing.T) {
	req.ID = &ret.ID
	err := UpdateGoodUser(context.Background(), req)
	assert.Nil(t, err)

	info, err := GetGoodUser(context.Background(), *req.EntID)
	if assert.Nil(t, err) {
		ret.UpdatedAt = info.UpdatedAt
		assert.Equal(t, info, ret)
	}

	req.ID = &ret.ID
	err = UpdateGoodUser(context.Background(), req)
	assert.Nil(t, err)

	info, err = GetGoodUser(context.Background(), *req.EntID)
	if assert.Nil(t, err) {
		ret.UpdatedAt = info.UpdatedAt
		assert.Equal(t, info, ret)
	}
}

func getGoodUser(t *testing.T) {
	info, err := GetGoodUser(context.Background(), ret.EntID)
	if assert.Nil(t, err) {
		assert.Equal(t, info, ret)
	}
}

func getGoodUsers(t *testing.T) {
	infos, total, err := GetGoodUsers(context.Background(), &npool.Conds{
		EntID: &basetypes.StringVal{Op: cruder.EQ, Value: ret.EntID},
	}, 0, 1)
	if assert.Nil(t, err) {
		assert.Equal(t, len(infos), 1)
		assert.Equal(t, total, uint32(1))
		assert.Equal(t, infos[0], ret)
	}
}

func deleteGoodUser(t *testing.T) {
	exist, err := ExistGoodUserConds(context.Background(), &npool.Conds{
		EntID: &basetypes.StringVal{
			Op:    cruder.EQ,
			Value: ret.EntID,
		},
	})
	if assert.Nil(t, err) {
		assert.Equal(t, true, exist)
	}

	err = DeleteGoodUser(context.Background(), ret.ID, ret.EntID)
	assert.Nil(t, err)

	exist, err = ExistGoodUserConds(context.Background(), &npool.Conds{
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

	t.Run("createRootUser", createRootUser)
	t.Run("createGoodUser", createGoodUser)
	t.Run("updateGoodUser", updateGoodUser)
	t.Run("getGoodUser", getGoodUser)
	t.Run("getGoodUsers", getGoodUsers)
	t.Run("deleteGoodUser", deleteGoodUser)
	t.Run("deleteRootUser", deleteRootUser)
}
