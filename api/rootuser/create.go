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

func (s *Server) CreateRootUser(ctx context.Context, in *npool.CreateRootUserRequest) (*npool.CreateRootUserResponse, error) {
	req := in.GetInfo()
	handler, err := rootuser.NewHandler(
		ctx,
		rootuser.WithID(req.ID, false),
		rootuser.WithEntID(req.EntID, false),
		rootuser.WithName(req.Name, true),
		rootuser.WithMiningpoolType(req.MiningpoolType, true),
		rootuser.WithEmail(req.Email, true),
		rootuser.WithAuthToken(req.AuthToken, true),
		rootuser.WithAuthed(req.Authed, true),
		rootuser.WithRemark(req.Remark, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateRootUser",
			"In", in,
			"Error", err,
		)
		return &npool.CreateRootUserResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := handler.CreateRootUser(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateRootUser",
			"In", in,
			"Error", err,
		)
		return &npool.CreateRootUserResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CreateRootUserResponse{
		Info: info,
	}, nil
}

func (s *Server) CreateRootUsers(ctx context.Context, in *npool.CreateRootUsersRequest) (*npool.CreateRootUsersResponse, error) {
	handler, err := rootuser.NewHandler(
		ctx,
		rootuser.WithReqs(in.GetInfos(), true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateRootUsers",
			"In", in,
			"Error", err,
		)
		return &npool.CreateRootUsersResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	infos, err := handler.CreateRootUsers(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateRootUsers",
			"In", in,
			"Error", err,
		)
		return &npool.CreateRootUsersResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CreateRootUsersResponse{
		Infos: infos,
	}, nil
}
