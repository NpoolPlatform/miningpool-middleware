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

func (s *Server) GetFractionRule(ctx context.Context, in *npool.GetFractionRuleRequest) (*npool.GetFractionRuleResponse, error) {
	handler, err := fractionrule.NewHandler(
		ctx,
		fractionrule.WithEntID(&in.EntID, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetFractionRule",
			"In", in,
			"Error", err,
		)
		return &npool.GetFractionRuleResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := handler.GetFractionRule(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetFractionRule",
			"In", in,
			"Error", err,
		)
		return &npool.GetFractionRuleResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetFractionRuleResponse{
		Info: info,
	}, nil
}

func (s *Server) GetFractionRules(ctx context.Context, in *npool.GetFractionRulesRequest) (*npool.GetFractionRulesResponse, error) {
	handler, err := fractionrule.NewHandler(
		ctx,
		fractionrule.WithConds(in.Conds),
		fractionrule.WithOffset(in.GetOffset()),
		fractionrule.WithLimit(in.GetLimit()),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetFractionRules",
			"In", in,
			"Error", err,
		)
		return &npool.GetFractionRulesResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	infos, total, err := handler.GetFractionRules(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetFractionRules",
			"In", in,
			"Error", err,
		)
		return &npool.GetFractionRulesResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetFractionRulesResponse{
		Infos: infos,
		Total: total,
	}, nil
}
