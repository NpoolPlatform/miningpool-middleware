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
		return &npool.DeleteOrderUserResponse{}, status.Error(codes.InvalidArgument, err.Error())
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
		return &npool.DeleteOrderUserResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	zeroProportion := "0"
	_req := &npool.OrderUserReq{
		ID:         req.ID,
		EntID:      req.EntID,
		Proportion: &zeroProportion,
	}
	_req.Proportion = &zeroProportion
	_, err = handleUpdateReq(ctx, _req)
	if err != nil {
		logger.Sugar().Errorw(
			"DeleteOrderUser",
			"In", in,
			"Error", err,
		)
		return &npool.DeleteOrderUserResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := handler.DeleteOrderUser(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"DeleteOrderUser",
			"In", in,
			"Error", err,
		)
		return &npool.DeleteOrderUserResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.DeleteOrderUserResponse{
		Info: info,
	}, nil
}
