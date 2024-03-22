package gooduser

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"

	"github.com/NpoolPlatform/go-service-framework/pkg/config"
	"github.com/NpoolPlatform/miningpool-middleware/pkg/testinit"
	"github.com/google/uuid"

	"bou.ke/monkey"
	grpc2 "github.com/NpoolPlatform/go-service-framework/pkg/grpc"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	cruder "github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/miningpool/v1"
	v1 "github.com/NpoolPlatform/message/npool/basetypes/v1"
	npool "github.com/NpoolPlatform/message/npool/miningpool/mw/v1/gooduser"
)

func init() {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}
	if err := testinit.Init(); err != nil {
		fmt.Printf("cannot init test stub: %v\n", err)
	}
}

var ret = &npool.GoodUser{
	EntID:          uuid.NewString(),
	Name:           "fffff",
	RootUserID:     rootUserRet.EntID,
	GoodID:         uuid.NewString(),
	MiningpoolType: basetypes.MiningpoolType_F2Pool,
	CoinType:       basetypes.CoinType_BitCoin,
	HashRate:       66,
	ReadPageLink:   "fffff",
	RevenueType:    basetypes.RevenueType_FPPS,
}

var req = &npool.GoodUserReq{
	EntID:          &ret.EntID,
	Name:           &ret.Name,
	RootUserID:     &ret.RootUserID,
	GoodID:         &ret.GoodID,
	MiningpoolType: &ret.MiningpoolType,
	CoinType:       &ret.CoinType,
	HashRate:       &ret.HashRate,
	ReadPageLink:   &ret.ReadPageLink,
	RevenueType:    &ret.RevenueType,
}

func createGoodUser(t *testing.T) {
	info, err := CreateGoodUser(context.Background(), req)
	if assert.Nil(t, err) {
		ret.Name = info.Name
		ret.ReadPageLink = info.ReadPageLink
		ret.CreatedAt = info.CreatedAt
		ret.MiningpoolTypeStr = info.MiningpoolTypeStr
		ret.CoinTypeStr = info.CoinTypeStr
		ret.RevenueTypeStr = info.RevenueTypeStr
		ret.UpdatedAt = info.UpdatedAt
		ret.ID = info.ID
		ret.EntID = info.EntID
		assert.Equal(t, ret, info)
	}
}

func updateGoodUser(t *testing.T) {
	ret.HashRate = 77
	req.ID = &ret.ID
	req.HashRate = &ret.HashRate
	info, err := UpdateGoodUser(context.Background(), req)
	if assert.Nil(t, err) {
		ret.UpdatedAt = info.UpdatedAt
		assert.Equal(t, info, ret)
	}

	ret.HashRate = 88
	req.ID = &ret.ID
	req.HashRate = &ret.HashRate
	info, err = UpdateGoodUser(context.Background(), req)
	if assert.Nil(t, err) {
		ret.UpdatedAt = info.UpdatedAt
		assert.Equal(t, info, ret)
	}
}

func getGoodUser(t *testing.T) {
	info, err := GetGoodUser(context.Background(), ret.EntID)
	if assert.Nil(t, err) {
		assert.Equal(t, info, ret)
	}
}

func getGoodUsers(t *testing.T) {
	infos, total, err := GetGoodUsers(context.Background(), &npool.Conds{
		EntID: &v1.StringVal{Op: cruder.EQ, Value: ret.EntID},
	}, 0, 1)
	if assert.Nil(t, err) {
		assert.Equal(t, len(infos), 1)
		assert.Equal(t, total, uint32(1))
		assert.Equal(t, infos[0], ret)
	}
}

func deleteGoodUser(t *testing.T) {
	exist, err := ExistGoodUserConds(context.Background(), &npool.Conds{
		EntID: &v1.StringVal{
			Op:    cruder.EQ,
			Value: ret.EntID,
		},
	})
	if assert.Nil(t, err) {
		assert.Equal(t, true, exist)
	}

	info, err := DeleteGoodUser(context.Background(), &npool.GoodUserReq{
		ID: &ret.ID,
	})
	if assert.Nil(t, err) {
		assert.Equal(t, info, ret)
	}

	exist, err = ExistGoodUserConds(context.Background(), &npool.Conds{
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

	t.Run("createGoodUser", createRootUser)
	t.Run("createGoodUser", createGoodUser)
	t.Run("updateGoodUser", updateGoodUser)
	t.Run("getGoodUser", getGoodUser)
	t.Run("getGoodUsers", getGoodUsers)
	t.Run("deleteGoodUser", deleteGoodUser)
	t.Run("deleteGoodUser", deleteRootUser)
}