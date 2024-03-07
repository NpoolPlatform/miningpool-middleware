package f2pool

import (
	"context"
	"fmt"
	"testing"
	"time"

	v1 "github.com/NpoolPlatform/message/npool/basetypes/miningpool/v1"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

var (
	accessToken = "7ecdq1fosdsfcruypom2otsn8hfr69azmqvh7v3zelol1ntsba85a1yvol66qp73"
)

func TestAuth(t *testing.T) {
	mgr, err := NewF2PoolManager(v1.CoinType_BitCoin, accessToken)
	assert.Nil(t, err)
	assert.NotNil(t, mgr)

	authed := mgr.CheckAuth(context.Background())
	assert.Equal(t, nil, authed)
}

func TestMininguser(t *testing.T) {
	mgr, err := NewF2PoolManager(v1.CoinType_BitCoin, accessToken)
	assert.Nil(t, err)
	assert.NotNil(t, mgr)

	ret, err := mgr.ExistMiningUser(context.Background(), "cococonut1")
	assert.Nil(t, err)
	assert.Equal(t, true, ret)

	time.Sleep(time.Second)

	ret, err = mgr.ExistMiningUser(context.Background(), uuid.NewString())
	assert.NotNil(t, err)
	assert.Equal(t, false, ret)
	time.Sleep(time.Second)

	username, link, err := mgr.AddMiningUser(context.Background())
	assert.Nil(t, err)
	time.Sleep(time.Second)

	ret, err = mgr.ExistMiningUser(context.Background(), username)
	assert.Nil(t, err)
	assert.Equal(t, true, ret)
	time.Sleep(time.Second)

	_link, err := mgr.GetReadPageLink(context.Background(), username)
	assert.Nil(t, err)
	assert.Equal(t, link, _link)
	time.Sleep(time.Second)

	// f2pool have not implment the method
	// err = mgr.DeleteMiningUser(context.Background(), username)
	// assert.Nil(t, err)
	// time.Sleep(time.Second)

	// ret, err = mgr.ExistMiningUser(context.Background(), username)
	// assert.NotNil(t, err)
	// assert.Equal(t, false, ret)
	// time.Sleep(time.Second)

	// _, err = mgr.GetReadPageLink(context.Background(), username)
	// assert.NotNil(t, err)
	// time.Sleep(time.Second)
}

func TestRevenueProportion(t *testing.T) {
	mgr, err := NewF2PoolManager(v1.CoinType_BitCoin, accessToken)
	assert.Nil(t, err)
	assert.NotNil(t, mgr)

	propotion := 0.75
	err = mgr.SetRevenueProportion(context.Background(), "cococonut2", "cococonut001", propotion)
	assert.Nil(t, err)
	time.Sleep(time.Second)

	_propotion, err := mgr.GetRevenueProportion(context.Background(), "cococonut2", "cococonut001")
	assert.Nil(t, err)
	assert.Equal(t, propotion, _propotion)
	time.Sleep(time.Second)
}

func TestRevenueAddress(t *testing.T) {
	mgr, err := NewF2PoolManager(v1.CoinType_BitCoin, accessToken)
	assert.Nil(t, err)
	assert.NotNil(t, mgr)

	addr := "1PWMfNSb3auXwZ1qhu96WRJL7BCgG4mGB4"
	err = mgr.SetRevenueAddress(context.Background(), "cococonut2", addr)
	fmt.Println(err)
	assert.Nil(t, err)
	time.Sleep(time.Second)

	_addr, err := mgr.GetRevenueAddress(context.Background(), "cococonut2")
	assert.Nil(t, err)
	assert.Equal(t, addr, _addr)
	time.Sleep(time.Second)
}

func TestPageLink(t *testing.T) {
	mgr, err := NewF2PoolManager(v1.CoinType_BitCoin, accessToken)
	assert.Nil(t, err)
	assert.NotNil(t, mgr)

	// link, err := mgr.AddReadPageLink(context.Background(), "cococonut2")
	// assert.Nil(t, err)
	// time.Sleep(time.Second)

	// _link, err := mgr.GetReadPageLink(context.Background(), "cococonut2")
	// assert.Nil(t, err)
	// assert.Equal(t, link, _link)
	// time.Sleep(time.Second)

	err = mgr.DeleteReadPageLink(context.Background(), "cococonut2")
	assert.Nil(t, err)
	time.Sleep(time.Second)

	// _link, err = mgr.GetReadPageLink(context.Background(), "cococonut2")
	// assert.Nil(t, err)
	// assert.NotEqual(t, link, _link)
	// time.Sleep(time.Second)
}

func TestPayment(t *testing.T) {
	mgr, err := NewF2PoolManager(v1.CoinType_BitCoin, accessToken)
	assert.Nil(t, err)
	assert.NotNil(t, mgr)

	paused, err := mgr.PausePayment(context.Background(), "cococonut2")
	assert.Nil(t, err)
	assert.Equal(t, true, paused)
	time.Sleep(time.Second)

	resumed, err := mgr.ResumePayment(context.Background(), "cococonut2")
	assert.Nil(t, err)
	assert.Equal(t, true, resumed)
	time.Sleep(time.Second)

	// _link, err = mgr.GetReadPageLink(context.Background(), "cococonut2")
	// assert.Nil(t, err)
	// assert.NotEqual(t, link, _link)
	// time.Sleep(time.Second)
}

func TestWithdrawPraction(t *testing.T) {
	time.Sleep(time.Second)
	mgr, err := NewF2PoolManager(v1.CoinType_BitCoin, accessToken)
	assert.Nil(t, err)
	assert.NotNil(t, mgr)

	tx_time, err := mgr.WithdrawPraction(context.Background(), "cococonut1")
	fmt.Println(tx_time)
	assert.Nil(t, err)
	assert.NotEqual(t, 0, tx_time)
	time.Sleep(time.Second)

	// _link, err = mgr.GetReadPageLink(context.Background(), "cococonut2")
	// assert.Nil(t, err)
	// assert.NotEqual(t, link, _link)
	// time.Sleep(time.Second)
}
