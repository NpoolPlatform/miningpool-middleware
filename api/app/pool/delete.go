package pool

import (
	"context"
	"fmt"

	npool "github.com/NpoolPlatform/message/npool/miningpool/mw/v1/app/pool"
	pool "github.com/NpoolPlatform/miningpool-middleware/pkg/mw/app/pool"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) DeletePool(ctx context.Context, in *npool.DeletePoolRequest) (*npool.DeletePoolResponse, error) {
	if in.GetInfo() == nil {
		err := fmt.Errorf("request is nil")
		logger.Sugar().Errorw(
			"DeletePool",
			"In", in,
			"Error", err,
		)
		return &npool.DeletePoolResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	req := in.GetInfo()
	handler, err := pool.NewHandler(
		ctx,
		pool.WithID(req.ID, false),
		pool.WithEntID(req.EntID, false),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"DeletePool",
			"In", in,
			"Error", err,
		)
		return &npool.DeletePoolResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := handler.DeletePool(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"DeletePool",
			"In", in,
			"Error", err,
		)
		return &npool.DeletePoolResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.DeletePoolResponse{
		Info: info,
	}, nil
}
