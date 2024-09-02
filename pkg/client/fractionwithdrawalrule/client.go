package fractionwithdrawalrule

import (
	"context"
	"fmt"
	"time"

	grpc2 "github.com/NpoolPlatform/go-service-framework/pkg/grpc"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	npool "github.com/NpoolPlatform/message/npool/miningpool/mw/v1/fractionwithdrawalrule"

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

func CreateFractionWithdrawalRule(ctx context.Context, in *npool.FractionWithdrawalRuleReq) error {
	_, err := do(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		_, err := cli.CreateFractionWithdrawalRule(ctx, &npool.CreateFractionWithdrawalRuleRequest{
			Info: in,
		})
		return nil, err
	})
	return err
}

func GetFractionWithdrawalRule(ctx context.Context, id string) (*npool.FractionWithdrawalRule, error) {
	info, err := do(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.GetFractionWithdrawalRule(ctx, &npool.GetFractionWithdrawalRuleRequest{
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
	return info.(*npool.FractionWithdrawalRule), nil
}

func GetFractionWithdrawalRules(ctx context.Context, conds *npool.Conds, offset, limit int32) ([]*npool.FractionWithdrawalRule, uint32, error) {
	var total uint32

	infos, err := do(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.GetFractionWithdrawalRules(ctx, &npool.GetFractionWithdrawalRulesRequest{
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
	return infos.([]*npool.FractionWithdrawalRule), total, nil
}

func UpdateFractionWithdrawalRule(ctx context.Context, in *npool.FractionWithdrawalRuleReq) error {
	_, err := do(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		_, err := cli.UpdateFractionWithdrawalRule(ctx, &npool.UpdateFractionWithdrawalRuleRequest{
			Info: in,
		})
		return nil, err
	})
	return err
}

func ExistFractionWithdrawalRule(ctx context.Context, id string) (bool, error) {
	info, err := do(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.ExistFractionWithdrawalRule(ctx, &npool.ExistFractionWithdrawalRuleRequest{
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

func ExistFractionWithdrawalRuleConds(ctx context.Context, conds *npool.Conds) (bool, error) {
	info, err := do(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.ExistFractionWithdrawalRuleConds(ctx, &npool.ExistFractionWithdrawalRuleCondsRequest{
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

func GetFractionWithdrawalRuleOnly(ctx context.Context, conds *npool.Conds) (*npool.FractionWithdrawalRule, error) {
	infos, err := do(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		const singleRowLimit = 2
		resp, err := cli.GetFractionWithdrawalRules(ctx, &npool.GetFractionWithdrawalRulesRequest{
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
	if len(infos.([]*npool.FractionWithdrawalRule)) == 0 {
		return nil, nil
	}
	if len(infos.([]*npool.FractionWithdrawalRule)) > 1 {
		return nil, fmt.Errorf("too many record")
	}
	return infos.([]*npool.FractionWithdrawalRule)[0], nil
}

func DeleteFractionWithdrawalRule(ctx context.Context, id uint32, entID string) error {
	_, err := do(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		_, err := cli.DeleteFractionWithdrawalRule(ctx, &npool.DeleteFractionWithdrawalRuleRequest{
			Info: &npool.FractionWithdrawalRuleReq{
				ID:    &id,
				EntID: &entID,
			},
		})
		return nil, err
	})
	return err
}
