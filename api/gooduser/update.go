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

func (s *Server) UpdateGoodUser(ctx context.Context, in *npool.UpdateGoodUserRequest) (*npool.UpdateGoodUserResponse, error) {
	if in.GetInfo() == nil {
		err := fmt.Errorf("request is nil")
		logger.Sugar().Errorw(
			"UpdateGoodUser",
			"In", in,
			"Error", err,
		)
		return &npool.UpdateGoodUserResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}
	req := in.GetInfo()
	handler, err := gooduser.NewHandler(
		ctx,
		gooduser.WithID(req.ID, false),
		gooduser.WithEntID(req.EntID, false),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"UpdateGoodUser",
			"In", in,
			"Error", err,
		)
		return &npool.UpdateGoodUserResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	err = handler.UpdateGoodUser(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"UpdateGoodUser",
			"In", in,
			"Error", err,
		)
		return &npool.UpdateGoodUserResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.UpdateGoodUserResponse{}, nil
}
