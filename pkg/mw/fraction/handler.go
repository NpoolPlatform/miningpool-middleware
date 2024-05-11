package fraction

import (
	"context"
	"fmt"

	basetypes "github.com/NpoolPlatform/message/npool/basetypes/miningpool/v1"
	npool "github.com/NpoolPlatform/message/npool/miningpool/mw/v1/fraction"
	constant "github.com/NpoolPlatform/miningpool-middleware/pkg/const"
	fractioncrud "github.com/NpoolPlatform/miningpool-middleware/pkg/crud/fraction"
	"github.com/NpoolPlatform/miningpool-middleware/pkg/mw/orderuser"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"

	"github.com/google/uuid"
)

type Handler struct {
	fractioncrud.Req
	Reqs   []*fractioncrud.Req
	Conds  *fractioncrud.Conds
	Offset int32
	Limit  int32
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

func WithAppID(appid *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if appid == nil {
			if must {
				return fmt.Errorf("invalid appid")
			}
			return nil
		}
		_id, err := uuid.Parse(*appid)
		if err != nil {
			return err
		}
		h.AppID = &_id
		return nil
	}
}

func WithUserID(userid *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if userid == nil {
			if must {
				return fmt.Errorf("invalid userid")
			}
			return nil
		}
		_id, err := uuid.Parse(*userid)
		if err != nil {
			return err
		}
		h.UserID = &_id
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

		ouH, err := orderuser.NewHandler(ctx, orderuser.WithEntID(orderuserid, true))
		if err != nil {
			return err
		}
		exists, err := ouH.ExistOrderUser(ctx)
		if err != nil {
			return err
		}

		if !exists {
			return fmt.Errorf("invalid orderuserid")
		}

		h.OrderUserID = ouH.EntID
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
		if *withdrawstate == basetypes.WithdrawState_DefaultWithdrawState {
			return fmt.Errorf("invalid withdrawstate,not allow be default type")
		}
		h.WithdrawState = withdrawstate
		return nil
	}
}

func WithWithdrawAt(withdrawtime *uint32, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if withdrawtime == nil {
			if must {
				return fmt.Errorf("invalid withdrawtime")
			}
			return nil
		}
		h.WithdrawAt = withdrawtime
		return nil
	}
}

func WithPromisePayAt(PromisePayAt *uint32, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if PromisePayAt == nil {
			if must {
				return fmt.Errorf("invalid PromisePayAt")
			}
			return nil
		}
		h.PromisePayAt = PromisePayAt
		return nil
	}
}

func WithMsg(msg *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if msg == nil {
			if must {
				return fmt.Errorf("invalid msg")
			}
			return nil
		}
		h.Msg = msg
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
		if conds.AppID != nil {
			id, err := uuid.Parse(conds.GetAppID().GetValue())
			if err != nil {
				return err
			}
			h.Conds.AppID = &cruder.Cond{
				Op:  conds.GetAppID().GetOp(),
				Val: id,
			}
		}
		if conds.UserID != nil {
			id, err := uuid.Parse(conds.GetUserID().GetValue())
			if err != nil {
				return err
			}
			h.Conds.UserID = &cruder.Cond{
				Op:  conds.GetUserID().GetOp(),
				Val: id,
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
				Val: basetypes.WithdrawState(conds.GetWithdrawState().GetValue()),
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
