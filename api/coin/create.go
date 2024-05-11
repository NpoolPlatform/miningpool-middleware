//nolint:nolintlint,dupl
package coin

import (
	"context"
	"fmt"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/miningpool/mw/v1/coin"
	coin "github.com/NpoolPlatform/miningpool-middleware/pkg/mw/coin"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) CreateCoin(ctx context.Context, in *npool.CreateCoinRequest) (*npool.CreateCoinResponse, error) {
	if in.GetInfo() == nil {
		err := fmt.Errorf("request is nil")
		logger.Sugar().Errorw(
			"CreateCoin",
			"In", in,
			"Error", err,
		)
		return &npool.CreateCoinResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	req := in.GetInfo()
	handler, err := coin.NewHandler(
		ctx,
		coin.WithEntID(req.EntID, false),
		coin.WithPoolID(req.PoolID, true),
		coin.WithCoinTypeID(req.CoinTypeID, true),
		coin.WithCoinType(req.CoinType, true),
		coin.WithRevenueType(req.RevenueType, true),
		coin.WithFeeRatio(req.FeeRatio, true),
		coin.WithFixedRevenueAble(req.FixedRevenueAble, true),
		coin.WithLeastTransferAmount(req.LeastTransferAmount, true),
		coin.WithBenefitIntervalSeconds(req.BenefitIntervalSeconds, true),
		coin.WithRemark(req.Remark, false),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateCoin",
			"In", in,
			"Error", err,
		)
		return &npool.CreateCoinResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	err = handler.CreateCoin(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateCoin",
			"In", in,
			"Error", err,
		)
		return &npool.CreateCoinResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CreateCoinResponse{}, nil
}
