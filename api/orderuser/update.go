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
	req := in.GetInfo()
	handler, err := orderuser.NewHandler(
		ctx,
		orderuser.WithID(req.ID, false),
		orderuser.WithEntID(req.EntID, false),
		orderuser.WithOrderID(req.OrderID, false),
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

func (s *Server) SetupProportion(ctx context.Context, in *npool.SetupProportionRequest) (*npool.SetupProportionResponse, error) {
	var err error
	defer func() {
		if err != nil {
			logger.Sugar().Errorw(
				"SetupProportion",
				"In", in,
				"Error", err,
			)
		}
	}()
	baseInfo, err := getBaseInfo(ctx, &in.EntID)
	if err != nil {
		return &npool.SetupProportionResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	mgr, err := pools.NewPoolManager(baseInfo.MiningpoolType, baseInfo.CoinType, baseInfo.AuthToken)
	if err != nil {
		return &npool.SetupProportionResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	err = mgr.SetRevenueProportion(ctx, baseInfo.Distributor, baseInfo.Recipient, float64(in.Proportion))
	if err != nil {
		return &npool.SetupProportionResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	handler, err := orderuser.NewHandler(
		ctx,
		orderuser.WithID(&baseInfo.OrderUserID, true),
		orderuser.WithProportion(&in.Proportion, true),
	)
	if err != nil {
		return &npool.SetupProportionResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := handler.UpdateOrderUser(ctx)
	if err != nil {
		return &npool.SetupProportionResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.SetupProportionResponse{
		Info: info,
	}, nil
}

func (s *Server) SetupRevenueAddress(ctx context.Context, in *npool.SetupRevenueAddressRequest) (*npool.SetupRevenueAddressResponse, error) {
	var err error
	defer func() {
		if err != nil {
			logger.Sugar().Errorw(
				"SetupProportion",
				"In", in,
				"Error", err,
			)
		}
	}()
	baseInfo, err := getBaseInfo(ctx, &in.EntID)
	if err != nil {
		return &npool.SetupRevenueAddressResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	mgr, err := pools.NewPoolManager(baseInfo.MiningpoolType, baseInfo.CoinType, baseInfo.AuthToken)
	if err != nil {
		return &npool.SetupRevenueAddressResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	err = mgr.SetRevenueAddress(ctx, baseInfo.Recipient, in.RevenueAddress)
	if err != nil {
		return &npool.SetupRevenueAddressResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	handler, err := orderuser.NewHandler(
		ctx,
		orderuser.WithID(&baseInfo.OrderUserID, true),
		orderuser.WithRevenueAddress(&in.RevenueAddress, true),
	)
	if err != nil {
		return &npool.SetupRevenueAddressResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := handler.UpdateOrderUser(ctx)
	if err != nil {
		return &npool.SetupRevenueAddressResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.SetupRevenueAddressResponse{
		Info: info,
	}, nil
}

func (s *Server) SetupAutoPay(ctx context.Context, in *npool.SetupAutoPayRequest) (*npool.SetupAutoPayResponse, error) {
	var err error
	defer func() {
		if err != nil {
			logger.Sugar().Errorw(
				"SetupProportion",
				"In", in,
				"Error", err,
			)
		}
	}()
	baseInfo, err := getBaseInfo(ctx, &in.EntID)
	if err != nil {
		return &npool.SetupAutoPayResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	mgr, err := pools.NewPoolManager(baseInfo.MiningpoolType, baseInfo.CoinType, baseInfo.AuthToken)
	if err != nil {
		return &npool.SetupAutoPayResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	autoPay := in.AutoPay
	paused := true
	if in.AutoPay {
		autoPay, err = mgr.ResumePayment(ctx, baseInfo.Recipient)
	} else {
		paused, err = mgr.PausePayment(ctx, baseInfo.Recipient)
	}
	if err != nil {
		return &npool.SetupAutoPayResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}
	if !paused {
		autoPay = false
	}

	handler, err := orderuser.NewHandler(
		ctx,
		orderuser.WithID(&baseInfo.OrderUserID, true),
		orderuser.WithAutoPay(&autoPay, true),
	)
	if err != nil {
		return &npool.SetupAutoPayResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := handler.UpdateOrderUser(ctx)
	if err != nil {
		return &npool.SetupAutoPayResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.SetupAutoPayResponse{
		Info: info,
	}, nil
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
	fmt.Println(1)

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
	fmt.Println(2)

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
	fmt.Println(3)
	return &baseInfo{
		OrderUserID:    orderUser.ID,
		MiningpoolType: orderUser.MiningpoolType,
		CoinType:       orderUser.CoinType,
		Distributor:    goodUser.Name,
		Recipient:      orderUser.Name,
		AuthToken:      rootUser.AuthToken,
	}, nil
}