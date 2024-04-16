//nolint:nolintlint,dupl
package fraction

import (
	"context"
	"fmt"
	"time"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	v1 "github.com/NpoolPlatform/message/npool/basetypes/miningpool/v1"
	npool "github.com/NpoolPlatform/message/npool/miningpool/mw/v1/fraction"
	fraction "github.com/NpoolPlatform/miningpool-middleware/pkg/mw/fraction"
	"github.com/NpoolPlatform/miningpool-middleware/pkg/mw/orderuser"
	"github.com/NpoolPlatform/miningpool-middleware/pkg/mw/rootuser"
	"github.com/NpoolPlatform/miningpool-middleware/pkg/pools"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) CreateFraction(ctx context.Context, in *npool.CreateFractionRequest) (*npool.CreateFractionResponse, error) {
	if in.GetInfo() == nil {
		err := fmt.Errorf("request is nil")
		logger.Sugar().Errorw(
			"CreateFraction",
			"In", in,
			"Error", err,
		)
		return &npool.CreateFractionResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	req, err := fractionInPool(ctx, in.GetInfo())
	if err != nil {
		logger.Sugar().Errorw(
			"CreateFraction",
			"In", in,
			"Error", err,
		)
		return &npool.CreateFractionResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}
	handler, err := fraction.NewHandler(
		ctx,
		fraction.WithEntID(req.EntID, false),
		fraction.WithAppID(req.AppID, true),
		fraction.WithUserID(req.UserID, true),
		fraction.WithOrderUserID(req.OrderUserID, true),
		fraction.WithWithdrawState(req.WithdrawState, true),
		fraction.WithWithdrawTime(req.WithdrawTime, true),
		fraction.WithWithdrawTime(req.PayTime, false),
		fraction.WithMsg(req.Msg, false),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateFraction",
			"In", in,
			"Error", err,
		)
		return &npool.CreateFractionResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := handler.CreateFraction(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateFraction",
			"In", in,
			"Error", err,
		)
		return &npool.CreateFractionResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CreateFractionResponse{
		Info: info,
	}, nil
}

func fractionInPool(ctx context.Context, info *npool.FractionReq) (*npool.FractionReq, error) {
	if info.OrderUserID == nil {
		return nil, fmt.Errorf("invalid orderuserid")
	}
	orderuserH, err := orderuser.NewHandler(ctx, orderuser.WithEntID(info.OrderUserID, true))
	if err != nil {
		return nil, err
	}
	orderUser, err := orderuserH.GetOrderUser(ctx)
	if err != nil {
		return nil, err
	}
	if orderUser == nil {
		return info, fmt.Errorf("have no orderuser,entid: %v", info.OrderUserID)
	}
	rootuserH, err := rootuser.NewHandler(ctx, rootuser.WithEntID(&orderUser.RootUserID, true))
	if err != nil {
		return nil, err
	}
	rootUser, err := rootuserH.GetRootUser(ctx)
	if err != nil {
		return nil, err
	}
	if rootUser == nil {
		return info, fmt.Errorf("have no rootuser,entid: %v", orderUser.RootUserID)
	}

	mgr, err := pools.NewPoolManager(orderUser.MiningpoolType, orderUser.CoinType, rootUser.AuthToken)
	if err != nil {
		return nil, err
	}
	withdrawTime := uint32(time.Now().Unix())
	info.WithdrawTime = &withdrawTime
	_paytime, err := mgr.WithdrawPraction(ctx, orderUser.Name)
	paytime := uint32(_paytime)
	if err != nil {
		errMsg := err.Error()
		info.Msg = &errMsg
		info.WithdrawState = v1.WithdrawState_WithdrawStateFailed.Enum()
	} else if paytime == 0 {
		info.PayTime = &paytime
		info.WithdrawState = v1.WithdrawState_WithdrawStateFailed.Enum()
	} else {
		info.PayTime = &paytime
		info.WithdrawState = v1.WithdrawState_WithdrawStateSuccess.Enum()
	}
	return info, nil
}
