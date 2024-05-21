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

func (s *Server) GetFraction(ctx context.Context, in *npool.GetFractionRequest) (*npool.GetFractionResponse, error) {
	handler, err := fraction.NewHandler(
		ctx,
		fraction.WithEntID(&in.EntID, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetFraction",
			"In", in,
			"Error", err,
		)
		return &npool.GetFractionResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := handler.GetFraction(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetFraction",
			"In", in,
			"Error", err,
		)
		return &npool.GetFractionResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetFractionResponse{
		Info: info,
	}, nil
}

func (s *Server) GetFractions(ctx context.Context, in *npool.GetFractionsRequest) (*npool.GetFractionsResponse, error) {
	handler, err := fraction.NewHandler(
		ctx,
		fraction.WithConds(in.Conds),
		fraction.WithOffset(in.GetOffset()),
		fraction.WithLimit(in.GetLimit()),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetFractions",
			"In", in,
			"Error", err,
		)
		return &npool.GetFractionsResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	infos, total, err := handler.GetFractions(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetFractions",
			"In", in,
			"Error", err,
		)
		return &npool.GetFractionsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetFractionsResponse{
		Infos: infos,
		Total: total,
	}, nil
}
