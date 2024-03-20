package coin

import (
	"context"
	"fmt"

	basetypes "github.com/NpoolPlatform/message/npool/basetypes/miningpool/v1"
	npool "github.com/NpoolPlatform/message/npool/miningpool/mw/v1/coin"
	constant "github.com/NpoolPlatform/miningpool-middleware/pkg/const"
	coincrud "github.com/NpoolPlatform/miningpool-middleware/pkg/crud/coin"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"

	"github.com/google/uuid"
)

type Handler struct {
	ID               *uint32
	EntID            *uuid.UUID
	CoinType         *basetypes.CoinType
	MiningpoolType   *basetypes.MiningpoolType
	RevenueTypes     *[]basetypes.RevenueType
	Site             *string
	FeeRate          *float32
	FixedRevenueAble *bool
	Threshold        *float32
	Remark           *string
	Reqs             []*coincrud.Req
	Conds            *coincrud.Conds
	Offset           int32
	Limit            int32
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

func WithRevenueTypes(revenuetypes *[]basetypes.RevenueType, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if revenuetypes == nil {
			if must {
				return fmt.Errorf("invalid revenuetypes")
			}
			return nil
		}
		h.RevenueTypes = revenuetypes
		return nil
	}
}

func WithSite(site *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if site == nil {
			if must {
				return fmt.Errorf("invalid site")
			}
			return nil
		}
		h.Site = site
		return nil
	}
}

func WithFeeRate(feerate *float32, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if feerate == nil {
			if must {
				return fmt.Errorf("invalid feerate")
			}
			return nil
		}
		h.FeeRate = feerate
		return nil
	}
}

func WithFixedRevenueAble(fixedrevenueable *bool, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if fixedrevenueable == nil {
			if must {
				return fmt.Errorf("invalid fixedrevenueable")
			}
			return nil
		}
		h.FixedRevenueAble = fixedrevenueable
		return nil
	}
}

func WithThreshold(threshold *float32, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if threshold == nil {
			if must {
				return fmt.Errorf("invalid threshold")
			}
			return nil
		}
		h.Threshold = threshold
		return nil
	}
}

func WithRemark(remark *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if remark == nil {
			if must {
				return fmt.Errorf("invalid remark")
			}
			return nil
		}
		h.Remark = remark
		return nil
	}
}

// nolint:gocognit
func WithReqs(reqs []*npool.CoinReq, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		_reqs := []*coincrud.Req{}
		for _, req := range reqs {
			_req := &coincrud.Req{}
			if req.EntID != nil {
				id, err := uuid.Parse(req.GetEntID())
				if err != nil {
					return err
				}
				_req.EntID = &id
			}
			if req.CoinType != nil {
				if req.CoinType == basetypes.CoinType_DefaultCoinType.Enum() {
					return fmt.Errorf("invalid cointype,not allow be default type")
				}
				_req.CoinType = req.CoinType
			}
			if req.MiningpoolType != nil {
				if req.MiningpoolType == basetypes.MiningpoolType_DefaultMiningpoolType.Enum() {
					return fmt.Errorf("invalid miningpooltype,not allow be default type")
				}
				_req.MiningpoolType = req.MiningpoolType
			}
			if req.RevenueTypes != nil {
				_req.RevenueTypes = &req.RevenueTypes
			}
			if req.Site != nil {
				_req.Site = req.Site
			}
			if req.FeeRate != nil {
				_req.FeeRate = req.FeeRate
			}
			if req.FixedRevenueAble != nil {
				_req.FixedRevenueAble = req.FixedRevenueAble
			}
			if req.Threshold != nil {
				_req.Threshold = req.Threshold
			}
			if req.Remark != nil {
				_req.Remark = req.Remark
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
		h.Conds = &coincrud.Conds{}
		if conds == nil {
			return nil
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
		if conds.CoinType != nil {
			h.Conds.CoinType = &cruder.Cond{
				Op:  conds.GetCoinType().GetOp(),
				Val: basetypes.CoinType(conds.GetCoinType().GetValue()),
			}
		}
		if conds.MiningpoolType != nil {
			h.Conds.MiningpoolType = &cruder.Cond{
				Op:  conds.GetMiningpoolType().GetOp(),
				Val: basetypes.MiningpoolType(conds.GetMiningpoolType().GetValue()),
			}
		}
		if conds.Site != nil {
			h.Conds.Site = &cruder.Cond{
				Op:  conds.GetSite().GetOp(),
				Val: conds.GetSite().GetValue(),
			}
		}
		if conds.FixedRevenueAble != nil {
			h.Conds.FixedRevenueAble = &cruder.Cond{
				Op:  conds.GetFixedRevenueAble().GetOp(),
				Val: conds.GetFixedRevenueAble().GetValue(),
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
