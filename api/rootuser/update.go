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

func (s *Server) UpdateRootUser(ctx context.Context, in *npool.UpdateRootUserRequest) (*npool.UpdateRootUserResponse, error) {
	if in.GetInfo() == nil {
		err := fmt.Errorf("request is nil")
		logger.Sugar().Errorw(
			"UpdateRootUser",
			"In", in,
			"Error", err,
		)
		return &npool.UpdateRootUserResponse{}, status.Error(codes.Internal, "internal server err")
	}

	req := in.GetInfo()
	handler, err := rootuser.NewHandler(
		ctx,
		rootuser.WithID(req.ID, false),
		rootuser.WithEntID(req.EntID, false),
		rootuser.WithName(req.Name, false),
		rootuser.WithEmail(req.Email, false),
		rootuser.WithAuthToken(req.AuthToken, false),
		rootuser.WithRemark(req.Remark, false),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"UpdateRootUser",
			"In", in,
			"Error", err,
		)
		return &npool.UpdateRootUserResponse{}, status.Error(codes.Internal, "internal server err")
	}

	err = handler.UpdateRootUser(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"UpdateRootUser",
			"In", in,
			"Error", err,
		)
		return &npool.UpdateRootUserResponse{}, status.Error(codes.Internal, "internal server err")
	}

	return &npool.UpdateRootUserResponse{}, nil
}
