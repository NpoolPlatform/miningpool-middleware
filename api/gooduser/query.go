//nolint:nolintlint,dupl
package gooduser

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/miningpool/mw/v1/gooduser"
	gooduser "github.com/NpoolPlatform/miningpool-middleware/pkg/mw/gooduser"
	"github.com/NpoolPlatform/miningpool-middleware/pkg/utils"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) GetGoodUser(ctx context.Context, in *npool.GetGoodUserRequest) (*npool.GetGoodUserResponse, error) {
	handler, err := gooduser.NewHandler(
		ctx,
		gooduser.WithEntID(&in.EntID, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetGoodUser",
			"In", in,
			"Error", err,
		)
		return &npool.GetGoodUserResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := handler.GetGoodUser(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetGoodUser",
			"In", in,
			"Error", err,
		)
		return &npool.GetGoodUserResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetGoodUserResponse{
		Info: info,
	}, nil
}

func (s *Server) GetGoodUsers(ctx context.Context, in *npool.GetGoodUsersRequest) (*npool.GetGoodUsersResponse, error) {
	handler, err := gooduser.NewHandler(
		ctx,
		gooduser.WithConds(in.Conds),
		gooduser.WithOffset(in.GetOffset()),
		gooduser.WithLimit(in.GetLimit()),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetGoodUsers",
			"In", in,
			"Error", err,
		)
		return &npool.GetGoodUsersResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}
	logger.Sugar().Error(utils.PrettyStruct(handler))
	infos, total, err := handler.GetGoodUsers(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetGoodUsers",
			"In", in,
			"Error", err,
		)
		return &npool.GetGoodUsersResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetGoodUsersResponse{
		Infos: infos,
		Total: total,
	}, nil
}
