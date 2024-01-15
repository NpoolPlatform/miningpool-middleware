//nolint:dupl
package client

import (
	"context"

	"github.com/NpoolPlatform/miningpool-middleware/pkg/pools/f2pool/types"
)

func (cli *Client) RevenueDistributionInfo(ctx context.Context, req *types.RevenueDistributionInfoReq) (*types.RevenueDistributionInfoResp, error) {
	resp := &types.RevenueDistributionInfoResp{}
	err := cli.post(ctx, types.RevenueDistributionInfo, req, resp)
	return resp, err
}

func (cli *Client) RevenueDistributionAdd(ctx context.Context, req *types.RevenueDistributionAddReq) (*types.RevenueDistributionAddResp, error) {
	resp := &types.RevenueDistributionAddResp{}
	err := cli.post(ctx, types.RevenueDistributionAdd, req, resp)
	return resp, err
}

func (cli *Client) RevenueDistributionDelete(ctx context.Context, req *types.RevenueDistributionDeleteReq) (*types.RevenueDistributionDeleteResp, error) {
	resp := &types.RevenueDistributionDeleteResp{}
	err := cli.post(ctx, types.RevenueDistributionDelete, req, resp)
	return resp, err
}
