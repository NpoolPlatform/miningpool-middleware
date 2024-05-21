package pool

import (
	pool "github.com/NpoolPlatform/message/npool/miningpool/mw/v1/pool"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type Server struct {
	pool.UnimplementedMiddlewareServer
}

func Register(server grpc.ServiceRegistrar) {
	pool.RegisterMiddlewareServer(server, &Server{})
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	return nil
}
