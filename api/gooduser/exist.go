//nolint:nolintlint,dupl
package gooduser

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/miningpool/mw/v1/gooduser"
	gooduser "github.com/NpoolPlatform/miningpool-middleware/pkg/mw/gooduser"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) ExistGoodUser(ctx context.Context, in *npool.ExistGoodUserRequest) (*npool.ExistGoodUserResponse, error) {
	handler, err := gooduser.NewHandler(
		ctx,
		gooduser.WithEntID(&in.EntID, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"ExistGoodUser",
			"In", in,
			"Error", err,
		)
		return &npool.ExistGoodUserResponse{}, status.Error(codes.Internal, "internal server err")
	}

	info, err := handler.ExistGoodUser(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"ExistGoodUser",
			"In", in,
			"Error", err,
		)
		return &npool.ExistGoodUserResponse{}, status.Error(codes.Internal, "internal server err")
	}

	return &npool.ExistGoodUserResponse{
		Info: info,
	}, nil
}

func (s *Server) ExistGoodUserConds(ctx context.Context, in *npool.ExistGoodUserCondsRequest) (*npool.ExistGoodUserCondsResponse, error) {
	handler, err := gooduser.NewHandler(
		ctx,
		gooduser.WithConds(in.Conds),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"ExistGoodUserConds",
			"In", in,
			"Error", err,
		)
		return &npool.ExistGoodUserCondsResponse{}, status.Error(codes.Internal, "internal server err")
	}

	info, err := handler.ExistGoodUserConds(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"ExistGoodUserConds",
			"In", in,
			"Error", err,
		)
		return &npool.ExistGoodUserCondsResponse{}, status.Error(codes.Internal, "internal server err")
	}

	return &npool.ExistGoodUserCondsResponse{
		Info: info,
	}, nil
}
