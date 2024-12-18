//nolint:nolintlint,dupl
package rootuser

import (
	"context"
	"fmt"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/miningpool/mw/v1/rootuser"
	rootuser "github.com/NpoolPlatform/miningpool-middleware/pkg/mw/rootuser"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) CreateRootUser(ctx context.Context, in *npool.CreateRootUserRequest) (*npool.CreateRootUserResponse, error) {
	if in.GetInfo() == nil {
		err := fmt.Errorf("request is nil")
		logger.Sugar().Errorw(
			"CreateRootUser",
			"In", in,
			"Error", err,
		)
		return &npool.CreateRootUserResponse{}, status.Error(codes.Internal, "internal server err")
	}

	req := in.GetInfo()
	handler, err := rootuser.NewHandler(
		ctx,
		rootuser.WithEntID(req.EntID, false),
		rootuser.WithPoolID(req.PoolID, true),
		rootuser.WithName(req.Name, true),
		rootuser.WithEmail(req.Email, true),
		rootuser.WithAuthToken(req.AuthToken, true),
		rootuser.WithRemark(req.Remark, false),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateRootUser",
			"In", in,
			"Error", err,
		)
		return &npool.CreateRootUserResponse{}, status.Error(codes.Internal, "internal server err")
	}

	err = handler.CreateRootUser(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateRootUser",
			"In", in,
			"Error", err,
		)
		return &npool.CreateRootUserResponse{}, status.Error(codes.Internal, "internal server err")
	}

	return &npool.CreateRootUserResponse{}, nil
}
