package f2pool

import (
	"context"
	"crypto/rand"
	"fmt"
	"math"
	"math/big"
	"strings"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	basetype "github.com/NpoolPlatform/message/npool/basetypes/miningpool/v1"
	"github.com/NpoolPlatform/miningpool-middleware/pkg/pools/f2pool/client"
	"github.com/NpoolPlatform/miningpool-middleware/pkg/pools/f2pool/types"
	"github.com/mr-tron/base58"
)

type Manager struct {
	currency  string
	authToken string
	cli       *client.Client
	threshold float64
}

var (
	MiningPoolType                                 = basetype.MiningpoolType_F2Pool
	CoinType2Currency map[basetype.CoinType]string = map[basetype.CoinType]string{
		basetype.CoinType_BitCoin: "bitcoin",
	}
	CoinType2Threshold map[basetype.CoinType]float32 = map[basetype.CoinType]float32{
		basetype.CoinType_BitCoin: 0.005,
	}
	CoinType2FeeRate map[basetype.CoinType]float32 = map[basetype.CoinType]float32{
		basetype.CoinType_BitCoin: 0.04,
	}
	F2PoolAPI     = "https://api.f2pool.com"
	F2PoolBaseURL = "https://f2pool.com"
	MaxRetries    = 10
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
		return nil, fmt.Errorf("have no pool manager for %v-%v", MiningPoolType, coinType)
	}
	mgr := &Manager{
		currency:  currency,
		authToken: auth,
		cli:       client.NewClient(F2PoolAPI, auth),
	}
	if err := mgr.CheckAuth(context.Background()); err != nil {
		return nil, err
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
			return "", "", err
		}
		resp, err = mgr.cli.MiningUserAdd(ctx, &types.MiningUserAddReq{MiningUserName: userName})
		if err != nil && !strings.Contains(err.Error(), "mining user name already exists") {
			return "", "", err
		}

		if err == nil {
			break
		}
	}

	if err != nil {
		return "", "", fmt.Errorf("add mining user retries exhausted,err: %v", err)
	}

	if resp == nil {
		return "", "", fmt.Errorf("failed to add mining user,have nil response")
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
		return false, err
	}
	return true, nil
}

// not implement
func (mgr *Manager) DeleteMiningUser(ctx context.Context, name string) error {
	return fmt.Errorf("f2pool has not yet implemented this method")
}

func (mgr *Manager) AddReadPageLink(ctx context.Context, name string) (string, error) {
	pageName, err := RandomF2PoolUser(MiningUserPageNameLen)
	if err != nil {
		return "", err
	}

	addResp, err := mgr.cli.MiningUserReadOnlyPageAdd(ctx, &types.MiningUserReadOnlyPageAddReq{
		MiningUserName: name,
		PageName:       pageName,
		Permissions:    DefaultPermissions,
	})
	if err != nil {
		return "", err
	}

	if addResp == nil {
		return "", fmt.Errorf("failed to add read page link,have nil response")
	}

	return getReadPageLink(addResp.Page.Key, addResp.MiningUserName), nil
}

func (mgr *Manager) GetReadPageLink(ctx context.Context, name string) (string, error) {
	getResp, err := mgr.cli.MiningUserGet(ctx, &types.MiningUserGetReq{MiningUserName: name})
	if err != nil {
		return "", fmt.Errorf("have no user name %v or %v", name, err)
	}

	if getResp == nil {
		return "", fmt.Errorf("have no user name %v", name)
	}

	if len(getResp.Pages) > 0 {
		return getReadPageLink(getResp.Pages[0].Key, getResp.MiningUserName), nil
	}

	return "", fmt.Errorf("have no read page link")
}

// not implement
// f2pool cannot delete page link
func (mgr *Manager) DeleteReadPageLink(ctx context.Context, name string) error {
	getResp, err := mgr.cli.MiningUserGet(ctx, &types.MiningUserGetReq{MiningUserName: name})
	if err != nil {
		return fmt.Errorf("have no user name %v or %v", name, err)
	}

	if getResp == nil {
		return fmt.Errorf("have no user name %v", name)
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
			return fmt.Errorf("failed to delete read page link of %v, error: %v", name, err)
		}
	}

	return nil
}

func (mgr *Manager) SetRevenueProportion(ctx context.Context, distributor, recipient string, proportion float64) error {
	infoResp, err := mgr.cli.RevenueDistributionInfo(ctx, &types.RevenueDistributionInfoReq{
		Currency:    mgr.currency,
		Distributor: distributor,
		Recipient:   recipient,
	})

	if err == nil && infoResp != nil {
		for _, v := range infoResp.Data {
			if v.Distributor != distributor || v.Currency != mgr.currency {
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

	addResp, err := mgr.cli.RevenueDistributionAdd(ctx, &types.RevenueDistributionAddReq{
		Currency:    mgr.currency,
		Distributor: distributor,
		Recipient:   recipient,
		Proportion:  proportion,
	})
	if err != nil {
		return err
	}

	if addResp == nil {
		return fmt.Errorf("failed to add revenue proportion,have nil response")
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
		return 0, err
	}

	if getResp == nil {
		return 0, fmt.Errorf("failed to get revenue proportion info,have nil response")
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
						Currency:  mgr.currency,
						Address:   address,
						Threshold: mgr.threshold,
					},
				},
			},
		},
	})
	if err != nil {
		return err
	}

	if setResp == nil {
		return fmt.Errorf("failed to set revenue address,have nil response")
	}
	return nil
}

func (mgr *Manager) GetRevenueAddress(ctx context.Context, name string) (string, error) {
	getResp, err := mgr.cli.MiningUserGet(ctx, &types.MiningUserGetReq{
		MiningUserName: name,
	})
	if err != nil {
		return "", err
	}

	if getResp == nil {
		return "", fmt.Errorf("failed to get revenue address,have nil response")
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
		return false, err
	}

	if pauseResp == nil {
		return false, fmt.Errorf("failed to pause payment,have nil response")
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
		return false, err
	}

	if resumeResp == nil {
		return false, fmt.Errorf("failed to resume payment,have nil response")
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
		return 0, err
	}

	if resumeResp == nil {
		return 0, fmt.Errorf("failed to resume payment,have nil response")
	}

	return resumeResp.PaidTime, nil
}

func getReadPageLink(key, userName string) string {
	return fmt.Sprintf("%v/mining-user/%v?user_name=%v", F2PoolBaseURL, key, userName)
}

// can only be a combination of lowercase characters and numbers
// start with letter
func RandomF2PoolUser(n int) (string, error) {
	startLetters := []rune("abcdefghijklmnopqretuvwxyz")
	randn, err := rand.Int(rand.Reader, big.NewInt(int64(len(startLetters))))
	if err != nil {
		return "", err
	}
	target := string(startLetters[randn.Int64()])
	for {
		randn, err := rand.Int(rand.Reader, big.NewInt(math.MaxInt64))
		if err != nil {
			return "", err
		}
		if len(target) >= n {
			return strings.ToLower(target[:n]), nil
		}
		target += base58.Encode(randn.Bytes())
	}
}
