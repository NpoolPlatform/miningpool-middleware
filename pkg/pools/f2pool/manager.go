package f2pool

import (
	"context"
	"crypto/rand"
	"fmt"
	"math"
	"math/big"
	"strings"

	basetype "github.com/NpoolPlatform/message/npool/basetypes/miningpool/v1"
	"github.com/NpoolPlatform/miningpool-middleware/pkg/pools/f2pool/client"
	"github.com/NpoolPlatform/miningpool-middleware/pkg/pools/f2pool/types"
	"github.com/mr-tron/base58"
)

type F2PoolMgr struct {
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
	CoinType2Threshold map[basetype.CoinType]float64 = map[basetype.CoinType]float64{
		basetype.CoinType_BitCoin: 0.05,
	}
	F2PoolAPI  = "https://api.f2pool.com"
	MaxRetries = 10
)

const (
	//2-15 characters
	MiningUserLen         = 15
	MiningUserPageNameLen = 20
	DefaultPermissions    = "1,2"
)

func NewF2PoolManager(coinType basetype.CoinType, auth string) (*F2PoolMgr, error) {
	currency := ""
	if k, ok := CoinType2Currency[coinType]; ok {
		currency = k
	} else {
		return nil, fmt.Errorf("have no pool manager for %v-%v", MiningPoolType, coinType)
	}
	return &F2PoolMgr{
		currency:  currency,
		authToken: auth,
		cli:       client.NewClient(F2PoolAPI, auth),
	}, nil
}

func (mgr *F2PoolMgr) CheckAuth(ctx context.Context) bool {
	return false
}

func (mgr *F2PoolMgr) AddMiningUser(ctx context.Context) (string, string, error) {
	var resp *types.MiningUserAddResp
	var err error
	for i := 0; i < MaxRetries; i++ {
		userName := RandomF2PoolUser(MiningUserLen)
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

	userName := resp.MiningUserName
	readPageLink := ""
	if len(resp.Pages) > 0 {
		readPageLink = getReadPageLink(resp.Pages[0].Key, resp.MiningUserName)
	}
	return userName, readPageLink, nil
}

// just judge wheather mining user in the root-user of the auth token
func (mgr *F2PoolMgr) ExistMiningUser(ctx context.Context, name string) (bool, error) {
	_, err := mgr.cli.MiningUserGet(ctx, &types.MiningUserGetReq{MiningUserName: name})
	if err != nil {
		return true, err
	}
	return false, nil
}

func (mgr *F2PoolMgr) DeleteMiningUser(ctx context.Context, name string) error {
	return fmt.Errorf("f2pool has not yet implemented this method")
}

func (mgr *F2PoolMgr) AddReadPageLink(ctx context.Context, name string) (string, error) {
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

	pageName := RandomF2PoolUser(MiningUserPageNameLen)
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

func (mgr *F2PoolMgr) GetReadPageLink(ctx context.Context, name string) (string, error) {
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

func (mgr *F2PoolMgr) DeleteReadPageLink(ctx context.Context, name string) error {
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

	_, err = mgr.cli.MiningUserReadOnlyPageDelete(ctx, &types.MiningUserReadOnlyPageDeleteReq{
		MiningUserName: getResp.MiningUserName,
		Key:            getResp.Pages[0].Key,
	})

	if err != nil {
		return fmt.Errorf("failed to delete read page link of %v, error: %v", name, err)
	}
	return nil
}

func (mgr *F2PoolMgr) SetRevenueProportion(ctx context.Context, distributor string, recipient string, proportion float64) error {
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

func (mgr *F2PoolMgr) GetRevenueProportion(ctx context.Context, distributor string, recipient string) (float64, error) {
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
		if v.Currency == mgr.currency && v.Description == distributor && v.Recipient == recipient {
			return v.Proportion, nil
		}
	}

	return 0, nil
}

func (mgr *F2PoolMgr) SetRevenueAddress(ctx context.Context, name string, address string) error {
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

func (mgr *F2PoolMgr) GetRevenueAddress(ctx context.Context, name string) (string, error) {
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

func (mgr *F2PoolMgr) PausePayment(ctx context.Context, name string) (bool, error) {
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

func (mgr *F2PoolMgr) ResumePayment(ctx context.Context, name string) (bool, error) {
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

func getReadPageLink(key, user_name string) string {
	return fmt.Sprintf("%v/mining-user/%v?user_name=%v", F2PoolAPI, key, user_name)
}

// can only be a combination of lowercase characters and numbers
// start with letter
func RandomF2PoolUser(n int) string {
	startLetters := []rune("abcdefghijklmnopqretuvwxyz")
	randn, _ := rand.Int(rand.Reader, big.NewInt(int64(len(startLetters))))
	target := string(startLetters[randn.Int64()])
	for {
		randn, _ := rand.Int(rand.Reader, big.NewInt(math.MaxInt64))
		if len(target) >= n {
			return strings.ToLower(target[:n])
		}
		target += base58.Encode(randn.Bytes())
	}
}
