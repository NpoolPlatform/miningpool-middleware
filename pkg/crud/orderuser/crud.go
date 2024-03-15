package orderuser

import (
	"fmt"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/miningpool/v1"
	"github.com/NpoolPlatform/miningpool-middleware/pkg/db/ent"
	orderuserent "github.com/NpoolPlatform/miningpool-middleware/pkg/db/ent/orderuser"

	"github.com/google/uuid"
)

type Req struct {
	EntID          *uuid.UUID
	RootUserID     *uuid.UUID
	GoodUserID     *uuid.UUID
	OrderID        *uuid.UUID
	Name           *string
	MiningpoolType *basetypes.MiningpoolType
	CoinType       *basetypes.CoinType
	Proportion     *float32
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
	if req.RootUserID != nil {
		c.SetRootUserID(*req.RootUserID)
	}
	if req.GoodUserID != nil {
		c.SetGoodUserID(*req.GoodUserID)
	}
	if req.OrderID != nil {
		c.SetOrderID(*req.OrderID)
	}
	if req.MiningpoolType != nil {
		c.SetMiningpoolType(req.MiningpoolType.String())
	}
	if req.CoinType != nil {
		c.SetCoinType(req.CoinType.String())
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
	if req.RootUserID != nil {
		u = u.SetRootUserID(*req.RootUserID)
	}
	if req.GoodUserID != nil {
		u = u.SetGoodUserID(*req.GoodUserID)
	}
	if req.OrderID != nil {
		u = u.SetOrderID(*req.OrderID)
	}
	if req.MiningpoolType != nil {
		u = u.SetMiningpoolType(req.MiningpoolType.String())
	}
	if req.CoinType != nil {
		u = u.SetCoinType(req.CoinType.String())
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
	EntID          *cruder.Cond
	Name           *cruder.Cond
	RootUserID     *cruder.Cond
	GoodUserID     *cruder.Cond
	OrderID        *cruder.Cond
	MiningpoolType *cruder.Cond
	CoinType       *cruder.Cond
	RevenueAddress *cruder.Cond
	AutoPay        *cruder.Cond
	EntIDs         *cruder.Cond
}

func SetQueryConds(q *ent.OrderUserQuery, conds *Conds) (*ent.OrderUserQuery, error) { //nolint
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
	if conds.RootUserID != nil {
		id, ok := conds.RootUserID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid rootuserid")
		}
		switch conds.RootUserID.Op {
		case cruder.EQ:
			q.Where(orderuserent.RootUserID(id))
		default:
			return nil, fmt.Errorf("invalid rootuserid field")
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
	if conds.OrderID != nil {
		id, ok := conds.OrderID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid orderid")
		}
		switch conds.OrderID.Op {
		case cruder.EQ:
			q.Where(orderuserent.OrderID(id))
		default:
			return nil, fmt.Errorf("invalid orderid field")
		}
	}
	if conds.MiningpoolType != nil {
		miningpooltype, ok := conds.MiningpoolType.Val.(basetypes.MiningpoolType)
		if !ok {
			return nil, fmt.Errorf("invalid miningpooltype")
		}
		switch conds.MiningpoolType.Op {
		case cruder.EQ:
			q.Where(orderuserent.MiningpoolType(miningpooltype.String()))
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
			q.Where(orderuserent.CoinType(cointype.String()))
		default:
			return nil, fmt.Errorf("invalid cointype field")
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
