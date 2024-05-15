package fractionrule

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"

	npool "github.com/NpoolPlatform/message/npool/miningpool/mw/v1/fractionrule"
	testinit "github.com/NpoolPlatform/miningpool-middleware/pkg/testinit"
	"github.com/google/uuid"

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

var ret = &npool.FractionRule{
	EntID:            uuid.NewString(),
	PoolCoinTypeID:   uuid.NewString(),
	WithdrawInterval: 1000,
	MinAmount:        "2.0",
	WithdrawRate:     "3.0",
}

var req = &npool.FractionRuleReq{
	EntID:            &ret.EntID,
	PoolCoinTypeID:   &ret.PoolCoinTypeID,
	WithdrawInterval: &ret.WithdrawInterval,
	MinAmount:        &ret.MinAmount,
	WithdrawRate:     &ret.WithdrawRate,
}

func create(t *testing.T) {
	getH, err := NewHandler(context.Background(), WithConds(nil), WithOffset(0), WithLimit(2))
	assert.Nil(t, err)

	infos, _, err := getH.GetFractionRules(context.Background())
	assert.Nil(t, err)
	if len(infos) == 0 {
		return
	}

	ret.PoolCoinTypeID = infos[0].PoolCoinTypeID
	req.PoolCoinTypeID = &infos[0].PoolCoinTypeID
	ret.CoinType = infos[0].CoinType
	ret.EntID = infos[0].EntID
	req.EntID = &infos[0].EntID

	handler, err := NewHandler(
		context.Background(),
		WithEntID(req.EntID, true),
		WithPoolCoinTypeID(req.PoolCoinTypeID, true),
		WithWithdrawInterval(req.WithdrawInterval, true),
		WithMinAmount(req.MinAmount, true),
		WithWithdrawRate(req.WithdrawRate, true),
	)
	assert.Nil(t, err)

	err = handler.CreateFractionRule(context.Background())
	assert.NotNil(t, err)
}

func TestFractionRule(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	t.Run("create", create)
}
