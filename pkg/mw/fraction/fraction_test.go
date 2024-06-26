package fraction

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"

	"bou.ke/monkey"
	"github.com/NpoolPlatform/go-service-framework/pkg/config"
	grpc2 "github.com/NpoolPlatform/go-service-framework/pkg/grpc"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	npool "github.com/NpoolPlatform/message/npool/miningpool/mw/v1/fraction"
	"github.com/NpoolPlatform/miningpool-middleware/pkg/pools"
	testinit "github.com/NpoolPlatform/miningpool-middleware/pkg/testinit"
	"github.com/google/uuid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

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

var ret = &npool.Fraction{
	EntID:       uuid.NewString(),
	OrderUserID: orderuserRet.EntID,
	AppID:       orderuserRet.AppID,
	UserID:      orderuserRet.UserID,
}

var req = &npool.FractionReq{
	EntID:         &ret.EntID,
	AppID:         &ret.AppID,
	UserID:        &ret.UserID,
	OrderUserID:   &ret.OrderUserID,
	WithdrawState: &ret.WithdrawState,
	WithdrawAt:    &ret.WithdrawAt,
	PromisePayAt:  &ret.PromisePayAt,
}

func create(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithEntID(req.EntID, true),
		WithAppID(req.AppID, true),
		WithUserID(req.UserID, true),
		WithOrderUserID(req.OrderUserID, true),
	)
	assert.Nil(t, err)

	err = handler.CreateFraction(context.Background())
	assert.Nil(t, err)

	info, err := handler.GetFraction(context.Background())
	if assert.Nil(t, err) {
		ret.UpdatedAt = info.UpdatedAt
		ret.CreatedAt = info.CreatedAt
		ret.WithdrawStateStr = info.WithdrawStateStr
		ret.ID = info.ID
		ret.EntID = info.EntID
		ret.WithdrawState = info.WithdrawState
		ret.WithdrawAt = info.WithdrawAt
		ret.PromisePayAt = info.PromisePayAt
		ret.Msg = info.Msg
		assert.Equal(t, info, ret)
	}
}

func update(t *testing.T) {
	ret.WithdrawState = basetypes.WithdrawState_WithdrawStateSuccess
	ret.WithdrawAt = 8888

	handler, err := NewHandler(
		context.Background(),
		WithID(&ret.ID, true),
		WithWithdrawState(&ret.WithdrawState, false),
		WithWithdrawAt(&ret.WithdrawAt, false),
		WithPromisePayAt(nil, false),
	)
	assert.Nil(t, err)

	err = handler.UpdateFraction(context.Background())
	assert.Nil(t, err)

	info, err := handler.GetFraction(context.Background())
	if assert.Nil(t, err) {
		ret.WithdrawStateStr = info.WithdrawStateStr
		ret.UpdatedAt = info.UpdatedAt
		assert.Equal(t, info, ret)
	}

	ret.WithdrawStateStr = basetypes.WithdrawState_WithdrawStateFailed.String()
	ret.PromisePayAt = 9999

	handler, err = NewHandler(
		context.Background(),
		WithID(&ret.ID, true),
		WithWithdrawState(&ret.WithdrawState, false),
		WithWithdrawAt(&ret.WithdrawAt, false),
		WithPromisePayAt(&ret.PromisePayAt, false),
	)
	assert.Nil(t, err)

	err = handler.UpdateFraction(context.Background())
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

	infos, total, err := handler.GetFractions(context.Background())
	if assert.Nil(t, err) {
		assert.Equal(t, uint32(1), total)
		ret.WithdrawStateStr = infos[0].WithdrawStateStr
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

	exist, err := handler.ExistFraction(context.Background())
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
	err = handler.DeleteFraction(context.Background())
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

	exist, err = handler.ExistFractionConds(context.Background())
	if assert.Nil(t, err) {
		assert.Equal(t, false, exist)
	}
}

func TestFraction(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}
	gport := config.GetIntValueWithNameSpace("", config.KeyGRPCPort)

	monkey.Patch(grpc2.GetGRPCConn, func(service string, tags ...string) (*grpc.ClientConn, error) {
		return grpc.Dial(fmt.Sprintf("localhost:%v", gport), grpc.WithTransportCredentials(insecure.NewCredentials()))
	})
	monkey.Patch(grpc2.GetGRPCConnV1, func(service string, recvMsgBytes int, tags ...string) (*grpc.ClientConn, error) {
		return grpc.Dial(fmt.Sprintf("localhost:%v", gport), grpc.WithTransportCredentials(insecure.NewCredentials()))
	})

	pools.InitTestInfo(context.Background())
	t.Run("create", createRootUser)
	t.Run("create", createGoodUser)
	t.Run("create", createOrderUser)
	t.Run("create", create)
	t.Run("update", update)
	t.Run("deleteRow", deleteRow)
	t.Run("deleteRow", deleteOrderUser)
	t.Run("deleteRow", deleteGoodUser)
	t.Run("deleteRow", deleteRootUser)
	pools.CleanTestInfo(context.Background())
}
