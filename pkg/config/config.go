package config

import (
	"os"
)

var (
	F2PoolAPI    = "https://api.f2pool.com"
	F2PoolAPIKey = "F2POOL_API"

	F2PoolBaseURL    = "https://f2pool.com"
	F2PoolBaseURLKey = "F2POOL_BASE_URL"
)

func init() {
	if v, ok := os.LookupEnv(F2PoolAPIKey); ok {
		F2PoolAPI = v
	}
	if v, ok := os.LookupEnv(F2PoolBaseURLKey); ok {
		F2PoolBaseURL = v
	}
}
