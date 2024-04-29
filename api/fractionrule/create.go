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

func (s *Server) CreateFractionRule(ctx context.Context, in *npool.CreateFractionRuleRequest) (*npool.CreateFractionRuleResponse, error) {
	if in.GetInfo() == nil {
		err := fmt.Errorf("request is nil")
		logger.Sugar().Errorw(
			"CreateFractionRule",
			"In", in,
			"Error", err,
		)
		return &npool.CreateFractionRuleResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}
	req := in.GetInfo()
	handler, err := fractionrule.NewHandler(
		ctx,
		fractionrule.WithEntID(req.EntID, false),
		fractionrule.WithCoinID(req.CoinID, true),
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

	err = handler.CreateFractionRule(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateFractionRule",
			"In", in,
			"Error", err,
		)
		return &npool.CreateFractionRuleResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CreateFractionRuleResponse{}, nil
}
