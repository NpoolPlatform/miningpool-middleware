package pool

import (
	"context"
	"fmt"
	"time"

	grpc2 "github.com/NpoolPlatform/go-service-framework/pkg/grpc"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	npool "github.com/NpoolPlatform/message/npool/miningpool/mw/v1/pool"

	servicename "github.com/NpoolPlatform/miningpool-middleware/pkg/servicename"
)

var timeout = 10 * time.Second

type handler func(context.Context, npool.MiddlewareClient) (cruder.Any, error)

func do(ctx context.Context, handler handler) (cruder.Any, error) {
	_ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	conn, err := grpc2.GetGRPCConn(servicename.ServiceDomain, grpc2.GRPCTAG)
	if err != nil {
		return nil, err
	}

	defer conn.Close()

	cli := npool.NewMiddlewareClient(conn)

	return handler(_ctx, cli)
}

func CreatePool(ctx context.Context, in *npool.PoolReq) error {
	_, err := do(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		_, err := cli.CreatePool(ctx, &npool.CreatePoolRequest{
			Info: in,
		})
		return nil, err
	})
	return err
}

func GetPool(ctx context.Context, id string) (*npool.Pool, error) {
	info, err := do(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.GetPool(ctx, &npool.GetPoolRequest{
			EntID: id,
		})
		if err != nil {
			return nil, err
		}
		return resp.Info, nil
	})
	if err != nil {
		return nil, err
	}
	return info.(*npool.Pool), nil
}

func GetPools(ctx context.Context, conds *npool.Conds, offset, limit int32) ([]*npool.Pool, uint32, error) {
	var total uint32

	infos, err := do(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.GetPools(ctx, &npool.GetPoolsRequest{
			Conds:  conds,
			Offset: offset,
			Limit:  limit,
		})
		if err != nil {
			return nil, err
		}

		total = resp.Total

		return resp.Infos, nil
	})
	if err != nil {
		return nil, 0, err
	}
	return infos.([]*npool.Pool), total, nil
}

func UpdatePool(ctx context.Context, in *npool.PoolReq) error {
	_, err := do(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		_, err := cli.UpdatePool(ctx, &npool.UpdatePoolRequest{
			Info: in,
		})
		return nil, err
	})
	return err
}

func ExistPool(ctx context.Context, id string) (bool, error) {
	info, err := do(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.ExistPool(ctx, &npool.ExistPoolRequest{
			EntID: id,
		})
		if err != nil {
			return nil, err
		}
		return resp.Info, nil
	})
	if err != nil {
		return false, err
	}
	return info.(bool), nil
}

func ExistPoolConds(ctx context.Context, conds *npool.Conds) (bool, error) {
	info, err := do(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.ExistPoolConds(ctx, &npool.ExistPoolCondsRequest{
			Conds: conds,
		})
		if err != nil {
			return nil, err
		}
		return resp.Info, nil
	})
	if err != nil {
		return false, err
	}
	return info.(bool), nil
}

func GetPoolOnly(ctx context.Context, conds *npool.Conds) (*npool.Pool, error) {
	infos, err := do(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		const singleRowLimit = 2
		resp, err := cli.GetPools(ctx, &npool.GetPoolsRequest{
			Conds:  conds,
			Offset: 0,
			Limit:  singleRowLimit,
		})
		if err != nil {
			return nil, err
		}
		return resp.Infos, nil
	})
	if err != nil {
		return nil, err
	}
	if len(infos.([]*npool.Pool)) == 0 {
		return nil, nil
	}
	if len(infos.([]*npool.Pool)) > 1 {
		return nil, fmt.Errorf("too many record")
	}
	return infos.([]*npool.Pool)[0], nil
}

func DeletePool(ctx context.Context, id uint32, entID string) error {
	_, err := do(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		_, err := cli.DeletePool(ctx, &npool.DeletePoolRequest{
			Info: &npool.PoolReq{
				ID:    &id,
				EntID: &entID,
			},
		})
		return nil, err
	})
	return err
}
