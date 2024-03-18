package pools

import (
	"fmt"

	basetype "github.com/NpoolPlatform/message/npool/basetypes/miningpool/v1"
	"github.com/NpoolPlatform/miningpool-middleware/pkg/pools/f2pool"
)

func NewPoolManager(poolType basetype.MiningpoolType, coinType basetype.CoinType, auth string) (PoolManager, error) {
	if poolType == basetype.MiningpoolType_F2Pool {
		return f2pool.NewF2PoolManager(coinType, auth)
	}
	return nil, fmt.Errorf("has not implemented for %v-%v ", poolType, coinType)
}
