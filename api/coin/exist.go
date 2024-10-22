//nolint:nolintlint,dupl
package coin

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/miningpool/mw/v1/coin"
	coin "github.com/NpoolPlatform/miningpool-middleware/pkg/mw/coin"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) ExistCoin(ctx context.Context, in *npool.ExistCoinRequest) (*npool.ExistCoinResponse, error) {
	handler, err := coin.NewHandler(
		ctx,
		coin.WithEntID(&in.EntID, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"ExistCoin",
			"In", in,
			"Error", err,
		)
		return &npool.ExistCoinResponse{}, status.Error(codes.Internal, "internal server err")
	}

	info, err := handler.ExistCoin(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"ExistCoin",
			"In", in,
			"Error", err,
		)
		return &npool.ExistCoinResponse{}, status.Error(codes.Internal, "internal server err")
	}

	return &npool.ExistCoinResponse{
		Info: info,
	}, nil
}

func (s *Server) ExistCoinConds(ctx context.Context, in *npool.ExistCoinCondsRequest) (*npool.ExistCoinCondsResponse, error) {
	handler, err := coin.NewHandler(
		ctx,
		coin.WithConds(in.Conds),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"ExistCoinConds",
			"In", in,
			"Error", err,
		)
		return &npool.ExistCoinCondsResponse{}, status.Error(codes.Internal, "internal server err")
	}

	info, err := handler.ExistCoinConds(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"ExistCoinConds",
			"In", in,
			"Error", err,
		)
		return &npool.ExistCoinCondsResponse{}, status.Error(codes.Internal, "internal server err")
	}

	return &npool.ExistCoinCondsResponse{
		Info: info,
	}, nil
}
