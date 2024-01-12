package main

import (
	"fmt"

	"github.com/NpoolPlatform/miningpool-middleware/pkg/utils"
)

func main() {
	url := "https://api.f2pool.com/v2/mining_user/get"

	payload := `{
		"mining_user_name": "coocoo"
	}`

	authHead := make(map[string]string)
	authHead["Content-Type"] = "application/json"
	authHead["F2P-API-SECRET"] = "wyrhuvsac5iaej9s3q1qx3l2lwvuoso1sdxvxzx1rju6tr27bqiujey9sj5ng546"

	resp, err := utils.PostJSON(url, []byte(payload), authHead)
	fmt.Println(utils.PrettyStruct(resp))
	fmt.Println(string(resp.Body))
	fmt.Println(utils.PrettyStruct(err))
}
