//nolint:nolintlint,dupl
package orderuser

import (
	"context"
	"fmt"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/miningpool/mw/v1/orderuser"
	"github.com/NpoolPlatform/miningpool-middleware/pkg/mw/gooduser"
	orderuser "github.com/NpoolPlatform/miningpool-middleware/pkg/mw/orderuser"
	"github.com/NpoolPlatform/miningpool-middleware/pkg/mw/rootuser"
	"github.com/NpoolPlatform/miningpool-middleware/pkg/pools"

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
	req, err := newOrderUserInPool(ctx, req)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateOrderUser",
			"In", in,
			"Error", err,
		)
		return &npool.CreateOrderUserResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	handler, err := orderuser.NewHandler(
		ctx,
		orderuser.WithEntID(req.EntID, false),
		orderuser.WithName(req.Name, true),
		orderuser.WithRootUserID(req.RootUserID, true),
		orderuser.WithGoodUserID(req.GoodUserID, true),
		orderuser.WithAppID(req.AppID, true),
		orderuser.WithUserID(req.UserID, true),
		orderuser.WithMiningpoolType(req.MiningpoolType, true),
		orderuser.WithCoinType(req.CoinType, true),
		orderuser.WithReadPageLink(req.ReadPageLink, true),
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

func newOrderUserInPool(ctx context.Context, req *npool.OrderUserReq) (*npool.OrderUserReq, error) {
	rootuserH, err := rootuser.NewHandler(ctx, rootuser.WithEntID(req.RootUserID, true))
	if err != nil {
		return req, err
	}
	rootUser, err := rootuserH.GetAuthToken(ctx)
	if err != nil {
		return req, err
	}
	if rootUser == nil {
		return req, fmt.Errorf("have no rootuser,entid: %v", req.RootUserID)
	}

	gooduserH, err := gooduser.NewHandler(ctx, gooduser.WithEntID(req.GoodUserID, true))
	if err != nil {
		return req, err
	}
	goodUser, err := gooduserH.GetGoodUser(ctx)
	if err != nil {
		return req, err
	}
	if goodUser == nil {
		return req, fmt.Errorf("have no gooduser,entid: %v", req.GoodUserID)
	}

	req.MiningpoolType = &goodUser.MiningpoolType
	req.CoinType = &goodUser.CoinType

	mgr, err := pools.NewPoolManager(*req.MiningpoolType, *req.CoinType, rootUser.AuthTokenPlain)
	if err != nil {
		return req, err
	}

	name, pagelink, err := mgr.AddMiningUser(ctx)
	if err != nil {
		return req, err
	}
	req.Name = &name
	req.ReadPageLink = &pagelink

	paused, err := mgr.PausePayment(ctx, name)
	if err != nil {
		return req, err
	}
	autopay := !paused
	req.AutoPay = &autopay

	return req, nil
}
