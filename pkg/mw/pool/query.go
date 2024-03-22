package pool

import (
	"context"
	"fmt"

	basetypes "github.com/NpoolPlatform/message/npool/basetypes/miningpool/v1"
	npool "github.com/NpoolPlatform/message/npool/miningpool/mw/v1/pool"

	"github.com/NpoolPlatform/miningpool-middleware/pkg/db"
	"github.com/NpoolPlatform/miningpool-middleware/pkg/db/ent"
	poolent "github.com/NpoolPlatform/miningpool-middleware/pkg/db/ent/pool"

	poolcrud "github.com/NpoolPlatform/miningpool-middleware/pkg/crud/pool"
)

type queryHandler struct {
	*Handler
	stm   *ent.PoolSelect
	infos []*npool.Pool
	total uint32
}

func (h *queryHandler) selectPool(stm *ent.PoolQuery) {
	h.stm = stm.Select(
		poolent.FieldID,
		poolent.FieldEntID,
		poolent.FieldMiningpoolType,
		poolent.FieldName,
		poolent.FieldSite,
		poolent.FieldDescription,
		poolent.FieldCreatedAt,
		poolent.FieldUpdatedAt,
	)
}

func (h *queryHandler) queryPool(cli *ent.Client) error {
	if h.ID == nil && h.EntID == nil {
		return fmt.Errorf("invalid id")
	}
	stm := cli.Pool.Query().Where(poolent.DeletedAt(0))
	if h.ID != nil {
		stm.Where(poolent.ID(*h.ID))
	}
	if h.EntID != nil {
		stm.Where(poolent.EntID(*h.EntID))
	}
	h.selectPool(stm)
	return nil
}

func (h *queryHandler) queryPools(ctx context.Context, cli *ent.Client) error {
	stm, err := poolcrud.SetQueryConds(cli.Pool.Query(), h.Conds)
	if err != nil {
		return err
	}

	total, err := stm.Count(ctx)
	if err != nil {
		return err
	}
	h.total = uint32(total)

	h.selectPool(stm)
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

func (h *Handler) GetPool(ctx context.Context) (*npool.Pool, error) {
	handler := &queryHandler{
		Handler: h,
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.queryPool(cli); err != nil {
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

func (h *Handler) GetPools(ctx context.Context) ([]*npool.Pool, uint32, error) {
	handler := &queryHandler{
		Handler: h,
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.queryPools(ctx, cli); err != nil {
			return err
		}
		handler.stm.
			Offset(int(h.Offset)).
			Limit(int(h.Limit)).
			Order(ent.Desc(poolent.FieldUpdatedAt))
		return handler.scan(_ctx)
	})
	if err != nil {
		return nil, 0, err
	}

	handler.formalize()
	return handler.infos, handler.total, nil
}
