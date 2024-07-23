// nolint
package f2pool

import (
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"time"

	"github.com/NpoolPlatform/miningpool-middleware/pkg/pools/f2pool/types"
	"github.com/google/uuid"
)

func httpHandler(path string, req interface{}, handle func() (interface{}, error)) {
	handler := func(w http.ResponseWriter, r *http.Request) {
		var reqBytes []byte
		var respBytes []byte
		var err error
		defer func() {
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				w.Write([]byte(err.Error()))
				return
			}
			w.Header().Add("Content-Type", "application/json")
			w.Write(respBytes)
		}()
		reqBytes, err = io.ReadAll(r.Body)
		if err != nil {
			return
		}

		err = json.Unmarshal(reqBytes, req)
		if err != nil {
			return
		}

		resp, err := handle()
		if err != nil {
			return
		}

		respBytes, err = json.Marshal(resp)
		if err != nil {
			return
		}
	}
	http.HandleFunc(path, handler)
}

var (
	users        map[string]struct{}
	proportions  map[string]float64
	revenueAddrs map[string]map[string]string
	defaultPage  types.ReadOnlyPage

	knownUsers = []string{"Bob123", "Lily123", "Hasiky123"}
)

func init() {
	users = make(map[string]struct{})
	for _, k := range knownUsers {
		users[k] = struct{}{}
	}

	proportions = make(map[string]float64)
	revenueAddrs = make(map[string]map[string]string)
	defaultPage = types.ReadOnlyPage{
		Key:         "mock_key",
		PageName:    "mock_name",
		Permissions: "1,2",
	}
}

func mockMiningUserGet() {
	req := &types.MiningUserGetReq{}
	httpHandler(types.MiningUserGet, req, func() (interface{}, error) {
		if _, ok := users[req.MiningUserName]; !ok {
			fmt.Println("error sssssssssdfasdf", req.MiningUserName)
			return nil, fmt.Errorf("have no user")
		}

		wallets := []types.Wallet{}
		if _, ok := revenueAddrs[req.MiningUserName]; ok {
			for k, v := range revenueAddrs[req.MiningUserName] {
				wallets = append(wallets, types.Wallet{
					Currency:            k,
					Address:             v,
					LeastTransferAmount: 0.005,
				})
			}
		}

		resp := &types.MiningUserGetResp{
			MiningUserName: req.MiningUserName,
			Pages: []types.ReadOnlyPage{
				defaultPage,
			},
			Wallets:     wallets,
			Description: "mock data",
		}
		return resp, nil
	})
}

func mockMiningUserAdd() {
	req := &types.MiningUserAddReq{}
	httpHandler(types.MiningUserAdd, req, func() (interface{}, error) {
		users[req.MiningUserName] = struct{}{}

		resp := &types.MiningUserAddResp{
			MiningUserName: req.MiningUserName,
			Pages: []types.ReadOnlyPage{
				defaultPage,
			},
		}
		return resp, nil
	})
}

func mockMiningUserList() {
	req := &types.MiningUserListReq{}
	httpHandler(types.MiningUserList, req, func() (interface{}, error) {
		userlist := []types.MiningUserInfo{}
		for username := range users {
			userlist = append(userlist, types.MiningUserInfo{
				MiningUserName: username,
				Pages: []types.ReadOnlyPage{
					defaultPage,
				},
				Wallets:     []types.Wallet{},
				Description: "mock data",
			})
		}

		resp := &types.MiningUserListResp{
			MiningUserList: userlist,
		}
		return resp, nil
	})
}

func mockMiningUserReadOnlyPageDelete() {
	req := &types.MiningUserReadOnlyPageDeleteReq{}
	httpHandler(types.MiningUserReadOnlyPageDelete, req, func() (interface{}, error) {
		resp := &types.MiningUserReadOnlyPageDeleteResp{
			MiningUserName: req.MiningUserName,
		}
		return resp, nil
	})
}

func mockRevenueDistributionInfo() {
	req := &types.RevenueDistributionInfoReq{}
	httpHandler(types.RevenueDistributionInfo, req, func() (interface{}, error) {
		proportionKey := fmt.Sprintf("%v-%v-%v", req.Currency, req.Distributor, req.Recipient)

		if _, ok := proportions[proportionKey]; !ok {
			return nil, fmt.Errorf("have no info")
		}
		resp := &types.RevenueDistributionInfoResp{
			Data: []types.RevenueDistribution{
				{
					ID:          rand.Int63(),
					Currency:    req.Currency,
					Distributor: req.Distributor,
					Recipient:   req.Recipient,
					Description: "mock data",
					Proportion:  proportions[proportionKey],
				},
			},
		}
		return resp, nil
	})
}

func mockRevenueDistributionAdd() {
	req := &types.RevenueDistributionAddReq{}
	httpHandler(types.RevenueDistributionAdd, req, func() (interface{}, error) {
		proportionKey := fmt.Sprintf("%v-%v-%v", req.Currency, req.Distributor, req.Recipient)
		proportions[proportionKey] = req.Proportion
		resp := &types.RevenueDistributionAddResp{
			ID:          uuid.NewString(),
			Currency:    req.Currency,
			Distributor: req.Distributor,
			Recipient:   req.Recipient,
			Description: "mock data",
			Proportion:  req.Proportion,
		}
		return resp, nil
	})
}

func mockRevenueDistributionDelete() {
	req := &types.RevenueDistributionDeleteReq{}
	httpHandler(types.RevenueDistributionDelete, req, func() (interface{}, error) {
		resp := &types.RevenueDistributionDeleteResp{
			Reason: "mock data",
		}
		return resp, nil
	})
}

func mockMiningUserWalletUpdate() {
	req := &types.MiningUserWalletUpdateReq{}
	httpHandler(types.MiningUserWalletUpdate, req, func() (interface{}, error) {
		name := req.Params[0].MiningUserName
		currency := req.Params[0].Wallets[0].Currency
		addr := req.Params[0].Wallets[0].Address

		if _, ok := revenueAddrs[name]; !ok {
			revenueAddrs[name] = make(map[string]string)
		}
		revenueAddrs[name][currency] = addr
		ret := make(map[string]types.MiningUserWalletResp)
		ret[name] = types.MiningUserWalletResp{
			MiningUserName: name,
			WalletResult: []types.WalletResp{
				{
					Currency: currency,
					Address:  addr,
					Result:   "ok",
				},
			},
		}
		resp := &types.MiningUserWalletUpdateResp{
			Results: ret,
		}
		return resp, nil
	})
}

func mockMiningUserPaymentPause() {
	req := &types.MiningUserPaymentPauseReq{}
	httpHandler(types.MiningUserPaymentPause, req, func() (interface{}, error) {
		ret := make(map[string]string)
		for _, v := range req.MiningUserNames {
			ret[v] = "ok"
		}
		resp := &types.MiningUserPaymentPauseResp{
			Results: ret,
		}
		return resp, nil
	})
}

func mockMiningUserPaymentResume() {
	req := &types.MiningUserPaymentResumeReq{}
	httpHandler(types.MiningUserPaymentResume, req, func() (interface{}, error) {
		ret := make(map[string]string)
		for _, v := range req.MiningUserNames {
			ret[v] = "ok"
		}
		resp := &types.MiningUserPaymentResumeResp{
			Results: ret,
		}
		return resp, nil
	})
}

func mockMiningUserBalanceWithdraw() {
	req := &types.MiningUserBalanceWithdrawReq{}
	httpHandler(types.MiningUserBalanceWithdraw, req, func() (interface{}, error) {
		resp := &types.MiningUserBalanceWithdrawResp{
			PaidTime: time.Now().Unix(),
		}
		return resp, nil
	})
}

func Init(httpPort uint16) {
	mockMiningUserGet()
	mockMiningUserAdd()
	mockMiningUserList()
	mockMiningUserReadOnlyPageDelete()
	mockRevenueDistributionInfo()
	mockRevenueDistributionAdd()
	mockRevenueDistributionDelete()
	mockMiningUserWalletUpdate()
	mockMiningUserPaymentPause()
	mockMiningUserPaymentResume()
	mockMiningUserBalanceWithdraw()

	httpServer := fmt.Sprintf(":%v", httpPort)
	fmt.Printf("start mock f2pool server,address: %v\n", httpServer)
	http.ListenAndServe(httpServer, nil)
}
