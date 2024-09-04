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
	if in.GetInfo() == nil {
		err := fmt.Errorf("request is nil")
		logger.Sugar().Errorw(
			"DeleteRootUser",
			"In", in,
			"Error", err,
		)
		return &npool.DeleteRootUserResponse{}, status.Error(codes.Internal, "internal server err")
	}

	req := in.GetInfo()
	handler, err := rootuser.NewHandler(
		ctx,
		rootuser.WithID(req.ID, false),
		rootuser.WithEntID(req.EntID, false),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"DeleteRootUser",
			"In", in,
			"Error", err,
		)
		return &npool.DeleteRootUserResponse{}, status.Error(codes.Internal, "internal server err")
	}

	err = handler.DeleteRootUser(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"DeleteRootUser",
			"In", in,
			"Error", err,
		)
		return &npool.DeleteRootUserResponse{}, status.Error(codes.Internal, "internal server err")
	}

	return &npool.DeleteRootUserResponse{}, nil
}
