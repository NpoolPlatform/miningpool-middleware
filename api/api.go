package api

import (
	"context"

	chainmw "github.com/NpoolPlatform/message/npool/miningpool/mw/v1"

	rootuser "github.com/NpoolPlatform/miningpool-middleware/api/rootuser"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type Server struct {
	chainmw.UnimplementedMiddlewareServer
}

func Register(server grpc.ServiceRegistrar) {
	rootuser.Register(server)
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	if err := chainmw.RegisterMiddlewareHandlerFromEndpoint(context.Background(), mux, endpoint, opts); err != nil {
		return err
	}
	return nil
}
