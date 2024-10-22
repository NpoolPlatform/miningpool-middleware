package fractionwithdrawalrule

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"

	"github.com/NpoolPlatform/miningpool-middleware/pkg/pools/registetestinfo"
	"github.com/NpoolPlatform/miningpool-middleware/pkg/testinit"

	"github.com/stretchr/testify/assert"

	mpbasetypes "github.com/NpoolPlatform/message/npool/basetypes/miningpool/v1"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	npool "github.com/NpoolPlatform/message/npool/miningpool/mw/v1/fractionwithdrawalrule"

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

var ret = &npool.FractionWithdrawalRule{
	EntID:                 uuid.NewString(),
	MiningPoolType:        mpbasetypes.MiningPoolType_F2Pool,
	CoinType:              basetypes.CoinType_CoinTypeBitCoin,
	WithdrawInterval:      3,
	PayoutThreshold:       "5.0",
	LeastWithdrawalAmount: "2.0",
	WithdrawFee:           "2.0",
}

var req = &npool.FractionWithdrawalRuleReq{
	EntID:                 &ret.EntID,
	WithdrawInterval:      &ret.WithdrawInterval,
	PayoutThreshold:       &ret.PayoutThreshold,
	LeastWithdrawalAmount: &ret.LeastWithdrawalAmount,
	WithdrawFee:           &ret.WithdrawFee,
}

func createFractionWithdrawalRule(t *testing.T) {
	infos, _, err := GetFractionWithdrawalRules(context.Background(), &npool.Conds{}, 0, 2)
	assert.Nil(t, err)
	if len(infos) == 0 {
		return
	}

	ret.MiningPoolType = infos[0].MiningPoolType
	ret.CoinType = infos[0].CoinType

	err = CreateFractionWithdrawalRule(context.Background(), req)
	assert.NotNil(t, err)

	info, err := GetFractionWithdrawalRule(context.Background(), infos[0].EntID)
	if assert.Nil(t, err) {
		assert.Equal(t, infos[0], info)
	}
}

func TestClient(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}
	registetestinfo.InitTestInfo(context.Background())
	t.Run("createFractionWithdrawalRule", createFractionWithdrawalRule)
	registetestinfo.CleanTestInfo(context.Background())
}
