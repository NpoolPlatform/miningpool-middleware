//nolint:nolintlint,dupl
package fractionwithdrawal

import (
	"context"
	"fmt"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/miningpool/mw/v1/fractionwithdrawal"
	fractionwithdrawal "github.com/NpoolPlatform/miningpool-middleware/pkg/mw/fractionwithdrawal"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) UpdateFractionWithdrawal(ctx context.Context, in *npool.UpdateFractionWithdrawalRequest) (*npool.UpdateFractionWithdrawalResponse, error) {
	if in.GetInfo() == nil {
		err := fmt.Errorf("request is nil")
		logger.Sugar().Errorw(
			"UpdateFractionWithdrawal",
			"In", in,
			"Error", err,
		)
		return &npool.UpdateFractionWithdrawalResponse{}, status.Error(codes.Internal, "internal server err")
	}
	req := in.GetInfo()

	handler, err := fractionwithdrawal.NewHandler(
		ctx,
		fractionwithdrawal.WithID(req.ID, false),
		fractionwithdrawal.WithEntID(req.EntID, false),
		fractionwithdrawal.WithMsg(req.Msg, false),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"UpdateFractionWithdrawal",
			"In", in,
			"Error", err,
		)
		return &npool.UpdateFractionWithdrawalResponse{}, status.Error(codes.Internal, "internal server err")
	}

	err = handler.UpdateFractionWithdrawal(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"UpdateFractionWithdrawal",
			"In", in,
			"Error", err,
		)
		return &npool.UpdateFractionWithdrawalResponse{}, status.Error(codes.Internal, "internal server err")
	}

	return &npool.UpdateFractionWithdrawalResponse{}, nil
}
