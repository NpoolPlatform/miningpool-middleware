package gooduser

import (
	"context"
	"fmt"

	npool "github.com/NpoolPlatform/message/npool/miningpool/mw/v1/gooduser"
	constant "github.com/NpoolPlatform/miningpool-middleware/pkg/const"
	goodusercrud "github.com/NpoolPlatform/miningpool-middleware/pkg/crud/gooduser"
	"github.com/NpoolPlatform/miningpool-middleware/pkg/mw/coin"
	"github.com/NpoolPlatform/miningpool-middleware/pkg/mw/rootuser"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"

	"github.com/google/uuid"
)

type Handler struct {
	goodusercrud.Req
	Reqs   []*goodusercrud.Req
	Conds  *goodusercrud.Conds
	Offset int32
	Limit  int32
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

func WithPoolCoinTypeID(poolcointypeid *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if poolcointypeid == nil {
			if must {
				return fmt.Errorf("invalid poolcointypeid")
			}
			return nil
		}

		coinH, err := coin.NewHandler(ctx, coin.WithEntID(poolcointypeid, true))
		if err != nil {
			return err
		}

		exist, err := coinH.ExistCoin(ctx)
		if err != nil {
			return err
		}
		if !exist {
			return fmt.Errorf("invalid poolcointypeid")
		}
		h.PoolCoinTypeID = coinH.EntID
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
		rootuserH, err := rootuser.NewHandler(ctx, rootuser.WithEntID(rootuserid, true))
		if err != nil {
			return err
		}
		exist, err := rootuserH.ExistRootUser(ctx)
		if err != nil {
			return err
		}
		if !exist {
			return fmt.Errorf("invalid rootuserid")
		}
		h.RootUserID = rootuserH.EntID
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

//nolint:gocognit
func WithConds(conds *npool.Conds) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.Conds = &goodusercrud.Conds{}
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
		if conds.PoolCoinTypeID != nil {
			id, err := uuid.Parse(conds.GetPoolCoinTypeID().GetValue())
			if err != nil {
				return err
			}
			h.Conds.PoolCoinTypeID = &cruder.Cond{
				Op:  conds.GetPoolCoinTypeID().GetOp(),
				Val: id,
			}
		}
		if conds.Name != nil {
			h.Conds.Name = &cruder.Cond{
				Op:  conds.GetName().GetOp(),
				Val: conds.GetName().GetValue(),
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
