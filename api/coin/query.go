//nolint:nolintlint,dupl
package coin

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/miningpool/mw/v1/coin"
	coin "github.com/NpoolPlatform/miningpool-middleware/pkg/mw/coin"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) GetCoin(ctx context.Context, in *npool.GetCoinRequest) (*npool.GetCoinResponse, error) {
	handler, err := coin.NewHandler(
		ctx,
		coin.WithEntID(&in.EntID, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetCoin",
			"In", in,
			"Error", err,
		)
		return &npool.GetCoinResponse{}, status.Error(codes.Internal, "internal server err")
	}

	info, err := handler.GetCoin(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetCoin",
			"In", in,
			"Error", err,
		)
		return &npool.GetCoinResponse{}, status.Error(codes.Internal, "internal server err")
	}

	return &npool.GetCoinResponse{
		Info: info,
	}, nil
}

func (s *Server) GetCoins(ctx context.Context, in *npool.GetCoinsRequest) (*npool.GetCoinsResponse, error) {
	handler, err := coin.NewHandler(
		ctx,
		coin.WithConds(in.Conds),
		coin.WithOffset(in.GetOffset()),
		coin.WithLimit(in.GetLimit()),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetCoins",
			"In", in,
			"Error", err,
		)
		return &npool.GetCoinsResponse{}, status.Error(codes.Internal, "internal server err")
	}

	infos, total, err := handler.GetCoins(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetCoins",
			"In", in,
			"Error", err,
		)
		return &npool.GetCoinsResponse{}, status.Error(codes.Internal, "internal server err")
	}

	return &npool.GetCoinsResponse{
		Infos: infos,
		Total: total,
	}, nil
}
