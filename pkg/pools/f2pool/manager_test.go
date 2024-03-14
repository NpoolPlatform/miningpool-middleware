package f2pool

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"

	v1 "github.com/NpoolPlatform/message/npool/basetypes/miningpool/v1"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

var (
	accessToken = "7ecdq1fosdsfcruypom2otsn8hfr69azmqvh7v3zelol1ntsba85a1yvol66qp73"
)

func init() {
	//nolint
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}
}

func TestAuth(t *testing.T) {
	mgr, err := NewF2PoolManager(v1.CoinType_BitCoin, accessToken)
	assert.Nil(t, err)
	assert.NotNil(t, mgr)

	authed := mgr.CheckAuth(context.Background())
	assert.Equal(t, nil, authed)
}

//nolint:gocritic
func TestMininguser(t *testing.T) {
	mgr, err := NewF2PoolManager(v1.CoinType_BitCoin, accessToken)
	assert.Nil(t, err)
	assert.NotNil(t, mgr)

	ret, err := mgr.ExistMiningUser(context.Background(), "cococonut1")
	assert.Nil(t, err)
	assert.Equal(t, true, ret)

	ret, err = mgr.ExistMiningUser(context.Background(), uuid.NewString())
	assert.NotNil(t, err)
	assert.Equal(t, false, ret)

	username, link, err := mgr.AddMiningUser(context.Background())
	assert.Nil(t, err)

	ret, err = mgr.ExistMiningUser(context.Background(), username)
	assert.Nil(t, err)
	assert.Equal(t, true, ret)

	_link, err := mgr.GetReadPageLink(context.Background(), username)
	assert.Nil(t, err)
	assert.Equal(t, link, _link)

	// f2pool have not implment the method
	// err = mgr.DeleteMiningUser(context.Background(), username)
	// assert.Nil(t, err)
	// ret, err = mgr.ExistMiningUser(context.Background(), username)
	// assert.NotNil(t, err)
	// assert.Equal(t, false, ret)
	// _, err = mgr.GetReadPageLink(context.Background(), username)
	// assert.NotNil(t, err)
}

func TestRevenueProportion(t *testing.T) {
	mgr, err := NewF2PoolManager(v1.CoinType_BitCoin, accessToken)
	assert.Nil(t, err)
	assert.NotNil(t, mgr)

	propotion := 0.75
	err = mgr.SetRevenueProportion(context.Background(), "cococonut2", "cococonut001", propotion)
	assert.Nil(t, err)

	_propotion, err := mgr.GetRevenueProportion(context.Background(), "cococonut2", "cococonut001")
	assert.Nil(t, err)
	assert.Equal(t, propotion, _propotion)
}

func TestRevenueAddress(t *testing.T) {
	mgr, err := NewF2PoolManager(v1.CoinType_BitCoin, accessToken)
	assert.Nil(t, err)
	assert.NotNil(t, mgr)

	addr := "1PWMfNSb3auXwZ1qhu96WRJL7BCgG4mGB4"
	err = mgr.SetRevenueAddress(context.Background(), "cococonut2", addr)
	fmt.Println(err)
	assert.Nil(t, err)

	_addr, err := mgr.GetRevenueAddress(context.Background(), "cococonut2")
	assert.Nil(t, err)
	assert.Equal(t, addr, _addr)
}

//nolint:gocritic
func TestPageLink(t *testing.T) {
	mgr, err := NewF2PoolManager(v1.CoinType_BitCoin, accessToken)
	assert.Nil(t, err)
	assert.NotNil(t, mgr)

	// f2pool api have bug,
	// link, err := mgr.AddReadPageLink(context.Background(), "cococonut2")
	// assert.Nil(t, err)

	// _link, err := mgr.GetReadPageLink(context.Background(), "cococonut2")
	// assert.Nil(t, err)
	// assert.Equal(t, link, _link)

	err = mgr.DeleteReadPageLink(context.Background(), "cococonut2")
	assert.Nil(t, err)

	// _link, err = mgr.GetReadPageLink(context.Background(), "cococonut2")
	// assert.Nil(t, err)
	// assert.NotEqual(t, link, _link)
}

func TestPayment(t *testing.T) {
	mgr, err := NewF2PoolManager(v1.CoinType_BitCoin, accessToken)
	assert.Nil(t, err)
	assert.NotNil(t, mgr)

	paused, err := mgr.PausePayment(context.Background(), "cococonut2")
	assert.Nil(t, err)
	assert.Equal(t, true, paused)

	resumed, err := mgr.ResumePayment(context.Background(), "cococonut2")
	assert.Nil(t, err)
	assert.Equal(t, true, resumed)

	// _link, err = mgr.GetReadPageLink(context.Background(), "cococonut2")
	// assert.Nil(t, err)
	// assert.NotEqual(t, link, _link)
}

func TestWithdrawPraction(t *testing.T) {
	mgr, err := NewF2PoolManager(v1.CoinType_BitCoin, accessToken)
	assert.Nil(t, err)
	assert.NotNil(t, mgr)

	tx_time, err := mgr.WithdrawPraction(context.Background(), "cococonut1")
	fmt.Println(tx_time)
	assert.Nil(t, err)
	assert.NotEqual(t, 0, tx_time)

	// _link, err = mgr.GetReadPageLink(context.Background(), "cococonut2")
	// assert.Nil(t, err)
	// assert.NotEqual(t, link, _link)
}
