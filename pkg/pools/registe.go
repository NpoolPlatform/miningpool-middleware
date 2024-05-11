package pools

import (
	mpbasetype "github.com/NpoolPlatform/message/npool/basetypes/miningpool/v1"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	coinmw "github.com/NpoolPlatform/message/npool/miningpool/mw/v1/coin"
	fractionrulemw "github.com/NpoolPlatform/message/npool/miningpool/mw/v1/fractionrule"
	poolmw "github.com/NpoolPlatform/message/npool/miningpool/mw/v1/pool"

	"github.com/NpoolPlatform/miningpool-middleware/pkg/config"
	"github.com/NpoolPlatform/miningpool-middleware/pkg/const/time"
	"github.com/NpoolPlatform/miningpool-middleware/pkg/pools/f2pool"
)

var (
	poolInfos = []*poolmw.Pool{
		{
			MiningpoolType: mpbasetype.MiningpoolType_F2Pool,
			Name:           mpbasetype.MiningpoolType_F2Pool.String(),
			Site:           config.F2PoolSite,
			Logo:           "https://static.f2pool.com/static/images/icon-logo-min-blue.png",
			Description:    "",
		},
	}

	coinInfos = []*coinmw.Coin{
		{
			MiningpoolType:         mpbasetype.MiningpoolType_F2Pool,
			CoinType:               basetypes.CoinType_CoinTypeBitCoin,
			RevenueType:            mpbasetype.RevenueType_FPPS,
			FeeRatio:               f2pool.CoinType2FeeRatio[basetypes.CoinType_CoinTypeBitCoin],
			FixedRevenueAble:       false,
			LeastTransferAmount:    f2pool.CoinType2LeastTransferAmount[basetypes.CoinType_CoinTypeBitCoin],
			BenefitIntervalSeconds: f2pool.CoinType2BenefitIntervalSeconds[basetypes.CoinType_CoinTypeBitCoin],
		},
	}

	fractionruleInfos = []*fractionrulemw.FractionRule{
		{
			MiningpoolType: mpbasetype.MiningpoolType_F2Pool,
			CoinType:       basetypes.CoinType_CoinTypeBitCoin,
			// 30Day
			WithdrawInterval: time.SecondsPerDay * 30,
			MinAmount:        "0.0005",
			WithdrawRate:     "0",
		},
	}
)
