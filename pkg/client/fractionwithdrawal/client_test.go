package fractionwithdrawal

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"

	coinmwpb "github.com/NpoolPlatform/message/npool/miningpool/mw/v1/coin"
	"github.com/NpoolPlatform/miningpool-middleware/pkg/client/coin"
	"github.com/NpoolPlatform/miningpool-middleware/pkg/pools/registetestinfo"
	"github.com/NpoolPlatform/miningpool-middleware/pkg/testinit"

	"github.com/stretchr/testify/assert"

	cruder "github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	v1 "github.com/NpoolPlatform/message/npool/basetypes/v1"
	npool "github.com/NpoolPlatform/message/npool/miningpool/mw/v1/fractionwithdrawal"

	"github.com/google/uuid"
)

func init() {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}
	if err := testinit.Init(); err != nil {
		fmt.Printf("cannot init test stub: %v\n", err)
	}
}

var ret = &npool.FractionWithdrawal{
	EntID:       uuid.NewString(),
	AppID:       orderserRet.AppID,
	UserID:      orderserRet.UserID,
	OrderUserID: orderserRet.EntID,
}

var req = &npool.FractionWithdrawalReq{
	EntID:       &ret.EntID,
	AppID:       &ret.AppID,
	UserID:      &ret.UserID,
	OrderUserID: &ret.OrderUserID,
	CoinTypeID:  &ret.CoinTypeID,
}

func createFractionWithdrawal(t *testing.T) {
	coinInfos, _, err := coin.GetCoins(context.Background(), &coinmwpb.Conds{
		PoolID: &v1.StringVal{Op: cruder.EQ, Value: orderserRet.PoolID},
	}, 0, 1)
	assert.Nil(t, err)
	if !assert.NotEqual(t, 0, len(coinInfos)) {
		return
	}

	req.CoinTypeID = &coinInfos[0].CoinTypeID
	ret.CoinTypeID = coinInfos[0].CoinTypeID
	err = CreateFractionWithdrawal(context.Background(), req)
	assert.Nil(t, err)
	info, err := GetFractionWithdrawal(context.Background(), ret.EntID)
	if assert.Nil(t, err) {
		ret.CreatedAt = info.CreatedAt
		ret.FractionWithdrawStateStr = info.FractionWithdrawStateStr
		ret.FractionWithdrawState = info.FractionWithdrawState
		ret.WithdrawAt = info.WithdrawAt
		ret.PromisePayAt = info.PromisePayAt
		ret.Msg = info.Msg
		ret.UpdatedAt = info.UpdatedAt
		ret.ID = info.ID
		ret.EntID = info.EntID
		assert.Equal(t, ret, info)
	}
}

func updateFractionWithdrawal(t *testing.T) {
	ret.Msg = "test"
	req.Msg = &ret.Msg
	req.ID = &ret.ID
	err := UpdateFractionWithdrawal(context.Background(), req)
	assert.Nil(t, err)

	info, err := GetFractionWithdrawal(context.Background(), *req.EntID)
	if assert.Nil(t, err) {
		ret.UpdatedAt = info.UpdatedAt
		assert.Equal(t, ret, info)
	}

	ret.Msg = "test1"
	req.Msg = &ret.Msg
	err = UpdateFractionWithdrawal(context.Background(), req)
	assert.Nil(t, err)

	info, err = GetFractionWithdrawal(context.Background(), *req.EntID)
	if assert.Nil(t, err) {
		ret.UpdatedAt = info.UpdatedAt
		assert.Equal(t, ret, info)
	}
}

func getFractionWithdrawal(t *testing.T) {
	info, err := GetFractionWithdrawal(context.Background(), ret.EntID)
	if assert.Nil(t, err) {
		assert.Equal(t, ret, info)
	}
}

func getFractionWithdrawals(t *testing.T) {
	infos, total, err := GetFractionWithdrawals(context.Background(), &npool.Conds{
		EntID: &v1.StringVal{Op: cruder.EQ, Value: ret.EntID},
	}, 0, 1)
	if assert.Nil(t, err) {
		assert.Equal(t, len(infos), 1)
		assert.Equal(t, total, uint32(1))
		assert.Equal(t, infos[0], ret)
	}
}

func deleteFractionWithdrawal(t *testing.T) {
	exist, err := ExistFractionWithdrawalConds(context.Background(), &npool.Conds{
		EntID: &v1.StringVal{
			Op:    cruder.EQ,
			Value: ret.EntID,
		},
	})
	if assert.Nil(t, err) {
		assert.Equal(t, true, exist)
	}

	err = DeleteFractionWithdrawal(context.Background(), ret.ID, ret.EntID)
	assert.Nil(t, err)

	exist, err = ExistFractionWithdrawalConds(context.Background(), &npool.Conds{
		EntID: &v1.StringVal{
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

	registetestinfo.InitTestInfo(context.Background())
	// Here won't pass test due to we always test with localhost
	t.Run("createRootUser", createRootUser)
	t.Run("createGoodUser", createGoodUser)
	t.Run("createOrderUser", createOrderUser)
	t.Run("createFractionWithdrawal", createFractionWithdrawal)
	t.Run("updateFractionWithdrawal", updateFractionWithdrawal)
	t.Run("getFractionWithdrawal", getFractionWithdrawal)
	t.Run("getFractionWithdrawals", getFractionWithdrawals)
	t.Run("deleteFractionWithdrawal", deleteFractionWithdrawal)
	t.Run("deleteOrderUser", deleteOrderUser)
	t.Run("deleteGoodUser", deleteGoodUser)
	t.Run("deleteRootUser", deleteRootUser)
	registetestinfo.CleanTestInfo(context.Background())
}
