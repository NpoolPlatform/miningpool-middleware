package fractionrule

import (
	"fmt"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/miningpool/v1"
	"github.com/NpoolPlatform/miningpool-middleware/pkg/db/ent"
	fractionruleent "github.com/NpoolPlatform/miningpool-middleware/pkg/db/ent/fractionrule"

	"github.com/google/uuid"
)

type Req struct {
	EntID            *uuid.UUID
	MiningpoolType   *basetypes.MiningpoolType
	CoinType         *basetypes.CoinType
	WithdrawInterval *uint32
	MinAmount        *float32
	WithdrawRate     *float32
	DeletedAt        *uint32
}

func CreateSet(c *ent.FractionRuleCreate, req *Req) *ent.FractionRuleCreate {
	if req.EntID != nil {
		c.SetEntID(*req.EntID)
	}
	if req.MiningpoolType != nil {
		c.SetMiningpoolType(req.MiningpoolType.String())
	}
	if req.CoinType != nil {
		c.SetCoinType(req.CoinType.String())
	}
	if req.WithdrawInterval != nil {
		c.SetWithdrawInterval(*req.WithdrawInterval)
	}
	if req.MinAmount != nil {
		c.SetMinAmount(*req.MinAmount)
	}
	if req.WithdrawRate != nil {
		c.SetWithdrawRate(*req.WithdrawRate)
	}
	return c
}

func UpdateSet(u *ent.FractionRuleUpdateOne, req *Req) (*ent.FractionRuleUpdateOne, error) {
	if req.MiningpoolType != nil {
		u = u.SetMiningpoolType(req.MiningpoolType.String())
	}
	if req.CoinType != nil {
		u = u.SetCoinType(req.CoinType.String())
	}
	if req.WithdrawInterval != nil {
		u = u.SetWithdrawInterval(*req.WithdrawInterval)
	}
	if req.MinAmount != nil {
		u = u.SetMinAmount(*req.MinAmount)
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
	CoinType         *cruder.Cond
	WithdrawInterval *cruder.Cond
	EntIDs           *cruder.Cond
}

func SetQueryConds(q *ent.FractionRuleQuery, conds *Conds) (*ent.FractionRuleQuery, error) { //nolint
	if conds == nil {
		return nil, fmt.Errorf("have no any conds")
	}
	if conds.ID != nil {
		id, ok := conds.ID.Val.(uint32)
		if !ok {
			return nil, fmt.Errorf("invalid id")
		}
		switch conds.ID.Op {
		case cruder.EQ:
			q.Where(fractionruleent.ID(id))
		default:
			return nil, fmt.Errorf("invalid id field")
		}
	}
	if conds.EntID != nil {
		id, ok := conds.EntID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid entid")
		}
		switch conds.EntID.Op {
		case cruder.EQ:
			q.Where(fractionruleent.EntID(id))
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
			q.Where(fractionruleent.EntIDIn(ids...))
		default:
			return nil, fmt.Errorf("invalid entids field")
		}
	}
	if conds.MiningpoolType != nil {
		miningpooltype, ok := conds.MiningpoolType.Val.(basetypes.MiningpoolType)
		if !ok {
			return nil, fmt.Errorf("invalid miningpooltype")
		}
		switch conds.MiningpoolType.Op {
		case cruder.EQ:
			q.Where(fractionruleent.MiningpoolType(miningpooltype.String()))
		default:
			return nil, fmt.Errorf("invalid miningpooltype field")
		}
	}
	if conds.CoinType != nil {
		cointype, ok := conds.CoinType.Val.(basetypes.CoinType)
		if !ok {
			return nil, fmt.Errorf("invalid cointype")
		}
		switch conds.CoinType.Op {
		case cruder.EQ:
			q.Where(fractionruleent.CoinType(cointype.String()))
		default:
			return nil, fmt.Errorf("invalid cointype field")
		}
	}
	if conds.WithdrawInterval != nil {
		withdrawinterval, ok := conds.WithdrawInterval.Val.(uint32)
		if !ok {
			return nil, fmt.Errorf("invalid withdrawinterval")
		}
		switch conds.WithdrawInterval.Op {
		case cruder.EQ:
			q.Where(fractionruleent.WithdrawInterval(withdrawinterval))
		default:
			return nil, fmt.Errorf("invalid withdrawinterval field")
		}
	}
	return q, nil
}
