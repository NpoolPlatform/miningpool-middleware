//nolint:dupl
package client

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/NpoolPlatform/miningpool-middleware/pkg/pools/f2pool/types"
	"github.com/NpoolPlatform/miningpool-middleware/pkg/utils"
)

const (
	F2PoolRps = 1
)

type Client struct {
	BaseURL         string
	AccessToken     string
	TimeTokenBucket chan struct{}
}

var client *Client

func NewClient(baseURL, accessToken string) *Client {
	if client == nil {
		client = &Client{BaseURL: baseURL, AccessToken: accessToken}
		client.TimeTokenBucket = make(chan struct{})
		go func() {
			for {
				client.TimeTokenBucket <- struct{}{}
				time.Sleep(time.Second / F2PoolRps)
			}
		}()
	}
	return client
}

func (cli *Client) post(ctx context.Context, path string, req, resp interface{}) error {
	reqBody, err := json.Marshal(req)
	if err != nil {
		return err
	}
	<-cli.TimeTokenBucket

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
