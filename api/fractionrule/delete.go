package fractionrule

import (
	"context"
	"fmt"

	npool "github.com/NpoolPlatform/message/npool/miningpool/mw/v1/fractionrule"
	fractionrule "github.com/NpoolPlatform/miningpool-middleware/pkg/mw/fractionrule"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) DeleteFractionRule(ctx context.Context, in *npool.DeleteFractionRuleRequest) (*npool.DeleteFractionRuleResponse, error) {
	if in.GetInfo() == nil {
		err := fmt.Errorf("request is nil")
		logger.Sugar().Errorw(
			"DeleteFractionRule",
			"In", in,
			"Error", err,
		)
		return &npool.DeleteFractionRuleResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}
	req := in.GetInfo()
	if req == nil || req.ID == nil {
		err := fmt.Errorf("wrong id")
		return &npool.DeleteFractionRuleResponse{}, status.Error(codes.Internal, err.Error())
	}
	handler, err := fractionrule.NewHandler(
		ctx,
		fractionrule.WithID(req.ID, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"DeleteFractionRule",
			"In", in,
			"Error", err,
		)
		return &npool.DeleteFractionRuleResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := handler.DeleteFractionRule(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"DeleteFractionRule",
			"In", in,
			"Error", err,
		)
		return &npool.DeleteFractionRuleResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.DeleteFractionRuleResponse{
		Info: info,
	}, nil
}
