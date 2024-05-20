package fractionrule

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	mpbasetypes "github.com/NpoolPlatform/message/npool/basetypes/miningpool/v1"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	npool "github.com/NpoolPlatform/message/npool/miningpool/mw/v1/fractionrule"

	"github.com/NpoolPlatform/miningpool-middleware/pkg/db"
	"github.com/NpoolPlatform/miningpool-middleware/pkg/db/ent"
	"github.com/NpoolPlatform/miningpool-middleware/pkg/db/ent/coin"
	fractionruleent "github.com/NpoolPlatform/miningpool-middleware/pkg/db/ent/fractionrule"
	"github.com/NpoolPlatform/miningpool-middleware/pkg/db/ent/pool"

	fractionrulecrud "github.com/NpoolPlatform/miningpool-middleware/pkg/crud/fractionrule"
)

type queryHandler struct {
	*Handler
	stm   *ent.FractionRuleSelect
	infos []*npool.FractionRule
	total uint32
}

func (h *queryHandler) selectFractionRule(stm *ent.FractionRuleQuery) {
	h.stm = stm.Select(
		fractionruleent.FieldID,
		fractionruleent.FieldCreatedAt,
		fractionruleent.FieldUpdatedAt,
		fractionruleent.FieldEntID,
		fractionruleent.FieldPoolCoinTypeID,
		fractionruleent.FieldWithdrawInterval,
		fractionruleent.FieldMinAmount,
		fractionruleent.FieldWithdrawRate,
	)
}

func (h *queryHandler) queryFractionRule(cli *ent.Client) error {
	if h.ID == nil && h.EntID == nil {
		return fmt.Errorf("invalid id")
	}
	stm := cli.FractionRule.Query().Where(fractionruleent.DeletedAt(0))
	if h.ID != nil {
		stm.Where(fractionruleent.ID(*h.ID))
	}
	if h.EntID != nil {
		stm.Where(fractionruleent.EntID(*h.EntID))
	}
	h.selectFractionRule(stm)
	return nil
}

func (h *queryHandler) queryFractionRules(ctx context.Context, cli *ent.Client) error {
	stm, err := fractionrulecrud.SetQueryConds(cli.FractionRule.Query(), h.Conds)
	if err != nil {
		return err
	}

	total, err := stm.Count(ctx)
	if err != nil {
		return err
	}
	h.total = uint32(total)

	h.selectFractionRule(stm)
	return nil
}

func (h *queryHandler) queryJoin() {
	h.stm.Modify(
		h.queryJoinCoinAndPool,
	)
}

func (h *queryHandler) queryJoinCoinAndPool(s *sql.Selector) {
	coinT := sql.Table(coin.Table)
	poolT := sql.Table(pool.Table)
	s.LeftJoin(coinT).On(
		s.C(fractionruleent.FieldPoolCoinTypeID),
		coinT.C(coin.FieldEntID),
	).OnP(
		sql.EQ(coinT.C(coin.FieldDeletedAt), 0),
	).LeftJoin(poolT).On(
		coinT.C(coin.FieldPoolID),
		poolT.C(pool.FieldEntID),
	).OnP(
		sql.EQ(poolT.C(pool.FieldDeletedAt), 0),
	).AppendSelect(
		coinT.C(coin.FieldCoinType),
		coinT.C(coin.FieldPoolID),
		poolT.C(pool.FieldMiningpoolType),
	)
}

func (h *queryHandler) scan(ctx context.Context) error {
	return h.stm.Scan(ctx, &h.infos)
}

func (h *queryHandler) formalize() {
	for _, info := range h.infos {
		info.MiningpoolType = mpbasetypes.MiningpoolType(mpbasetypes.MiningpoolType_value[info.MiningpoolTypeStr])
		info.CoinType = basetypes.CoinType(basetypes.CoinType_value[info.CoinTypeStr])
	}
}

func (h *Handler) GetFractionRule(ctx context.Context) (*npool.FractionRule, error) {
	handler := &queryHandler{
		Handler: h,
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.queryFractionRule(cli); err != nil {
			return err
		}
		handler.queryJoin()
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

func (h *Handler) GetFractionRules(ctx context.Context) ([]*npool.FractionRule, uint32, error) {
	handler := &queryHandler{
		Handler: h,
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.queryFractionRules(ctx, cli); err != nil {
			return err
		}
		handler.queryJoin()
		handler.stm.
			Offset(int(h.Offset)).
			Limit(int(h.Limit)).
			Order(ent.Desc(fractionruleent.FieldUpdatedAt))
		return handler.scan(_ctx)
	})
	if err != nil {
		return nil, 0, err
	}

	handler.formalize()
	return handler.infos, handler.total, nil
}
