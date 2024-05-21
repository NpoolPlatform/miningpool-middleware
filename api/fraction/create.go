//nolint:nolintlint,dupl
package fraction

import (
	"context"
	"fmt"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/miningpool/mw/v1/fraction"
	fraction "github.com/NpoolPlatform/miningpool-middleware/pkg/mw/fraction"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) CreateFraction(ctx context.Context, in *npool.CreateFractionRequest) (*npool.CreateFractionResponse, error) {
	if in.GetInfo() == nil {
		err := fmt.Errorf("request is nil")
		logger.Sugar().Errorw(
			"CreateFraction",
			"In", in,
			"Error", err,
		)
		return &npool.CreateFractionResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	req := in.GetInfo()
	handler, err := fraction.NewHandler(
		ctx,
		fraction.WithEntID(req.EntID, false),
		fraction.WithAppID(req.AppID, true),
		fraction.WithUserID(req.UserID, true),
		fraction.WithOrderUserID(req.OrderUserID, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateFraction",
			"In", in,
			"Error", err,
		)
		return &npool.CreateFractionResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	err = handler.CreateFraction(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateFraction",
			"In", in,
			"Error", err,
		)
		return &npool.CreateFractionResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CreateFractionResponse{}, nil
}
