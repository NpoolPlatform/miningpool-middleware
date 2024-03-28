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
		gooduser.WithID(req.ID, false),
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

	info, err := handler.CreateGoodUser(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateGoodUser",
			"In", in,
			"Error", err,
		)
		return &npool.CreateGoodUserResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CreateGoodUserResponse{
		Info: info,
	}, nil
}

func newGoodUserInPool(ctx context.Context, req *npool.GoodUserReq) (*npool.GoodUserReq, error) {
	h, err := rootuser.NewHandler(ctx, rootuser.WithEntID(req.RootUserID, true))
	if err != nil {
		return req, err
	}
	rootUser, err := h.GetRootUser(ctx)
	if err != nil {
		return req, err
	}
	if rootUser == nil {
		return req, fmt.Errorf("have no rootuser,entid: %v", req.RootUserID)
	}
	req.MiningpoolType = &rootUser.MiningpoolType
	mgr, err := pools.NewPoolManager(*req.MiningpoolType, *req.CoinType, rootUser.AuthToken)
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

func (s *Server) CreateGoodUsers(ctx context.Context, in *npool.CreateGoodUsersRequest) (*npool.CreateGoodUsersResponse, error) {
	reqs := in.GetInfos()
	var err error
	for i, req := range reqs {
		req, err = newGoodUserInPool(ctx, req)
		reqs[i] = req
		if err != nil {
			logger.Sugar().Errorw(
				"CreateGoodUsers",
				"In", in,
				"Error", err,
			)
			return &npool.CreateGoodUsersResponse{}, status.Error(codes.InvalidArgument, err.Error())
		}
	}
	handler, err := gooduser.NewHandler(
		ctx,
		gooduser.WithReqs(in.GetInfos(), true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateGoodUsers",
			"In", in,
			"Error", err,
		)
		return &npool.CreateGoodUsersResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	infos, err := handler.CreateGoodUsers(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateGoodUsers",
			"In", in,
			"Error", err,
		)
		return &npool.CreateGoodUsersResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CreateGoodUsersResponse{
		Infos: infos,
	}, nil
}
