//nolint:nolintlint,dupl
package pool

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/miningpool/mw/v1/app/pool"
	pool "github.com/NpoolPlatform/miningpool-middleware/pkg/mw/app/pool"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) ExistPoolConds(ctx context.Context, in *npool.ExistPoolCondsRequest) (*npool.ExistPoolCondsResponse, error) {
	handler, err := pool.NewHandler(
		ctx,
		pool.WithConds(in.Conds),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"ExistPoolConds",
			"In", in,
			"Error", err,
		)
		return &npool.ExistPoolCondsResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := handler.ExistPoolConds(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"ExistPoolConds",
			"In", in,
			"Error", err,
		)
		return &npool.ExistPoolCondsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.ExistPoolCondsResponse{
		Info: info,
	}, nil
}
