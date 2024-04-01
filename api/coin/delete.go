package coin

import (
	"context"
	"fmt"

	npool "github.com/NpoolPlatform/message/npool/miningpool/mw/v1/coin"
	coin "github.com/NpoolPlatform/miningpool-middleware/pkg/mw/coin"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) DeleteCoin(ctx context.Context, in *npool.DeleteCoinRequest) (*npool.DeleteCoinResponse, error) {
	req := in.GetInfo()
	if req == nil || req.ID == nil {
		err := fmt.Errorf("wrong id")
		return &npool.DeleteCoinResponse{}, status.Error(codes.Internal, err.Error())
	}
	handler, err := coin.NewHandler(
		ctx,
		coin.WithID(req.ID, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"DeleteCoin",
			"In", in,
			"Error", err,
		)
		return &npool.DeleteCoinResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := handler.DeleteCoin(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"DeleteCoin",
			"In", in,
			"Error", err,
		)
		return &npool.DeleteCoinResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.DeleteCoinResponse{
		Info: info,
	}, nil
}
