package orderuser

import (
	"context"

	npool "github.com/NpoolPlatform/message/npool/miningpool/mw/v1/orderuser"
	orderusercrud "github.com/NpoolPlatform/miningpool-middleware/pkg/crud/orderuser"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	"github.com/NpoolPlatform/miningpool-middleware/pkg/db"
	"github.com/NpoolPlatform/miningpool-middleware/pkg/db/ent"

	"github.com/google/uuid"
)

func (h *Handler) CreateOrderUser(ctx context.Context) (*npool.OrderUser, error) {
	id := uuid.New()
	if h.EntID == nil {
		h.EntID = &id
	}

	err := db.WithClient(ctx, func(ctx context.Context, cli *ent.Client) error {
		info, err := orderusercrud.CreateSet(
			cli.OrderUser.Create(),
			&orderusercrud.Req{
				EntID:          h.EntID,
				Name:           h.Name,
				RootUserID:     h.RootUserID,
				GoodUserID:     h.GoodUserID,
				OrderID:        h.OrderID,
				MiningpoolType: h.MiningpoolType,
				CoinType:       h.CoinType,
				Proportion:     h.Proportion,
				RevenueAddress: h.RevenueAddress,
				ReadPageLink:   h.ReadPageLink,
				AutoPay:        h.AutoPay,
			},
		).Save(ctx)
		if err != nil {
			return err
		}
		h.ID = &info.ID
		return nil
	})
	if err != nil {
		return nil, err
	}

	return h.GetOrderUser(ctx)
}

func (h *Handler) CreateOrderUsers(ctx context.Context) ([]*npool.OrderUser, error) {
	ids := []uuid.UUID{}

	err := db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		for _, req := range h.Reqs {
			info, err := orderusercrud.CreateSet(tx.OrderUser.Create(), req).Save(_ctx)
			if err != nil {
				return err
			}
			ids = append(ids, info.EntID)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	h.Conds = &orderusercrud.Conds{
		EntIDs: &cruder.Cond{Op: cruder.IN, Val: ids},
	}
	h.Offset = 0
	h.Limit = int32(len(ids))

	infos, _, err := h.GetOrderUsers(ctx)
	return infos, err
}
