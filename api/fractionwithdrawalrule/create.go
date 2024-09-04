//nolint:nolintlint,dupl
package fractionwithdrawalrule

import (
	"context"
	"fmt"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/miningpool/mw/v1/fractionwithdrawalrule"
	fractionwithdrawalrule "github.com/NpoolPlatform/miningpool-middleware/pkg/mw/fractionwithdrawalrule"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) CreateFractionWithdrawalRule(ctx context.Context, in *npool.CreateFractionWithdrawalRuleRequest) (*npool.CreateFractionWithdrawalRuleResponse, error) {
	if in.GetInfo() == nil {
		err := fmt.Errorf("request is nil")
		logger.Sugar().Errorw(
			"CreateFractionWithdrawalRule",
			"In", in,
			"Error", err,
		)
		return &npool.CreateFractionWithdrawalRuleResponse{}, status.Error(codes.Internal, "internal server err")
	}
	req := in.GetInfo()
	handler, err := fractionwithdrawalrule.NewHandler(
		ctx,
		fractionwithdrawalrule.WithEntID(req.EntID, false),
		fractionwithdrawalrule.WithPoolCoinTypeID(req.PoolCoinTypeID, true),
		fractionwithdrawalrule.WithWithdrawInterval(req.WithdrawInterval, true),
		fractionwithdrawalrule.WithPayoutThreshold(req.PayoutThreshold, true),
		fractionwithdrawalrule.WithLeastWithdrawalAmount(req.LeastWithdrawalAmount, true),
		fractionwithdrawalrule.WithWithdrawFee(req.WithdrawFee, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateFractionWithdrawalRule",
			"In", in,
			"Error", err,
		)
		return &npool.CreateFractionWithdrawalRuleResponse{}, status.Error(codes.Internal, "internal server err")
	}

	err = handler.CreateFractionWithdrawalRule(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateFractionWithdrawalRule",
			"In", in,
			"Error", err,
		)
		return &npool.CreateFractionWithdrawalRuleResponse{}, status.Error(codes.Internal, "internal server err")
	}

	return &npool.CreateFractionWithdrawalRuleResponse{}, nil
}
