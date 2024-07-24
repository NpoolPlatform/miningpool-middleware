package registetestinfo

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	mpbasetype "github.com/NpoolPlatform/message/npool/basetypes/miningpool/v1"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	coinmw "github.com/NpoolPlatform/message/npool/miningpool/mw/v1/coin"
	fractionrulemw "github.com/NpoolPlatform/message/npool/miningpool/mw/v1/fractionrule"
	poolmw "github.com/NpoolPlatform/message/npool/miningpool/mw/v1/pool"
	"github.com/google/uuid"

	"github.com/NpoolPlatform/miningpool-middleware/pkg/config"
	"github.com/NpoolPlatform/miningpool-middleware/pkg/const/time"
	"github.com/NpoolPlatform/miningpool-middleware/pkg/mw/coin"
	"github.com/NpoolPlatform/miningpool-middleware/pkg/mw/fractionrule"
	"github.com/NpoolPlatform/miningpool-middleware/pkg/mw/pool"
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
	poolEntIDs         = []*uuid.UUID{}
	coinEntIDs         = []*uuid.UUID{}
	fractionruleEntIDs = []*uuid.UUID{}

	coinInfos = []*coinmw.Coin{
		{
			MiningpoolType:         mpbasetype.MiningpoolType_F2Pool,
			CoinType:               basetypes.CoinType_CoinTypeBitCoin,
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

func registePool(ctx context.Context) {
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

		info.EntID = uuid.NewString()
		poolH, err = pool.NewHandler(ctx,
			pool.WithEntID(&info.EntID, true),
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
		poolEntIDs = append(poolEntIDs, poolH.EntID)
	}
}

func registeCoinInfo(ctx context.Context) {
	for _, info := range coinInfos {
		// check if exist
		coinH, err := coin.NewHandler(ctx, coin.WithConds(&coinmw.Conds{
			CoinType: &basetypes.Uint32Val{
				Op:    cruder.EQ,
				Value: uint32(info.CoinType),
			},
			MiningpoolType: &basetypes.Uint32Val{
				Op:    cruder.EQ,
				Value: uint32(info.MiningpoolType),
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

		info.EntID = uuid.NewString()
		info.CoinTypeID = uuid.NewString()
		// create coin
		coinH, err = coin.NewHandler(ctx,
			coin.WithEntID(&info.EntID, true),
			coin.WithCoinTypeID(&info.CoinTypeID, true),
			coin.WithPoolID(&poolInfos[0].EntID, true),
			coin.WithCoinType(&info.CoinType, true),
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
		coinEntIDs = append(coinEntIDs, coinH.EntID)
	}
}

func registeFractionRule(ctx context.Context) {
	for _, info := range fractionruleInfos {
		coinH, err := coin.NewHandler(ctx,
			coin.WithConds(&coinmw.Conds{
				CoinType: &basetypes.Uint32Val{
					Op:    cruder.EQ,
					Value: uint32(info.CoinType),
				},
				MiningpoolType: &basetypes.Uint32Val{
					Op:    cruder.EQ,
					Value: uint32(info.MiningpoolType),
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
			continue
		}
		// check if exist
		fractionruleH, err := fractionrule.NewHandler(ctx, fractionrule.WithConds(
			&fractionrulemw.Conds{
				PoolCoinTypeID: &basetypes.StringVal{
					Op:    cruder.EQ,
					Value: coinInfos[0].EntID,
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

		info.EntID = uuid.NewString()
		// create fraction rule
		fractionruleH, err = fractionrule.NewHandler(ctx,
			fractionrule.WithEntID(&info.EntID, true),
			fractionrule.WithPoolCoinTypeID(&coinInfos[0].EntID, true),
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
		fractionruleEntIDs = append(fractionruleEntIDs, fractionruleH.EntID)
	}
}

func cleanPools(ctx context.Context) {
	for _, entid := range poolEntIDs {
		_entid := entid.String()
		poolH, err := pool.NewHandler(ctx, pool.WithEntID(&_entid, true))
		if err != nil {
			logger.Sugar().Error(err)
		}
		err = poolH.DeletePool(ctx)
		if err != nil {
			logger.Sugar().Error(err)
		}
	}
}

func cleanCoins(ctx context.Context) {
	for _, entid := range coinEntIDs {
		_entid := entid.String()
		coinH, err := coin.NewHandler(ctx, coin.WithEntID(&_entid, true))
		if err != nil {
			logger.Sugar().Error(err)
		}
		err = coinH.DeleteCoin(ctx)
		if err != nil {
			logger.Sugar().Error(err)
		}
	}
}

func cleanFractionrules(ctx context.Context) {
	for _, entid := range fractionruleEntIDs {
		_entid := entid.String()
		fractionruleH, err := fractionrule.NewHandler(ctx, fractionrule.WithEntID(&_entid, true))
		if err != nil {
			logger.Sugar().Error(err)
		}
		err = fractionruleH.DeleteFractionRule(ctx)
		if err != nil {
			logger.Sugar().Error(err)
		}
	}
}

func InitTestInfo(ctx context.Context) {
	registePool(ctx)
	registeCoinInfo(ctx)
	registeFractionRule(ctx)
}

func CleanTestInfo(ctx context.Context) {
	cleanFractionrules(ctx)
	cleanCoins(ctx)
	cleanPools(ctx)
}
