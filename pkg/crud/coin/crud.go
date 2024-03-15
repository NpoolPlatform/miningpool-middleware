package coin

import (
	"fmt"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/miningpool/v1"
	"github.com/NpoolPlatform/miningpool-middleware/pkg/db/ent"
	coinent "github.com/NpoolPlatform/miningpool-middleware/pkg/db/ent/coin"

	"github.com/google/uuid"
)

type Req struct {
	EntID            *uuid.UUID
	MiningpoolType   *basetypes.MiningpoolType
	CoinType         *basetypes.CoinType
	RevenueTypes     *[]basetypes.RevenueType
	Site             *string
	FeeRate          *float32
	FixedRevenueAble *bool
	Threshold        *float32
	Remark           *string
	DeletedAt        *uint32
}

func CreateSet(c *ent.CoinCreate, req *Req) *ent.CoinCreate {
	if req.EntID != nil {
		c.SetEntID(*req.EntID)
	}
	if req.MiningpoolType != nil {
		c.SetMiningpoolType(req.MiningpoolType.String())
	}
	if req.CoinType != nil {
		c.SetCoinType(req.CoinType.String())
	}
	if req.RevenueTypes != nil {
		revenueTypes := []string{}
		for _, v := range *req.RevenueTypes {
			revenueTypes = append(revenueTypes, v.String())
		}
		c.SetRevenueTypes(revenueTypes)
	}
	if req.Site != nil {
		c.SetSite(*req.Site)
	}
	if req.FeeRate != nil {
		c.SetFeeRate(*req.FeeRate)
	}
	if req.FixedRevenueAble != nil {
		c.SetFixedRevenueAble(*req.FixedRevenueAble)
	}
	if req.Threshold != nil {
		c.SetThreshold(*req.Threshold)
	}
	if req.Remark != nil {
		c.SetRemark(*req.Remark)
	}
	return c
}

func UpdateSet(u *ent.CoinUpdateOne, req *Req) (*ent.CoinUpdateOne, error) {
	if req.MiningpoolType != nil {
		u = u.SetMiningpoolType(req.MiningpoolType.String())
	}
	if req.CoinType != nil {
		u = u.SetCoinType(req.CoinType.String())
	}
	if req.RevenueTypes != nil {
		revenueTypes := []string{}
		for _, v := range *req.RevenueTypes {
			revenueTypes = append(revenueTypes, v.String())
		}
		u = u.SetRevenueTypes(revenueTypes)
	}
	if req.Site != nil {
		u = u.SetSite(*req.Site)
	}
	if req.FeeRate != nil {
		u = u.SetFeeRate(*req.FeeRate)
	}
	if req.FixedRevenueAble != nil {
		u = u.SetFixedRevenueAble(*req.FixedRevenueAble)
	}
	if req.Threshold != nil {
		u = u.SetThreshold(*req.Threshold)
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
	EntID            *cruder.Cond
	MiningpoolType   *cruder.Cond
	CoinType         *cruder.Cond
	Site             *cruder.Cond
	FixedRevenueAble *cruder.Cond
	EntIDs           *cruder.Cond
}

func SetQueryConds(q *ent.CoinQuery, conds *Conds) (*ent.CoinQuery, error) { //nolint
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
	if conds.MiningpoolType != nil {
		miningpooltype, ok := conds.MiningpoolType.Val.(basetypes.MiningpoolType)
		if !ok {
			return nil, fmt.Errorf("invalid miningpooltype")
		}
		switch conds.MiningpoolType.Op {
		case cruder.EQ:
			q.Where(coinent.MiningpoolType(miningpooltype.String()))
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
			q.Where(coinent.CoinType(cointype.String()))
		default:
			return nil, fmt.Errorf("invalid cointype field")
		}
	}
	if conds.Site != nil {
		site, ok := conds.Site.Val.(string)
		if !ok {
			return nil, fmt.Errorf("invalid site")
		}
		switch conds.Site.Op {
		case cruder.EQ:
			q.Where(coinent.Site(site))
		default:
			return nil, fmt.Errorf("invalid site field")
		}
	}
	if conds.FixedRevenueAble != nil {
		fixedrevenueable, ok := conds.FixedRevenueAble.Val.(string)
		if !ok {
			return nil, fmt.Errorf("invalid fixedrevenueable")
		}
		switch conds.FixedRevenueAble.Op {
		case cruder.EQ:
			q.Where(coinent.Site(fixedrevenueable))
		default:
			return nil, fmt.Errorf("invalid fixedrevenueable field")
		}
	}
	return q, nil
}
