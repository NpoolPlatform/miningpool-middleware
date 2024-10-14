package fractionwithdrawalrule

import (
	"entgo.io/ent/dialect/sql"
	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	"github.com/google/uuid"

	"github.com/NpoolPlatform/miningpool-middleware/pkg/db/ent"
	"github.com/NpoolPlatform/miningpool-middleware/pkg/db/ent/coin"
	fractionwithdrawalruleent "github.com/NpoolPlatform/miningpool-middleware/pkg/db/ent/fractionwithdrawalrule"
	"github.com/NpoolPlatform/miningpool-middleware/pkg/db/ent/pool"
)

func (h *Handler) queryJoin(stm *ent.FractionWithdrawalRuleQuery) {
	stm.Modify(
		h.queryJoinCoinAndPool,
	)
}

func (h *Handler) queryJoinCoinAndPool(s *sql.Selector) {
	coinT := sql.Table(coin.Table)
	poolT := sql.Table(pool.Table)
	s.Join(coinT).On(
		s.C(fractionwithdrawalruleent.FieldPoolCoinTypeID),
		coinT.C(coin.FieldEntID),
	).OnP(
		sql.EQ(coinT.C(coin.FieldDeletedAt), 0),
	).Join(poolT).On(
		coinT.C(coin.FieldPoolID),
		poolT.C(pool.FieldEntID),
	).OnP(
		sql.EQ(poolT.C(pool.FieldDeletedAt), 0),
	).AppendSelect(
		coinT.C(coin.FieldCoinType),
		coinT.C(coin.FieldPoolID),
		coinT.C(coin.FieldCoinTypeID),
		poolT.C(pool.FieldMiningPoolType),
	)

	if h.PoolConds != nil && h.PoolConds.EntID != nil {
		_id, ok := h.PoolConds.EntID.Val.(uuid.UUID)
		if !ok {
			logger.Sugar().Error("invalid poolid")
			return
		}

		if h.PoolConds.EntID.Op == cruder.EQ {
			s.OnP(
				sql.EQ(poolT.C(pool.FieldEntID), _id.String()),
			)
		}
	}
}
