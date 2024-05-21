package fraction

import (
	"context"
	"fmt"

	basetypes "github.com/NpoolPlatform/message/npool/basetypes/miningpool/v1"
	fractioncrud "github.com/NpoolPlatform/miningpool-middleware/pkg/crud/fraction"

	"github.com/NpoolPlatform/miningpool-middleware/pkg/db"
	"github.com/NpoolPlatform/miningpool-middleware/pkg/db/ent"
	fractionent "github.com/NpoolPlatform/miningpool-middleware/pkg/db/ent/fraction"
)

type updateHandler struct {
	*Handler
}

//nolint:gocyclo
func (h *updateHandler) validateState(info *ent.Fraction) error {
	if info.WithdrawState == basetypes.WithdrawState_DefaultWithdrawState.String() {
		return fmt.Errorf("invalid withdrawstate")
	}
	return nil
}

func (h *Handler) UpdateFraction(ctx context.Context) error {
	info, err := h.GetFraction(ctx)
	if err != nil {
		return fmt.Errorf("invalid id or ent_id")
	}
	if info == nil {
		return fmt.Errorf("invalid id or ent_id")
	}

	h.ID = &info.ID

	handler := &updateHandler{
		Handler: h,
	}

	return db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		info, err := cli.
			Fraction.
			Query().
			Where(
				fractionent.ID(*h.ID),
			).
			Only(_ctx)
		if err != nil {
			return err
		}

		if err := handler.validateState(info); err != nil {
			return err
		}

		stm, err := fractioncrud.UpdateSet(
			info.Update(),
			&fractioncrud.Req{
				AppID:         h.AppID,
				UserID:        h.UserID,
				OrderUserID:   h.OrderUserID,
				WithdrawState: h.WithdrawState,
				WithdrawAt:    h.WithdrawAt,
				PromisePayAt:  h.PromisePayAt,
				Msg:           h.Msg,
			},
		)
		if err != nil {
			return err
		}
		if _, err := stm.Save(_ctx); err != nil {
			return err
		}
		return nil
	})
}
