//nolint:dupl
package client

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/NpoolPlatform/miningpool-middleware/pkg/pools/f2pool/types"
	"github.com/NpoolPlatform/miningpool-middleware/pkg/utils"
)

type Client struct {
	BaseURL     string
	AccessToken string
}

func NewClient(baseURL, accessToken string) *Client {
	return &Client{BaseURL: baseURL, AccessToken: accessToken}
}

func (cli *Client) post(ctx context.Context, path string, req, resp interface{}) error {
	reqBody, err := json.Marshal(req)
	if err != nil {
		return err
	}

	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"
	headers["F2P-API-SECRET"] = cli.AccessToken

	url := fmt.Sprintf("%v%v", cli.BaseURL, path)
	postResp, err := utils.PostJSON(ctx, url, reqBody, headers)
	if err != nil {
		return err
	}

	if postResp.StatusCode != utils.SuccessStatusCode {
		return fmt.Errorf("wrong response,status code: %v,response: %v", postResp.StatusCode, string(postResp.Body))
	}

	errResp := &types.ErrorResponse{}
	err = json.Unmarshal(postResp.Body, errResp)

	if err == nil && errResp.Code != 0 {
		return fmt.Errorf("request api %v error,status_code: %v,err: %v", url, errResp.Code, errResp.Msg)
	}

	return json.Unmarshal(postResp.Body, resp)
}
