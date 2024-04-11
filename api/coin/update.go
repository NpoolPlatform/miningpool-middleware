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

func (s *Server) UpdateCoin(ctx context.Context, in *npool.UpdateCoinRequest) (*npool.UpdateCoinResponse, error) {
	if in.GetInfo() == nil {
		err := fmt.Errorf("request is nil")
		logger.Sugar().Errorw(
			"UpdateCoin",
			"In", in,
			"Error", err,
		)
		return &npool.UpdateCoinResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	req := in.GetInfo()
	handler, err := coin.NewHandler(
		ctx,
		coin.WithID(req.ID, false),
		coin.WithEntID(req.EntID, false),
		coin.WithMiningpoolType(req.MiningpoolType, true),
		coin.WithCoinType(req.CoinType, true),
		coin.WithRevenueTypes(&req.RevenueTypes, true),
		coin.WithFeeRate(req.FeeRate, true),
		coin.WithFixedRevenueAble(req.FixedRevenueAble, true),
		coin.WithThreshold(req.Threshold, true),
		coin.WithRemark(req.Remark, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"UpdateCoin",
			"In", in,
			"Error", err,
		)
		return &npool.UpdateCoinResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := handler.UpdateCoin(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"UpdateCoin",
			"In", in,
			"Error", err,
		)
		return &npool.UpdateCoinResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.UpdateCoinResponse{
		Info: info,
	}, nil
}
