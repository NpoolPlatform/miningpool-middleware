package gooduser

import (
	"context"
	"fmt"

	"github.com/NpoolPlatform/miningpool-middleware/pkg/db"
	"github.com/NpoolPlatform/miningpool-middleware/pkg/db/ent"
	"github.com/NpoolPlatform/miningpool-middleware/pkg/mw/rootuser"
	"github.com/NpoolPlatform/miningpool-middleware/pkg/pools"

	"github.com/google/uuid"
)

func (h *Handler) CreateGoodUser(ctx context.Context) error {
	rootuserID := h.RootUserID.String()
	rootuserH, err := rootuser.NewHandler(ctx, rootuser.WithEntID(&rootuserID, true))
	if err != nil {
		return err
	}

	info, err := rootuserH.GetRootUser(ctx)
	if info == nil || err != nil {
		return fmt.Errorf("invalid rootuserid")
	}

	if h.EntID == nil {
		h.EntID = func() *uuid.UUID { uid := uuid.New(); return &uid }()
	}

	err = h.newGoodUserInPool(ctx)
	if err != nil {
		return err
	}

	return db.WithTx(ctx, func(ctx context.Context, tx *ent.Tx) error {
		sql, err := h.genCreateSQL()
		if err != nil {
			return err
		}
		rc, err := tx.ExecContext(ctx, sql)
		if err != nil {
			return err
		}
		if n, err := rc.RowsAffected(); err != nil || n != 1 {
			return fmt.Errorf("fail create gooduser: %v", err)
		}
		return nil
	})
}

func (h *Handler) newGoodUserInPool(ctx context.Context) error {
	rootuserID := h.RootUserID.String()
	rootuserH, err := rootuser.NewHandler(ctx, rootuser.WithEntID(&rootuserID, true))
	if err != nil {
		return err
	}
	rootUser, err := rootuserH.GetAuthToken(ctx)
	if err != nil {
		return err
	}
	if rootUser == nil {
		return fmt.Errorf("have no rootuser,entid: %v", rootuserID)
	}

	if h.MiningpoolType.String() != rootUser.MiningpoolType.String() {
		return fmt.Errorf("the miningpool type is different from type of rootuser")
	}

	mgr, err := pools.NewPoolManager(*h.MiningpoolType, *h.CoinType, rootUser.AuthTokenPlain)
	if err != nil {
		return err
	}
	name, pageLink, err := mgr.AddMiningUser(ctx)
	if err != nil {
		return err
	}
	h.Name = &name
	h.ReadPageLink = &pageLink
	return nil
}
