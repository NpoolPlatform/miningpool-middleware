//nolint:nolintlint,dupl
package orderuser

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/miningpool/mw/v1/orderuser"
	orderuser "github.com/NpoolPlatform/miningpool-middleware/pkg/mw/orderuser"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) ExistOrderUser(ctx context.Context, in *npool.ExistOrderUserRequest) (*npool.ExistOrderUserResponse, error) {
	handler, err := orderuser.NewHandler(
		ctx,
		orderuser.WithEntID(&in.EntID, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"ExistOrderUser",
			"In", in,
			"Error", err,
		)
		return &npool.ExistOrderUserResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := handler.ExistOrderUser(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"ExistOrderUser",
			"In", in,
			"Error", err,
		)
		return &npool.ExistOrderUserResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.ExistOrderUserResponse{
		Info: info,
	}, nil
}

func (s *Server) ExistOrderUserConds(ctx context.Context, in *npool.ExistOrderUserCondsRequest) (*npool.ExistOrderUserCondsResponse, error) {
	handler, err := orderuser.NewHandler(
		ctx,
		orderuser.WithConds(in.Conds),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"ExistOrderUserConds",
			"In", in,
			"Error", err,
		)
		return &npool.ExistOrderUserCondsResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := handler.ExistOrderUserConds(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"ExistOrderUserConds",
			"In", in,
			"Error", err,
		)
		return &npool.ExistOrderUserCondsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.ExistOrderUserCondsResponse{
		Info: info,
	}, nil
}
