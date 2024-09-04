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

func (s *Server) UpdateOrderUser(ctx context.Context, in *npool.UpdateOrderUserRequest) (*npool.UpdateOrderUserResponse, error) {
	if in.GetInfo() == nil {
		err := fmt.Errorf("request is nil")
		logger.Sugar().Errorw(
			"UpdateOrderUser",
			"In", in,
			"Error", err,
		)
		return &npool.UpdateOrderUserResponse{}, status.Error(codes.Internal, "internal server err")
	}
	req := in.GetInfo()
	handler, err := orderuser.NewHandler(
		ctx,
		orderuser.WithID(req.ID, false),
		orderuser.WithEntID(req.EntID, false),
		orderuser.WithCoinTypeID(req.CoinTypeID, false),
		orderuser.WithProportion(req.Proportion, false),
		orderuser.WithRevenueAddress(req.RevenueAddress, false),
		orderuser.WithAutoPay(req.AutoPay, false),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"UpdateOrderUser",
			"In", in,
			"Error", err,
		)
		return &npool.UpdateOrderUserResponse{}, status.Error(codes.Internal, "internal server err")
	}

	err = handler.UpdateOrderUser(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"UpdateOrderUser",
			"In", in,
			"Error", err,
		)
		return &npool.UpdateOrderUserResponse{}, status.Error(codes.Internal, "internal server err")
	}

	return &npool.UpdateOrderUserResponse{}, nil
}
