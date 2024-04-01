package rootuser

import (
	"context"
	"fmt"

	npool "github.com/NpoolPlatform/message/npool/miningpool/mw/v1/rootuser"
	rootuser "github.com/NpoolPlatform/miningpool-middleware/pkg/mw/rootuser"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) DeleteRootUser(ctx context.Context, in *npool.DeleteRootUserRequest) (*npool.DeleteRootUserResponse, error) {
	req := in.GetInfo()
	if req == nil || req.ID == nil {
		err := fmt.Errorf("wrong id")
		return &npool.DeleteRootUserResponse{}, status.Error(codes.Internal, err.Error())
	}

	handler, err := rootuser.NewHandler(
		ctx,
		rootuser.WithID(req.ID, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"DeleteRootUser",
			"In", in,
			"Error", err,
		)
		return &npool.DeleteRootUserResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := handler.DeleteRootUser(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"DeleteRootUser",
			"In", in,
			"Error", err,
		)
		return &npool.DeleteRootUserResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.DeleteRootUserResponse{
		Info: info,
	}, nil
}
