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

func (s *Server) CreateFraction(ctx context.Context, in *npool.CreateFractionRequest) (*npool.CreateFractionResponse, error) {
	req := in.GetInfo()
	handler, err := fraction.NewHandler(
		ctx,
		fraction.WithID(req.ID, false),
		fraction.WithEntID(req.EntID, false),
		fraction.WithOrderUserID(req.OrderUserID, true),
		fraction.WithWithdrawState(req.WithdrawState, true),
		fraction.WithWithdrawTime(req.WithdrawTime, true),
		fraction.WithPayTime(req.PayTime, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateFraction",
			"In", in,
			"Error", err,
		)
		return &npool.CreateFractionResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := handler.CreateFraction(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateFraction",
			"In", in,
			"Error", err,
		)
		return &npool.CreateFractionResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CreateFractionResponse{
		Info: info,
	}, nil
}

func (s *Server) CreateFractions(ctx context.Context, in *npool.CreateFractionsRequest) (*npool.CreateFractionsResponse, error) {
	handler, err := fraction.NewHandler(
		ctx,
		fraction.WithReqs(in.GetInfos(), true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateFractions",
			"In", in,
			"Error", err,
		)
		return &npool.CreateFractionsResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	infos, err := handler.CreateFractions(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateFractions",
			"In", in,
			"Error", err,
		)
		return &npool.CreateFractionsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CreateFractionsResponse{
		Infos: infos,
	}, nil
}
