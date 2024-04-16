//nolint:nolintlint,dupl
package fractionrule

import (
	"context"
	"fmt"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/miningpool/mw/v1/fractionrule"
	fractionrule "github.com/NpoolPlatform/miningpool-middleware/pkg/mw/fractionrule"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) UpdateFractionRule(ctx context.Context, in *npool.UpdateFractionRuleRequest) (*npool.UpdateFractionRuleResponse, error) {
	if in.GetInfo() == nil {
		err := fmt.Errorf("request is nil")
		logger.Sugar().Errorw(
			"UpdateFractionRule",
			"In", in,
			"Error", err,
		)
		return &npool.UpdateFractionRuleResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	req := in.GetInfo()
	handler, err := fractionrule.NewHandler(
		ctx,
		fractionrule.WithID(req.ID, true),
		fractionrule.WithEntID(req.EntID, false),
		fractionrule.WithMiningpoolType(req.MiningpoolType, true),
		fractionrule.WithCoinType(req.CoinType, true),
		fractionrule.WithWithdrawInterval(req.WithdrawInterval, true),
		fractionrule.WithMinAmount(req.MinAmount, true),
		fractionrule.WithWithdrawRate(req.WithdrawRate, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"UpdateFractionRule",
			"In", in,
			"Error", err,
		)
		return &npool.UpdateFractionRuleResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	err = handler.UpdateFractionRule(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"UpdateFractionRule",
			"In", in,
			"Error", err,
		)
		return &npool.UpdateFractionRuleResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.UpdateFractionRuleResponse{}, nil
}
