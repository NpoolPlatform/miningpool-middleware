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

func (s *Server) CreateFractionRule(ctx context.Context, in *npool.CreateFractionRuleRequest) (*npool.CreateFractionRuleResponse, error) {
	req := in.GetInfo()
	handler, err := fractionrule.NewHandler(
		ctx,
		fractionrule.WithID(req.ID, false),
		fractionrule.WithEntID(req.EntID, false),
		fractionrule.WithMiningpoolType(req.MiningpoolType, true),
		fractionrule.WithCoinType(req.CoinType, true),
		fractionrule.WithWithdrawInterval(req.WithdrawInterval, true),
		fractionrule.WithMinAmount(req.MinAmount, true),
		fractionrule.WithWithdrawRate(req.WithdrawRate, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateFractionRule",
			"In", in,
			"Error", err,
		)
		return &npool.CreateFractionRuleResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := handler.CreateFractionRule(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateFractionRule",
			"In", in,
			"Error", err,
		)
		return &npool.CreateFractionRuleResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CreateFractionRuleResponse{
		Info: info,
	}, nil
}

func (s *Server) CreateFractionRules(ctx context.Context, in *npool.CreateFractionRulesRequest) (*npool.CreateFractionRulesResponse, error) {
	handler, err := fractionrule.NewHandler(
		ctx,
		fractionrule.WithReqs(in.GetInfos(), true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateFractionRules",
			"In", in,
			"Error", err,
		)
		return &npool.CreateFractionRulesResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	infos, err := handler.CreateFractionRules(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateFractionRules",
			"In", in,
			"Error", err,
		)
		return &npool.CreateFractionRulesResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CreateFractionRulesResponse{
		Infos: infos,
	}, nil
}
