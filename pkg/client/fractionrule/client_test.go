package fractionrule

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"

	"github.com/NpoolPlatform/miningpool-middleware/pkg/testinit"

	"github.com/stretchr/testify/assert"

	basetypes "github.com/NpoolPlatform/message/npool/basetypes/miningpool/v1"
	npool "github.com/NpoolPlatform/message/npool/miningpool/mw/v1/fractionrule"

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

var ret = &npool.FractionRule{
	EntID:            uuid.NewString(),
	MiningpoolType:   basetypes.MiningpoolType_F2Pool,
	CoinType:         basetypes.CoinType_BitCoin,
	WithdrawInterval: 3,
	MinAmount:        "2.0",
	WithdrawRate:     "2.0",
}

var req = &npool.FractionRuleReq{
	EntID:            &ret.EntID,
	WithdrawInterval: &ret.WithdrawInterval,
	MinAmount:        &ret.MinAmount,
	WithdrawRate:     &ret.WithdrawRate,
}

func createFractionRule(t *testing.T) {
	infos, _, err := GetFractionRules(context.Background(), &npool.Conds{}, 0, 2)
	assert.Nil(t, err)
	if len(infos) == 0 {
		return
	}

	ret.MiningpoolType = infos[0].MiningpoolType
	ret.CoinType = infos[0].CoinType

	err = CreateFractionRule(context.Background(), req)
	assert.NotNil(t, err)

	info, err := GetFractionRule(context.Background(), infos[0].EntID)
	if assert.Nil(t, err) {
		assert.Equal(t, infos[0], info)
	}
}

func TestClient(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	t.Run("createFractionRule", createFractionRule)
}
