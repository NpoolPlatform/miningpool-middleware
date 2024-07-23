package fraction

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"

	"github.com/NpoolPlatform/miningpool-middleware/pkg/pools/registetestinfo"
	"github.com/NpoolPlatform/miningpool-middleware/pkg/testinit"

	"github.com/stretchr/testify/assert"

	cruder "github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	v1 "github.com/NpoolPlatform/message/npool/basetypes/v1"
	npool "github.com/NpoolPlatform/message/npool/miningpool/mw/v1/fraction"

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

var ret = &npool.Fraction{
	EntID:       uuid.NewString(),
	AppID:       orderserRet.AppID,
	UserID:      orderserRet.UserID,
	OrderUserID: orderserRet.EntID,
}

var req = &npool.FractionReq{
	EntID:       &ret.EntID,
	AppID:       &ret.AppID,
	UserID:      &ret.UserID,
	OrderUserID: &ret.OrderUserID,
}

func createFraction(t *testing.T) {
	err := CreateFraction(context.Background(), req)
	assert.Nil(t, err)
	info, err := GetFraction(context.Background(), ret.EntID)
	if assert.Nil(t, err) {
		ret.CreatedAt = info.CreatedAt
		ret.WithdrawStateStr = info.WithdrawStateStr
		ret.WithdrawState = info.WithdrawState
		ret.WithdrawAt = info.WithdrawAt
		ret.PromisePayAt = info.PromisePayAt
		ret.Msg = info.Msg
		ret.UpdatedAt = info.UpdatedAt
		ret.ID = info.ID
		ret.EntID = info.EntID
		assert.Equal(t, ret, info)
	}
}

func updateFraction(t *testing.T) {
	ret.Msg = "test"
	req.Msg = &ret.Msg
	req.ID = &ret.ID
	err := UpdateFraction(context.Background(), req)
	assert.Nil(t, err)

	info, err := GetFraction(context.Background(), *req.EntID)
	if assert.Nil(t, err) {
		ret.UpdatedAt = info.UpdatedAt
		assert.Equal(t, ret, info)
	}

	ret.Msg = "test1"
	req.Msg = &ret.Msg
	err = UpdateFraction(context.Background(), req)
	assert.Nil(t, err)

	info, err = GetFraction(context.Background(), *req.EntID)
	if assert.Nil(t, err) {
		ret.UpdatedAt = info.UpdatedAt
		assert.Equal(t, ret, info)
	}
}

func getFraction(t *testing.T) {
	info, err := GetFraction(context.Background(), ret.EntID)
	if assert.Nil(t, err) {
		assert.Equal(t, ret, info)
	}
}

func getFractions(t *testing.T) {
	infos, total, err := GetFractions(context.Background(), &npool.Conds{
		EntID: &v1.StringVal{Op: cruder.EQ, Value: ret.EntID},
	}, 0, 1)
	if assert.Nil(t, err) {
		assert.Equal(t, len(infos), 1)
		assert.Equal(t, total, uint32(1))
		assert.Equal(t, infos[0], ret)
	}
}

func deleteFraction(t *testing.T) {
	exist, err := ExistFractionConds(context.Background(), &npool.Conds{
		EntID: &v1.StringVal{
			Op:    cruder.EQ,
			Value: ret.EntID,
		},
	})
	if assert.Nil(t, err) {
		assert.Equal(t, true, exist)
	}

	err = DeleteFraction(context.Background(), ret.ID, ret.EntID)
	assert.Nil(t, err)

	exist, err = ExistFractionConds(context.Background(), &npool.Conds{
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
	t.Run("createFraction", createFraction)
	t.Run("updateFraction", updateFraction)
	t.Run("getFraction", getFraction)
	t.Run("getFractions", getFractions)
	t.Run("deleteFraction", deleteFraction)
	t.Run("deleteOrderUser", deleteOrderUser)
	t.Run("deleteGoodUser", deleteGoodUser)
	t.Run("deleteRootUser", deleteRootUser)
	registetestinfo.CleanTestInfo(context.Background())
}
