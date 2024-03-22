package pool

import (
	"context"
	"fmt"

	basetypes "github.com/NpoolPlatform/message/npool/basetypes/miningpool/v1"
	npool "github.com/NpoolPlatform/message/npool/miningpool/mw/v1/pool"
	constant "github.com/NpoolPlatform/miningpool-middleware/pkg/const"
	poolcrud "github.com/NpoolPlatform/miningpool-middleware/pkg/crud/pool"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"

	"github.com/google/uuid"
)

type Handler struct {
	ID             *uint32
	EntID          *uuid.UUID
	MiningpoolType *basetypes.MiningpoolType
	Name           *string
	Site           *string
	Description    *string
	Reqs           []*poolcrud.Req
	Conds          *poolcrud.Conds
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

func WithDescription(description *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if description == nil {
			if must {
				return fmt.Errorf("invalid description")
			}
			return nil
		}
		h.Description = description
		return nil
	}
}

// nolint:gocognit
func WithReqs(reqs []*npool.PoolReq, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		_reqs := []*poolcrud.Req{}
		for _, req := range reqs {
			_req := &poolcrud.Req{}
			if req.EntID != nil {
				id, err := uuid.Parse(req.GetEntID())
				if err != nil {
					return err
				}
				_req.EntID = &id
			}
			if req.MiningpoolType != nil {
				if req.MiningpoolType == basetypes.MiningpoolType_DefaultMiningpoolType.Enum() {
					return fmt.Errorf("invalid miningpooltype,not allow be default type")
				}
				_req.MiningpoolType = req.MiningpoolType
			}
			if req.Name != nil {
				_req.Name = req.Name
			}
			if req.Site != nil {
				_req.Site = req.Site
			}
			if req.Description != nil {
				_req.Description = req.Description
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
		h.Conds = &poolcrud.Conds{}
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
		if conds.MiningpoolType != nil {
			h.Conds.MiningpoolType = &cruder.Cond{
				Op:  conds.GetMiningpoolType().GetOp(),
				Val: basetypes.MiningpoolType(conds.GetMiningpoolType().GetValue()),
			}
		}
		if conds.Name != nil {
			h.Conds.Name = &cruder.Cond{
				Op:  conds.GetName().GetOp(),
				Val: conds.GetName().GetValue(),
			}
		}
		if conds.Site != nil {
			h.Conds.Site = &cruder.Cond{
				Op:  conds.GetSite().GetOp(),
				Val: conds.GetSite().GetValue(),
			}
		}
		if conds.Description != nil {
			h.Conds.Description = &cruder.Cond{
				Op:  conds.GetDescription().GetOp(),
				Val: conds.GetDescription().GetValue(),
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
