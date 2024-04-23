//nolint:nolintlint,dupl
package orderuser

import (
	"context"
	"fmt"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/miningpool/mw/v1/orderuser"
	orderuser "github.com/NpoolPlatform/miningpool-middleware/pkg/mw/orderuser"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) CreateOrderUser(ctx context.Context, in *npool.CreateOrderUserRequest) (*npool.CreateOrderUserResponse, error) {
	if in.GetInfo() == nil {
		err := fmt.Errorf("request is nil")
		logger.Sugar().Errorw(
			"CreateOrderUser",
			"In", in,
			"Error", err,
		)
		return &npool.CreateOrderUserResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	req := in.GetInfo()
	handler, err := orderuser.NewHandler(
		ctx,
		orderuser.WithEntID(req.EntID, false),
		orderuser.WithRootUserID(req.RootUserID, true),
		orderuser.WithGoodUserID(req.GoodUserID, true),
		orderuser.WithAppID(req.AppID, true),
		orderuser.WithUserID(req.UserID, true),
		orderuser.WithMiningpoolType(req.MiningpoolType, false),
		orderuser.WithCoinType(req.CoinType, false),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateOrderUser",
			"In", in,
			"Error", err,
		)
		return &npool.CreateOrderUserResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	err = handler.CreateOrderUser(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateOrderUser",
			"In", in,
			"Error", err,
		)
		return &npool.CreateOrderUserResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CreateOrderUserResponse{}, nil
}
