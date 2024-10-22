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

func (s *Server) CreatePool(ctx context.Context, in *npool.CreatePoolRequest) (*npool.CreatePoolResponse, error) {
	if in.GetInfo() == nil {
		err := fmt.Errorf("request is nil")
		logger.Sugar().Errorw(
			"CreatePool",
			"In", in,
			"Error", err,
		)
		return &npool.CreatePoolResponse{}, status.Error(codes.Internal, "internal server err")
	}

	req := in.GetInfo()
	handler, err := pool.NewHandler(
		ctx,
		pool.WithEntID(req.EntID, false),
		pool.WithAppID(req.AppID, true),
		pool.WithPoolID(req.PoolID, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"CreatePool",
			"In", in,
			"Error", err,
		)
		return &npool.CreatePoolResponse{}, status.Error(codes.Internal, "internal server err")
	}

	err = handler.CreatePool(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"CreatePool",
			"In", in,
			"Error", err,
		)
		return &npool.CreatePoolResponse{}, status.Error(codes.Internal, "internal server err")
	}

	return &npool.CreatePoolResponse{}, nil
}
