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

func (s *Server) CreateFractionWithdrawal(ctx context.Context, in *npool.CreateFractionWithdrawalRequest) (*npool.CreateFractionWithdrawalResponse, error) {
	if in.GetInfo() == nil {
		err := fmt.Errorf("request is nil")
		logger.Sugar().Errorw(
			"CreateFractionWithdrawal",
			"In", in,
			"Error", err,
		)
		return &npool.CreateFractionWithdrawalResponse{}, status.Error(codes.Internal, "internal server err")
	}

	req := in.GetInfo()
	handler, err := fractionwithdrawal.NewHandler(
		ctx,
		fractionwithdrawal.WithEntID(req.EntID, false),
		fractionwithdrawal.WithAppID(req.AppID, true),
		fractionwithdrawal.WithUserID(req.UserID, true),
		fractionwithdrawal.WithOrderUserID(req.OrderUserID, true),
		fractionwithdrawal.WithCoinTypeID(req.CoinTypeID, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateFractionWithdrawal",
			"In", in,
			"Error", err,
		)
		return &npool.CreateFractionWithdrawalResponse{}, status.Error(codes.Internal, "internal server err")
	}

	err = handler.CreateFractionWithdrawal(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateFractionWithdrawal",
			"In", in,
			"Error", err,
		)
		return &npool.CreateFractionWithdrawalResponse{}, status.Error(codes.Internal, "internal server err")
	}

	return &npool.CreateFractionWithdrawalResponse{}, nil
}
