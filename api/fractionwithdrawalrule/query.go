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

func (s *Server) GetFractionWithdrawalRule(ctx context.Context, in *npool.GetFractionWithdrawalRuleRequest) (*npool.GetFractionWithdrawalRuleResponse, error) {
	handler, err := fractionwithdrawalrule.NewHandler(
		ctx,
		fractionwithdrawalrule.WithEntID(&in.EntID, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetFractionWithdrawalRule",
			"In", in,
			"Error", err,
		)
		return &npool.GetFractionWithdrawalRuleResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := handler.GetFractionWithdrawalRule(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetFractionWithdrawalRule",
			"In", in,
			"Error", err,
		)
		return &npool.GetFractionWithdrawalRuleResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetFractionWithdrawalRuleResponse{
		Info: info,
	}, nil
}

func (s *Server) GetFractionWithdrawalRules(ctx context.Context, in *npool.GetFractionWithdrawalRulesRequest) (*npool.GetFractionWithdrawalRulesResponse, error) {
	handler, err := fractionwithdrawalrule.NewHandler(
		ctx,
		fractionwithdrawalrule.WithConds(in.Conds),
		fractionwithdrawalrule.WithOffset(in.GetOffset()),
		fractionwithdrawalrule.WithLimit(in.GetLimit()),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetFractionWithdrawalRules",
			"In", in,
			"Error", err,
		)
		return &npool.GetFractionWithdrawalRulesResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	infos, total, err := handler.GetFractionWithdrawalRules(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetFractionWithdrawalRules",
			"In", in,
			"Error", err,
		)
		return &npool.GetFractionWithdrawalRulesResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetFractionWithdrawalRulesResponse{
		Infos: infos,
		Total: total,
	}, nil
}
