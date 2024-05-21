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

func (s *Server) GetPool(ctx context.Context, in *npool.GetPoolRequest) (*npool.GetPoolResponse, error) {
	handler, err := pool.NewHandler(
		ctx,
		pool.WithEntID(&in.EntID, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetPool",
			"In", in,
			"Error", err,
		)
		return &npool.GetPoolResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := handler.GetPool(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetPool",
			"In", in,
			"Error", err,
		)
		return &npool.GetPoolResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetPoolResponse{
		Info: info,
	}, nil
}

func (s *Server) GetPools(ctx context.Context, in *npool.GetPoolsRequest) (*npool.GetPoolsResponse, error) {
	handler, err := pool.NewHandler(
		ctx,
		pool.WithConds(in.Conds),
		pool.WithOffset(in.GetOffset()),
		pool.WithLimit(in.GetLimit()),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetPools",
			"In", in,
			"Error", err,
		)
		return &npool.GetPoolsResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	infos, total, err := handler.GetPools(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetPools",
			"In", in,
			"Error", err,
		)
		return &npool.GetPoolsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetPoolsResponse{
		Infos: infos,
		Total: total,
	}, nil
}
