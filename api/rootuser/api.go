package rootuser

import (
	rootuser "github.com/NpoolPlatform/message/npool/miningpool/mw/v1/rootuser"

	"google.golang.org/grpc"
)

type Server struct {
	rootuser.UnimplementedMiddlewareServer
}

func Register(server grpc.ServiceRegistrar) {
	rootuser.RegisterMiddlewareServer(server, &Server{})
}
