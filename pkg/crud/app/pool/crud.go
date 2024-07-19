package apppool

import (
	"github.com/NpoolPlatform/go-service-framework/pkg/wlog"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	"github.com/NpoolPlatform/miningpool-middleware/pkg/db/ent"
	apppoolent "github.com/NpoolPlatform/miningpool-middleware/pkg/db/ent/apppool"

	"github.com/google/uuid"
)

type Req struct {
	ID        *uint32
	EntID     *uuid.UUID
	AppID     *uuid.UUID
	PoolID    *uuid.UUID
	DeletedAt *uint32
}

func CreateSet(c *ent.AppPoolCreate, req *Req) *ent.AppPoolCreate {
	if req.EntID != nil {
		c.SetEntID(*req.EntID)
	}
	if req.AppID != nil {
		c.SetAppID(*req.AppID)
	}
	if req.PoolID != nil {
		c.SetPoolID(*req.PoolID)
	}
	return c
}

func UpdateSet(u *ent.AppPoolUpdateOne, req *Req) (*ent.AppPoolUpdateOne, error) {
	if req.AppID != nil {
		u = u.SetAppID(*req.AppID)
	}
	if req.PoolID != nil {
		u = u.SetPoolID(*req.PoolID)
	}
	if req.DeletedAt != nil {
		u = u.SetDeletedAt(*req.DeletedAt)
	}
	return u, nil
}

type Conds struct {
	ID     *cruder.Cond
	EntID  *cruder.Cond
	AppID  *cruder.Cond
	PoolID *cruder.Cond
	EntIDs *cruder.Cond
}

func SetQueryConds(q *ent.AppPoolQuery, conds *Conds) (*ent.AppPoolQuery, error) { //nolint
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
			q.Where(apppoolent.ID(id))
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
			q.Where(apppoolent.EntID(id))
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
			q.Where(apppoolent.EntIDIn(ids...))
		default:
			return nil, wlog.Errorf("invalid entids field")
		}
	}
	if conds.AppID != nil {
		id, ok := conds.AppID.Val.(uuid.UUID)
		if !ok {
			return nil, wlog.Errorf("invalid appid")
		}
		switch conds.AppID.Op {
		case cruder.EQ:
			q.Where(apppoolent.AppID(id))
		default:
			return nil, wlog.Errorf("invalid appid field")
		}
	}
	if conds.PoolID != nil {
		id, ok := conds.PoolID.Val.(uuid.UUID)
		if !ok {
			return nil, wlog.Errorf("invalid poolid")
		}
		switch conds.PoolID.Op {
		case cruder.EQ:
			q.Where(apppoolent.PoolID(id))
		default:
			return nil, wlog.Errorf("invalid poolid field")
		}
	}
	return q, nil
}
