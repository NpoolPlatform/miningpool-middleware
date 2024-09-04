package pools

import (
	"github.com/NpoolPlatform/go-service-framework/pkg/wlog"
	mpbasetype "github.com/NpoolPlatform/message/npool/basetypes/miningpool/v1"
	basetype "github.com/NpoolPlatform/message/npool/basetypes/v1"
	"github.com/NpoolPlatform/miningpool-middleware/pkg/pools/f2pool"
	"github.com/NpoolPlatform/miningpool-middleware/pkg/pools/types"
)

// TODO: support use default coinType
func NewPoolManager(poolType mpbasetype.MiningPoolType, coinType *basetype.CoinType, auth string) (types.PoolManager, error) {
	if poolType == mpbasetype.MiningPoolType_F2Pool {
		return f2pool.NewF2PoolManager(coinType, auth)
	}
	return nil, wlog.Errorf("has not implemented for %v-%v ", poolType, coinType)
}
