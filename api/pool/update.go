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
		return &npool.UpdatePoolResponse{}, status.Error(codes.Internal, "internal server err")
	}

	req := in.GetInfo()
	handler, err := pool.NewHandler(
		ctx,
		pool.WithID(req.ID, false),
		pool.WithEntID(req.EntID, false),
		pool.WithName(req.Name, false),
		pool.WithSite(req.Site, false),
		pool.WithLogo(req.Logo, false),
		pool.WithDescription(req.Description, false),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"UpdatePool",
			"In", in,
			"Error", err,
		)
		return &npool.UpdatePoolResponse{}, status.Error(codes.Internal, "internal server err")
	}

	err = handler.UpdatePool(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"UpdatePool",
			"In", in,
			"Error", err,
		)
		return &npool.UpdatePoolResponse{}, status.Error(codes.Internal, "internal server err")
	}

	return &npool.UpdatePoolResponse{}, nil
}
