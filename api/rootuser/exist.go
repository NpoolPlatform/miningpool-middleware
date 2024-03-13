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

func (s *Server) ExistRootUserConds(ctx context.Context, in *npool.ExistRootUserCondsRequest) (*npool.ExistRootUserCondsResponse, error) {
	handler, err := rootuser.NewHandler(
		ctx,
		rootuser.WithConds(in.Conds),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"ExistRootUserConds",
			"In", in,
			"Error", err,
		)
		return &npool.ExistRootUserCondsResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := handler.ExistRootUserConds(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"ExistRootUserConds",
			"In", in,
			"Error", err,
		)
		return &npool.ExistRootUserCondsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.ExistRootUserCondsResponse{
		Info: info,
	}, nil
}
