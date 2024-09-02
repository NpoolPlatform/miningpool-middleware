//nolint:nolintlint,dupl
package fractionwithdrawal

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/miningpool/mw/v1/fractionwithdrawal"
	fractionwithdrawal "github.com/NpoolPlatform/miningpool-middleware/pkg/mw/fractionwithdrawal"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) GetFractionWithdrawal(ctx context.Context, in *npool.GetFractionWithdrawalRequest) (*npool.GetFractionWithdrawalResponse, error) {
	handler, err := fractionwithdrawal.NewHandler(
		ctx,
		fractionwithdrawal.WithEntID(&in.EntID, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetFractionWithdrawal",
			"In", in,
			"Error", err,
		)
		return &npool.GetFractionWithdrawalResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := handler.GetFractionWithdrawal(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetFractionWithdrawal",
			"In", in,
			"Error", err,
		)
		return &npool.GetFractionWithdrawalResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetFractionWithdrawalResponse{
		Info: info,
	}, nil
}

func (s *Server) GetFractionWithdrawals(ctx context.Context, in *npool.GetFractionWithdrawalsRequest) (*npool.GetFractionWithdrawalsResponse, error) {
	handler, err := fractionwithdrawal.NewHandler(
		ctx,
		fractionwithdrawal.WithConds(in.Conds),
		fractionwithdrawal.WithOffset(in.GetOffset()),
		fractionwithdrawal.WithLimit(in.GetLimit()),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetFractionWithdrawals",
			"In", in,
			"Error", err,
		)
		return &npool.GetFractionWithdrawalsResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	infos, total, err := handler.GetFractionWithdrawals(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetFractionWithdrawals",
			"In", in,
			"Error", err,
		)
		return &npool.GetFractionWithdrawalsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetFractionWithdrawalsResponse{
		Infos: infos,
		Total: total,
	}, nil
}
