//nolint:nolintlint,dupl
package fraction

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/miningpool/mw/v1/fraction"
	fraction "github.com/NpoolPlatform/miningpool-middleware/pkg/mw/fraction"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) ExistFractionConds(ctx context.Context, in *npool.ExistFractionCondsRequest) (*npool.ExistFractionCondsResponse, error) {
	handler, err := fraction.NewHandler(
		ctx,
		fraction.WithConds(in.Conds),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"ExistFractionConds",
			"In", in,
			"Error", err,
		)
		return &npool.ExistFractionCondsResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := handler.ExistFractionConds(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"ExistFractionConds",
			"In", in,
			"Error", err,
		)
		return &npool.ExistFractionCondsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.ExistFractionCondsResponse{
		Info: info,
	}, nil
}
