//nolint:nolintlint,dupl
package pool

import (
	"context"
	"fmt"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/miningpool/mw/v1/app/pool"
	pool "github.com/NpoolPlatform/miningpool-middleware/pkg/mw/app/pool"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) UpdatePool(ctx context.Context, in *npool.UpdatePoolRequest) (*npool.UpdatePoolResponse, error) {
	req := in.GetInfo()
	if req == nil {
		err := fmt.Errorf("request is nil")
		logger.Sugar().Errorw(
			"UpdatePool",
			"In", in,
			"Error", err,
		)
		return &npool.UpdatePoolResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}
	handler, err := pool.NewHandler(
		ctx,
		pool.WithID(req.ID, false),
		pool.WithEntID(req.EntID, false),
		pool.WithAppID(req.AppID, true),
		pool.WithPoolID(req.PoolID, true),
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
