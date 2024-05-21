package pool

import (
	pool "github.com/NpoolPlatform/message/npool/miningpool/mw/v1/app/pool"

	"google.golang.org/grpc"
)

type Server struct {
	pool.UnimplementedMiddlewareServer
}

func Register(server grpc.ServiceRegistrar) {
	pool.RegisterMiddlewareServer(server, &Server{})
}
