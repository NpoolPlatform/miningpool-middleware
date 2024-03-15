//nolint:nolintlint,dupl
package orderuser

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/miningpool/mw/v1/orderuser"
	orderuser "github.com/NpoolPlatform/miningpool-middleware/pkg/mw/orderuser"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) UpdateOrderUser(ctx context.Context, in *npool.UpdateOrderUserRequest) (*npool.UpdateOrderUserResponse, error) {
	req := in.GetInfo()
	handler, err := orderuser.NewHandler(
		ctx,
		orderuser.WithID(req.ID, false),
		orderuser.WithEntID(req.EntID, false),
		orderuser.WithName(req.Name, false),
		orderuser.WithRootUserID(req.RootUserID, false),
		orderuser.WithGoodUserID(req.GoodUserID, false),
		orderuser.WithOrderID(req.OrderID, false),
		orderuser.WithMiningpoolType(req.MiningpoolType, false),
		orderuser.WithCoinType(req.CoinType, false),
		orderuser.WithProportion(req.Proportion, false),
		orderuser.WithRevenueAddress(req.RevenueAddress, false),
		orderuser.WithReadPageLink(req.ReadPageLink, false),
		orderuser.WithAutoPay(req.AutoPay, false),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"UpdateOrderUser",
			"In", in,
			"Error", err,
		)
		return &npool.UpdateOrderUserResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := handler.UpdateOrderUser(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"UpdateOrderUser",
			"In", in,
			"Error", err,
		)
		return &npool.UpdateOrderUserResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.UpdateOrderUserResponse{
		Info: info,
	}, nil
}
