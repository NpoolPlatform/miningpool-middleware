package api

import (
	"context"

	mw "github.com/NpoolPlatform/message/npool/miningpool/mw/v1"

	apppool "github.com/NpoolPlatform/miningpool-middleware/api/app/pool"
	coin "github.com/NpoolPlatform/miningpool-middleware/api/coin"
	fraction "github.com/NpoolPlatform/miningpool-middleware/api/fraction"
	fractionrule "github.com/NpoolPlatform/miningpool-middleware/api/fractionrule"
	gooduser "github.com/NpoolPlatform/miningpool-middleware/api/gooduser"
	orderuser "github.com/NpoolPlatform/miningpool-middleware/api/orderuser"
	pool "github.com/NpoolPlatform/miningpool-middleware/api/pool"
	rootuser "github.com/NpoolPlatform/miningpool-middleware/api/rootuser"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type Server struct {
	mw.UnimplementedMiddlewareServer
}

func Register(server grpc.ServiceRegistrar) {
	mw.RegisterMiddlewareServer(server, &Server{})
	pool.Register(server)
	coin.Register(server)
	rootuser.Register(server)
	fractionrule.Register(server)
	fraction.Register(server)
	gooduser.Register(server)
	orderuser.Register(server)
	apppool.Register(server)
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	if err := mw.RegisterMiddlewareHandlerFromEndpoint(context.Background(), mux, endpoint, opts); err != nil {
		return err
	}
	return nil
}
