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

func (s *Server) CreatePool(ctx context.Context, in *npool.CreatePoolRequest) (*npool.CreatePoolResponse, error) {
	req := in.GetInfo()
	handler, err := pool.NewHandler(
		ctx,
		pool.WithID(req.ID, false),
		pool.WithEntID(req.EntID, false),
		pool.WithMiningpoolType(req.MiningpoolType, true),
		pool.WithName(req.Name, true),
		pool.WithSite(req.Site, true),
		pool.WithDescription(req.Description, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"CreatePool",
			"In", in,
			"Error", err,
		)
		return &npool.CreatePoolResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := handler.CreatePool(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"CreatePool",
			"In", in,
			"Error", err,
		)
		return &npool.CreatePoolResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CreatePoolResponse{
		Info: info,
	}, nil
}