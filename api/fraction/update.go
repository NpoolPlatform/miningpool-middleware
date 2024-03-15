//nolint:nolintlint,dupl
package fraction

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/miningpool/mw/v1/fraction"
	fraction "github.com/NpoolPlatform/miningpool-middleware/pkg/mw/fraction"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) UpdateFraction(ctx context.Context, in *npool.UpdateFractionRequest) (*npool.UpdateFractionResponse, error) {
	req := in.GetInfo()
	handler, err := fraction.NewHandler(
		ctx,
		fraction.WithID(req.ID, true),
		fraction.WithEntID(req.EntID, false),
		fraction.WithOrderUserID(req.OrderUserID, true),
		fraction.WithWithdrawState(req.WithdrawState, true),
		fraction.WithWithdrawTime(req.WithdrawTime, true),
		fraction.WithPayTime(req.PayTime, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"UpdateFraction",
			"In", in,
			"Error", err,
		)
		return &npool.UpdateFractionResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := handler.UpdateFraction(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"UpdateFraction",
			"In", in,
			"Error", err,
		)
		return &npool.UpdateFractionResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.UpdateFractionResponse{
		Info: info,
	}, nil
}
