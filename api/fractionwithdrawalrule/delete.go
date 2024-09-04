package fractionwithdrawalrule

import (
	"context"
	"fmt"

	npool "github.com/NpoolPlatform/message/npool/miningpool/mw/v1/fractionwithdrawalrule"
	fractionwithdrawalrule "github.com/NpoolPlatform/miningpool-middleware/pkg/mw/fractionwithdrawalrule"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) DeleteFractionWithdrawalRule(ctx context.Context, in *npool.DeleteFractionWithdrawalRuleRequest) (*npool.DeleteFractionWithdrawalRuleResponse, error) {
	if in.GetInfo() == nil {
		err := fmt.Errorf("request is nil")
		logger.Sugar().Errorw(
			"DeleteFractionWithdrawalRule",
			"In", in,
			"Error", err,
		)
		return &npool.DeleteFractionWithdrawalRuleResponse{}, status.Error(codes.Internal, "internal server err")
	}

	req := in.GetInfo()
	handler, err := fractionwithdrawalrule.NewHandler(
		ctx,
		fractionwithdrawalrule.WithID(req.ID, false),
		fractionwithdrawalrule.WithEntID(req.EntID, false),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"DeleteFractionWithdrawalRule",
			"In", in,
			"Error", err,
		)
		return &npool.DeleteFractionWithdrawalRuleResponse{}, status.Error(codes.Internal, "internal server err")
	}

	err = handler.DeleteFractionWithdrawalRule(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"DeleteFractionWithdrawalRule",
			"In", in,
			"Error", err,
		)
		return &npool.DeleteFractionWithdrawalRuleResponse{}, status.Error(codes.Internal, "internal server err")
	}

	return &npool.DeleteFractionWithdrawalRuleResponse{}, nil
}
