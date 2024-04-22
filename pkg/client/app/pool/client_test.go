package pool

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"

	"github.com/NpoolPlatform/miningpool-middleware/pkg/client/pool"
	"github.com/NpoolPlatform/miningpool-middleware/pkg/testinit"

	"github.com/stretchr/testify/assert"

	cruder "github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	v1 "github.com/NpoolPlatform/message/npool/basetypes/v1"
	npool "github.com/NpoolPlatform/message/npool/miningpool/mw/v1/app/pool"

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

var ret = &npool.Pool{
	EntID:  uuid.NewString(),
	AppID:  uuid.NewString(),
	PoolID: uuid.NewString(),
}

var req = &npool.PoolReq{
	EntID:  &ret.EntID,
	AppID:  &ret.AppID,
	PoolID: &ret.PoolID,
}

func createPool(t *testing.T) {
	infos, _, err := pool.GetPools(context.Background(), nil, 0, 2)
	assert.Nil(t, err)
	fmt.Println(err)
	if len(infos) == 0 {
		return
	}

	ret.PoolID = infos[0].EntID
	req.PoolID = &infos[0].EntID

	err = CreatePool(context.Background(), req)
	assert.Nil(t, err)

	info, err := GetPool(context.Background(), *req.EntID)
	if assert.Nil(t, err) {
		ret.CreatedAt = info.CreatedAt
		ret.UpdatedAt = info.UpdatedAt
		ret.ID = info.ID
		ret.EntID = info.EntID
		assert.Equal(t, ret, info)
	}
}

func getPool(t *testing.T) {
	info, err := GetPool(context.Background(), ret.EntID)
	if assert.Nil(t, err) {
		assert.Equal(t, info, ret)
	}
}

func getPools(t *testing.T) {
	infos, total, err := GetPools(context.Background(), &npool.Conds{
		EntID: &v1.StringVal{Op: cruder.EQ, Value: ret.EntID},
	}, 0, 1)
	if assert.Nil(t, err) {
		assert.Equal(t, len(infos), 1)
		assert.Equal(t, total, uint32(1))
		assert.Equal(t, infos[0], ret)
	}
}

func deletePool(t *testing.T) {
	exist, err := ExistPoolConds(context.Background(), &npool.Conds{
		EntID: &v1.StringVal{
			Op:    cruder.EQ,
			Value: ret.EntID,
		},
	})
	if assert.Nil(t, err) {
		assert.Equal(t, true, exist)
	}

	err = DeletePool(context.Background(), ret.ID, ret.EntID)
	assert.Nil(t, err)

	exist, err = ExistPoolConds(context.Background(), &npool.Conds{
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
	t.Run("createPool", createPool)
	t.Run("getPool", getPool)
	t.Run("getPools", getPools)
	t.Run("deletePool", deletePool)
}
