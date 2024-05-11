package gooduser

import (
	"fmt"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	"github.com/NpoolPlatform/miningpool-middleware/pkg/db/ent"
	gooduserent "github.com/NpoolPlatform/miningpool-middleware/pkg/db/ent/gooduser"

	"github.com/google/uuid"
)

type Req struct {
	ID           *uint32
	EntID        *uuid.UUID
	CoinID       *uuid.UUID
	RootUserID   *uuid.UUID
	Name         *string
	ReadPageLink *string
	DeletedAt    *uint32
}

func CreateSet(c *ent.GoodUserCreate, req *Req) *ent.GoodUserCreate {
	if req.EntID != nil {
		c.SetEntID(*req.EntID)
	}
	if req.CoinID != nil {
		c.SetCoinID(*req.CoinID)
	}
	if req.RootUserID != nil {
		c.SetRootUserID(*req.RootUserID)
	}
	if req.Name != nil {
		c.SetName(*req.Name)
	}
	if req.ReadPageLink != nil {
		c.SetReadPageLink(*req.ReadPageLink)
	}
	return c
}

func UpdateSet(u *ent.GoodUserUpdateOne, req *Req) (*ent.GoodUserUpdateOne, error) {
	if req.CoinID != nil {
		u = u.SetCoinID(*req.CoinID)
	}
	if req.RootUserID != nil {
		u = u.SetRootUserID(*req.RootUserID)
	}
	if req.Name != nil {
		u = u.SetName(*req.Name)
	}
	if req.ReadPageLink != nil {
		u = u.SetReadPageLink(*req.ReadPageLink)
	}
	if req.DeletedAt != nil {
		u = u.SetDeletedAt(*req.DeletedAt)
	}
	return u, nil
}

type Conds struct {
	ID         *cruder.Cond
	EntID      *cruder.Cond
	Name       *cruder.Cond
	RootUserID *cruder.Cond
	CoinID     *cruder.Cond
	EntIDs     *cruder.Cond
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
	if conds.CoinID != nil {
		id, ok := conds.CoinID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid coinid")
		}
		switch conds.CoinID.Op {
		case cruder.EQ:
			q.Where(gooduserent.CoinID(id))
		default:
			return nil, fmt.Errorf("invalid coinid field")
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
	return q, nil
}
