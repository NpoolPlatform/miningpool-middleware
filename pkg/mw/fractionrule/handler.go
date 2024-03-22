package fractionrule

import (
	"context"
	"fmt"

	basetypes "github.com/NpoolPlatform/message/npool/basetypes/miningpool/v1"
	npool "github.com/NpoolPlatform/message/npool/miningpool/mw/v1/fractionrule"
	constant "github.com/NpoolPlatform/miningpool-middleware/pkg/const"
	fractionrulecrud "github.com/NpoolPlatform/miningpool-middleware/pkg/crud/fractionrule"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"

	"github.com/google/uuid"
)

type Handler struct {
	ID               *uint32
	EntID            *uuid.UUID
	MiningpoolType   *basetypes.MiningpoolType
	CoinType         *basetypes.CoinType
	WithdrawInterval *uint32
	MinAmount        *float32
	WithdrawRate     *float32
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

func WithMiningpoolType(miningpooltype *basetypes.MiningpoolType, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if miningpooltype == nil {
			if must {
				return fmt.Errorf("invalid miningpooltype")
			}
			return nil
		}
		if miningpooltype == basetypes.MiningpoolType_DefaultMiningpoolType.Enum() {
			return fmt.Errorf("invalid miningpooltype,not allow be default type")
		}
		h.MiningpoolType = miningpooltype
		return nil
	}
}

func WithCoinType(cointype *basetypes.CoinType, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if cointype == nil {
			if must {
				return fmt.Errorf("invalid cointype")
			}
			return nil
		}
		if cointype == basetypes.CoinType_DefaultCoinType.Enum() {
			return fmt.Errorf("invalid cointype,not allow be default type")
		}
		h.CoinType = cointype
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

func WithMinAmount(minamount *float32, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if minamount == nil {
			if must {
				return fmt.Errorf("invalid minamount")
			}
			return nil
		}
		h.MinAmount = minamount
		return nil
	}
}

func WithWithdrawRate(withdrawrate *float32, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if withdrawrate == nil {
			if must {
				return fmt.Errorf("invalid withdrawrate")
			}
			return nil
		}
		h.WithdrawRate = withdrawrate
		return nil
	}
}

// nolint:gocognit
func WithReqs(reqs []*npool.FractionRuleReq, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		_reqs := []*fractionrulecrud.Req{}
		for _, req := range reqs {
			_req := &fractionrulecrud.Req{}
			if req.EntID != nil {
				id, err := uuid.Parse(req.GetEntID())
				if err != nil {
					return err
				}
				_req.EntID = &id
			}
			if req.MiningpoolType != nil {
				if req.MiningpoolType == basetypes.MiningpoolType_DefaultMiningpoolType.Enum() {
					return fmt.Errorf("invalid miningpooltype,not allow be default type")
				}
				_req.MiningpoolType = req.MiningpoolType
			}
			if req.CoinType != nil {
				if req.CoinType == basetypes.CoinType_DefaultCoinType.Enum() {
					return fmt.Errorf("invalid cointype,not allow be default type")
				}
				_req.CoinType = req.CoinType
			}
			if req.WithdrawInterval != nil {
				_req.WithdrawInterval = req.WithdrawInterval
			}
			if req.MinAmount != nil {
				_req.MinAmount = req.MinAmount
			}
			if req.WithdrawRate != nil {
				_req.WithdrawRate = req.WithdrawRate
			}
			_reqs = append(_reqs, _req)
		}
		h.Reqs = _reqs
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
		if conds.MiningpoolType != nil {
			h.Conds.MiningpoolType = &cruder.Cond{
				Op:  conds.GetMiningpoolType().GetOp(),
				Val: basetypes.MiningpoolType(conds.GetMiningpoolType().GetValue()),
			}
		}
		if conds.CoinType != nil {
			h.Conds.CoinType = &cruder.Cond{
				Op:  conds.GetCoinType().GetOp(),
				Val: basetypes.CoinType(conds.GetCoinType().GetValue()),
			}
		}
		if conds.WithdrawInterval != nil {
			h.Conds.WithdrawInterval = &cruder.Cond{
				Op:  conds.GetWithdrawInterval().GetOp(),
				Val: conds.GetWithdrawInterval().GetValue(),
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
