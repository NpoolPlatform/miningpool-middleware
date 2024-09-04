//nolint:nolintlint,dupl
package rootuser

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/miningpool/mw/v1/rootuser"
	rootuser "github.com/NpoolPlatform/miningpool-middleware/pkg/mw/rootuser"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) GetRootUser(ctx context.Context, in *npool.GetRootUserRequest) (*npool.GetRootUserResponse, error) {
	handler, err := rootuser.NewHandler(
		ctx,
		rootuser.WithEntID(&in.EntID, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetRootUser",
			"In", in,
			"Error", err,
		)
		return &npool.GetRootUserResponse{}, status.Error(codes.Internal, "internal server err")
	}

	info, err := handler.GetRootUser(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetRootUser",
			"In", in,
			"Error", err,
		)
		return &npool.GetRootUserResponse{}, status.Error(codes.Internal, "internal server err")
	}

	return &npool.GetRootUserResponse{
		Info: info,
	}, nil
}

func (s *Server) GetRootUsers(ctx context.Context, in *npool.GetRootUsersRequest) (*npool.GetRootUsersResponse, error) {
	handler, err := rootuser.NewHandler(
		ctx,
		rootuser.WithConds(in.Conds),
		rootuser.WithOffset(in.GetOffset()),
		rootuser.WithLimit(in.GetLimit()),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetRootUsers",
			"In", in,
			"Error", err,
		)
		return &npool.GetRootUsersResponse{}, status.Error(codes.Internal, "internal server err")
	}

	infos, total, err := handler.GetRootUsers(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetRootUsers",
			"In", in,
			"Error", err,
		)
		return &npool.GetRootUsersResponse{}, status.Error(codes.Internal, "internal server err")
	}

	return &npool.GetRootUsersResponse{
		Infos: infos,
		Total: total,
	}, nil
}
