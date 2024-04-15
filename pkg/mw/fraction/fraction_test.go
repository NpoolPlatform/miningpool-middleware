package fraction

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	npool "github.com/NpoolPlatform/message/npool/miningpool/mw/v1/fraction"
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

var ret = &npool.Fraction{
	OrderUserID:   uuid.NewString(),
	AppID:         uuid.NewString(),
	UserID:        uuid.NewString(),
	WithdrawState: basetypes.WithdrawState_WithdrawStateProccessing,
	WithdrawTime:  5555,
	PayTime:       7777,
	Msg:           "7777",
}

var req = &npool.FractionReq{
	AppID:         &ret.AppID,
	UserID:        &ret.UserID,
	OrderUserID:   &ret.OrderUserID,
	WithdrawState: &ret.WithdrawState,
	WithdrawTime:  &ret.WithdrawTime,
	PayTime:       &ret.PayTime,
	Msg:           &ret.Msg,
}

func create(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithAppID(req.AppID, true),
		WithUserID(req.UserID, true),
		WithOrderUserID(req.OrderUserID, true),
		WithWithdrawState(req.WithdrawState, true),
		WithWithdrawTime(req.WithdrawTime, true),
		WithPayTime(req.PayTime, true),
		WithMsg(req.Msg, true),
	)
	assert.Nil(t, err)

	info, err := handler.CreateFraction(context.Background())
	if assert.Nil(t, err) {
		ret.UpdatedAt = info.UpdatedAt
		ret.CreatedAt = info.CreatedAt
		ret.WithdrawStateStr = info.WithdrawStateStr
		ret.ID = info.ID
		ret.EntID = info.EntID
		assert.Equal(t, info, ret)
	}
}

func update(t *testing.T) {
	ret.WithdrawState = basetypes.WithdrawState_WithdrawStateSuccess
	ret.WithdrawTime = 8888

	handler, err := NewHandler(
		context.Background(),
		WithID(&ret.ID, true),
		WithWithdrawState(&ret.WithdrawState, false),
		WithWithdrawTime(&ret.WithdrawTime, false),
		WithPayTime(nil, false),
	)
	assert.Nil(t, err)

	info, err := handler.UpdateFraction(context.Background())
	if assert.Nil(t, err) {
		ret.WithdrawStateStr = info.WithdrawStateStr
		ret.UpdatedAt = info.UpdatedAt
		assert.Equal(t, info, ret)
	}

	ret.WithdrawStateStr = basetypes.WithdrawState_WithdrawStateFailed.String()
	ret.PayTime = 9999

	handler, err = NewHandler(
		context.Background(),
		WithID(&ret.ID, true),
		WithWithdrawState(&ret.WithdrawState, false),
		WithWithdrawTime(&ret.WithdrawTime, false),
		WithPayTime(&ret.PayTime, false),
	)
	assert.Nil(t, err)

	_, err = handler.UpdateFraction(context.Background())
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
	deletedItem, err := handler.DeleteFraction(context.Background())
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

	exist, err = handler.ExistFractionConds(context.Background())
	if assert.Nil(t, err) {
		assert.Equal(t, false, exist)
	}
}

func TestFraction(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	t.Run("create", create)
	t.Run("update", update)
	t.Run("deleteRow", deleteRow)
}
