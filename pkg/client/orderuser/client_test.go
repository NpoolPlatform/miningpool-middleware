package orderuser

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
	npool "github.com/NpoolPlatform/message/npool/miningpool/mw/v1/orderuser"
)

func init() {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}
	if err := testinit.Init(); err != nil {
		fmt.Printf("cannot init test stub: %v\n", err)
	}
}

var ret = &npool.OrderUser{
	EntID:          uuid.NewString(),
	Name:           "fffff",
	RootUserID:     uuid.NewString(),
	GoodUserID:     uuid.NewString(),
	OrderID:        uuid.NewString(),
	MiningpoolType: basetypes.MiningpoolType_AntPool,
	CoinType:       basetypes.CoinType_BitCoin,
	Proportion:     66,
	RevenueAddress: "fffff",
	ReadPageLink:   "fffff",
	AutoPay:        false,
}

var req = &npool.OrderUserReq{
	EntID:          &ret.EntID,
	Name:           &ret.Name,
	RootUserID:     &ret.RootUserID,
	GoodUserID:     &ret.GoodUserID,
	OrderID:        &ret.OrderID,
	MiningpoolType: &ret.MiningpoolType,
	CoinType:       &ret.CoinType,
	Proportion:     &ret.Proportion,
	RevenueAddress: &ret.RevenueAddress,
	ReadPageLink:   &ret.ReadPageLink,
	AutoPay:        &ret.AutoPay,
}

func createOrderUser(t *testing.T) {
	info, err := CreateOrderUser(context.Background(), req)
	if assert.Nil(t, err) {
		ret.CreatedAt = info.CreatedAt
		ret.MiningpoolTypeStr = info.MiningpoolTypeStr
		ret.CoinTypeStr = info.CoinTypeStr
		ret.UpdatedAt = info.UpdatedAt
		ret.ID = info.ID
		ret.EntID = info.EntID
		assert.Equal(t, ret, info)
	}
}

func updateOrderUser(t *testing.T) {
	mtype := basetypes.MiningpoolType_AntPool
	ret.MiningpoolTypeStr = mtype.String()
	ret.MiningpoolType = mtype
	req.ID = &ret.ID
	req.MiningpoolType = &mtype

	info, err := UpdateOrderUser(context.Background(), req)
	if assert.Nil(t, err) {
		ret.UpdatedAt = info.UpdatedAt
		assert.Equal(t, info, ret)
	}

	mtype = basetypes.MiningpoolType_F2Pool
	ret.MiningpoolType = mtype
	ret.MiningpoolTypeStr = mtype.String()
	req.MiningpoolType = &mtype
	info, err = UpdateOrderUser(context.Background(), req)
	if assert.Nil(t, err) {
		ret.UpdatedAt = info.UpdatedAt
		assert.Equal(t, info, ret)
	}
}

func getOrderUser(t *testing.T) {
	info, err := GetOrderUser(context.Background(), ret.EntID)
	if assert.Nil(t, err) {
		assert.Equal(t, info, ret)
	}
}

func getOrderUsers(t *testing.T) {
	infos, total, err := GetOrderUsers(context.Background(), &npool.Conds{
		EntID: &v1.StringVal{Op: cruder.EQ, Value: ret.EntID},
	}, 0, 1)
	if assert.Nil(t, err) {
		assert.Equal(t, len(infos), 1)
		assert.Equal(t, total, uint32(1))
		assert.Equal(t, infos[0], ret)
	}
}

func deleteOrderUser(t *testing.T) {
	exist, err := ExistOrderUserConds(context.Background(), &npool.Conds{
		EntID: &v1.StringVal{
			Op:    cruder.EQ,
			Value: ret.EntID,
		},
	})
	if assert.Nil(t, err) {
		assert.Equal(t, true, exist)
	}

	info, err := DeleteOrderUser(context.Background(), &npool.OrderUserReq{
		ID: &ret.ID,
	})
	if assert.Nil(t, err) {
		assert.Equal(t, info, ret)
	}

	exist, err = ExistOrderUserConds(context.Background(), &npool.Conds{
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

	t.Run("createOrderUser", createOrderUser)
	t.Run("updateOrderUser", updateOrderUser)
	t.Run("getOrderUser", getOrderUser)
	t.Run("getOrderUsers", getOrderUsers)
	t.Run("deleteOrderUser", deleteOrderUser)
}
