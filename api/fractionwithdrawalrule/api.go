package fractionwithdrawalrule

import (
	fractionwithdrawalrule "github.com/NpoolPlatform/message/npool/miningpool/mw/v1/fractionwithdrawalrule"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type Server struct {
	fractionwithdrawalrule.UnimplementedMiddlewareServer
}

func Register(server grpc.ServiceRegistrar) {
	fractionwithdrawalrule.RegisterMiddlewareServer(server, &Server{})
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	return nil
}
