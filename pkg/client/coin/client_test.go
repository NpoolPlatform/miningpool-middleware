package coin

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"

	"github.com/NpoolPlatform/miningpool-middleware/pkg/testinit"

	"github.com/stretchr/testify/assert"

	basetypes "github.com/NpoolPlatform/message/npool/basetypes/miningpool/v1"
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
	FeeRate:          "2.0",
	FixedRevenueAble: false,
	Threshold:        "3.0",
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
	infos, _, err := GetCoins(context.Background(), &npool.Conds{}, 0, 2)
	assert.Nil(t, err)
	if len(infos) == 0 {
		return
	}

	ret.MiningpoolType = infos[0].MiningpoolType
	req.MiningpoolType = &infos[0].MiningpoolType
	ret.CoinType = infos[0].CoinType
	req.CoinType = &infos[0].CoinType

	_, err = CreateCoin(context.Background(), req)
	assert.NotNil(t, err)

	info, err := GetCoin(context.Background(), infos[0].EntID)
	if assert.Nil(t, err) {
		assert.Equal(t, infos[0], info)
	}
}

func TestClient(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}
	t.Run("createCoin", createCoin)
}
