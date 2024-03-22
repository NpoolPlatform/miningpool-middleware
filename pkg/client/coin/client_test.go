package coin

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
	npool "github.com/NpoolPlatform/message/npool/miningpool/mw/v1/coin"

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

var ret = &npool.Coin{
	EntID:            uuid.NewString(),
	MiningpoolType:   basetypes.MiningpoolType_F2Pool,
	CoinType:         basetypes.CoinType_BitCoin,
	RevenueTypes:     []basetypes.RevenueType{basetypes.RevenueType_FPPS, basetypes.RevenueType_PPLNS},
	FeeRate:          2.0,
	FixedRevenueAble: false,
	Threshold:        3.0,
	Remark:           "sdfasdfasdf",
}

var req = &npool.CoinReq{
	EntID:            &ret.EntID,
	MiningpoolType:   &ret.MiningpoolType,
	CoinType:         &ret.CoinType,
	RevenueTypes:     ret.RevenueTypes,
	FeeRate:          &ret.FeeRate,
	FixedRevenueAble: &ret.FixedRevenueAble,
	Threshold:        &ret.Threshold,
	Remark:           &ret.Remark,
}

func createCoin(t *testing.T) {
	info, err := CreateCoin(context.Background(), req)
	if assert.Nil(t, err) {
		ret.CreatedAt = info.CreatedAt
		ret.MiningpoolTypeStr = info.MiningpoolTypeStr
		ret.CoinTypeStr = info.CoinTypeStr
		ret.RevenueTypesStr = info.RevenueTypesStr
		ret.UpdatedAt = info.UpdatedAt
		ret.ID = info.ID
		ret.EntID = info.EntID
		assert.Equal(t, ret, info)
	}
}

func updateCoin(t *testing.T) {
	pooltype := basetypes.MiningpoolType_AntPool
	ret.MiningpoolTypeStr = pooltype.String()
	ret.MiningpoolType = pooltype
	req.ID = &ret.ID
	req.MiningpoolType = &pooltype

	info, err := UpdateCoin(context.Background(), req)
	if assert.Nil(t, err) {
		ret.UpdatedAt = info.UpdatedAt
		assert.Equal(t, info, ret)
	}

	pooltype = basetypes.MiningpoolType_F2Pool
	ret.MiningpoolType = pooltype
	ret.MiningpoolTypeStr = pooltype.String()
	req.MiningpoolType = &pooltype
	info, err = UpdateCoin(context.Background(), req)
	if assert.Nil(t, err) {
		ret.UpdatedAt = info.UpdatedAt
		assert.Equal(t, info, ret)
	}
}

func getCoin(t *testing.T) {
	info, err := GetCoin(context.Background(), ret.EntID)
	if assert.Nil(t, err) {
		assert.Equal(t, info, ret)
	}
}

func getCoins(t *testing.T) {
	infos, total, err := GetCoins(context.Background(), &npool.Conds{
		EntID: &v1.StringVal{Op: cruder.EQ, Value: ret.EntID},
	}, 0, 1)
	if assert.Nil(t, err) {
		assert.Equal(t, len(infos), 1)
		assert.Equal(t, total, uint32(1))
		assert.Equal(t, infos[0], ret)
	}
}

func deleteCoin(t *testing.T) {
	exist, err := ExistCoinConds(context.Background(), &npool.Conds{
		EntID: &v1.StringVal{
			Op:    cruder.EQ,
			Value: ret.EntID,
		},
	})
	if assert.Nil(t, err) {
		assert.Equal(t, true, exist)
	}

	info, err := DeleteCoin(context.Background(), &npool.CoinReq{
		ID: &ret.ID,
	})
	if assert.Nil(t, err) {
		assert.Equal(t, info, ret)
	}

	exist, err = ExistCoinConds(context.Background(), &npool.Conds{
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

	t.Run("createCoin", createCoin)
	t.Run("updateCoin", updateCoin)
	t.Run("getCoin", getCoin)
	t.Run("getCoins", getCoins)
	t.Run("deleteCoin", deleteCoin)
}
