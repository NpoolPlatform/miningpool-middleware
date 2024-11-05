package fractionwithdrawal

import (
	"context"
	"fmt"
	"time"

	grpc2 "github.com/NpoolPlatform/go-service-framework/pkg/grpc"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	npool "github.com/NpoolPlatform/message/npool/miningpool/mw/v1/fractionwithdrawal"

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

func CreateFractionWithdrawal(ctx context.Context, in *npool.FractionWithdrawalReq) error {
	_, err := do(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		_, err := cli.CreateFractionWithdrawal(ctx, &npool.CreateFractionWithdrawalRequest{
			Info: in,
		})
		return nil, err
	})
	return err
}

func GetFractionWithdrawal(ctx context.Context, id string) (*npool.FractionWithdrawal, error) {
	info, err := do(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.GetFractionWithdrawal(ctx, &npool.GetFractionWithdrawalRequest{
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
	return info.(*npool.FractionWithdrawal), nil
}

func GetFractionWithdrawals(ctx context.Context, conds *npool.Conds, offset, limit int32) ([]*npool.FractionWithdrawal, uint32, error) {
	var total uint32

	infos, err := do(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.GetFractionWithdrawals(ctx, &npool.GetFractionWithdrawalsRequest{
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
	return infos.([]*npool.FractionWithdrawal), total, nil
}

func UpdateFractionWithdrawal(ctx context.Context, in *npool.FractionWithdrawalReq) error {
	_, err := do(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		_, err := cli.UpdateFractionWithdrawal(ctx, &npool.UpdateFractionWithdrawalRequest{
			Info: in,
		})
		return nil, err
	})
	return err
}

func ExistFractionWithdrawal(ctx context.Context, id string) (bool, error) {
	info, err := do(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.ExistFractionWithdrawal(ctx, &npool.ExistFractionWithdrawalRequest{
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

func ExistFractionWithdrawalConds(ctx context.Context, conds *npool.Conds) (bool, error) {
	info, err := do(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.ExistFractionWithdrawalConds(ctx, &npool.ExistFractionWithdrawalCondsRequest{
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

func GetFractionWithdrawalOnly(ctx context.Context, conds *npool.Conds) (*npool.FractionWithdrawal, error) {
	infos, err := do(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		const singleRowLimit = 2
		resp, err := cli.GetFractionWithdrawals(ctx, &npool.GetFractionWithdrawalsRequest{
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
	if len(infos.([]*npool.FractionWithdrawal)) == 0 {
		return nil, nil
	}
	if len(infos.([]*npool.FractionWithdrawal)) > 1 {
		return nil, fmt.Errorf("too many record")
	}
	return infos.([]*npool.FractionWithdrawal)[0], nil
}

func DeleteFractionWithdrawal(ctx context.Context, id uint32, entID string) error {
	_, err := do(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		_, err := cli.DeleteFractionWithdrawal(ctx, &npool.DeleteFractionWithdrawalRequest{
			Info: &npool.FractionWithdrawalReq{
				ID:    &id,
				EntID: &entID,
			},
		})
		return nil, err
	})
	return err
}
