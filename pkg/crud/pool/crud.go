package pool

import (
	"fmt"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/miningpool/v1"
	"github.com/NpoolPlatform/miningpool-middleware/pkg/db/ent"
	poolent "github.com/NpoolPlatform/miningpool-middleware/pkg/db/ent/pool"

	"github.com/google/uuid"
)

type Req struct {
	EntID          *uuid.UUID
	MiningpoolType *basetypes.MiningpoolType
	Name           *string
	Site           *string
	Description    *string
	DeletedAt      *uint32
}

func CreateSet(c *ent.PoolCreate, req *Req) *ent.PoolCreate {
	if req.EntID != nil {
		c.SetEntID(*req.EntID)
	}
	if req.MiningpoolType != nil {
		c.SetMiningpoolType(req.MiningpoolType.String())
	}
	if req.Name != nil {
		c.SetName(*req.Name)
	}
	if req.Site != nil {
		c.SetSite(*req.Site)
	}
	if req.Description != nil {
		c.SetDescription(*req.Description)
	}
	c.Mutation().Where()
	return c
}

func UpdateSet(u *ent.PoolUpdateOne, req *Req) (*ent.PoolUpdateOne, error) {
	if req.MiningpoolType != nil {
		u = u.SetMiningpoolType(req.MiningpoolType.String())
	}
	if req.Name != nil {
		u = u.SetName(*req.Name)
	}
	if req.Site != nil {
		u = u.SetSite(*req.Site)
	}
	if req.Description != nil {
		u = u.SetDescription(*req.Description)
	}
	if req.DeletedAt != nil {
		u = u.SetDeletedAt(*req.DeletedAt)
	}
	return u, nil
}

type Conds struct {
	ID             *cruder.Cond
	EntID          *cruder.Cond
	MiningpoolType *cruder.Cond
	Name           *cruder.Cond
	Site           *cruder.Cond
	Description    *cruder.Cond
	EntIDs         *cruder.Cond
}

func SetQueryConds(q *ent.PoolQuery, conds *Conds) (*ent.PoolQuery, error) { //nolint
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
			q.Where(poolent.ID(id))
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
			q.Where(poolent.EntID(id))
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
			q.Where(poolent.EntIDIn(ids...))
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
			q.Where(poolent.MiningpoolType(miningpooltype.String()))
		default:
			return nil, fmt.Errorf("invalid miningpooltype field")
		}
	}
	if conds.Name != nil {
		name, ok := conds.Name.Val.(string)
		if !ok {
			return nil, fmt.Errorf("invalid name")
		}
		switch conds.Name.Op {
		case cruder.EQ:
			q.Where(poolent.Name(name))
		default:
			return nil, fmt.Errorf("invalid name field")
		}
	}
	if conds.Site != nil {
		site, ok := conds.Site.Val.(string)
		if !ok {
			return nil, fmt.Errorf("invalid site")
		}
		switch conds.Site.Op {
		case cruder.EQ:
			q.Where(poolent.Site(site))
		default:
			return nil, fmt.Errorf("invalid site field")
		}
	}
	if conds.Description != nil {
		description, ok := conds.Description.Val.(string)
		if !ok {
			return nil, fmt.Errorf("invalid description")
		}
		switch conds.Description.Op {
		case cruder.EQ:
			q.Where(poolent.Description(description))
		default:
			return nil, fmt.Errorf("invalid description field")
		}
	}
	return q, nil
}
