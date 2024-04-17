//nolint:nolintlint,dupl
package gooduser

import (
	"context"
	"fmt"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/miningpool/mw/v1/gooduser"
	gooduser "github.com/NpoolPlatform/miningpool-middleware/pkg/mw/gooduser"
	"github.com/NpoolPlatform/miningpool-middleware/pkg/mw/rootuser"
	"github.com/NpoolPlatform/miningpool-middleware/pkg/pools"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) CreateGoodUser(ctx context.Context, in *npool.CreateGoodUserRequest) (*npool.CreateGoodUserResponse, error) {
	if in.GetInfo() == nil {
		err := fmt.Errorf("request is nil")
		logger.Sugar().Errorw(
			"CreateGoodUser",
			"In", in,
			"Error", err,
		)
		return &npool.CreateGoodUserResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}
	req := in.GetInfo()
	req, err := newGoodUserInPool(ctx, req)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateGoodUser",
			"In", in,
			"Error", err,
		)
		return &npool.CreateGoodUserResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	handler, err := gooduser.NewHandler(
		ctx,
		gooduser.WithEntID(req.EntID, false),
		gooduser.WithName(req.Name, true),
		gooduser.WithRootUserID(req.RootUserID, true),
		gooduser.WithMiningpoolType(req.MiningpoolType, true),
		gooduser.WithCoinType(req.CoinType, true),
		gooduser.WithHashRate(req.HashRate, false),
		gooduser.WithReadPageLink(req.ReadPageLink, true),
		gooduser.WithRevenueType(req.RevenueType, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateGoodUser",
			"In", in,
			"Error", err,
		)
		return &npool.CreateGoodUserResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	err = handler.CreateGoodUser(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateGoodUser",
			"In", in,
			"Error", err,
		)
		return &npool.CreateGoodUserResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CreateGoodUserResponse{}, nil
}

func newGoodUserInPool(ctx context.Context, req *npool.GoodUserReq) (*npool.GoodUserReq, error) {
	if req.CoinType == nil {
		return req, fmt.Errorf("invalid cointype")
	}
	h, err := rootuser.NewHandler(ctx, rootuser.WithEntID(req.RootUserID, true))
	if err != nil {
		return req, err
	}
	rootUser, err := h.GetAuthToken(ctx)
	if err != nil {
		return req, err
	}
	if rootUser == nil {
		return req, fmt.Errorf("have no rootuser,entid: %v", req.RootUserID)
	}
	req.MiningpoolType = &rootUser.MiningpoolType
	mgr, err := pools.NewPoolManager(*req.MiningpoolType, *req.CoinType, rootUser.AuthTokenPlain)
	if err != nil {
		return req, err
	}
	name, pageLink, err := mgr.AddMiningUser(ctx)
	if err != nil {
		return req, err
	}
	req.Name = &name
	req.ReadPageLink = &pageLink
	return req, nil
}
