package fraction

import (
	"context"
	"fmt"

	basetypes "github.com/NpoolPlatform/message/npool/basetypes/miningpool/v1"
	npool "github.com/NpoolPlatform/message/npool/miningpool/mw/v1/fraction"
	constant "github.com/NpoolPlatform/miningpool-middleware/pkg/const"
	fractioncrud "github.com/NpoolPlatform/miningpool-middleware/pkg/crud/fraction"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"

	"github.com/google/uuid"
)

type Handler struct {
	ID            *uint32
	EntID         *uuid.UUID
	OrderUserID   *uuid.UUID
	WithdrawState *basetypes.WithdrawState
	WithdrawTime  *uint32
	PayTime       *uint32
	Reqs          []*fractioncrud.Req
	Conds         *fractioncrud.Conds
	Offset        int32
	Limit         int32
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

func WithOrderUserID(orderuserid *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if orderuserid == nil {
			if must {
				return fmt.Errorf("invalid orderuserid")
			}
			return nil
		}
		_id, err := uuid.Parse(*orderuserid)
		if err != nil {
			return err
		}
		h.OrderUserID = &_id
		return nil
	}
}

func WithWithdrawState(withdrawstate *basetypes.WithdrawState, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if withdrawstate == nil {
			if must {
				return fmt.Errorf("invalid withdrawstate")
			}
			return nil
		}
		if withdrawstate == basetypes.WithdrawState_DefaultWithdrawState.Enum() {
			return fmt.Errorf("invalid withdrawstate,not allow be default type")
		}
		h.WithdrawState = withdrawstate
		return nil
	}
}

func WithWithdrawTime(withdrawtime *uint32, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if withdrawtime == nil {
			if must {
				return fmt.Errorf("invalid withdrawtime")
			}
			return nil
		}
		h.WithdrawTime = withdrawtime
		return nil
	}
}

func WithPayTime(paytime *uint32, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if paytime == nil {
			if must {
				return fmt.Errorf("invalid paytime")
			}
			return nil
		}
		h.PayTime = paytime
		return nil
	}
}

// nolint:gocognit
func WithReqs(reqs []*npool.FractionReq, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		_reqs := []*fractioncrud.Req{}
		for _, req := range reqs {
			_req := &fractioncrud.Req{}
			if req.EntID != nil {
				id, err := uuid.Parse(req.GetEntID())
				if err != nil {
					return err
				}
				_req.EntID = &id
			}
			if req.OrderUserID != nil {
				id, err := uuid.Parse(req.GetEntID())
				if err != nil {
					return err
				}
				_req.OrderUserID = &id
			}
			if req.WithdrawState != nil {
				if req.WithdrawState == basetypes.WithdrawState_DefaultWithdrawState.Enum() {
					return fmt.Errorf("invalid withdrawstate,not allow be default type")
				}
				_req.WithdrawState = req.WithdrawState
			}
			if req.WithdrawTime != nil {
				_req.WithdrawTime = req.WithdrawTime
			}
			if req.PayTime != nil {
				_req.PayTime = req.PayTime
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
		h.Conds = &fractioncrud.Conds{}
		if conds == nil {
			return nil
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
		if conds.OrderUserID != nil {
			id, err := uuid.Parse(conds.GetOrderUserID().GetValue())
			if err != nil {
				return err
			}
			h.Conds.OrderUserID = &cruder.Cond{
				Op:  conds.GetOrderUserID().GetOp(),
				Val: id,
			}
		}
		if conds.WithdrawState != nil {
			h.Conds.WithdrawState = &cruder.Cond{
				Op:  conds.GetWithdrawState().GetOp(),
				Val: conds.GetWithdrawState().GetValue(),
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
