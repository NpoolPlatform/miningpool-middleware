package pool

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"

	npool "github.com/NpoolPlatform/message/npool/miningpool/mw/v1/pool"
	testinit "github.com/NpoolPlatform/miningpool-middleware/pkg/testinit"
	"github.com/google/uuid"

	basetypes "github.com/NpoolPlatform/message/npool/basetypes/miningpool/v1"

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

var ret = &npool.Pool{
	EntID:          uuid.NewString(),
	MiningpoolType: basetypes.MiningpoolType_AntPool,
	Name:           "5.4",
	Site:           "false.com",
	Description:    "asdfaf",
}

var req = &npool.PoolReq{
	EntID:          &ret.EntID,
	MiningpoolType: &ret.MiningpoolType,
	Name:           &ret.Name,
	Site:           &ret.Site,
	Description:    &ret.Description,
}

func create(t *testing.T) {
	getH, err := NewHandler(
		context.Background(),
		WithConds(nil),
		WithOffset(0),
		WithLimit(2),
	)
	assert.Nil(t, err)

	infos, _, err := getH.GetPools(context.Background())
	if len(infos) == 0 || err != nil {
		return
	}

	ret.MiningpoolType = infos[0].MiningpoolType
	req.MiningpoolType = &infos[0].MiningpoolType

	handler, err := NewHandler(
		context.Background(),
		WithEntID(req.EntID, true),
		WithMiningpoolType(req.MiningpoolType, true),
		WithName(req.Name, true),
		WithSite(req.Site, true),
		WithDescription(req.Description, true),
	)
	assert.Nil(t, err)

	err = handler.CreatePool(context.Background())
	assert.NotNil(t, err)
}

func TestPool(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	t.Run("create", create)
}
