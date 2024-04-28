package orderuser

import (
	"fmt"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	"github.com/NpoolPlatform/miningpool-middleware/pkg/db/ent"
	orderuserent "github.com/NpoolPlatform/miningpool-middleware/pkg/db/ent/orderuser"
	"github.com/shopspring/decimal"

	"github.com/google/uuid"
)

type Req struct {
	EntID          *uuid.UUID
	GoodUserID     *uuid.UUID
	AppID          *uuid.UUID
	UserID         *uuid.UUID
	Name           *string
	Proportion     *decimal.Decimal
	RevenueAddress *string
	ReadPageLink   *string
	AutoPay        *bool
	DeletedAt      *uint32
}

func CreateSet(c *ent.OrderUserCreate, req *Req) *ent.OrderUserCreate {
	if req.EntID != nil {
		c.SetEntID(*req.EntID)
	}
	if req.Name != nil {
		c.SetName(*req.Name)
	}
	if req.GoodUserID != nil {
		c.SetGoodUserID(*req.GoodUserID)
	}
	if req.AppID != nil {
		c.SetAppID(*req.AppID)
	}
	if req.UserID != nil {
		c.SetUserID(*req.UserID)
	}
	if req.Proportion != nil {
		c.SetProportion(*req.Proportion)
	}
	if req.RevenueAddress != nil {
		c.SetRevenueAddress(*req.RevenueAddress)
	}
	if req.ReadPageLink != nil {
		c.SetReadPageLink(*req.ReadPageLink)
	}
	if req.AutoPay != nil {
		c.SetAutoPay(*req.AutoPay)
	}
	return c
}

func UpdateSet(u *ent.OrderUserUpdateOne, req *Req) (*ent.OrderUserUpdateOne, error) {
	if req.Name != nil {
		u = u.SetName(*req.Name)
	}
	if req.GoodUserID != nil {
		u = u.SetGoodUserID(*req.GoodUserID)
	}
	if req.AppID != nil {
		u = u.SetAppID(*req.AppID)
	}
	if req.UserID != nil {
		u = u.SetUserID(*req.UserID)
	}
	if req.Proportion != nil {
		u = u.SetProportion(*req.Proportion)
	}
	if req.RevenueAddress != nil {
		u = u.SetRevenueAddress(*req.RevenueAddress)
	}
	if req.ReadPageLink != nil {
		u = u.SetReadPageLink(*req.ReadPageLink)
	}
	if req.AutoPay != nil {
		u = u.SetAutoPay(*req.AutoPay)
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
	GoodUserID     *cruder.Cond
	AppID          *cruder.Cond
	UserID         *cruder.Cond
	RevenueAddress *cruder.Cond
	AutoPay        *cruder.Cond
	EntIDs         *cruder.Cond
}

func SetQueryConds(q *ent.OrderUserQuery, conds *Conds) (*ent.OrderUserQuery, error) { //nolint
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
			q.Where(orderuserent.ID(id))
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
			q.Where(orderuserent.EntID(id))
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
			q.Where(orderuserent.EntIDIn(ids...))
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
			q.Where(orderuserent.Name(name))
		default:
			return nil, fmt.Errorf("invalid name field")
		}
	}
	if conds.GoodUserID != nil {
		id, ok := conds.GoodUserID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid gooduserid")
		}
		switch conds.GoodUserID.Op {
		case cruder.EQ:
			q.Where(orderuserent.GoodUserID(id))
		default:
			return nil, fmt.Errorf("invalid gooduserid field")
		}
	}
	if conds.AppID != nil {
		id, ok := conds.AppID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid appid")
		}
		switch conds.AppID.Op {
		case cruder.EQ:
			q.Where(orderuserent.AppID(id))
		default:
			return nil, fmt.Errorf("invalid appid field")
		}
	}
	if conds.UserID != nil {
		id, ok := conds.UserID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid userid")
		}
		switch conds.UserID.Op {
		case cruder.EQ:
			q.Where(orderuserent.UserID(id))
		default:
			return nil, fmt.Errorf("invalid userid field")
		}
	}
	if conds.RevenueAddress != nil {
		revenueaddress, ok := conds.RevenueAddress.Val.(string)
		if !ok {
			return nil, fmt.Errorf("invalid revenueaddress")
		}
		switch conds.RevenueAddress.Op {
		case cruder.EQ:
			q.Where(orderuserent.RevenueAddress(revenueaddress))
		default:
			return nil, fmt.Errorf("invalid revenueaddress field")
		}
	}
	if conds.AutoPay != nil {
		autopay, ok := conds.AutoPay.Val.(bool)
		if !ok {
			return nil, fmt.Errorf("invalid autopay")
		}
		switch conds.AutoPay.Op {
		case cruder.EQ:
			q.Where(orderuserent.AutoPay(autopay))
		default:
			return nil, fmt.Errorf("invalid autopay field")
		}
	}
	return q, nil
}
