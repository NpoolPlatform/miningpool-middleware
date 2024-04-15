package rootuser

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	npool "github.com/NpoolPlatform/message/npool/miningpool/mw/v1/rootuser"
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

var ret = &npool.RootUser{
	EntID:          uuid.NewString(),
	Name:           "mm_mw_rootuser_test_data",
	MiningpoolType: basetypes.MiningpoolType_AntPool,
	Email:          "gggo@go.go",
	AuthToken:      "7ecdq1fosdsfcruypom2otsn8hfr69azmqvh7v3zelol1ntsba85a1yvol66qp73",
	Authed:         true,
	Remark:         "asdfaf",
}

var req = &npool.RootUserReq{
	EntID:          &ret.EntID,
	Name:           &ret.Name,
	MiningpoolType: &ret.MiningpoolType,
	Email:          &ret.Email,
	AuthToken:      &ret.AuthToken,
	Authed:         &ret.Authed,
	Remark:         &ret.Remark,
}

func create(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithName(req.Name, true),
		WithMiningpoolType(req.MiningpoolType, true),
		WithEmail(req.Email, true),
		WithAuthToken(req.AuthToken, true),
		WithAuthed(req.Authed, true),
		WithRemark(req.Remark, true),
	)
	assert.Nil(t, err)

	err = handler.CreateRootUser(context.Background())
	assert.Nil(t, err)

	info, err := handler.GetRootUser(context.Background())
	if assert.Nil(t, err) {
		ret.UpdatedAt = info.UpdatedAt
		ret.CreatedAt = info.CreatedAt
		ret.MiningpoolTypeStr = info.MiningpoolTypeStr
		ret.ID = info.ID
		ret.EntID = info.EntID
		assert.Equal(t, info, ret)
	}
}

func update(t *testing.T) {
	ret.MiningpoolType = basetypes.MiningpoolType_F2Pool

	handler, err := NewHandler(
		context.Background(),
		WithID(&ret.ID, true),
		WithMiningpoolType(&ret.MiningpoolType, false),
		WithEmail(nil, false),
	)
	assert.Nil(t, err)

	err = handler.UpdateRootUser(context.Background())
	assert.Nil(t, err)

	info, err := handler.GetRootUser(context.Background())
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
		WithEmail(nil, false),
	)
	assert.Nil(t, err)

	err = handler.UpdateRootUser(context.Background())
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

	infos, total, err := handler.GetRootUsers(context.Background())
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

	exist, err := handler.ExistRootUser(context.Background())
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
	deletedItem, err := handler.DeleteRootUser(context.Background())
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

	exist, err = handler.ExistRootUserConds(context.Background())
	if assert.Nil(t, err) {
		assert.Equal(t, false, exist)
	}
}

func TestRootUser(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	t.Run("create", create)
	t.Run("update", update)
	t.Run("deleteRow", deleteRow)
}
