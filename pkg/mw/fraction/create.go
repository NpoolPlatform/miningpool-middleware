package fraction

import (
	"context"

	npool "github.com/NpoolPlatform/message/npool/miningpool/mw/v1/fraction"
	fractioncrud "github.com/NpoolPlatform/miningpool-middleware/pkg/crud/fraction"

	"github.com/NpoolPlatform/miningpool-middleware/pkg/db"
	"github.com/NpoolPlatform/miningpool-middleware/pkg/db/ent"

	"github.com/google/uuid"
)

func (h *Handler) CreateFraction(ctx context.Context) (*npool.Fraction, error) {
	id := uuid.New()
	if h.EntID == nil {
		h.EntID = &id
	}

	err := db.WithClient(ctx, func(ctx context.Context, cli *ent.Client) error {
		info, err := fractioncrud.CreateSet(
			cli.Fraction.Create(),
			&fractioncrud.Req{
				EntID:         h.EntID,
				AppID:         h.AppID,
				UserID:        h.UserID,
				OrderUserID:   h.OrderUserID,
				WithdrawState: h.WithdrawState,
				WithdrawTime:  h.WithdrawTime,
				PayTime:       h.PayTime,
				Msg:           h.Msg,
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

	return h.GetFraction(ctx)
}
