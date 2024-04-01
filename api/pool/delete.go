package pool

import (
	"context"
	"fmt"

	npool "github.com/NpoolPlatform/message/npool/miningpool/mw/v1/pool"
	pool "github.com/NpoolPlatform/miningpool-middleware/pkg/mw/pool"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) DeletePool(ctx context.Context, in *npool.DeletePoolRequest) (*npool.DeletePoolResponse, error) {
	req := in.GetInfo()
	if req == nil || req.ID == nil {
		err := fmt.Errorf("wrong id")
		return &npool.DeletePoolResponse{}, status.Error(codes.Internal, err.Error())
	}

	handler, err := pool.NewHandler(
		ctx,
		pool.WithID(req.ID, true),
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
