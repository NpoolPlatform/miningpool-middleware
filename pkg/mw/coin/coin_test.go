package coin

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"

	npool "github.com/NpoolPlatform/message/npool/miningpool/mw/v1/coin"
	testinit "github.com/NpoolPlatform/miningpool-middleware/pkg/testinit"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"

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

var ret = &npool.Coin{
	EntID:            uuid.NewString(),
	CoinType:         basetypes.CoinType_BitCoin,
	MiningpoolType:   basetypes.MiningpoolType_AntPool,
	RevenueTypes:     []basetypes.RevenueType{basetypes.RevenueType_FPPS, basetypes.RevenueType_PPLNS},
	FeeRate:          decimal.NewFromFloat(5.4).String(),
	FixedRevenueAble: false,
	Threshold:        decimal.NewFromFloat(5.4).String(),
	Remark:           "asdfaf",
}

var req = &npool.CoinReq{
	EntID:            &ret.EntID,
	CoinType:         &ret.CoinType,
	MiningpoolType:   &ret.MiningpoolType,
	RevenueTypes:     ret.RevenueTypes,
	FeeRate:          &ret.FeeRate,
	FixedRevenueAble: &ret.FixedRevenueAble,
	Threshold:        &ret.Threshold,
	Remark:           &ret.Remark,
}

func create(t *testing.T) {
	getH, err := NewHandler(context.Background(), WithConds(nil), WithOffset(0), WithLimit(2))
	assert.Nil(t, err)

	infos, _, err := getH.GetCoins(context.Background())
	assert.Nil(t, err)

	if len(infos) == 0 {
		return
	}

	ret.MiningpoolType = infos[0].MiningpoolType
	req.MiningpoolType = &infos[0].MiningpoolType
	ret.CoinType = infos[0].CoinType
	req.CoinType = &infos[0].CoinType

	handler, err := NewHandler(
		context.Background(),
		WithEntID(req.EntID, true),
		WithCoinType(req.CoinType, true),
		WithMiningpoolType(req.MiningpoolType, true),
		WithRevenueTypes(&req.RevenueTypes, true),
		WithFeeRate(req.FeeRate, true),
		WithFixedRevenueAble(req.FixedRevenueAble, true),
		WithThreshold(req.Threshold, true),
		WithRemark(req.Remark, true),
	)
	assert.Nil(t, err)

	err = handler.CreateCoin(context.Background())
	assert.NotNil(t, err)
}

func TestCoin(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	t.Run("create", create)
}
