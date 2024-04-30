package coin

import (
	"fmt"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/miningpool/v1"
	"github.com/NpoolPlatform/miningpool-middleware/pkg/db/ent"
	coinent "github.com/NpoolPlatform/miningpool-middleware/pkg/db/ent/coin"
	"github.com/shopspring/decimal"

	"github.com/google/uuid"
)

type Req struct {
	ID                     *uint32
	EntID                  *uuid.UUID
	PoolID                 *uuid.UUID
	CoinType               *basetypes.CoinType
	RevenueType            *basetypes.RevenueType
	FeeRatio               *decimal.Decimal
	FixedRevenueAble       *bool
	LeastTransferAmount    *decimal.Decimal
	BenefitIntervalSeconds *uint32
	Remark                 *string
	DeletedAt              *uint32
}

func CreateSet(c *ent.CoinCreate, req *Req) *ent.CoinCreate {
	if req.EntID != nil {
		c.SetEntID(*req.EntID)
	}
	if req.PoolID != nil {
		c.SetPoolID(*req.PoolID)
	}
	if req.CoinType != nil {
		c.SetCoinType(req.CoinType.String())
	}
	if req.RevenueType != nil {
		c.SetCoinType(req.CoinType.String())
	}
	if req.FeeRatio != nil {
		c.SetFeeRatio(*req.FeeRatio)
	}
	if req.FixedRevenueAble != nil {
		c.SetFixedRevenueAble(*req.FixedRevenueAble)
	}
	if req.LeastTransferAmount != nil {
		c.SetLeastTransferAmount(*req.LeastTransferAmount)
	}
	if req.BenefitIntervalSeconds != nil {
		c.SetBenefitIntervalSeconds(*req.BenefitIntervalSeconds)
	}
	if req.Remark != nil {
		c.SetRemark(*req.Remark)
	}
	return c
}

func UpdateSet(u *ent.CoinUpdateOne, req *Req) (*ent.CoinUpdateOne, error) {
	if req.PoolID != nil {
		u = u.SetPoolID(*req.PoolID)
	}
	if req.CoinType != nil {
		u = u.SetCoinType(req.CoinType.String())
	}
	if req.RevenueType != nil {
		u = u.SetRevenueType(req.RevenueType.String())
	}
	if req.FeeRatio != nil {
		u = u.SetFeeRatio(*req.FeeRatio)
	}
	if req.FixedRevenueAble != nil {
		u = u.SetFixedRevenueAble(*req.FixedRevenueAble)
	}
	if req.LeastTransferAmount != nil {
		u = u.SetLeastTransferAmount(*req.LeastTransferAmount)
	}
	if req.BenefitIntervalSeconds != nil {
		u = u.SetBenefitIntervalSeconds(*req.BenefitIntervalSeconds)
	}
	if req.Remark != nil {
		u = u.SetRemark(*req.Remark)
	}
	if req.DeletedAt != nil {
		u = u.SetDeletedAt(*req.DeletedAt)
	}
	return u, nil
}

type Conds struct {
	ID               *cruder.Cond
	EntID            *cruder.Cond
	PoolID           *cruder.Cond
	CoinType         *cruder.Cond
	MiningpoolType   *cruder.Cond
	RevenueType      *cruder.Cond
	FixedRevenueAble *cruder.Cond
	EntIDs           *cruder.Cond
}

func SetQueryConds(q *ent.CoinQuery, conds *Conds) (*ent.CoinQuery, error) { //nolint
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
			q.Where(coinent.ID(id))
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
			q.Where(coinent.EntID(id))
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
			q.Where(coinent.EntIDIn(ids...))
		default:
			return nil, fmt.Errorf("invalid entids field")
		}
	}
	if conds.PoolID != nil {
		poolid, ok := conds.PoolID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid poolid")
		}
		switch conds.PoolID.Op {
		case cruder.EQ:
			q.Where(coinent.PoolID(poolid))
		default:
			return nil, fmt.Errorf("invalid poolid field")
		}
	}
	if conds.CoinType != nil {
		cointype, ok := conds.CoinType.Val.(basetypes.CoinType)
		if !ok {
			return nil, fmt.Errorf("invalid cointype")
		}
		switch conds.CoinType.Op {
		case cruder.EQ:
			q.Where(coinent.CoinType(cointype.String()))
		default:
			return nil, fmt.Errorf("invalid cointype field")
		}
	}
	if conds.RevenueType != nil {
		revenuetype, ok := conds.RevenueType.Val.(basetypes.RevenueType)
		if !ok {
			return nil, fmt.Errorf("invalid revenuetype")
		}
		switch conds.RevenueType.Op {
		case cruder.EQ:
			q.Where(coinent.RevenueType(revenuetype.String()))
		default:
			return nil, fmt.Errorf("invalid revenuetype field")
		}
	}
	if conds.FixedRevenueAble != nil {
		fixedrevenueable, ok := conds.FixedRevenueAble.Val.(bool)
		if !ok {
			return nil, fmt.Errorf("invalid fixedrevenueable")
		}
		switch conds.FixedRevenueAble.Op {
		case cruder.EQ:
			q.Where(coinent.FixedRevenueAble(fixedrevenueable))
		default:
			return nil, fmt.Errorf("invalid fixedrevenueable field")
		}
	}
	return q, nil
}
