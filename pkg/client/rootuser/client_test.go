package rootuser

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"

	"github.com/NpoolPlatform/miningpool-middleware/pkg/client/pool"
	"github.com/NpoolPlatform/miningpool-middleware/pkg/pools"
	"github.com/NpoolPlatform/miningpool-middleware/pkg/pools/f2pool"
	"github.com/NpoolPlatform/miningpool-middleware/pkg/testinit"

	poolmw "github.com/NpoolPlatform/message/npool/miningpool/mw/v1/pool"
	"github.com/stretchr/testify/assert"

	cruder "github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/miningpool/v1"
	v1 "github.com/NpoolPlatform/message/npool/basetypes/v1"
	npool "github.com/NpoolPlatform/message/npool/miningpool/mw/v1/rootuser"

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

var ret = &npool.RootUser{
	EntID:          uuid.NewString(),
	MiningpoolType: basetypes.MiningpoolType_F2Pool,
	Email:          "sssss@ss.com",
	AuthToken:      "7ecdq1fosdsfcruypom2otsn7hfr69azmqvh7v3zelol1ntsba85a1yvol66qp73",
	Authed:         true,
	Remark:         "sdfasdfasdf",
}

var req = &npool.RootUserReq{
	EntID:     &ret.EntID,
	Email:     &ret.Email,
	AuthToken: &ret.AuthToken,
	Authed:    &ret.Authed,
	Remark:    &ret.Remark,
}

func createRootUser(t *testing.T) {
	poolInfos, _, err := pool.GetPools(context.Background(), &poolmw.Conds{
		MiningpoolType: &v1.Uint32Val{
			Op:    cruder.EQ,
			Value: uint32(ret.MiningpoolType),
		},
	}, 0, 2)
	assert.Nil(t, err)
	assert.NotEqual(t, 0, len(poolInfos))

	ret.PoolID = poolInfos[0].EntID
	req.PoolID = &poolInfos[0].EntID

	name, err := f2pool.RandomF2PoolUser(7)
	assert.Nil(t, err)

	ret.Name = name
	req.Name = &name

	err = CreateRootUser(context.Background(), req)
	assert.Nil(t, err)

	info, err := GetRootUser(context.Background(), *req.EntID)
	if assert.Nil(t, err) {
		ret.CreatedAt = info.CreatedAt
		ret.MiningpoolTypeStr = info.MiningpoolTypeStr
		ret.UpdatedAt = info.UpdatedAt
		ret.ID = info.ID
		ret.AuthToken = info.AuthToken
		ret.EntID = info.EntID
		assert.Equal(t, ret, info)
	}
}

func updateRootUser(t *testing.T) {
	name, err := f2pool.RandomF2PoolUser(7)
	assert.Nil(t, err)

	ret.Name = name
	req.Name = &name
	req.ID = &ret.ID
	req.AuthToken = nil

	err = UpdateRootUser(context.Background(), req)
	assert.Nil(t, err)

	info, err := GetRootUser(context.Background(), *req.EntID)
	if assert.Nil(t, err) {
		ret.UpdatedAt = info.UpdatedAt
		assert.Equal(t, ret, info)
	}

	ret.Email = "test_api_client@ss.com"
	req.Email = &ret.Email
	err = UpdateRootUser(context.Background(), req)
	assert.Nil(t, err)

	info, err = GetRootUser(context.Background(), *req.EntID)
	if assert.Nil(t, err) {
		ret.UpdatedAt = info.UpdatedAt
		assert.Equal(t, ret, info)
	}
}

func getRootUser(t *testing.T) {
	info, err := GetRootUser(context.Background(), ret.EntID)
	if assert.Nil(t, err) {
		assert.Equal(t, ret, info)
	}
}

func getRootUsers(t *testing.T) {
	infos, total, err := GetRootUsers(context.Background(), &npool.Conds{
		EntID: &v1.StringVal{Op: cruder.EQ, Value: ret.EntID},
	}, 0, 1)
	if assert.Nil(t, err) {
		assert.Equal(t, len(infos), 1)
		assert.Equal(t, total, uint32(1))
		assert.Equal(t, infos[0], ret)
	}
}

func deleteRootUser(t *testing.T) {
	exist, err := ExistRootUserConds(context.Background(), &npool.Conds{
		EntID: &v1.StringVal{
			Op:    cruder.EQ,
			Value: ret.EntID,
		},
	})
	if assert.Nil(t, err) {
		assert.Equal(t, true, exist)
	}

	err = DeleteRootUser(context.Background(), ret.ID, ret.EntID)
	assert.Nil(t, err)

	exist, err = ExistRootUserConds(context.Background(), &npool.Conds{
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

	pools.InitTestInfo(context.Background())
	t.Run("createRootUser", createRootUser)
	t.Run("updateRootUser", updateRootUser)
	t.Run("getRootUser", getRootUser)
	t.Run("getRootUsers", getRootUsers)
	t.Run("deleteRootUser", deleteRootUser)
	pools.CleanTestInfo(context.Background())
}
