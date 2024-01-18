package main

import (
	"context"
	"crypto/rand"
	"fmt"
	"math"
	"math/big"
	"strings"

	"github.com/NpoolPlatform/miningpool-middleware/pkg/pools/f2pool/client"
	"github.com/NpoolPlatform/miningpool-middleware/pkg/pools/f2pool/types"
	"github.com/NpoolPlatform/miningpool-middleware/pkg/utils"
	"github.com/mr-tron/base58"
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
