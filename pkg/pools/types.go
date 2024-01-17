package pools

import "context"

type PoolManager interface {
	CheckAuth(context.Context) bool
	AddMiningUser(context.Context) (string, string, error)
	ExistMiningUser(context.Context, string) (bool, error)
	DeleteMiningUser(context.Context, string) error

	AddReadPageLink(context.Context, string) (string, error)
	GetReadPageLink(context.Context, string) (string, error)
	DeleteReadPageLink(context.Context, string) error

	SetRevenueProportion(context.Context, string, string, float64) error
	GetRevenueProportion(context.Context, string, string) (float64, error)

	SetRevenueAddress(context.Context, string, string) error
	GetRevenueAddress(context.Context, string) (string, error)
	// auto pay
	// praction
}
