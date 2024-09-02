package fractionwithdrawal

import (
	"context"
	"fmt"

	npool "github.com/NpoolPlatform/message/npool/miningpool/mw/v1/fractionwithdrawal"
	fractionwithdrawal "github.com/NpoolPlatform/miningpool-middleware/pkg/mw/fractionwithdrawal"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) DeleteFractionWithdrawal(ctx context.Context, in *npool.DeleteFractionWithdrawalRequest) (*npool.DeleteFractionWithdrawalResponse, error) {
	if in.GetInfo() == nil {
		err := fmt.Errorf("request is nil")
		logger.Sugar().Errorw(
			"DeleteFractionWithdrawal",
			"In", in,
			"Error", err,
		)
		return &npool.DeleteFractionWithdrawalResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	req := in.GetInfo()
	handler, err := fractionwithdrawal.NewHandler(
		ctx,
		fractionwithdrawal.WithID(req.ID, false),
		fractionwithdrawal.WithEntID(req.EntID, false),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"DeleteFractionWithdrawal",
			"In", in,
			"Error", err,
		)
		return &npool.DeleteFractionWithdrawalResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	err = handler.DeleteFractionWithdrawal(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"DeleteFractionWithdrawal",
			"In", in,
			"Error", err,
		)
		return &npool.DeleteFractionWithdrawalResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.DeleteFractionWithdrawalResponse{}, nil
}
