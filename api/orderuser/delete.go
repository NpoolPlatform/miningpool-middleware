package orderuser

import (
	"context"
	"fmt"

	npool "github.com/NpoolPlatform/message/npool/miningpool/mw/v1/orderuser"
	orderuser "github.com/NpoolPlatform/miningpool-middleware/pkg/mw/orderuser"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) DeleteOrderUser(ctx context.Context, in *npool.DeleteOrderUserRequest) (*npool.DeleteOrderUserResponse, error) {
	if in.GetInfo() == nil {
		err := fmt.Errorf("request is nil")
		logger.Sugar().Errorw(
			"DeleteOrderUser",
			"In", in,
			"Error", err,
		)
		return &npool.DeleteOrderUserResponse{}, status.Error(codes.Internal, "internal server err")
	}

	req := in.GetInfo()
	handler, err := orderuser.NewHandler(
		ctx,
		orderuser.WithID(req.ID, false),
		orderuser.WithEntID(req.EntID, false),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"DeleteOrderUser",
			"In", in,
			"Error", err,
		)
		return &npool.DeleteOrderUserResponse{}, status.Error(codes.Internal, "internal server err")
	}

	err = handler.DeleteOrderUser(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"DeleteOrderUser",
			"In", in,
			"Error", err,
		)
		return &npool.DeleteOrderUserResponse{}, status.Error(codes.Internal, "internal server err")
	}

	return &npool.DeleteOrderUserResponse{}, nil
}
