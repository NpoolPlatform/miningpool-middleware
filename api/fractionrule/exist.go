//nolint:nolintlint,dupl
package fractionrule

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/miningpool/mw/v1/fractionrule"
	fractionrule "github.com/NpoolPlatform/miningpool-middleware/pkg/mw/fractionrule"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) ExistFractionRule(ctx context.Context, in *npool.ExistFractionRuleRequest) (*npool.ExistFractionRuleResponse, error) {
	handler, err := fractionrule.NewHandler(
		ctx,
		fractionrule.WithEntID(&in.EntID, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"ExistFractionRule",
			"In", in,
			"Error", err,
		)
		return &npool.ExistFractionRuleResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := handler.ExistFractionRule(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"ExistFractionRule",
			"In", in,
			"Error", err,
		)
		return &npool.ExistFractionRuleResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.ExistFractionRuleResponse{
		Info: info,
	}, nil
}

func (s *Server) ExistFractionRuleConds(ctx context.Context, in *npool.ExistFractionRuleCondsRequest) (*npool.ExistFractionRuleCondsResponse, error) {
	handler, err := fractionrule.NewHandler(
		ctx,
		fractionrule.WithConds(in.Conds),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"ExistFractionRuleConds",
			"In", in,
			"Error", err,
		)
		return &npool.ExistFractionRuleCondsResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := handler.ExistFractionRuleConds(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"ExistFractionRuleConds",
			"In", in,
			"Error", err,
		)
		return &npool.ExistFractionRuleCondsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.ExistFractionRuleCondsResponse{
		Info: info,
	}, nil
}
