package tx

import (
	"context"
	"fmt"

	basetypes "github.com/NpoolPlatform/message/npool/basetypes/miningpool/v1"
	npool "github.com/NpoolPlatform/message/npool/miningpool/mw/v1/rootuser"
	rootusercrud "github.com/NpoolPlatform/miningpool-middleware/pkg/crud/rootuser"

	"github.com/NpoolPlatform/miningpool-middleware/pkg/db"
	"github.com/NpoolPlatform/miningpool-middleware/pkg/db/ent"
	rootuserent "github.com/NpoolPlatform/miningpool-middleware/pkg/db/ent/rootuser"
)

type updateHandler struct {
	*Handler
}

//nolint:gocyclo
func (h *updateHandler) validateState(info *ent.RootUser) error {
	if h.MiningpoolType == nil {
		return nil
	}
	if h.MiningpoolType == basetypes.MiningpoolType_DefaultMiningpoolType.Enum() {
		return fmt.Errorf("invalid miningpooltype")
	}
	return nil
}

func (h *Handler) UpdateRootUser(ctx context.Context) (*npool.RootUser, error) {
	if h.ID == nil {
		return nil, fmt.Errorf("invalid id")
	}

	handler := &updateHandler{
		Handler: h,
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		info, err := cli.
			RootUser.
			Query().
			Where(
				rootuserent.ID(*h.ID),
			).
			Only(_ctx)
		if err != nil {
			return err
		}

		if err := handler.validateState(info); err != nil {
			return err
		}

		stm, err := rootusercrud.UpdateSet(
			info.Update(),
			&rootusercrud.Req{
				Name:           h.Name,
				MiningpoolType: h.MiningpoolType,
				Email:          h.Email,
				AuthToken:      h.AuthToken,
				Authed:         h.Authed,
				Remark:         h.Remark,
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
	if err != nil {
		return nil, err
	}

	return h.GetRootUser(ctx)
}
