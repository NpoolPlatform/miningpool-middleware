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
			"CreateOrderUser",
			"In", in,
			"Error", err,
		)
		return &npool.CreateOrderUserResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := handler.CreateOrderUser(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateOrderUser",
			"In", in,
			"Error", err,
		)
		return &npool.CreateOrderUserResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CreateOrderUserResponse{
		Info: info,
	}, nil
}

func newOrderUserInPool(ctx context.Context, req *npool.OrderUserReq) (*npool.OrderUserReq, error) {
	rootuserH, err := rootuser.NewHandler(ctx, rootuser.WithEntID(req.RootUserID, true))
	if err != nil {
		return req, err
	}
	rootUser, err := rootuserH.GetRootUser(ctx)
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

	mgr, err := pools.NewPoolManager(*req.MiningpoolType, *req.CoinType, rootUser.AuthToken)
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
	req.MiningpoolType = &goodUser.MiningpoolType
	req.CoinType = &goodUser.CoinType
	return req, nil
}

func (s *Server) CreateOrderUsers(ctx context.Context, in *npool.CreateOrderUsersRequest) (*npool.CreateOrderUsersResponse, error) {
	reqs := in.GetInfos()
	for i, req := range reqs {
		_req, err := newOrderUserInPool(ctx, req)
		reqs[i] = _req
		if err != nil {
			logger.Sugar().Errorw(
				"CreateOrderUsers",
				"In", in,
				"Error", err,
			)
			return &npool.CreateOrderUsersResponse{}, status.Error(codes.InvalidArgument, err.Error())
		}
	}
	handler, err := orderuser.NewHandler(
		ctx,
		orderuser.WithReqs(in.GetInfos(), true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateOrderUsers",
			"In", in,
			"Error", err,
		)
		return &npool.CreateOrderUsersResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	infos, err := handler.CreateOrderUsers(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateOrderUsers",
			"In", in,
			"Error", err,
		)
		return &npool.CreateOrderUsersResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CreateOrderUsersResponse{
		Infos: infos,
	}, nil
}
