package pools

import (
	"fmt"

	mpbasetype "github.com/NpoolPlatform/message/npool/basetypes/miningpool/v1"
	basetype "github.com/NpoolPlatform/message/npool/basetypes/v1"
	"github.com/NpoolPlatform/miningpool-middleware/pkg/pools/f2pool"
)

// TODO: support use default coinType
func NewPoolManager(poolType mpbasetype.MiningpoolType, coinType basetype.CoinType, auth string) (PoolManager, error) {
	if poolType == mpbasetype.MiningpoolType_F2Pool {
		return f2pool.NewF2PoolManager(coinType, auth)
	}
	return nil, fmt.Errorf("has not implemented for %v-%v ", poolType, coinType)
}
