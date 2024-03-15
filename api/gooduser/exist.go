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
		return &npool.ExistGoodUserCondsResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := handler.ExistGoodUserConds(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"ExistGoodUserConds",
			"In", in,
			"Error", err,
		)
		return &npool.ExistGoodUserCondsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.ExistGoodUserCondsResponse{
		Info: info,
	}, nil
}
