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

func (s *Server) ExistFractionWithdrawal(ctx context.Context, in *npool.ExistFractionWithdrawalRequest) (*npool.ExistFractionWithdrawalResponse, error) {
	handler, err := fractionwithdrawal.NewHandler(
		ctx,
		fractionwithdrawal.WithEntID(&in.EntID, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"ExistFractionWithdrawal",
			"In", in,
			"Error", err,
		)
		return &npool.ExistFractionWithdrawalResponse{}, status.Error(codes.Internal, "internal server err")
	}

	info, err := handler.ExistFractionWithdrawal(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"ExistFractionWithdrawal",
			"In", in,
			"Error", err,
		)
		return &npool.ExistFractionWithdrawalResponse{}, status.Error(codes.Internal, "internal server err")
	}

	return &npool.ExistFractionWithdrawalResponse{
		Info: info,
	}, nil
}

func (s *Server) ExistFractionWithdrawalConds(ctx context.Context, in *npool.ExistFractionWithdrawalCondsRequest) (*npool.ExistFractionWithdrawalCondsResponse, error) {
	handler, err := fractionwithdrawal.NewHandler(
		ctx,
		fractionwithdrawal.WithConds(in.Conds),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"ExistFractionWithdrawalConds",
			"In", in,
			"Error", err,
		)
		return &npool.ExistFractionWithdrawalCondsResponse{}, status.Error(codes.Internal, "internal server err")
	}

	info, err := handler.ExistFractionWithdrawalConds(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"ExistFractionWithdrawalConds",
			"In", in,
			"Error", err,
		)
		return &npool.ExistFractionWithdrawalCondsResponse{}, status.Error(codes.Internal, "internal server err")
	}

	return &npool.ExistFractionWithdrawalCondsResponse{
		Info: info,
	}, nil
}
