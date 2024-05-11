package rootuser

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	poolmw "github.com/NpoolPlatform/message/npool/miningpool/mw/v1/pool"
	npool "github.com/NpoolPlatform/message/npool/miningpool/mw/v1/rootuser"
	"github.com/NpoolPlatform/miningpool-middleware/pkg/mw/pool"
	"github.com/NpoolPlatform/miningpool-middleware/pkg/pools"
	"github.com/NpoolPlatform/miningpool-middleware/pkg/pools/f2pool"
	testinit "github.com/NpoolPlatform/miningpool-middleware/pkg/testinit"
	"github.com/google/uuid"

	basetype "github.com/NpoolPlatform/message/npool/basetypes/miningpool/v1"
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
	PoolID:         uuid.NewString(),
	Email:          "gggo@go.go",
	MiningpoolType: basetype.MiningpoolType_F2Pool,
	AuthToken:      "7ecdq1fosdsfcruypom2otsn8hfr69azmqvh7v3zelol1ntsba85a1yvol66qp73",
	Authed:         true,
	Remark:         "asdfaf",
}

var req = &npool.RootUserReq{
	EntID:     &ret.EntID,
	PoolID:    &ret.PoolID,
	Email:     &ret.Email,
	AuthToken: &ret.AuthToken,
	Authed:    &ret.Authed,
	Remark:    &ret.Remark,
}

func create(t *testing.T) {
	name, err := f2pool.RandomF2PoolUser(8)
	assert.Nil(t, err)

	ret.Name = name
	req.Name = &name
	handler, err := NewHandler(
		context.Background(),
		WithName(req.Name, true),
		WithPoolID(req.PoolID, true),
		WithEmail(req.Email, true),
		WithAuthToken(req.AuthToken, true),
		WithAuthed(req.Authed, true),
		WithRemark(req.Remark, true),
	)
	assert.Nil(t, handler)
	assert.NotNil(t, err)

	poolH, err := pool.NewHandler(context.Background(),
		pool.WithConds(&poolmw.Conds{
			MiningpoolType: &v1.Uint32Val{
				Op:    cruder.EQ,
				Value: uint32(ret.MiningpoolType),
			},
		}),
		pool.WithLimit(2),
		pool.WithOffset(0),
	)
	assert.Nil(t, err)

	poolInfos, _, err := poolH.GetPools(context.Background())
	assert.Nil(t, err)
	if !assert.NotEqual(t, len(poolInfos), 0) {
		return
	}

	ret.PoolID = poolInfos[0].EntID
	req.PoolID = &poolInfos[0].EntID

	handler, err = NewHandler(
		context.Background(),
		WithName(req.Name, true),
		WithPoolID(req.PoolID, true),
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
		ret.AuthToken = info.AuthToken
		ret.ID = info.ID
		ret.EntID = info.EntID
		assert.Equal(t, info, ret)
	}
}

func update(t *testing.T) {
	ret.Email = "ok@true.right"
	req.Email = &ret.Email

	handler, err := NewHandler(
		context.Background(),
		WithID(&ret.ID, true),
		WithEmail(&ret.Email, false),
	)
	assert.Nil(t, err)

	err = handler.UpdateRootUser(context.Background())
	assert.Nil(t, err)

	info, err := handler.GetRootUser(context.Background())
	if assert.Nil(t, err) {
		ret.UpdatedAt = info.UpdatedAt
		assert.Equal(t, info, ret)
	}
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
	err = handler.DeleteRootUser(context.Background())
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

	exist, err = handler.ExistRootUserConds(context.Background())
	if assert.Nil(t, err) {
		assert.Equal(t, false, exist)
	}
}

func TestRootUser(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	pools.InitTestInfo(context.Background())
	t.Run("create", create)
	t.Run("update", update)
	t.Run("deleteRow", deleteRow)
	pools.CleanTestInfo(context.Background())
}
