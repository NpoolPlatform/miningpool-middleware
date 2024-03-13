package main

import (
	"context"
	"fmt"

	"github.com/NpoolPlatform/miningpool-middleware/pkg/pools/f2pool/client"
	"github.com/NpoolPlatform/miningpool-middleware/pkg/pools/f2pool/types"
	"github.com/NpoolPlatform/miningpool-middleware/pkg/utils"
)

func main() {
	cli := client.NewClient("https://api.f2pool.com", "t8jbn1rqsoah9wirc0msgqmm24vdvg65t4r15ge758u7ut1g46xgqkk8bn5w2jc8")

	resp, err := cli.MiningUserPaymentResume(context.Background(), &types.MiningUserPaymentResumeReq{
		Currency:        "bitcoin",
		MiningUserNames: []string{"coocoo", "cococonut001"},
	})
	fmt.Println(err)
	fmt.Println(utils.PrettyStruct(resp))
}
