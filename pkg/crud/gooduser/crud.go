package gooduser

import (
	"fmt"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/miningpool/v1"
	"github.com/NpoolPlatform/miningpool-middleware/pkg/db/ent"
	gooduserent "github.com/NpoolPlatform/miningpool-middleware/pkg/db/ent/gooduser"

	"github.com/google/uuid"
)

type Req struct {
	EntID          *uuid.UUID
	RootUserID     *uuid.UUID
	Name           *string
	MiningpoolType *basetypes.MiningpoolType
	CoinType       *basetypes.CoinType
	HashRate       *float32
	ReadPageLink   *string
	RevenueType    *basetypes.RevenueType
	DeletedAt      *uint32
}

func CreateSet(c *ent.GoodUserCreate, req *Req) *ent.GoodUserCreate {
	if req.EntID != nil {
		c.SetEntID(*req.EntID)
	}
	if req.RootUserID != nil {
		c.SetRootUserID(*req.RootUserID)
	}
	if req.Name != nil {
		c.SetName(*req.Name)
	}
	if req.MiningpoolType != nil {
		c.SetMiningpoolType(req.MiningpoolType.String())
	}
	if req.CoinType != nil {
		c.SetCoinType(req.CoinType.String())
	}
	if req.HashRate != nil {
		c.SetHashRate(*req.HashRate)
	}
	if req.ReadPageLink != nil {
		c.SetReadPageLink(*req.ReadPageLink)
	}
	if req.RevenueType != nil {
		c.SetRevenueType(req.RevenueType.String())
	}
	return c
}

func UpdateSet(u *ent.GoodUserUpdateOne, req *Req) (*ent.GoodUserUpdateOne, error) {
	if req.RootUserID != nil {
		u = u.SetRootUserID(*req.RootUserID)
	}
	if req.Name != nil {
		u = u.SetName(*req.Name)
	}
	if req.MiningpoolType != nil {
		u = u.SetMiningpoolType(req.MiningpoolType.String())
	}
	if req.CoinType != nil {
		u = u.SetCoinType(req.CoinType.String())
	}
	if req.HashRate != nil {
		u = u.SetHashRate(*req.HashRate)
	}
	if req.ReadPageLink != nil {
		u = u.SetReadPageLink(*req.ReadPageLink)
	}
	if req.RevenueType != nil {
		u = u.SetRevenueType(req.RevenueType.String())
	}
	if req.DeletedAt != nil {
		u = u.SetDeletedAt(*req.DeletedAt)
	}
	return u, nil
}

type Conds struct {
	ID             *cruder.Cond
	EntID          *cruder.Cond
	Name           *cruder.Cond
	RootUserID     *cruder.Cond
	MiningpoolType *cruder.Cond
	CoinType       *cruder.Cond
	RevenueType    *cruder.Cond
	EntIDs         *cruder.Cond
}

func SetQueryConds(q *ent.GoodUserQuery, conds *Conds) (*ent.GoodUserQuery, error) { //nolint
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
			q.Where(gooduserent.ID(id))
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
			q.Where(gooduserent.EntID(id))
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
			q.Where(gooduserent.EntIDIn(ids...))
		default:
			return nil, fmt.Errorf("invalid entids field")
		}
	}
	if conds.RootUserID != nil {
		id, ok := conds.RootUserID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid rootuserid")
		}
		switch conds.RootUserID.Op {
		case cruder.EQ:
			q.Where(gooduserent.RootUserID(id))
		default:
			return nil, fmt.Errorf("invalid rootuserid field")
		}
	}
	if conds.Name != nil {
		name, ok := conds.Name.Val.(string)
		if !ok {
			return nil, fmt.Errorf("invalid name")
		}
		switch conds.Name.Op {
		case cruder.EQ:
			q.Where(gooduserent.Name(name))
		default:
			return nil, fmt.Errorf("invalid name field")
		}
	}
	if conds.MiningpoolType != nil {
		miningpooltype, ok := conds.MiningpoolType.Val.(basetypes.MiningpoolType)
		if !ok {
			return nil, fmt.Errorf("invalid miningpooltype")
		}
		switch conds.MiningpoolType.Op {
		case cruder.EQ:
			q.Where(gooduserent.MiningpoolType(miningpooltype.String()))
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
			q.Where(gooduserent.CoinType(cointype.String()))
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
			q.Where(gooduserent.RevenueType(revenuetype.String()))
		default:
			return nil, fmt.Errorf("invalid revenuetype field")
		}
	}
	return q, nil
}
