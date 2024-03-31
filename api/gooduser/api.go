package gooduser

import (
	"context"

	gooduser "github.com/NpoolPlatform/message/npool/miningpool/mw/v1/gooduser"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type Server struct {
	gooduser.UnimplementedMiddlewareServer
}

func Register(server grpc.ServiceRegistrar) {
	gooduser.RegisterMiddlewareServer(server, &Server{})
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	return gooduser.RegisterMiddlewareHandlerFromEndpoint(context.Background(), mux, endpoint, opts)
}
