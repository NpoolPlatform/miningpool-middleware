//nolint:nolintlint,dupl
package orderuser

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/miningpool/mw/v1/orderuser"
	orderuser "github.com/NpoolPlatform/miningpool-middleware/pkg/mw/orderuser"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) GetOrderUser(ctx context.Context, in *npool.GetOrderUserRequest) (*npool.GetOrderUserResponse, error) {
	handler, err := orderuser.NewHandler(
		ctx,
		orderuser.WithEntID(&in.EntID, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetOrderUser",
			"In", in,
			"Error", err,
		)
		return &npool.GetOrderUserResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := handler.GetOrderUser(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetOrderUser",
			"In", in,
			"Error", err,
		)
		return &npool.GetOrderUserResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetOrderUserResponse{
		Info: info,
	}, nil
}

func (s *Server) GetOrderUsers(ctx context.Context, in *npool.GetOrderUsersRequest) (*npool.GetOrderUsersResponse, error) {
	handler, err := orderuser.NewHandler(
		ctx,
		orderuser.WithConds(in.Conds),
		orderuser.WithOffset(in.GetOffset()),
		orderuser.WithLimit(in.GetLimit()),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetOrderUsers",
			"In", in,
			"Error", err,
		)
		return &npool.GetOrderUsersResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	infos, total, err := handler.GetOrderUsers(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetOrderUsers",
			"In", in,
			"Error", err,
		)
		return &npool.GetOrderUsersResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetOrderUsersResponse{
		Infos: infos,
		Total: total,
	}, nil
}

func (s *Server) GetOrderUserProportion(ctx context.Context, in *npool.GetOrderUserProportionRequest) (*npool.GetOrderUserProportionResponse, error) {
	handler, err := orderuser.NewHandler(
		ctx,
		orderuser.WithEntID(&in.EntID, true),
		orderuser.WithCoinTypeID(&in.CoinTypeID, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetOrderUserProportion",
			"In", in,
			"Error", err,
		)
		return &npool.GetOrderUserProportionResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	proportion, err := handler.GetOrderUserProportion(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetOrderUserProportion",
			"In", in,
			"Error", err,
		)
		return &npool.GetOrderUserProportionResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetOrderUserProportionResponse{
		Proportion: proportion,
	}, nil
}

func (s *Server) GetOrderUserBalance(ctx context.Context, in *npool.GetOrderUserBalanceRequest) (*npool.GetOrderUserBalanceResponse, error) {
	handler, err := orderuser.NewHandler(
		ctx,
		orderuser.WithEntID(&in.EntID, true),
		orderuser.WithCoinTypeID(&in.CoinTypeID, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetOrderUserBalance",
			"In", in,
			"Error", err,
		)
		return &npool.GetOrderUserBalanceResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	assetsBalance, err := handler.GetOrderUserBalance(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetOrderUserBalance",
			"In", in,
			"Error", err,
		)
		return &npool.GetOrderUserBalanceResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetOrderUserBalanceResponse{
		Balance:              assetsBalance.Balance,
		Paid:                 assetsBalance.Paid,
		TotalIncome:          assetsBalance.TotalIncome,
		YesterdayIncome:      assetsBalance.YesterdayIncome,
		EstimatedTodayIncome: assetsBalance.EstimatedTodayIncome,
	}, nil
}
