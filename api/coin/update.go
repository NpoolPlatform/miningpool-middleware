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
		coin.WithRevenueType(req.RevenueType, false),
		coin.WithFeeRatio(req.FeeRatio, false),
		coin.WithFixedRevenueAble(req.FixedRevenueAble, false),
		coin.WithLeastTransferAmount(req.LeastTransferAmount, false),
		coin.WithBenefitIntervalSeconds(req.BenefitIntervalSeconds, false),
		coin.WithRemark(req.Remark, false),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"UpdateCoin",
			"In", in,
			"Error", err,
		)
		return &npool.UpdateCoinResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	err = handler.UpdateCoin(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"UpdateCoin",
			"In", in,
			"Error", err,
		)
		return &npool.UpdateCoinResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.UpdateCoinResponse{}, nil
}
