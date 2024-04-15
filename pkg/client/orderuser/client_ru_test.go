package orderuser

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"

	rootuserclient "github.com/NpoolPlatform/miningpool-middleware/pkg/client/rootuser"
	"github.com/NpoolPlatform/miningpool-middleware/pkg/testinit"
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

var rootUserRet = &npool.RootUser{
	EntID:          uuid.NewString(),
	Name:           "mininpool_middleware_client_rootuser_test",
	MiningpoolType: basetypes.MiningpoolType_F2Pool,
	Email:          "sssss",
	AuthToken:      "7ecdq1fosdsfcruypom2otsn8hfr69azmqvh7v3zelol1ntsba85a1yvol66qp73",
	Authed:         true,
	Remark:         "sdfasdfasdf",
}

var rootUserReq = &npool.RootUserReq{
	EntID:          &rootUserRet.EntID,
	Name:           &rootUserRet.Name,
	MiningpoolType: &rootUserRet.MiningpoolType,
	Email:          &rootUserRet.Email,
	AuthToken:      &rootUserRet.AuthToken,
	Authed:         &rootUserRet.Authed,
	Remark:         &rootUserRet.Remark,
}

func createRootUser(t *testing.T) {
	_, err := rootuserclient.CreateRootUser(context.Background(), rootUserReq)
	assert.Nil(t, err)

	info, err := rootuserclient.GetRootUser(context.Background(), *rootUserReq.EntID)
	if assert.Nil(t, err) {
		rootUserRet.CreatedAt = info.CreatedAt
		rootUserRet.MiningpoolTypeStr = info.MiningpoolTypeStr
		rootUserRet.UpdatedAt = info.UpdatedAt
		rootUserRet.ID = info.ID
		rootUserRet.EntID = info.EntID
		assert.Equal(t, rootUserRet, info)
	}
}

func deleteRootUser(t *testing.T) {
	exist, err := rootuserclient.ExistRootUserConds(context.Background(), &npool.Conds{
		EntID: &v1.StringVal{
			Op:    cruder.EQ,
			Value: rootUserRet.EntID,
		},
	})
	if assert.Nil(t, err) {
		assert.Equal(t, true, exist)
	}

	info, err := rootuserclient.DeleteRootUser(context.Background(), rootUserRet.ID)
	if assert.Nil(t, err) {
		assert.Equal(t, rootUserRet, info)
	}

	exist, err = rootuserclient.ExistRootUserConds(context.Background(), &npool.Conds{
		EntID: &v1.StringVal{
			Op:    cruder.EQ,
			Value: rootUserRet.EntID,
		},
	})
	if assert.Nil(t, err) {
		assert.Equal(t, false, exist)
	}
}
