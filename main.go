package main

import (
	"context"
	"fmt"

	"github.com/NpoolPlatform/miningpool-middleware/pkg/pools/f2pool/client"
	"github.com/NpoolPlatform/miningpool-middleware/pkg/pools/f2pool/types"
	"github.com/NpoolPlatform/miningpool-middleware/pkg/utils"
)

func main() {
	baseURL := "https://api.f2pool.com"

	accessToken := "wyrhuvsac5iaej9s3q1qx3l2lwvuoso1sdxvxzx1rju6tr27bqiujey9sj5ng546"

	cli := client.NewClient(baseURL, accessToken)
	resp, err := cli.MiningUserGet(context.Background(), &types.MiningUserGetReq{
		MiningUserName: "cococonut3",
	})

	fmt.Println(utils.PrettyStruct(resp))
	fmt.Println(err)
}
