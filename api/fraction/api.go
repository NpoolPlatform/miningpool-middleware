package fraction

import (
	"context"

	fraction "github.com/NpoolPlatform/message/npool/miningpool/mw/v1/fraction"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type Server struct {
	fraction.UnimplementedMiddlewareServer
}

func Register(server grpc.ServiceRegistrar) {
	fraction.RegisterMiddlewareServer(server, &Server{})
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	return fraction.RegisterMiddlewareHandlerFromEndpoint(context.Background(), mux, endpoint, opts)
}
