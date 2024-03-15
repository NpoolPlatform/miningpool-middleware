//nolint:nolintlint,dupl
package gooduser

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/miningpool/mw/v1/gooduser"
	gooduser "github.com/NpoolPlatform/miningpool-middleware/pkg/mw/gooduser"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) UpdateGoodUser(ctx context.Context, in *npool.UpdateGoodUserRequest) (*npool.UpdateGoodUserResponse, error) {
	req := in.GetInfo()
	handler, err := gooduser.NewHandler(
		ctx,
		gooduser.WithID(req.ID, false),
		gooduser.WithEntID(req.EntID, false),
		gooduser.WithName(req.Name, false),
		gooduser.WithRootUserID(req.RootUserID, false),
		gooduser.WithGoodID(req.GoodID, false),
		gooduser.WithMiningpoolType(req.MiningpoolType, false),
		gooduser.WithCoinType(req.CoinType, false),
		gooduser.WithHashRate(req.HashRate, false),
		gooduser.WithReadPageLink(req.ReadPageLink, false),
		gooduser.WithRevenueType(req.RevenueType, false),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"UpdateGoodUser",
			"In", in,
			"Error", err,
		)
		return &npool.UpdateGoodUserResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := handler.UpdateGoodUser(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"UpdateGoodUser",
			"In", in,
			"Error", err,
		)
		return &npool.UpdateGoodUserResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.UpdateGoodUserResponse{
		Info: info,
	}, nil
}
