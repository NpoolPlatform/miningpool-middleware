//nolint:nolintlint,dupl
package pool

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/miningpool/mw/v1/pool"
	pool "github.com/NpoolPlatform/miningpool-middleware/pkg/mw/pool"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) ExistPool(ctx context.Context, in *npool.ExistPoolRequest) (*npool.ExistPoolResponse, error) {
	handler, err := pool.NewHandler(
		ctx,
		pool.WithEntID(&in.EntID, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"ExistPool",
			"In", in,
			"Error", err,
		)
		return &npool.ExistPoolResponse{}, status.Error(codes.Internal, "internal server err")
	}

	info, err := handler.ExistPool(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"ExistPool",
			"In", in,
			"Error", err,
		)
		return &npool.ExistPoolResponse{}, status.Error(codes.Internal, "internal server err")
	}

	return &npool.ExistPoolResponse{
		Info: info,
	}, nil
}

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
		return &npool.ExistPoolCondsResponse{}, status.Error(codes.Internal, "internal server err")
	}

	info, err := handler.ExistPoolConds(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"ExistPoolConds",
			"In", in,
			"Error", err,
		)
		return &npool.ExistPoolCondsResponse{}, status.Error(codes.Internal, "internal server err")
	}

	return &npool.ExistPoolCondsResponse{
		Info: info,
	}, nil
}
