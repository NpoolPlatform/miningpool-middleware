package rootuser

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
	Name:           "mininpool_middleware_client_rootuser_test",
	MiningpoolType: basetypes.MiningpoolType_F2Pool,
	Email:          "sssss",
	AuthToken:      "sadfasdf",
	Authed:         false,
	Remark:         "sdfasdfasdf",
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

func createRootUser(t *testing.T) {
	info, err := CreateRootUser(context.Background(), req)
	if assert.Nil(t, err) {
		ret.CreatedAt = info.CreatedAt
		ret.MiningpoolTypeStr = info.MiningpoolTypeStr
		ret.UpdatedAt = info.UpdatedAt
		ret.ID = info.ID
		ret.EntID = info.EntID
		assert.Equal(t, ret, info)
	}
}

func updateRootUser(t *testing.T) {
	pooltype := basetypes.MiningpoolType_AntPool
	ret.MiningpoolTypeStr = pooltype.String()
	ret.MiningpoolType = pooltype
	req.ID = &ret.ID
	req.MiningpoolType = &pooltype

	info, err := UpdateRootUser(context.Background(), req)
	if assert.Nil(t, err) {
		ret.UpdatedAt = info.UpdatedAt
		assert.Equal(t, info, ret)
	}

	pooltype = basetypes.MiningpoolType_F2Pool
	ret.MiningpoolType = pooltype
	ret.MiningpoolTypeStr = pooltype.String()
	req.MiningpoolType = &pooltype
	info, err = UpdateRootUser(context.Background(), req)
	if assert.Nil(t, err) {
		ret.UpdatedAt = info.UpdatedAt
		assert.Equal(t, info, ret)
	}
}

func getRootUser(t *testing.T) {
	info, err := GetRootUser(context.Background(), ret.EntID)
	if assert.Nil(t, err) {
		assert.Equal(t, info, ret)
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

	t.Run("createRootUser", createRootUser)
	t.Run("updateRootUser", updateRootUser)
	t.Run("getRootUser", getRootUser)
	t.Run("getRootUsers", getRootUsers)
}
