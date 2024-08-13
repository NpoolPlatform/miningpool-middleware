package fractionrule

import (
	"github.com/NpoolPlatform/go-service-framework/pkg/wlog"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	"github.com/NpoolPlatform/miningpool-middleware/pkg/db/ent"
	fractionruleent "github.com/NpoolPlatform/miningpool-middleware/pkg/db/ent/fractionrule"
	"github.com/shopspring/decimal"

	"github.com/google/uuid"
)

type Req struct {
	ID               *uint32
	EntID            *uuid.UUID
	PoolCoinTypeID   *uuid.UUID
	WithdrawInterval *uint32
	PayoutThreshold  *decimal.Decimal
	MinAmount        *decimal.Decimal
	WithdrawRate     *decimal.Decimal
	DeletedAt        *uint32
}

func CreateSet(c *ent.FractionRuleCreate, req *Req) *ent.FractionRuleCreate {
	if req.EntID != nil {
		c.SetEntID(*req.EntID)
	}
	if req.PoolCoinTypeID != nil {
		c.SetPoolCoinTypeID(*req.PoolCoinTypeID)
	}
	if req.WithdrawInterval != nil {
		c.SetWithdrawInterval(*req.WithdrawInterval)
	}
	if req.MinAmount != nil {
		c.SetMinAmount(*req.MinAmount)
	}
	if req.PayoutThreshold != nil {
		c.SetPayoutThreshold(*req.PayoutThreshold)
	}
	if req.WithdrawRate != nil {
		c.SetWithdrawRate(*req.WithdrawRate)
	}
	return c
}

func UpdateSet(u *ent.FractionRuleUpdateOne, req *Req) (*ent.FractionRuleUpdateOne, error) {
	if req.PoolCoinTypeID != nil {
		u = u.SetPoolCoinTypeID(*req.PoolCoinTypeID)
	}
	if req.WithdrawInterval != nil {
		u = u.SetWithdrawInterval(*req.WithdrawInterval)
	}
	if req.MinAmount != nil {
		u = u.SetMinAmount(*req.MinAmount)
	}
	if req.PayoutThreshold != nil {
		u = u.SetPayoutThreshold(*req.PayoutThreshold)
	}
	if req.WithdrawRate != nil {
		u = u.SetWithdrawRate(*req.WithdrawRate)
	}
	if req.DeletedAt != nil {
		u = u.SetDeletedAt(*req.DeletedAt)
	}
	return u, nil
}

type Conds struct {
	ID               *cruder.Cond
	EntID            *cruder.Cond
	MiningpoolType   *cruder.Cond
	PoolCoinTypeID   *cruder.Cond
	WithdrawInterval *cruder.Cond
	EntIDs           *cruder.Cond
}

func SetQueryConds(q *ent.FractionRuleQuery, conds *Conds) (*ent.FractionRuleQuery, error) { //nolint
	if conds == nil {
		return nil, wlog.Errorf("have no any conds")
	}
	if conds.ID != nil {
		id, ok := conds.ID.Val.(uint32)
		if !ok {
			return nil, wlog.Errorf("invalid id")
		}
		switch conds.ID.Op {
		case cruder.EQ:
			q.Where(fractionruleent.ID(id))
		default:
			return nil, wlog.Errorf("invalid id field")
		}
	}
	if conds.EntID != nil {
		id, ok := conds.EntID.Val.(uuid.UUID)
		if !ok {
			return nil, wlog.Errorf("invalid entid")
		}
		switch conds.EntID.Op {
		case cruder.EQ:
			q.Where(fractionruleent.EntID(id))
		default:
			return nil, wlog.Errorf("invalid entid field")
		}
	}
	if conds.EntIDs != nil {
		ids, ok := conds.EntIDs.Val.([]uuid.UUID)
		if !ok {
			return nil, wlog.Errorf("invalid entids")
		}
		switch conds.EntIDs.Op {
		case cruder.IN:
			q.Where(fractionruleent.EntIDIn(ids...))
		default:
			return nil, wlog.Errorf("invalid entids field")
		}
	}
	if conds.PoolCoinTypeID != nil {
		poolcointypeid, ok := conds.PoolCoinTypeID.Val.(uuid.UUID)
		if !ok {
			return nil, wlog.Errorf("invalid poolcointypeid")
		}
		switch conds.PoolCoinTypeID.Op {
		case cruder.EQ:
			q.Where(fractionruleent.PoolCoinTypeID(poolcointypeid))
		default:
			return nil, wlog.Errorf("invalid poolcointypeid field")
		}
	}
	return q, nil
}
