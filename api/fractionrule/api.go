package fractionrule

import (
	fractionrule "github.com/NpoolPlatform/message/npool/miningpool/mw/v1/fractionrule"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type Server struct {
	fractionrule.UnimplementedMiddlewareServer
}

func Register(server grpc.ServiceRegistrar) {
	fractionrule.RegisterMiddlewareServer(server, &Server{})
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	return nil
}
