//nolint:nolintlint,dupl
package pool

import (
	"context"
	"fmt"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/miningpool/mw/v1/pool"
	pool "github.com/NpoolPlatform/miningpool-middleware/pkg/mw/pool"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) UpdatePool(ctx context.Context, in *npool.UpdatePoolRequest) (*npool.UpdatePoolResponse, error) {
	if in.GetInfo() == nil {
		err := fmt.Errorf("request is nil")
		logger.Sugar().Errorw(
			"UpdatePool",
			"In", in,
			"Error", err,
		)
		return &npool.UpdatePoolResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	req := in.GetInfo()
	handler, err := pool.NewHandler(
		ctx,
		pool.WithID(req.ID, false),
		pool.WithEntID(req.EntID, false),
		pool.WithMiningpoolType(req.MiningpoolType, false),
		pool.WithName(req.Name, false),
		pool.WithSite(req.Site, false),
		pool.WithDescription(req.Description, false),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"UpdatePool",
			"In", in,
			"Error", err,
		)
		return &npool.UpdatePoolResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := handler.UpdatePool(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"UpdatePool",
			"In", in,
			"Error", err,
		)
		return &npool.UpdatePoolResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.UpdatePoolResponse{
		Info: info,
	}, nil
}
