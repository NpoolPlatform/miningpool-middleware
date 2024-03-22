package orderuser

import (
	"context"
	"fmt"

	basetypes "github.com/NpoolPlatform/message/npool/basetypes/miningpool/v1"
	npool "github.com/NpoolPlatform/message/npool/miningpool/mw/v1/orderuser"
	constant "github.com/NpoolPlatform/miningpool-middleware/pkg/const"
	orderusercrud "github.com/NpoolPlatform/miningpool-middleware/pkg/crud/orderuser"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"

	"github.com/google/uuid"
)

type Handler struct {
	ID             *uint32
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
	Reqs           []*orderusercrud.Req
	Conds          *orderusercrud.Conds
	Offset         int32
	Limit          int32
}

func NewHandler(ctx context.Context, options ...func(context.Context, *Handler) error) (*Handler, error) {
	handler := &Handler{}
	for _, opt := range options {
		if err := opt(ctx, handler); err != nil {
			return nil, err
		}
	}
	return handler, nil
}

func WithID(u *uint32, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if u == nil {
			if must {
				return fmt.Errorf("invalid id")
			}
			return nil
		}
		h.ID = u
		return nil
	}
}

func WithEntID(id *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if id == nil {
			if must {
				return fmt.Errorf("invalid entid")
			}
			return nil
		}
		_id, err := uuid.Parse(*id)
		if err != nil {
			return err
		}
		h.EntID = &_id
		return nil
	}
}

func WithRootUserID(rootuserid *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if rootuserid == nil {
			if must {
				return fmt.Errorf("invalid rootuserid")
			}
			return nil
		}
		_id, err := uuid.Parse(*rootuserid)
		if err != nil {
			return err
		}
		h.RootUserID = &_id
		return nil
	}
}

func WithGoodUserID(gooduserid *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if gooduserid == nil {
			if must {
				return fmt.Errorf("invalid gooduserid")
			}
			return nil
		}
		_id, err := uuid.Parse(*gooduserid)
		if err != nil {
			return err
		}
		h.GoodUserID = &_id
		return nil
	}
}

func WithOrderID(orderid *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if orderid == nil {
			if must {
				return fmt.Errorf("invalid orderid")
			}
			return nil
		}
		_id, err := uuid.Parse(*orderid)
		if err != nil {
			return err
		}
		h.OrderID = &_id
		return nil
	}
}

func WithName(name *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if name == nil {
			if must {
				return fmt.Errorf("invalid name")
			}
			return nil
		}
		h.Name = name
		return nil
	}
}

func WithMiningpoolType(miningpooltype *basetypes.MiningpoolType, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if miningpooltype == nil {
			if must {
				return fmt.Errorf("invalid miningpooltype")
			}
			return nil
		}
		if miningpooltype == basetypes.MiningpoolType_DefaultMiningpoolType.Enum() {
			return fmt.Errorf("invalid miningpooltype,not allow be default type")
		}
		h.MiningpoolType = miningpooltype
		return nil
	}
}

func WithCoinType(cointype *basetypes.CoinType, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if cointype == nil {
			if must {
				return fmt.Errorf("invalid cointype")
			}
			return nil
		}
		if cointype == basetypes.CoinType_DefaultCoinType.Enum() {
			return fmt.Errorf("invalid cointype,not allow be default type")
		}
		h.CoinType = cointype
		return nil
	}
}

func WithProportion(proportion *float32, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if proportion == nil {
			if must {
				return fmt.Errorf("invalid proportion")
			}
			return nil
		}
		h.Proportion = proportion
		return nil
	}
}

func WithRevenueAddress(revenueaddress *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if revenueaddress == nil {
			if must {
				return fmt.Errorf("invalid revenueaddress")
			}
			return nil
		}
		h.RevenueAddress = revenueaddress
		return nil
	}
}

func WithReadPageLink(readpagelink *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if readpagelink == nil {
			if must {
				return fmt.Errorf("invalid readpagelink")
			}
			return nil
		}
		h.ReadPageLink = readpagelink
		return nil
	}
}

func WithAutoPay(autopay *bool, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if autopay == nil {
			if must {
				return fmt.Errorf("invalid autopay")
			}
			return nil
		}
		h.AutoPay = autopay
		return nil
	}
}

// nolint:gocognit
func WithReqs(reqs []*npool.OrderUserReq, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		_reqs := []*orderusercrud.Req{}
		for _, req := range reqs {
			_req := &orderusercrud.Req{}
			if req.EntID != nil {
				id, err := uuid.Parse(req.GetEntID())
				if err != nil {
					return err
				}
				_req.EntID = &id
			}
			if req.RootUserID != nil {
				id, err := uuid.Parse(req.GetEntID())
				if err != nil {
					return err
				}
				_req.RootUserID = &id
			}
			if req.GoodUserID != nil {
				id, err := uuid.Parse(req.GetEntID())
				if err != nil {
					return err
				}
				_req.GoodUserID = &id
			}
			if req.OrderID != nil {
				id, err := uuid.Parse(req.GetEntID())
				if err != nil {
					return err
				}
				_req.OrderID = &id
			}
			if req.Name != nil {
				_req.Name = req.Name
			}
			if req.MiningpoolType != nil {
				if req.MiningpoolType == basetypes.MiningpoolType_DefaultMiningpoolType.Enum() {
					return fmt.Errorf("invalid miningpooltype,not allow be default type")
				}
				_req.MiningpoolType = req.MiningpoolType
			}
			if req.CoinType != nil {
				if req.CoinType == basetypes.CoinType_DefaultCoinType.Enum() {
					return fmt.Errorf("invalid cointype,not allow be default type")
				}
				_req.CoinType = req.CoinType
			}
			if req.Proportion != nil {
				_req.Proportion = req.Proportion
			}
			if req.RevenueAddress != nil {
				_req.RevenueAddress = req.RevenueAddress
			}
			if req.ReadPageLink != nil {
				_req.ReadPageLink = req.ReadPageLink
			}
			if req.AutoPay != nil {
				_req.AutoPay = req.AutoPay
			}
			_reqs = append(_reqs, _req)
		}
		h.Reqs = _reqs
		return nil
	}
}

//nolint:gocognit
func WithConds(conds *npool.Conds) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.Conds = &orderusercrud.Conds{}
		if conds == nil {
			return nil
		}
		if conds.ID != nil {
			h.Conds.ID = &cruder.Cond{
				Op:  conds.GetID().GetOp(),
				Val: conds.GetID().GetValue(),
			}
		}
		if conds.EntID != nil {
			id, err := uuid.Parse(conds.GetEntID().GetValue())
			if err != nil {
				return err
			}
			h.Conds.EntID = &cruder.Cond{
				Op:  conds.GetEntID().GetOp(),
				Val: id,
			}
		}
		if conds.EntIDs != nil {
			ids := []uuid.UUID{}
			for _, id := range conds.GetEntIDs().GetValue() {
				_id, err := uuid.Parse(id)
				if err != nil {
					return err
				}
				ids = append(ids, _id)
			}
			h.Conds.EntIDs = &cruder.Cond{
				Op:  conds.GetEntIDs().GetOp(),
				Val: ids,
			}
		}
		if conds.Name != nil {
			h.Conds.Name = &cruder.Cond{
				Op:  conds.GetName().GetOp(),
				Val: conds.GetName().GetValue(),
			}
		}
		if conds.RootUserID != nil {
			id, err := uuid.Parse(conds.GetRootUserID().GetValue())
			if err != nil {
				return err
			}
			h.Conds.RootUserID = &cruder.Cond{
				Op:  conds.GetRootUserID().GetOp(),
				Val: id,
			}
		}
		if conds.GoodUserID != nil {
			id, err := uuid.Parse(conds.GetGoodUserID().GetValue())
			if err != nil {
				return err
			}
			h.Conds.GoodUserID = &cruder.Cond{
				Op:  conds.GetGoodUserID().GetOp(),
				Val: id,
			}
		}
		if conds.OrderID != nil {
			id, err := uuid.Parse(conds.GetOrderID().GetValue())
			if err != nil {
				return err
			}
			h.Conds.OrderID = &cruder.Cond{
				Op:  conds.GetOrderID().GetOp(),
				Val: id,
			}
		}
		if conds.MiningpoolType != nil {
			h.Conds.MiningpoolType = &cruder.Cond{
				Op:  conds.GetMiningpoolType().GetOp(),
				Val: basetypes.MiningpoolType(conds.GetMiningpoolType().GetValue()),
			}
		}
		if conds.CoinType != nil {
			h.Conds.CoinType = &cruder.Cond{
				Op:  conds.GetCoinType().GetOp(),
				Val: basetypes.CoinType(conds.GetCoinType().GetValue()),
			}
		}
		if conds.RevenueAddress != nil {
			h.Conds.RevenueAddress = &cruder.Cond{
				Op:  conds.GetRevenueAddress().GetOp(),
				Val: conds.GetRevenueAddress().GetValue(),
			}
		}
		if conds.AutoPay != nil {
			h.Conds.AutoPay = &cruder.Cond{
				Op:  conds.GetAutoPay().GetOp(),
				Val: conds.GetAutoPay().GetValue(),
			}
		}
		return nil
	}
}

func WithOffset(offset int32) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.Offset = offset
		return nil
	}
}

func WithLimit(limit int32) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if limit == 0 {
			limit = constant.DefaultRowLimit
		}
		h.Limit = limit
		return nil
	}
}
