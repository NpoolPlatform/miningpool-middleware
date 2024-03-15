package fraction

import (
	"fmt"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/miningpool/v1"
	"github.com/NpoolPlatform/miningpool-middleware/pkg/db/ent"
	fractionent "github.com/NpoolPlatform/miningpool-middleware/pkg/db/ent/fraction"

	"github.com/google/uuid"
)

type Req struct {
	EntID         *uuid.UUID
	OrderUserID   *uuid.UUID
	WithdrawState *basetypes.WithdrawState
	WithdrawTime  *uint32
	PayTime       *uint32
	DeletedAt     *uint32
}

func CreateSet(c *ent.FractionCreate, req *Req) *ent.FractionCreate {
	if req.EntID != nil {
		c.SetEntID(*req.EntID)
	}
	if req.OrderUserID != nil {
		c.SetOrderUserID(*req.OrderUserID)
	}
	if req.WithdrawState != nil {
		c.SetWithdrawState(req.WithdrawState.String())
	}
	if req.WithdrawTime != nil {
		c.SetWithdrawTime(*req.WithdrawTime)
	}
	if req.PayTime != nil {
		c.SetPayTime(*req.PayTime)
	}
	return c
}

func UpdateSet(u *ent.FractionUpdateOne, req *Req) (*ent.FractionUpdateOne, error) {
	if req.OrderUserID != nil {
		u = u.SetOrderUserID(*req.OrderUserID)
	}
	if req.WithdrawState != nil {
		u = u.SetWithdrawState(req.WithdrawState.String())
	}
	if req.WithdrawTime != nil {
		u = u.SetWithdrawTime(*req.WithdrawTime)
	}
	if req.PayTime != nil {
		u = u.SetPayTime(*req.PayTime)
	}
	if req.DeletedAt != nil {
		u = u.SetDeletedAt(*req.DeletedAt)
	}
	return u, nil
}

type Conds struct {
	EntID         *cruder.Cond
	OrderUserID   *cruder.Cond
	WithdrawState *cruder.Cond
	EntIDs        *cruder.Cond
}

func SetQueryConds(q *ent.FractionQuery, conds *Conds) (*ent.FractionQuery, error) { //nolint
	if conds == nil {
		return nil, fmt.Errorf("have no any conds")
	}
	if conds.EntID != nil {
		id, ok := conds.EntID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid entid")
		}
		switch conds.EntID.Op {
		case cruder.EQ:
			q.Where(fractionent.EntID(id))
		default:
			return nil, fmt.Errorf("invalid entid field")
		}
	}
	if conds.EntIDs != nil {
		ids, ok := conds.EntIDs.Val.([]uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid entids")
		}
		switch conds.EntIDs.Op {
		case cruder.IN:
			q.Where(fractionent.EntIDIn(ids...))
		default:
			return nil, fmt.Errorf("invalid entids field")
		}
	}
	if conds.OrderUserID != nil {
		orderuserid, ok := conds.OrderUserID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid orderuserid")
		}
		switch conds.OrderUserID.Op {
		case cruder.EQ:
			q.Where(fractionent.OrderUserID(orderuserid))
		default:
			return nil, fmt.Errorf("invalid orderuserid field")
		}
	}
	if conds.WithdrawState != nil {
		withdrawstate, ok := conds.WithdrawState.Val.(basetypes.WithdrawState)
		if !ok {
			return nil, fmt.Errorf("invalid withdrawstate")
		}
		switch conds.WithdrawState.Op {
		case cruder.EQ:
			q.Where(fractionent.WithdrawState(withdrawstate.String()))
		default:
			return nil, fmt.Errorf("invalid withdrawstate field")
		}
	}
	return q, nil
}
