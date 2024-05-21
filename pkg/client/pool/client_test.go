package pool

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"

	"github.com/NpoolPlatform/miningpool-middleware/pkg/pools"
	"github.com/NpoolPlatform/miningpool-middleware/pkg/testinit"

	"github.com/stretchr/testify/assert"

	basetypes "github.com/NpoolPlatform/message/npool/basetypes/miningpool/v1"
	npool "github.com/NpoolPlatform/message/npool/miningpool/mw/v1/pool"

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
	EntID:          uuid.NewString(),
	MiningpoolType: basetypes.MiningpoolType_F2Pool,
	Name:           "ssss",
	Site:           "sss",
	Description:    "asdfadf",
	Logo:           "sdfasdf",
}

var req = &npool.PoolReq{
	EntID:          &ret.EntID,
	MiningpoolType: &ret.MiningpoolType,
	Name:           &ret.Name,
	Site:           &ret.Site,
	Logo:           &ret.Logo,
	Description:    &ret.Description,
}

func createPool(t *testing.T) {
	infos, _, err := GetPools(context.Background(), &npool.Conds{}, 0, 2)
	assert.Nil(t, err)

	if len(infos) == 0 {
		return
	}

	err = CreatePool(context.Background(), req)
	assert.NotNil(t, err)

	info, err := GetPool(context.Background(), infos[0].EntID)
	if assert.Nil(t, err) {
		assert.Equal(t, infos[0], info)
	}
}

func TestClient(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	pools.InitTestInfo(context.Background())
	t.Run("createPool", createPool)
	pools.CleanTestInfo(context.Background())
}
