package main

import (
	"context"

	"github.com/NpoolPlatform/miningpool-middleware/api"
	"github.com/NpoolPlatform/miningpool-middleware/pkg/db"

	apicli "github.com/NpoolPlatform/basal-middleware/pkg/client/api"
	"github.com/NpoolPlatform/go-service-framework/pkg/action"
	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	cli "github.com/urfave/cli/v2"

	"google.golang.org/grpc"
)

var runCmd = &cli.Command{
	Name:    "run",
	Aliases: []string{"s"},
	Usage:   "Run the daemon",
	Action: func(c *cli.Context) error {
		err := action.Run(
			c.Context,
			run,
			rpcRegister,
			rpcGatewayRegister,
			watch,
		)
		return err
	},
}

func run(ctx context.Context) error {
	if err := db.Init(); err != nil {
		return err
	}
	return nil
}

func shutdown(ctx context.Context) {
	<-ctx.Done()
	logger.Sugar().Infow(
		"Watch",
		"State", "Done",
		"Error", ctx.Err(),
	)
}

func watch(ctx context.Context, cancel context.CancelFunc) error {
	go shutdown(ctx)
	return nil
}

func rpcRegister(server grpc.ServiceRegistrar) error {
	api.Register(server)
	apicli.RegisterGRPC(server)
	return nil
}

func rpcGatewayRegister(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	err := api.RegisterGateway(mux, endpoint, opts)
	if err != nil {
		return err
	}

	err = apicli.Register(mux)
	if err != nil {
		logger.Sugar().Warnw(
			"rpcGatewayRegister",
			"Error", err.Error(),
		)
	}

	return nil
}
