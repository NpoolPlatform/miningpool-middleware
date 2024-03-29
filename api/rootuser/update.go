//nolint:nolintlint,dupl
package rootuser

import (
	"context"
	"fmt"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	v1 "github.com/NpoolPlatform/message/npool/basetypes/miningpool/v1"
	npool "github.com/NpoolPlatform/message/npool/miningpool/mw/v1/rootuser"
	rootuser "github.com/NpoolPlatform/miningpool-middleware/pkg/mw/rootuser"
	"github.com/NpoolPlatform/miningpool-middleware/pkg/pools"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) UpdateRootUser(ctx context.Context, in *npool.UpdateRootUserRequest) (*npool.UpdateRootUserResponse, error) {
	req := in.GetInfo()
	req, err := checkUpdateAuthed(ctx, req)
	if err != nil {
		logger.Sugar().Warnf(
			"UpdateRootUser",
			"In", in,
			"Error", err,
		)
	}

	handler, err := rootuser.NewHandler(
		ctx,
		rootuser.WithID(req.ID, true),
		rootuser.WithEntID(req.EntID, false),
		rootuser.WithName(req.Name, false),
		rootuser.WithEmail(req.Email, false),
		rootuser.WithAuthToken(req.AuthToken, false),
		rootuser.WithAuthed(req.Authed, false),
		rootuser.WithRemark(req.Remark, false),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"UpdateRootUser",
			"In", in,
			"Error", err,
		)
		return &npool.UpdateRootUserResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := handler.UpdateRootUser(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"UpdateRootUser",
			"In", in,
			"Error", err,
		)
		return &npool.UpdateRootUserResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.UpdateRootUserResponse{
		Info: info,
	}, nil
}

func checkUpdateAuthed(ctx context.Context, req *npool.RootUserReq) (*npool.RootUserReq, error) {
	if req.AuthToken == nil && req.MiningpoolType == nil {
		return req, nil
	}

	defaultCoinType := v1.CoinType_BitCoin
	miningtype := req.MiningpoolType
	authtoken := req.AuthToken
	authed := false
	req.Authed = &authed

	if req.AuthToken == nil || req.MiningpoolType == nil {
		h, err := rootuser.NewHandler(ctx, rootuser.WithID(req.ID, true))
		if err != nil {
			return req, err
		}
		info, err := h.GetRootUser(ctx)
		if err != nil {
			return req, err
		}
		if req.AuthToken == nil {
			authtoken = &info.AuthToken
		}
		if req.MiningpoolType == nil {
			miningtype = &info.MiningpoolType
		}
	}

	mgr, err := pools.NewPoolManager(*miningtype, defaultCoinType, *authtoken)
	if err != nil {
		return req, err
	}
	err = mgr.CheckAuth(ctx)
	if err != nil {
		err = fmt.Errorf("have no permission to opreate pool, miningpool: %v, username: %v", req.MiningpoolType, req.Name)
		return req, err
	}
	authed = true
	return req, nil
}
