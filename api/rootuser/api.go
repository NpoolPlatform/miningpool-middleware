package rootuser

import (
	rootuser "github.com/NpoolPlatform/message/npool/miningpool/mw/v1/rootuser"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type Server struct {
	rootuser.UnimplementedMiddlewareServer
}

func Register(server grpc.ServiceRegistrar) {
	rootuser.RegisterMiddlewareServer(server, &Server{})
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	return nil
}
