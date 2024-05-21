package rootuser

import (
	"fmt"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	"github.com/NpoolPlatform/miningpool-middleware/pkg/db/ent"
	rootuserent "github.com/NpoolPlatform/miningpool-middleware/pkg/db/ent/rootuser"

	"github.com/google/uuid"
)

type Req struct {
	ID        *uint32
	EntID     *uuid.UUID
	Name      *string
	PoolID    *uuid.UUID
	Email     *string
	AuthToken *string
	Authed    *bool
	Remark    *string
	DeletedAt *uint32
}

func CreateSet(c *ent.RootUserCreate, req *Req) *ent.RootUserCreate {
	if req.EntID != nil {
		c.SetEntID(*req.EntID)
	}
	if req.Name != nil {
		c.SetName(*req.Name)
	}
	if req.PoolID != nil {
		c.SetPoolID(*req.PoolID)
	}
	if req.Email != nil {
		c.SetEmail(*req.Email)
	}
	if req.AuthToken != nil {
		c.SetAuthToken(*req.AuthToken)
	}
	if req.Authed != nil {
		c.SetAuthed(*req.Authed)
	}
	if req.Remark != nil {
		c.SetRemark(*req.Remark)
	}
	return c
}

func UpdateSet(u *ent.RootUserUpdateOne, req *Req) (*ent.RootUserUpdateOne, error) {
	if req.Name != nil {
		u = u.SetName(*req.Name)
	}
	if req.PoolID != nil {
		u = u.SetPoolID(*req.PoolID)
	}
	if req.Email != nil {
		u = u.SetEmail(*req.Email)
	}
	if req.AuthToken != nil {
		u = u.SetAuthToken(*req.AuthToken)
	}
	if req.Authed != nil {
		u = u.SetAuthed(*req.Authed)
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
	ID     *cruder.Cond
	EntID  *cruder.Cond
	Name   *cruder.Cond
	PoolID *cruder.Cond
	Email  *cruder.Cond
	Authed *cruder.Cond
	EntIDs *cruder.Cond
}

func SetQueryConds(q *ent.RootUserQuery, conds *Conds) (*ent.RootUserQuery, error) { //nolint
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
			q.Where(rootuserent.ID(id))
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
			q.Where(rootuserent.EntID(id))
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
			q.Where(rootuserent.EntIDIn(ids...))
		default:
			return nil, fmt.Errorf("invalid entids field")
		}
	}
	if conds.Name != nil {
		name, ok := conds.Name.Val.(string)
		if !ok {
			return nil, fmt.Errorf("invalid name")
		}
		switch conds.Name.Op {
		case cruder.EQ:
			q.Where(rootuserent.Name(name))
		default:
			return nil, fmt.Errorf("invalid name field")
		}
	}
	if conds.PoolID != nil {
		id, ok := conds.PoolID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid poolid")
		}
		switch conds.PoolID.Op {
		case cruder.EQ:
			q.Where(rootuserent.PoolID(id))
		default:
			return nil, fmt.Errorf("invalid poolid field")
		}
	}
	if conds.Email != nil {
		email, ok := conds.Email.Val.(string)
		if !ok {
			return nil, fmt.Errorf("invalid email")
		}
		switch conds.Email.Op {
		case cruder.EQ:
			q.Where(rootuserent.Email(email))
		default:
			return nil, fmt.Errorf("invalid email field")
		}
	}
	if conds.Authed != nil {
		authed, ok := conds.Authed.Val.(bool)
		if !ok {
			return nil, fmt.Errorf("invalid authed")
		}
		switch conds.Authed.Op {
		case cruder.EQ:
			q.Where(rootuserent.Authed(authed))
		default:
			return nil, fmt.Errorf("invalid authed field")
		}
	}
	return q, nil
}
