package orderuser

import (
	"context"
	"fmt"
	"time"

	grpc2 "github.com/NpoolPlatform/go-service-framework/pkg/grpc"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	npool "github.com/NpoolPlatform/message/npool/miningpool/mw/v1/orderuser"

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

func CreateOrderUser(ctx context.Context, in *npool.OrderUserReq) error {
	_, err := do(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		_, err := cli.CreateOrderUser(ctx, &npool.CreateOrderUserRequest{
			Info: in,
		})
		return nil, err
	})
	return err
}

func GetOrderUser(ctx context.Context, id string) (*npool.OrderUser, error) {
	info, err := do(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.GetOrderUser(ctx, &npool.GetOrderUserRequest{
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
	return info.(*npool.OrderUser), nil
}

func GetOrderUsers(ctx context.Context, conds *npool.Conds, offset, limit int32) ([]*npool.OrderUser, uint32, error) {
	var total uint32

	infos, err := do(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.GetOrderUsers(ctx, &npool.GetOrderUsersRequest{
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
	return infos.([]*npool.OrderUser), total, nil
}

func GetOrderUserProportion(ctx context.Context, id, cointypeid string) (*float64, error) {
	info, err := do(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.GetOrderUserProportion(ctx, &npool.GetOrderUserProportionRequest{
			EntID:      id,
			CoinTypeID: cointypeid,
		})
		if err != nil {
			return nil, err
		}
		return resp.Proportion, nil
	})
	if err != nil {
		return nil, err
	}
	return info.(*float64), nil
}

func GetOrderUserBalance(ctx context.Context, id, cointypeid string) (*npool.GetOrderUserBalanceResponse, error) {
	info, err := do(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.GetOrderUserBalance(ctx, &npool.GetOrderUserBalanceRequest{
			EntID:      id,
			CoinTypeID: cointypeid,
		})
		if err != nil {
			return nil, err
		}
		return resp, nil
	})
	if err != nil {
		return nil, err
	}
	return info.(*npool.GetOrderUserBalanceResponse), nil
}

func UpdateOrderUser(ctx context.Context, in *npool.OrderUserReq) error {
	_, err := do(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		_, err := cli.UpdateOrderUser(ctx, &npool.UpdateOrderUserRequest{
			Info: in,
		})
		return nil, err
	})
	return err
}

func ExistOrderUser(ctx context.Context, id string) (bool, error) {
	info, err := do(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.ExistOrderUser(ctx, &npool.ExistOrderUserRequest{
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

func ExistOrderUserConds(ctx context.Context, conds *npool.Conds) (bool, error) {
	info, err := do(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.ExistOrderUserConds(ctx, &npool.ExistOrderUserCondsRequest{
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

func GetOrderUserOnly(ctx context.Context, conds *npool.Conds) (*npool.OrderUser, error) {
	infos, err := do(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		const singleRowLimit = 2
		resp, err := cli.GetOrderUsers(ctx, &npool.GetOrderUsersRequest{
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
	if len(infos.([]*npool.OrderUser)) == 0 {
		return nil, nil
	}
	if len(infos.([]*npool.OrderUser)) > 1 {
		return nil, fmt.Errorf("too many record")
	}
	return infos.([]*npool.OrderUser)[0], nil
}

func DeleteOrderUser(ctx context.Context, id uint32, entID string) error {
	_, err := do(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		_, err := cli.DeleteOrderUser(ctx, &npool.DeleteOrderUserRequest{
			Info: &npool.OrderUserReq{
				ID:    &id,
				EntID: &entID,
			},
		})
		return nil, err
	})
	return err
}
