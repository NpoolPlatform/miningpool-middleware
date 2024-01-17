package main

import (
	"context"
	"fmt"

	"github.com/NpoolPlatform/miningpool-middleware/pkg/pools/f2pool/client"
	"github.com/NpoolPlatform/miningpool-middleware/pkg/pools/f2pool/types"
	"github.com/NpoolPlatform/miningpool-middleware/pkg/utils"
	"github.com/google/uuid"
)

func main() {
	baseURL := "https://api.f2pool.com"

	accessToken := "73wnf4csbcgu85rrlm3v5n0vxncv8ts7mvhdy6u65pfdw14len5q7lliol9zx7y0"
	// accessToken := "kvgmc1xfr5mqzbtqogoy6g5xfj4yp1odxf3wyby6c4pngnv9eylyohvduancc8hs"

	cli := client.NewClient(baseURL, accessToken)
	// resp, err := cli.RevenueDistributionAdd(context.Background(), &types.RevenueDistributionAddReq{
	// 	Currency:    "bitcoin",
	// 	Distributor: "coocoo",
	// 	Recipient:   "cococonut1",
	// 	Description: "test_distribution",
	// 	Proportion:  0.02,
	// })

	// resp, err := cli.RevenueDistributionInfo(context.Background(), &types.RevenueDistributionInfoReq{
	// 	Currency:    "bitcoin",
	// 	Distributor: "noonoo",
	// 	Recipient:   "noonoo2",
	// })

	resp, err := cli.MiningUserAdd(context.Background(), &types.MiningUserAddReq{
		MiningUserName: uuid.NewString(),
	})

	fmt.Println(utils.PrettyStruct(resp))
	fmt.Println(err)

	fmt.Println(len(uuid.NewString()))

}
