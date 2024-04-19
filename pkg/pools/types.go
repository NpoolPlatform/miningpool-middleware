package pools

import "context"

type PoolManager interface {
	CheckAuth(ctx context.Context) error
	// mining user
	AddMiningUser(ctx context.Context) (userName string, readPageLink string, err error)
	ExistMiningUser(ctx context.Context, userName string) (bool, error)
	DeleteMiningUser(ctx context.Context, userName string) error

	// read page link
	AddReadPageLink(ctx context.Context, userName string) (string, error)
	GetReadPageLink(ctx context.Context, userName string) (string, error)
	DeleteReadPageLink(ctx context.Context, userName string) error

	// revenue
	SetRevenueProportion(ctx context.Context, distributor string, recipient string, proportion string) error
	GetRevenueProportion(ctx context.Context, distributor string, recipient string) (float64, error)
	SetRevenueAddress(ctx context.Context, userName string, address string) error
	GetRevenueAddress(ctx context.Context, userName string) (string, error)

	// auto pay
	PausePayment(ctx context.Context, userName string) (bool, error)
	ResumePayment(ctx context.Context, userName string) (bool, error)

	// withdraw praction
	WithdrawPraction(ctx context.Context, userName string) (int64, error)
}
