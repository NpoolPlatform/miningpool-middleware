//nolint:nolintlint,dupl
package orderuser

import (
	"context"
	"fmt"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	v1 "github.com/NpoolPlatform/message/npool/basetypes/miningpool/v1"
	npool "github.com/NpoolPlatform/message/npool/miningpool/mw/v1/orderuser"
	"github.com/NpoolPlatform/miningpool-middleware/pkg/mw/gooduser"
	orderuser "github.com/NpoolPlatform/miningpool-middleware/pkg/mw/orderuser"
	"github.com/NpoolPlatform/miningpool-middleware/pkg/mw/rootuser"
	"github.com/NpoolPlatform/miningpool-middleware/pkg/pools"

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
		return &npool.UpdateOrderUserResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	req, err := handleUpdateReq(ctx, in.GetInfo())
	if err != nil {
		logger.Sugar().Errorw(
			"UpdateOrderUser",
			"In", in,
			"Error", err,
		)
		return &npool.UpdateOrderUserResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	handler, err := orderuser.NewHandler(
		ctx,
		orderuser.WithID(req.ID, false),
		orderuser.WithEntID(req.EntID, false),
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
		return &npool.UpdateOrderUserResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	err = handler.UpdateOrderUser(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"UpdateOrderUser",
			"In", in,
			"Error", err,
		)
		return &npool.UpdateOrderUserResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.UpdateOrderUserResponse{}, nil
}

func handleUpdateReq(ctx context.Context, req *npool.OrderUserReq) (*npool.OrderUserReq, error) {
	baseInfo, err := getBaseInfo(ctx, req.EntID)
	if err != nil {
		return nil, err
	}
	mgr, err := pools.NewPoolManager(baseInfo.MiningpoolType, baseInfo.CoinType, baseInfo.AuthToken)
	if err != nil {
		return nil, err
	}
	if req.Proportion != nil {
		err = mgr.SetRevenueProportion(ctx, baseInfo.Distributor, baseInfo.Recipient, float64(*req.Proportion))
		if err != nil {
			return nil, err
		}
	}

	if req.RevenueAddress != nil {
		err = mgr.SetRevenueAddress(ctx, baseInfo.Recipient, *req.RevenueAddress)
		if err != nil {
			return nil, err
		}
	}

	if req.AutoPay != nil {
		autoPay := *req.AutoPay
		paused := true

		if autoPay {
			autoPay, err = mgr.ResumePayment(ctx, baseInfo.Recipient)
		} else {
			paused, err = mgr.PausePayment(ctx, baseInfo.Recipient)
		}
		if err != nil {
			return nil, err
		}
		if !paused {
			autoPay = false
		}

		req.AutoPay = &autoPay
	}

	return req, nil
}

type baseInfo struct {
	OrderUserID    uint32
	MiningpoolType v1.MiningpoolType
	CoinType       v1.CoinType
	AuthToken      string
	Recipient      string
	Distributor    string
}

func getBaseInfo(ctx context.Context, entid *string) (*baseInfo, error) {
	orderuserH, err := orderuser.NewHandler(ctx, orderuser.WithEntID(entid, true))
	if err != nil {
		return nil, err
	}
	orderUser, err := orderuserH.GetOrderUser(ctx)
	if err != nil {
		return nil, err
	}
	if orderUser == nil {
		err = fmt.Errorf("have no record of orderuser with entid %v", entid)
		return nil, err
	}

	gooduserH, err := gooduser.NewHandler(ctx, gooduser.WithEntID(&orderUser.GoodUserID, true))
	if err != nil {
		return nil, err
	}
	goodUser, err := gooduserH.GetGoodUser(ctx)
	if err != nil {
		return nil, err
	}
	if goodUser == nil {
		err = fmt.Errorf("have no record of gooduser with entid %v", orderUser.GoodUserID)
		return nil, err
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
		err = fmt.Errorf("have no record of rootuser with entid %v", orderUser.RootUserID)
		return nil, err
	}
	return &baseInfo{
		OrderUserID:    orderUser.ID,
		MiningpoolType: orderUser.MiningpoolType,
		CoinType:       orderUser.CoinType,
		Distributor:    goodUser.Name,
		Recipient:      orderUser.Name,
		AuthToken:      rootUser.AuthToken,
	}, nil
}
