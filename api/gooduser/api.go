package gooduser

import (
	gooduser "github.com/NpoolPlatform/message/npool/miningpool/mw/v1/gooduser"

	"google.golang.org/grpc"
)

type Server struct {
	gooduser.UnimplementedMiddlewareServer
}

func Register(server grpc.ServiceRegistrar) {
	gooduser.RegisterMiddlewareServer(server, &Server{})
}
