package f2pool

import (
	"context"
	"crypto/rand"
	"fmt"
	"math"
	"math/big"
	"strings"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	"github.com/NpoolPlatform/go-service-framework/pkg/wlog"
	miningpoolbase "github.com/NpoolPlatform/message/npool/basetypes/miningpool/v1"
	basetype "github.com/NpoolPlatform/message/npool/basetypes/v1"
	"github.com/NpoolPlatform/miningpool-middleware/pkg/config"
	"github.com/NpoolPlatform/miningpool-middleware/pkg/const/time"
	"github.com/NpoolPlatform/miningpool-middleware/pkg/pools/f2pool/client"
	"github.com/NpoolPlatform/miningpool-middleware/pkg/pools/f2pool/types"
	"github.com/mr-tron/base58"
	"github.com/shopspring/decimal"
)

type Manager struct {
	currency            string
	authToken           string
	cli                 *client.Client
	leastTransferAmount float64
}

var (
	MiningPoolType                                 = miningpoolbase.MiningpoolType_F2Pool
	CoinType2Currency map[basetype.CoinType]string = map[basetype.CoinType]string{
		basetype.CoinType_CoinTypeBitCoin: "bitcoin",
	}
	CoinType2LeastTransferAmount map[basetype.CoinType]string = map[basetype.CoinType]string{
		basetype.CoinType_CoinTypeBitCoin: "0.005",
	}
	CoinType2FeeRatio map[basetype.CoinType]string = map[basetype.CoinType]string{
		basetype.CoinType_CoinTypeBitCoin: "0.04",
	}
	CoinType2BenefitIntervalSeconds map[basetype.CoinType]uint32 = map[basetype.CoinType]uint32{
		basetype.CoinType_CoinTypeBitCoin: time.SecondsPerDay,
	}
	MaxRetries    = 10
	MaxProportion = decimal.NewFromFloat(100)
	MinProportion = decimal.NewFromFloat(0.00)
	ProportionExp = int32(2)
)

const (
	// 2-15 characters
	MiningUserLen         = 15
	MiningUserPageNameLen = 20
	DefaultPermissions    = "1,2"
)

func NewF2PoolManager(coinType basetype.CoinType, auth string) (*Manager, error) {
	currency := ""
	if k, ok := CoinType2Currency[coinType]; ok {
		currency = k
	} else {
		return nil, wlog.Errorf("have no pool manager for %v-%v", MiningPoolType, coinType)
	}
	mgr := &Manager{
		currency:  currency,
		authToken: auth,
		cli:       client.NewClient(config.F2PoolAPI, auth),
	}
	return mgr, nil
}

func (mgr *Manager) CheckAuth(ctx context.Context) error {
	_, err := mgr.cli.MiningUserList(ctx, &types.MiningUserListReq{})
	return err
}

func (mgr *Manager) AddMiningUser(ctx context.Context) (userName, readPageLink string, err error) {
	var resp *types.MiningUserAddResp
	for i := 0; i < MaxRetries; i++ {
		userName, err := RandomF2PoolUser(MiningUserLen)
		if err != nil {
			return "", "", wlog.WrapError(err)
		}
		resp, err = mgr.cli.MiningUserAdd(ctx, &types.MiningUserAddReq{MiningUserName: userName})
		if err != nil && !strings.Contains(err.Error(), "mining user name already exists") {
			return "", "", wlog.WrapError(err)
		}

		if err == nil {
			break
		}
	}

	if resp == nil {
		return "", "", wlog.Errorf("failed to add mining user,have nil response")
	}

	userName = resp.MiningUserName
	readPageLink = ""
	if len(resp.Pages) > 0 {
		readPageLink = getReadPageLink(resp.Pages[0].Key, resp.MiningUserName)
	}
	return userName, readPageLink, nil
}

// just judge wheather mining user in the root-user of the auth token
func (mgr *Manager) ExistMiningUser(ctx context.Context, name string) (bool, error) {
	_, err := mgr.cli.MiningUserGet(ctx, &types.MiningUserGetReq{MiningUserName: name})
	if err != nil {
		return false, wlog.WrapError(err)
	}
	return true, nil
}

// not implement
func (mgr *Manager) DeleteMiningUser(ctx context.Context, name string) error {
	return wlog.Errorf("f2pool has not yet implemented this method")
}

func (mgr *Manager) AddReadPageLink(ctx context.Context, name string) (string, error) {
	pageName, err := RandomF2PoolUser(MiningUserPageNameLen)
	if err != nil {
		return "", wlog.WrapError(err)
	}

	addResp, err := mgr.cli.MiningUserReadOnlyPageAdd(ctx, &types.MiningUserReadOnlyPageAddReq{
		MiningUserName: name,
		PageName:       pageName,
		Permissions:    DefaultPermissions,
	})
	if err != nil {
		return "", wlog.WrapError(err)
	}

	if addResp == nil {
		return "", wlog.Errorf("failed to add read page link,have nil response")
	}

	return getReadPageLink(addResp.Page.Key, addResp.MiningUserName), nil
}

func (mgr *Manager) GetReadPageLink(ctx context.Context, name string) (string, error) {
	getResp, err := mgr.cli.MiningUserGet(ctx, &types.MiningUserGetReq{MiningUserName: name})
	if err != nil {
		return "", wlog.Errorf("have no user name %v or %v", name, err)
	}

	if getResp == nil {
		return "", wlog.Errorf("have no user name %v", name)
	}

	if len(getResp.Pages) > 0 {
		return getReadPageLink(getResp.Pages[0].Key, getResp.MiningUserName), nil
	}

	return "", wlog.Errorf("have no read page link")
}

// not implement
// f2pool cannot delete page link
func (mgr *Manager) DeleteReadPageLink(ctx context.Context, name string) error {
	getResp, err := mgr.cli.MiningUserGet(ctx, &types.MiningUserGetReq{MiningUserName: name})
	if err != nil {
		return wlog.Errorf("have no user name %v or %v", name, err)
	}

	if getResp == nil {
		return wlog.Errorf("have no user name %v", name)
	}

	if len(getResp.Pages) == 0 {
		return nil
	}

	for _, page := range getResp.Pages {
		_, err = mgr.cli.MiningUserReadOnlyPageDelete(ctx, &types.MiningUserReadOnlyPageDeleteReq{
			MiningUserName: getResp.MiningUserName,
			Key:            page.Key,
		})

		if err != nil {
			return wlog.Errorf("failed to delete read page link of %v, error: %v", name, err)
		}
	}

	return nil
}

//nolint:gocognit
func (mgr *Manager) SetRevenueProportion(ctx context.Context, distributor, recipient, proportion string) error {
	proportionDec, err := decimal.NewFromString(proportion)
	if err != nil {
		return wlog.WrapError(err)
	}
	if proportionDec.Cmp(MinProportion) < 0 || proportionDec.Cmp(MaxProportion) > 0 {
		return wlog.Errorf("wront proportion, please input[%v,%v]", MinProportion, MaxProportion)
	}
	if proportionDec.Truncate(ProportionExp).StringFixed(ProportionExp) != proportionDec.RoundCeil(ProportionExp).StringFixed(ProportionExp) {
		return wlog.Errorf("wront proportion precision, please enter %v decimal places", ProportionExp)
	}
	proportionDec = proportionDec.Truncate(ProportionExp)
	infoResp, err := mgr.cli.RevenueDistributionInfo(ctx, &types.RevenueDistributionInfoReq{
		Currency:    mgr.currency,
		Distributor: distributor,
		Recipient:   recipient,
	})

	if err == nil && infoResp != nil {
		for _, v := range infoResp.Data {
			if v.Distributor != distributor || v.Currency != mgr.currency || v.Recipient != recipient {
				continue
			}
			_, err = mgr.cli.RevenueDistributionDelete(ctx, &types.RevenueDistributionDeleteReq{
				ID:          v.ID,
				Distributor: v.Distributor,
			})
			if err != nil {
				logger.Sugar().Warn(err)
			}
		}
	}

	if proportionDec.Cmp(MinProportion) == 0 {
		return nil
	}

	_proportion := proportionDec.InexactFloat64()
	addResp, err := mgr.cli.RevenueDistributionAdd(ctx, &types.RevenueDistributionAddReq{
		Currency:    mgr.currency,
		Distributor: distributor,
		Recipient:   recipient,
		Proportion:  _proportion,
	})
	if err != nil {
		return wlog.WrapError(err)
	}

	if addResp == nil {
		return wlog.Errorf("failed to add revenue proportion,have nil response")
	}
	return nil
}

func (mgr *Manager) GetRevenueProportion(ctx context.Context, distributor, recipient string) (float64, error) {
	getResp, err := mgr.cli.RevenueDistributionInfo(ctx, &types.RevenueDistributionInfoReq{
		Distributor: distributor,
		Recipient:   recipient,
		Currency:    mgr.currency,
	})
	if err != nil {
		return 0, wlog.WrapError(err)
	}

	if getResp == nil {
		return 0, wlog.Errorf("failed to get revenue proportion info,have nil response")
	}

	for _, v := range getResp.Data {
		if v.Currency == mgr.currency && v.Distributor == distributor && v.Recipient == recipient {
			return v.Proportion, nil
		}
	}

	return 0, nil
}

func (mgr *Manager) SetRevenueAddress(ctx context.Context, name, address string) error {
	setResp, err := mgr.cli.MiningUserWalletUpdate(ctx, &types.MiningUserWalletUpdateReq{
		Params: []types.WalletParams{
			{
				MiningUserName: name,
				Wallets: []types.Wallet{
					{
						Currency:            mgr.currency,
						Address:             address,
						LeastTransferAmount: mgr.leastTransferAmount,
					},
				},
			},
		},
	})
	if err != nil {
		return wlog.WrapError(err)
	}

	if setResp == nil {
		return wlog.Errorf("failed to set revenue address,have nil response")
	}
	return nil
}

func (mgr *Manager) GetRevenueAddress(ctx context.Context, name string) (string, error) {
	getResp, err := mgr.cli.MiningUserGet(ctx, &types.MiningUserGetReq{
		MiningUserName: name,
	})
	if err != nil {
		return "", wlog.WrapError(err)
	}

	if getResp == nil {
		return "", wlog.Errorf("failed to get revenue address,have nil response")
	}

	for _, v := range getResp.Wallets {
		if v.Currency == mgr.currency {
			return v.Address, nil
		}
	}

	return "", nil
}

func (mgr *Manager) PausePayment(ctx context.Context, name string) (bool, error) {
	pauseResp, err := mgr.cli.MiningUserPaymentPause(ctx, &types.MiningUserPaymentPauseReq{
		MiningUserNames: []string{name},
		Currency:        mgr.currency,
	})
	if err != nil {
		return false, wlog.WrapError(err)
	}

	if pauseResp == nil {
		return false, wlog.Errorf("failed to pause payment,have nil response")
	}

	for k, v := range pauseResp.Results {
		if k == name && v == "ok" {
			return true, nil
		}
	}

	return false, nil
}

func (mgr *Manager) ResumePayment(ctx context.Context, name string) (bool, error) {
	resumeResp, err := mgr.cli.MiningUserPaymentResume(ctx, &types.MiningUserPaymentResumeReq{
		MiningUserNames: []string{name},
		Currency:        mgr.currency,
	})
	if err != nil {
		return false, wlog.WrapError(err)
	}

	if resumeResp == nil {
		return false, wlog.Errorf("failed to resume payment,have nil response")
	}

	for k, v := range resumeResp.Results {
		if k == name && v == "ok" {
			return true, nil
		}
	}

	return false, nil
}

func (mgr *Manager) WithdrawPraction(ctx context.Context, name string) (int64, error) {
	resumeResp, err := mgr.cli.MiningUserBalanceWithdraw(ctx, &types.MiningUserBalanceWithdrawReq{
		MiningUserName: name,
		Currency:       mgr.currency,
	})
	if err != nil {
		return 0, wlog.WrapError(err)
	}

	if resumeResp == nil {
		return 0, wlog.Errorf("failed to resume payment,have nil response")
	}

	return resumeResp.PaidTime, nil
}

func getReadPageLink(key, userName string) string {
	return fmt.Sprintf("%v/mining-user/%v?user_name=%v", config.F2PoolBaseURL, key, userName)
}

// can only be a combination of lowercase characters and numbers
// start with letter
func RandomF2PoolUser(n int) (string, error) {
	startLetters := []rune("abcdefghijklmnopqretuvwxyz")
	randn, err := rand.Int(rand.Reader, big.NewInt(int64(len(startLetters))))
	if err != nil {
		return "", wlog.WrapError(err)
	}
	target := string(startLetters[randn.Int64()])
	for {
		randn, err := rand.Int(rand.Reader, big.NewInt(math.MaxInt64))
		if err != nil {
			return "", wlog.WrapError(err)
		}
		if len(target) >= n {
			return strings.ToLower(target[:n]), nil
		}
		target += base58.Encode(randn.Bytes())
	}
}
