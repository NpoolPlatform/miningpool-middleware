package main

import (
	"context"
	"fmt"

	v1 "github.com/NpoolPlatform/message/npool/basetypes/miningpool/v1"
	"github.com/NpoolPlatform/miningpool-middleware/pkg/pools"
)

func main() {
	mgr, err := pools.NewPoolManager(v1.MiningpoolType_F2Pool, v1.CoinType_BitCoin, "7ecdq1fosdsfcruypom2otsn8hfr69azmqvh7v3zelol1ntsba85a1yvol66qp73")
	// mgr, err := pools.NewPoolManager(v1.MiningpoolType_F2Pool, v1.CoinType_BitCoin, "sssdfasdfa")
	fmt.Println(err)
	fmt.Println(mgr.CheckAuth(context.Background()))
}
