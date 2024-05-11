//nolint:dupl
package gooduser

import (
	"context"
	"fmt"

	npool "github.com/NpoolPlatform/message/npool/miningpool/mw/v1/gooduser"
	gooduser "github.com/NpoolPlatform/miningpool-middleware/pkg/mw/gooduser"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) DeleteGoodUser(ctx context.Context, in *npool.DeleteGoodUserRequest) (*npool.DeleteGoodUserResponse, error) {
	if in.GetInfo() == nil {
		err := fmt.Errorf("request is nil")
		logger.Sugar().Errorw(
			"DeleteGoodUser",
			"In", in,
			"Error", err,
		)
		return &npool.DeleteGoodUserResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	req := in.GetInfo()
	handler, err := gooduser.NewHandler(
		ctx,
		gooduser.WithID(req.ID, false),
		gooduser.WithEntID(req.EntID, false),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"DeleteGoodUser",
			"In", in,
			"Error", err,
		)
		return &npool.DeleteGoodUserResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	err = handler.DeleteGoodUser(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"DeleteGoodUser",
			"In", in,
			"Error", err,
		)
		return &npool.DeleteGoodUserResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.DeleteGoodUserResponse{}, nil
}
