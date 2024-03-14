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
		return &npool.ExistCoinCondsResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := handler.ExistCoinConds(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"ExistCoinConds",
			"In", in,
			"Error", err,
		)
		return &npool.ExistCoinCondsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.ExistCoinCondsResponse{
		Info: info,
	}, nil
}
