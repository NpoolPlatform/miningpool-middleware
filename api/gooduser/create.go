//nolint:nolintlint,dupl
package gooduser

import (
	"context"
	"fmt"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/miningpool/mw/v1/gooduser"
	gooduser "github.com/NpoolPlatform/miningpool-middleware/pkg/mw/gooduser"

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
		return &npool.CreateGoodUserResponse{}, status.Error(codes.Aborted, "aborted err")
	}
	req := in.GetInfo()
	handler, err := gooduser.NewHandler(
		ctx,
		gooduser.WithEntID(req.EntID, false),
		gooduser.WithRootUserID(req.RootUserID, true),
		gooduser.WithCoinTypeIDs(req.CoinTypeIDs, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateGoodUser",
			"In", in,
			"Error", err,
		)
		return &npool.CreateGoodUserResponse{}, status.Error(codes.Aborted, "aborted err")
	}

	err = handler.CreateGoodUser(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateGoodUser",
			"In", in,
			"Error", err,
		)
		return &npool.CreateGoodUserResponse{}, status.Error(codes.Aborted, "aborted err")
	}

	return &npool.CreateGoodUserResponse{}, nil
}
