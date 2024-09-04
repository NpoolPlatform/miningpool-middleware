//nolint:nolintlint,dupl
package fractionwithdrawalrule

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/miningpool/mw/v1/fractionwithdrawalrule"
	fractionwithdrawalrule "github.com/NpoolPlatform/miningpool-middleware/pkg/mw/fractionwithdrawalrule"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) ExistFractionWithdrawalRule(ctx context.Context, in *npool.ExistFractionWithdrawalRuleRequest) (*npool.ExistFractionWithdrawalRuleResponse, error) {
	handler, err := fractionwithdrawalrule.NewHandler(
		ctx,
		fractionwithdrawalrule.WithEntID(&in.EntID, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"ExistFractionWithdrawalRule",
			"In", in,
			"Error", err,
		)
		return &npool.ExistFractionWithdrawalRuleResponse{}, status.Error(codes.Internal, "internal server err")
	}

	info, err := handler.ExistFractionWithdrawalRule(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"ExistFractionWithdrawalRule",
			"In", in,
			"Error", err,
		)
		return &npool.ExistFractionWithdrawalRuleResponse{}, status.Error(codes.Internal, "internal server err")
	}

	return &npool.ExistFractionWithdrawalRuleResponse{
		Info: info,
	}, nil
}

func (s *Server) ExistFractionWithdrawalRuleConds(ctx context.Context, in *npool.ExistFractionWithdrawalRuleCondsRequest) (*npool.ExistFractionWithdrawalRuleCondsResponse, error) {
	handler, err := fractionwithdrawalrule.NewHandler(
		ctx,
		fractionwithdrawalrule.WithConds(in.Conds),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"ExistFractionWithdrawalRuleConds",
			"In", in,
			"Error", err,
		)
		return &npool.ExistFractionWithdrawalRuleCondsResponse{}, status.Error(codes.Internal, "internal server err")
	}

	info, err := handler.ExistFractionWithdrawalRuleConds(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"ExistFractionWithdrawalRuleConds",
			"In", in,
			"Error", err,
		)
		return &npool.ExistFractionWithdrawalRuleCondsResponse{}, status.Error(codes.Internal, "internal server err")
	}

	return &npool.ExistFractionWithdrawalRuleCondsResponse{
		Info: info,
	}, nil
}
