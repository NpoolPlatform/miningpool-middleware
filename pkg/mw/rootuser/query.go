package rootuser

import (
	"context"
	"fmt"

	basetypes "github.com/NpoolPlatform/message/npool/basetypes/miningpool/v1"
	npool "github.com/NpoolPlatform/message/npool/miningpool/mw/v1/rootuser"

	"github.com/NpoolPlatform/miningpool-middleware/pkg/db"
	"github.com/NpoolPlatform/miningpool-middleware/pkg/db/ent"
	rootuserent "github.com/NpoolPlatform/miningpool-middleware/pkg/db/ent/rootuser"

	rootusercrud "github.com/NpoolPlatform/miningpool-middleware/pkg/crud/rootuser"
)

type queryHandler struct {
	*Handler
	stm   *ent.RootUserSelect
	infos []*npool.RootUser
	total uint32
}

func (h *queryHandler) selectRootUser(stm *ent.RootUserQuery) {
	h.stm = stm.Select(
		rootuserent.FieldID,
		rootuserent.FieldEntID,
		rootuserent.FieldName,
		rootuserent.FieldMiningpoolType,
		rootuserent.FieldEmail,
		rootuserent.FieldAuthToken,
		rootuserent.FieldAuthed,
		rootuserent.FieldRemark,
		rootuserent.FieldCreatedAt,
		rootuserent.FieldUpdatedAt,
	)
}

func (h *queryHandler) queryRootUser(cli *ent.Client) error {
	if h.ID == nil && h.EntID == nil {
		return fmt.Errorf("invalid id")
	}
	stm := cli.RootUser.Query().Where(rootuserent.DeletedAt(0))
	if h.ID != nil {
		stm.Where(rootuserent.ID(*h.ID))
	}
	if h.EntID != nil {
		stm.Where(rootuserent.EntID(*h.EntID))
	}
	h.selectRootUser(stm)
	return nil
}

func (h *queryHandler) queryRootUsers(ctx context.Context, cli *ent.Client) error {
	stm, err := rootusercrud.SetQueryConds(cli.RootUser.Query(), h.Conds)
	if err != nil {
		return err
	}

	total, err := stm.Count(ctx)
	if err != nil {
		return err
	}
	h.total = uint32(total)

	h.selectRootUser(stm)
	return nil
}

func (h *queryHandler) scan(ctx context.Context) error {
	return h.stm.Scan(ctx, &h.infos)
}

func (h *queryHandler) formalize() {
	for _, info := range h.infos {
		info.MiningpoolType = basetypes.MiningpoolType(basetypes.MiningpoolType_value[info.MiningpoolTypeStr])
	}
}

func (h *Handler) GetRootUser(ctx context.Context) (*npool.RootUser, error) {
	handler := &queryHandler{
		Handler: h,
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.queryRootUser(cli); err != nil {
			return err
		}
		const singleRowLimit = 2
		handler.stm.Offset(0).Limit(singleRowLimit)
		return handler.scan(_ctx)
	})
	if err != nil {
		return nil, err
	}
	if len(handler.infos) == 0 {
		return nil, nil
	}
	if len(handler.infos) > 1 {
		return nil, fmt.Errorf("too many record")
	}

	handler.formalize()
	return handler.infos[0], nil
}

func (h *Handler) GetRootUsers(ctx context.Context) ([]*npool.RootUser, uint32, error) {
	handler := &queryHandler{
		Handler: h,
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.queryRootUsers(ctx, cli); err != nil {
			return err
		}
		handler.stm.
			Offset(int(h.Offset)).
			Limit(int(h.Limit)).
			Order(ent.Desc(rootuserent.FieldUpdatedAt))
		return handler.scan(_ctx)
	})
	if err != nil {
		return nil, 0, err
	}

	handler.formalize()
	return handler.infos, handler.total, nil
}