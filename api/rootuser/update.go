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

func (s *Server) UpdateRootUser(ctx context.Context, in *npool.UpdateRootUserRequest) (*npool.UpdateRootUserResponse, error) {
	req := in.GetInfo()
	handler, err := rootuser.NewHandler(
		ctx,
		rootuser.WithID(req.ID, true),
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
			"UpdateRootUser",
			"In", in,
			"Error", err,
		)
		return &npool.UpdateRootUserResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := handler.UpdateRootUser(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"UpdateRootUser",
			"In", in,
			"Error", err,
		)
		return &npool.UpdateRootUserResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.UpdateRootUserResponse{
		Info: info,
	}, nil
}
