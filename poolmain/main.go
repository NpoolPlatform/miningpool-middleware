package main

import (
	"context"
	"fmt"
	"os"
	"time"

	v1 "github.com/NpoolPlatform/message/npool/basetypes/miningpool/v1"
	"github.com/NpoolPlatform/miningpool-middleware/pkg/pools/f2pool"
)

func main() {
	mgr, err := f2pool.NewF2PoolManager(v1.CoinType_BitCoin, "7ecdq1fosdsfcruypom2otsn8hfr69azmqvh7v3zelol1ntsba85a1yvol66qp72")
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
	err = mgr.CheckAuth(context.Background())
	fmt.Println(err)

	time.Sleep(time.Second * 3)
}
