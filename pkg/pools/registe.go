package pools

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	v1 "github.com/NpoolPlatform/message/npool/basetypes/miningpool/v1"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	coinmw "github.com/NpoolPlatform/message/npool/miningpool/mw/v1/coin"
	fractionrulemw "github.com/NpoolPlatform/message/npool/miningpool/mw/v1/fractionrule"
	poolmw "github.com/NpoolPlatform/message/npool/miningpool/mw/v1/pool"

	"github.com/NpoolPlatform/miningpool-middleware/pkg/config"
	"github.com/NpoolPlatform/miningpool-middleware/pkg/const/time"
	"github.com/NpoolPlatform/miningpool-middleware/pkg/mw/coin"
	"github.com/NpoolPlatform/miningpool-middleware/pkg/mw/fractionrule"
	"github.com/NpoolPlatform/miningpool-middleware/pkg/mw/pool"
	"github.com/NpoolPlatform/miningpool-middleware/pkg/pools/f2pool"
)

func Init() {
	RegistePool(context.Background())
	RegisteCoinInfo(context.Background())
	RegisteFractionRule(context.Background())
}

var (
	poolInfos = []*poolmw.Pool{
		{
			MiningpoolType: v1.MiningpoolType_F2Pool,
			Name:           v1.MiningpoolType_F2Pool.String(),
			Site:           config.F2PoolSite,
			Logo:           "https://static.f2pool.com/static/images/icon-logo-min-blue.png",
			Description:    "",
		},
	}

	coinInfos = []*coinmw.Coin{
		{
			MiningpoolType:         v1.MiningpoolType_F2Pool,
			CoinType:               v1.CoinType_BitCoin,
			RevenueType:            v1.RevenueType_FPPS,
			FeeRatio:               f2pool.CoinType2FeeRatio[v1.CoinType_BitCoin],
			FixedRevenueAble:       false,
			LeastTransferAmount:    f2pool.CoinType2LeastTransferAmount[v1.CoinType_BitCoin],
			BenefitIntervalSeconds: f2pool.CoinType2BenefitIntervalSeconds[v1.CoinType_BitCoin],
		},
	}

	fractionruleInfos = []*fractionrulemw.FractionRule{
		{
			MiningpoolType: v1.MiningpoolType_F2Pool,
			CoinType:       v1.CoinType_BitCoin,
			// 30Day
			WithdrawInterval: time.SecondsPerDay * 30,
			MinAmount:        "0.0005",
			WithdrawRate:     "0",
		},
	}
)

func RegistePool(ctx context.Context) {
	for _, info := range poolInfos {
		poolH, err := pool.NewHandler(ctx, pool.WithConds(&poolmw.Conds{
			MiningpoolType: &basetypes.Uint32Val{
				Op:    cruder.EQ,
				Value: uint32(info.MiningpoolType),
			},
			Name: &basetypes.StringVal{
				Op:    cruder.EQ,
				Value: info.Name,
			},
		}))
		if err != nil {
			logger.Sugar().Error(err)
			continue
		}

		exist, err := poolH.ExistPoolConds(ctx)
		if err != nil {
			logger.Sugar().Error(err)
			continue
		}
		if exist {
			continue
		}

		poolH, err = pool.NewHandler(ctx,
			pool.WithMiningpoolType(&info.MiningpoolType, true),
			pool.WithName(&info.Name, true),
			pool.WithSite(&info.Site, true),
			pool.WithLogo(&info.Logo, true),
			pool.WithDescription(&info.Description, true),
		)
		if err != nil {
			logger.Sugar().Error(err)
			continue
		}
		err = poolH.CreatePool(ctx)
		if err != nil {
			logger.Sugar().Error(err)
		}
	}
}

func RegisteCoinInfo(ctx context.Context) {
	for _, info := range coinInfos {
		// check if exist
		coinH, err := coin.NewHandler(ctx, coin.WithConds(&coinmw.Conds{
			CoinType: &basetypes.Uint32Val{
				Op:    cruder.EQ,
				Value: uint32(info.CoinType),
			},
		}))
		if err != nil {
			logger.Sugar().Error(err)
			continue
		}

		exist, err := coinH.ExistCoinConds(ctx)
		if err != nil {
			logger.Sugar().Error(err)
			continue
		}
		if exist {
			continue
		}

		// get pools
		poolH, err := pool.NewHandler(ctx,
			pool.WithConds(&poolmw.Conds{
				MiningpoolType: &basetypes.Uint32Val{
					Op:    cruder.EQ,
					Value: uint32(info.MiningpoolType),
				},
			}),
			pool.WithOffset(0),
			pool.WithLimit(2),
		)
		if err != nil {
			logger.Sugar().Error(err)
			continue
		}
		poolInfos, _, err := poolH.GetPools(ctx)
		if err != nil {
			logger.Sugar().Error(err)
			continue
		}

		if len(poolInfos) == 0 {
			logger.Sugar().Errorf("have no pool of %v", info.MiningpoolType)
			continue
		}

		// create coin
		coinH, err = coin.NewHandler(ctx,
			coin.WithPoolID(&poolInfos[0].EntID, true),
			coin.WithCoinType(&info.CoinType, true),
			coin.WithRevenueType(&info.RevenueType, true),
			coin.WithFeeRatio(&info.FeeRatio, true),
			coin.WithFixedRevenueAble(&info.FixedRevenueAble, true),
			coin.WithLeastTransferAmount(&info.LeastTransferAmount, true),
			coin.WithBenefitIntervalSeconds(&info.BenefitIntervalSeconds, true),
		)
		if err != nil {
			logger.Sugar().Error(err)
			continue
		}
		err = coinH.CreateCoin(ctx)
		if err != nil {
			logger.Sugar().Error(err)
		}
	}
}

func RegisteFractionRule(ctx context.Context) {
	for _, info := range fractionruleInfos {

		coinH, err := coin.NewHandler(ctx,
			coin.WithConds(&coinmw.Conds{
				CoinType: &basetypes.Uint32Val{
					Op:    cruder.EQ,
					Value: uint32(info.CoinType),
				},
			}),
			coin.WithOffset(0),
			coin.WithLimit(2))
		if err != nil {
			logger.Sugar().Error(err)
			continue
		}

		// get coins
		coinInfos, _, err := coinH.GetCoins(ctx)
		if err != nil {
			logger.Sugar().Error(err)
			continue
		}
		if len(coinInfos) == 0 {
			logger.Sugar().Errorf("have no coin of %v", info.CoinType)
		}
		// check if exist
		fractionruleH, err := fractionrule.NewHandler(ctx, fractionrule.WithConds(
			&fractionrulemw.Conds{
				CoinID: &basetypes.StringVal{
					Op:    cruder.EQ,
					Value: string(coinInfos[0].EntID),
				},
			},
		))
		if err != nil {
			logger.Sugar().Error(err)
			continue
		}

		exist, err := fractionruleH.ExistFractionRuleConds(ctx)
		if err != nil {
			logger.Sugar().Error(err)
			continue
		}
		if exist {
			continue
		}
		// create fraction rule
		fractionruleH, err = fractionrule.NewHandler(ctx,
			fractionrule.WithCoinID(&coinInfos[0].EntID, true),
			fractionrule.WithWithdrawInterval(&info.WithdrawInterval, true),
			fractionrule.WithMinAmount(&info.MinAmount, true),
			fractionrule.WithWithdrawRate(&info.WithdrawRate, true),
		)
		if err != nil {
			logger.Sugar().Error(err)
			continue
		}
		err = fractionruleH.CreateFractionRule(ctx)
		if err != nil {
			logger.Sugar().Error(err)
		}
	}
}
