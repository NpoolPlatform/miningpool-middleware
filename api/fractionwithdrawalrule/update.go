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

func (s *Server) UpdateFractionWithdrawalRule(ctx context.Context, in *npool.UpdateFractionWithdrawalRuleRequest) (*npool.UpdateFractionWithdrawalRuleResponse, error) {
	if in.GetInfo() == nil {
		err := fmt.Errorf("request is nil")
		logger.Sugar().Errorw(
			"UpdateFractionWithdrawalRule",
			"In", in,
			"Error", err,
		)
		return &npool.UpdateFractionWithdrawalRuleResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	req := in.GetInfo()
	handler, err := fractionwithdrawalrule.NewHandler(
		ctx,
		fractionwithdrawalrule.WithID(req.ID, false),
		fractionwithdrawalrule.WithEntID(req.EntID, false),
		fractionwithdrawalrule.WithWithdrawInterval(req.WithdrawInterval, false),
		fractionwithdrawalrule.WithPayoutThreshold(req.PayoutThreshold, false),
		fractionwithdrawalrule.WithLeastWithdrawalAmount(req.LeastWithdrawalAmount, false),
		fractionwithdrawalrule.WithWithdrawFee(req.WithdrawFee, false),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"UpdateFractionWithdrawalRule",
			"In", in,
			"Error", err,
		)
		return &npool.UpdateFractionWithdrawalRuleResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	err = handler.UpdateFractionWithdrawalRule(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"UpdateFractionWithdrawalRule",
			"In", in,
			"Error", err,
		)
		return &npool.UpdateFractionWithdrawalRuleResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.UpdateFractionWithdrawalRuleResponse{}, nil
}
