package main

import (
	"context"
	"fmt"
	"os"

	v1 "github.com/NpoolPlatform/message/npool/basetypes/miningpool/v1"
	"github.com/NpoolPlatform/miningpool-middleware/pkg/pools/f2pool"
	"github.com/NpoolPlatform/miningpool-middleware/pkg/utils"
)

var (
	accessToken = "7ecdq1fosdsfcruypom2otsn8hfr69azmqvh7v3zelol1ntsba85a1yvol66qp73"
)

func main() {
	mgr, err := f2pool.NewF2PoolManager(v1.CoinType_BitCoin, accessToken)
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
	authed := mgr.CheckAuth(context.Background())
	fmt.Println(authed)
	ret, err := mgr.ExistMiningUser(context.Background(), "cococonut1")
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
	fmt.Println(utils.PrettyStruct(ret))

}
