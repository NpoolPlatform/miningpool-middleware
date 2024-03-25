package fraction

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"

	"github.com/NpoolPlatform/go-service-framework/pkg/config"
	"github.com/NpoolPlatform/miningpool-middleware/pkg/testinit"

	"bou.ke/monkey"
	grpc2 "github.com/NpoolPlatform/go-service-framework/pkg/grpc"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

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
	AppID:       uuid.NewString(),
	UserID:      uuid.NewString(),
	OrderUserID: orderserRet.EntID,
}

var req = &npool.FractionReq{
	EntID:       &ret.EntID,
	AppID:       &ret.AppID,
	UserID:      &ret.UserID,
	OrderUserID: &ret.OrderUserID,
}

func createFraction(t *testing.T) {
	info, err := CreateFraction(context.Background(), req)
	if assert.Nil(t, err) {
		ret.CreatedAt = info.CreatedAt
		ret.WithdrawStateStr = info.WithdrawStateStr
		ret.WithdrawState = info.WithdrawState
		ret.WithdrawTime = info.WithdrawTime
		ret.PayTime = info.PayTime
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
	info, err := UpdateFraction(context.Background(), req)
	if assert.Nil(t, err) {
		ret.UpdatedAt = info.UpdatedAt
		assert.Equal(t, ret, info)
	}

	ret.Msg = "test1"
	req.Msg = &ret.Msg
	info, err = UpdateFraction(context.Background(), req)
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

	info, err := DeleteFraction(context.Background(), ret.ID)
	if assert.Nil(t, err) {
		assert.Equal(t, info, ret)
	}

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
	// Here won't pass test due to we always test with localhost

	gport := config.GetIntValueWithNameSpace("", config.KeyGRPCPort)

	monkey.Patch(grpc2.GetGRPCConn, func(service string, tags ...string) (*grpc.ClientConn, error) {
		return grpc.Dial(fmt.Sprintf("localhost:%v", gport), grpc.WithTransportCredentials(insecure.NewCredentials()))
	})
	monkey.Patch(grpc2.GetGRPCConnV1, func(service string, recvMsgBytes int, tags ...string) (*grpc.ClientConn, error) {
		return grpc.Dial(fmt.Sprintf("localhost:%v", gport), grpc.WithTransportCredentials(insecure.NewCredentials()))
	})
	t.Run("createOrderUser", createRootUser)
	t.Run("createOrderUser", createGoodUser)
	t.Run("createOrderUser", createOrderUser)
	t.Run("createFraction", createFraction)
	t.Run("updateFraction", updateFraction)
	t.Run("getFraction", getFraction)
	t.Run("getFractions", getFractions)
	t.Run("deleteFraction", deleteFraction)
	t.Run("deleteOrderUser", deleteOrderUser)
	t.Run("deleteOrderUser", deleteGoodUser)
	t.Run("deleteOrderUser", deleteRootUser)
}
