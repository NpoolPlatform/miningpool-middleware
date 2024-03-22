package pool

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
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/miningpool/v1"
	v1 "github.com/NpoolPlatform/message/npool/basetypes/v1"
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
}

var req = &npool.PoolReq{
	EntID:          &ret.EntID,
	MiningpoolType: &ret.MiningpoolType,
	Name:           &ret.Name,
	Site:           &ret.Site,
	Description:    &ret.Description,
}

func createPool(t *testing.T) {
	info, err := CreatePool(context.Background(), req)
	if assert.Nil(t, err) {
		ret.CreatedAt = info.CreatedAt
		ret.MiningpoolTypeStr = info.MiningpoolTypeStr
		ret.UpdatedAt = info.UpdatedAt
		ret.ID = info.ID
		ret.EntID = info.EntID
		assert.Equal(t, ret, info)
	}
}

func updatePool(t *testing.T) {
	ret.Name = "zzzzzzzz"
	req.Name = &ret.Name
	req.ID = &ret.ID

	info, err := UpdatePool(context.Background(), req)
	if assert.Nil(t, err) {
		ret.UpdatedAt = info.UpdatedAt
		assert.Equal(t, info, ret)
	}

	ret.Site = "zzzzzzzz"
	req.Site = &ret.Site
	info, err = UpdatePool(context.Background(), req)
	if assert.Nil(t, err) {
		ret.UpdatedAt = info.UpdatedAt
		assert.Equal(t, info, ret)
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

	info, err := DeletePool(context.Background(), &npool.PoolReq{
		ID: &ret.ID,
	})
	if assert.Nil(t, err) {
		assert.Equal(t, info, ret)
	}

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
	// Here won't pass test due to we always test with localhost

	gport := config.GetIntValueWithNameSpace("", config.KeyGRPCPort)

	monkey.Patch(grpc2.GetGRPCConn, func(service string, tags ...string) (*grpc.ClientConn, error) {
		return grpc.Dial(fmt.Sprintf("localhost:%v", gport), grpc.WithTransportCredentials(insecure.NewCredentials()))
	})
	monkey.Patch(grpc2.GetGRPCConnV1, func(service string, recvMsgBytes int, tags ...string) (*grpc.ClientConn, error) {
		return grpc.Dial(fmt.Sprintf("localhost:%v", gport), grpc.WithTransportCredentials(insecure.NewCredentials()))
	})

	t.Run("createPool", createPool)
	t.Run("updatePool", updatePool)
	t.Run("getPool", getPool)
	t.Run("getPools", getPools)
	t.Run("deletePool", deletePool)
}
