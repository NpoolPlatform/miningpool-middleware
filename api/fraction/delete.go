package fraction

import (
	"context"
	"fmt"

	npool "github.com/NpoolPlatform/message/npool/miningpool/mw/v1/fraction"
	fraction "github.com/NpoolPlatform/miningpool-middleware/pkg/mw/fraction"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) DeleteFraction(ctx context.Context, in *npool.DeleteFractionRequest) (*npool.DeleteFractionResponse, error) {
	if in.GetInfo() == nil {
		err := fmt.Errorf("request is nil")
		logger.Sugar().Errorw(
			"DeleteFraction",
			"In", in,
			"Error", err,
		)
		return &npool.DeleteFractionResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	req := in.GetInfo()
	if req == nil || req.ID == nil {
		err := fmt.Errorf("wrong id")
		return &npool.DeleteFractionResponse{}, status.Error(codes.Internal, err.Error())
	}
	handler, err := fraction.NewHandler(
		ctx,
		fraction.WithID(req.ID, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"DeleteFraction",
			"In", in,
			"Error", err,
		)
		return &npool.DeleteFractionResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := handler.DeleteFraction(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"DeleteFraction",
			"In", in,
			"Error", err,
		)
		return &npool.DeleteFractionResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.DeleteFractionResponse{
		Info: info,
	}, nil
}
