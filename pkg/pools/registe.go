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

func RegistePool(ctx context.Context) {
	infos := []*poolmw.Pool{
		{
			MiningpoolType: v1.MiningpoolType_F2Pool,
			Name:           v1.MiningpoolType_F2Pool.String(),
			Site:           config.F2PoolSite,
			Description:    "",
		},
	}
	for _, info := range infos {
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
	infos := []*coinmw.Coin{
		{
			MiningpoolType:   v1.MiningpoolType_F2Pool,
			CoinType:         v1.CoinType_BitCoin,
			RevenueTypes:     []v1.RevenueType{v1.RevenueType_FPPS, v1.RevenueType_PPLNS},
			FeeRate:          f2pool.CoinType2FeeRate[v1.CoinType_BitCoin],
			FixedRevenueAble: false,
			Threshold:        f2pool.CoinType2Threshold[v1.CoinType_BitCoin],
		},
	}
	for _, info := range infos {
		coinH, err := coin.NewHandler(ctx, coin.WithConds(&coinmw.Conds{
			MiningpoolType: &basetypes.Uint32Val{
				Op:    cruder.EQ,
				Value: uint32(info.MiningpoolType),
			},
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

		coinH, err = coin.NewHandler(ctx,
			coin.WithMiningpoolType(&info.MiningpoolType, true),
			coin.WithCoinType(&info.CoinType, true),
			coin.WithRevenueTypes(&info.RevenueTypes, true),
			coin.WithFeeRate(&info.FeeRate, true),
			coin.WithFixedRevenueAble(&info.FixedRevenueAble, true),
			coin.WithThreshold(&info.Threshold, true),
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
	infos := []*fractionrulemw.FractionRule{
		{
			MiningpoolType: v1.MiningpoolType_F2Pool,
			CoinType:       v1.CoinType_BitCoin,
			// 30Day
			WithdrawInterval: 60 * 60 * 24 * 30,
			MinAmount:        0.0005,
			WithdrawRate:     0,
		},
	}
	for _, info := range infos {
		fractionruleH, err := fractionrule.NewHandler(ctx, fractionrule.WithConds(
			&fractionrulemw.Conds{
				MiningpoolType: &basetypes.Uint32Val{
					Op:    cruder.EQ,
					Value: uint32(info.MiningpoolType),
				},
				CoinType: &basetypes.Uint32Val{
					Op:    cruder.EQ,
					Value: uint32(info.CoinType),
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

		fractionruleH, err = fractionrule.NewHandler(ctx,
			fractionrule.WithMiningpoolType(&info.MiningpoolType, true),
			fractionrule.WithCoinType(&info.CoinType, true),
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
