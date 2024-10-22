package fractionwithdrawal

import (
	"context"

	fractionwithdrawal "github.com/NpoolPlatform/message/npool/miningpool/mw/v1/fractionwithdrawal"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type Server struct {
	fractionwithdrawal.UnimplementedMiddlewareServer
}

func Register(server grpc.ServiceRegistrar) {
	fractionwithdrawal.RegisterMiddlewareServer(server, &Server{})
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	return fractionwithdrawal.RegisterMiddlewareHandlerFromEndpoint(context.Background(), mux, endpoint, opts)
}
