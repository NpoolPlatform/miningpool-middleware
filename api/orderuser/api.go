package orderuser

import (
	orderuser "github.com/NpoolPlatform/message/npool/miningpool/mw/v1/orderuser"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type Server struct {
	orderuser.UnimplementedMiddlewareServer
}

func Register(server grpc.ServiceRegistrar) {
	orderuser.RegisterMiddlewareServer(server, &Server{})
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	return nil
}
