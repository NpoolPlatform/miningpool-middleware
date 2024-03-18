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

func (s *Server) CreateRootUser(ctx context.Context, in *npool.CreateRootUserRequest) (*npool.CreateRootUserResponse, error) {
	req := in.GetInfo()
	authed, err := checkCreateAuthed(ctx, req)
	if err != nil {
		logger.Sugar().Warnw(
			"CreateRootUser",
			"In", in,
			"Error", err,
		)
	}
	req.Authed = &authed

	handler, err := rootuser.NewHandler(
		ctx,
		rootuser.WithID(req.ID, false),
		rootuser.WithEntID(req.EntID, false),
		rootuser.WithName(req.Name, true),
		rootuser.WithMiningpoolType(req.MiningpoolType, true),
		rootuser.WithEmail(req.Email, false),
		rootuser.WithAuthToken(req.AuthToken, true),
		rootuser.WithAuthed(req.Authed, true),
		rootuser.WithRemark(req.Remark, false),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateRootUser",
			"In", in,
			"Error", err,
		)
		return &npool.CreateRootUserResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := handler.CreateRootUser(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateRootUser",
			"In", in,
			"Error", err,
		)
		return &npool.CreateRootUserResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CreateRootUserResponse{
		Info: info,
	}, nil
}

func checkCreateAuthed(ctx context.Context, req *npool.RootUserReq) (bool, error) {
	defaultCoinType := v1.CoinType_BitCoin
	mgr, err := pools.NewPoolManager(*req.MiningpoolType, defaultCoinType, *req.AuthToken)
	if err != nil {
		return false, err
	}

	err = mgr.CheckAuth(ctx)
	if err != nil {
		err = fmt.Errorf("have no permission to opreate pool, miningpool: %v, username: %v", req.MiningpoolType, req.Name)
		return false, err
	}
	return true, nil
}

func (s *Server) CreateRootUsers(ctx context.Context, in *npool.CreateRootUsersRequest) (*npool.CreateRootUsersResponse, error) {
	reqs := in.GetInfos()
	for _, req := range reqs {
		authed, err := checkCreateAuthed(ctx, req)
		if err != nil {
			logger.Sugar().Warnw(
				"CreateRootUsers",
				"In", in,
				"Error", err,
			)
		}
		req.Authed = &authed
	}

	handler, err := rootuser.NewHandler(
		ctx,
		rootuser.WithReqs(reqs, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateRootUsers",
			"In", in,
			"Error", err,
		)
		return &npool.CreateRootUsersResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	infos, err := handler.CreateRootUsers(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateRootUsers",
			"In", in,
			"Error", err,
		)
		return &npool.CreateRootUsersResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CreateRootUsersResponse{
		Infos: infos,
	}, nil
}
