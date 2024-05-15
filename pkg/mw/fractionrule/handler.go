package fractionrule

import (
	"context"
	"fmt"

	npool "github.com/NpoolPlatform/message/npool/miningpool/mw/v1/fractionrule"
	constant "github.com/NpoolPlatform/miningpool-middleware/pkg/const"
	fractionrulecrud "github.com/NpoolPlatform/miningpool-middleware/pkg/crud/fractionrule"
	"github.com/NpoolPlatform/miningpool-middleware/pkg/mw/coin"
	"github.com/shopspring/decimal"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"

	"github.com/google/uuid"
)

type Handler struct {
	ID               *uint32
	EntID            *uuid.UUID
	PoolCoinTypeID   *uuid.UUID
	WithdrawInterval *uint32
	MinAmount        *decimal.Decimal
	WithdrawRate     *decimal.Decimal
	Reqs             []*fractionrulecrud.Req
	Conds            *fractionrulecrud.Conds
	Offset           int32
	Limit            int32
}

func NewHandler(ctx context.Context, options ...func(context.Context, *Handler) error) (*Handler, error) {
	handler := &Handler{}
	for _, opt := range options {
		if err := opt(ctx, handler); err != nil {
			return nil, err
		}
	}
	return handler, nil
}

func WithID(u *uint32, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if u == nil {
			if must {
				return fmt.Errorf("invalid id")
			}
			return nil
		}
		h.ID = u
		return nil
	}
}

func WithEntID(id *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if id == nil {
			if must {
				return fmt.Errorf("invalid entid")
			}
			return nil
		}
		_id, err := uuid.Parse(*id)
		if err != nil {
			return err
		}
		h.EntID = &_id
		return nil
	}
}

func WithPoolCoinTypeID(poolcointypeid *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if poolcointypeid == nil {
			if must {
				return fmt.Errorf("invalid poolcointypeid")
			}
			return nil
		}

		coinH, err := coin.NewHandler(ctx, coin.WithEntID(poolcointypeid, true))
		if err != nil {
			return err
		}

		exist, err := coinH.ExistCoin(ctx)
		if err != nil {
			return err
		}
		if !exist {
			return fmt.Errorf("invalid poolcointypeid")
		}
		h.PoolCoinTypeID = coinH.EntID
		return nil
	}
}

func WithWithdrawInterval(withdrawinterval *uint32, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if withdrawinterval == nil {
			if must {
				return fmt.Errorf("invalid withdrawinterval")
			}
			return nil
		}
		h.WithdrawInterval = withdrawinterval
		return nil
	}
}

func WithMinAmount(minamount *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if minamount == nil {
			if must {
				return fmt.Errorf("invalid minamount")
			}
			return nil
		}
		_minamount, err := decimal.NewFromString(*minamount)
		if err != nil {
			return fmt.Errorf("invalid minamount,err: %v", err)
		}
		if _minamount.Sign() <= 0 {
			return fmt.Errorf("invalid minamount")
		}
		h.MinAmount = &_minamount
		return nil
	}
}

func WithWithdrawRate(withdrawrate *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if withdrawrate == nil {
			if must {
				return fmt.Errorf("invalid withdrawrate")
			}
			return nil
		}
		_withdrawrate, err := decimal.NewFromString(*withdrawrate)
		if err != nil {
			return fmt.Errorf("invalid withdrawrate,err: %v", err)
		}
		h.WithdrawRate = &_withdrawrate
		return nil
	}
}

//nolint:gocognit
func WithConds(conds *npool.Conds) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.Conds = &fractionrulecrud.Conds{}
		if conds == nil {
			return nil
		}
		if conds.ID != nil {
			h.Conds.ID = &cruder.Cond{
				Op:  conds.GetID().GetOp(),
				Val: conds.GetID().GetValue(),
			}
		}
		if conds.EntID != nil {
			id, err := uuid.Parse(conds.GetEntID().GetValue())
			if err != nil {
				return err
			}
			h.Conds.EntID = &cruder.Cond{
				Op:  conds.GetEntID().GetOp(),
				Val: id,
			}
		}
		if conds.EntIDs != nil {
			ids := []uuid.UUID{}
			for _, id := range conds.GetEntIDs().GetValue() {
				_id, err := uuid.Parse(id)
				if err != nil {
					return err
				}
				ids = append(ids, _id)
			}
			h.Conds.EntIDs = &cruder.Cond{
				Op:  conds.GetEntIDs().GetOp(),
				Val: ids,
			}
		}
		if conds.PoolCoinTypeID != nil {
			id, err := uuid.Parse(conds.GetPoolCoinTypeID().GetValue())
			if err != nil {
				return err
			}
			h.Conds.PoolCoinTypeID = &cruder.Cond{
				Op:  conds.GetPoolCoinTypeID().GetOp(),
				Val: id,
			}
		}
		return nil
	}
}

func WithOffset(offset int32) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.Offset = offset
		return nil
	}
}

func WithLimit(limit int32) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if limit == 0 {
			limit = constant.DefaultRowLimit
		}
		h.Limit = limit
		return nil
	}
}
