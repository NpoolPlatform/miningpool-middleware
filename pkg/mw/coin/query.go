package coin

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/miningpool/v1"
	npool "github.com/NpoolPlatform/message/npool/miningpool/mw/v1/coin"

	"github.com/NpoolPlatform/miningpool-middleware/pkg/db"
	"github.com/NpoolPlatform/miningpool-middleware/pkg/db/ent"
	coinent "github.com/NpoolPlatform/miningpool-middleware/pkg/db/ent/coin"

	coincrud "github.com/NpoolPlatform/miningpool-middleware/pkg/crud/coin"
)

type queryHandler struct {
	*Handler
	stm   *ent.CoinSelect
	infos []*npool.Coin
	total uint32
}

func (h *queryHandler) selectCoin(stm *ent.CoinQuery) {
	h.stm = stm.Select(
		coinent.FieldID,
		coinent.FieldEntID,
		coinent.FieldCoinType,
		coinent.FieldMiningpoolType,
		coinent.FieldRevenueTypes,
		coinent.FieldFeeRate,
		coinent.FieldFixedRevenueAble,
		coinent.FieldThreshold,
		coinent.FieldRemark,
		coinent.FieldCreatedAt,
		coinent.FieldUpdatedAt,
	)
}

func (h *queryHandler) queryCoin(cli *ent.Client) error {
	if h.ID == nil && h.EntID == nil {
		return fmt.Errorf("invalid id")
	}
	stm := cli.Coin.Query().Where(coinent.DeletedAt(0))
	if h.ID != nil {
		stm.Where(coinent.ID(*h.ID))
	}
	if h.EntID != nil {
		stm.Where(coinent.EntID(*h.EntID))
	}
	h.selectCoin(stm)
	return nil
}

func (h *queryHandler) queryCoins(ctx context.Context, cli *ent.Client) error {
	stm, err := coincrud.SetQueryConds(cli.Coin.Query(), h.Conds)
	if err != nil {
		return err
	}

	total, err := stm.Count(ctx)
	if err != nil {
		return err
	}
	h.total = uint32(total)

	h.selectCoin(stm)
	return nil
}

func (h *queryHandler) scan(ctx context.Context) error {
	return h.stm.Scan(ctx, &h.infos)
}

func (h *queryHandler) formalize() {
	for _, info := range h.infos {
		info.MiningpoolType = basetypes.MiningpoolType(basetypes.MiningpoolType_value[info.MiningpoolTypeStr])
		info.CoinType = basetypes.CoinType(basetypes.CoinType_value[info.CoinTypeStr])
		revenueTypes := []string{}
		err := json.Unmarshal([]byte(info.RevenueTypesStr), &revenueTypes)
		if err != nil {
			logger.Sugar().Warn(err)
		}
		for _, v := range revenueTypes {
			info.RevenueTypes = append(info.RevenueTypes, basetypes.RevenueType(basetypes.RevenueType_value[v]))
		}
	}
}

func (h *Handler) GetCoin(ctx context.Context) (*npool.Coin, error) {
	handler := &queryHandler{
		Handler: h,
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.queryCoin(cli); err != nil {
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

func (h *Handler) GetCoins(ctx context.Context) ([]*npool.Coin, uint32, error) {
	handler := &queryHandler{
		Handler: h,
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.queryCoins(ctx, cli); err != nil {
			return err
		}
		handler.stm.
			Offset(int(h.Offset)).
			Limit(int(h.Limit)).
			Order(ent.Desc(coinent.FieldUpdatedAt))
		return handler.scan(_ctx)
	})
	if err != nil {
		return nil, 0, err
	}

	handler.formalize()
	return handler.infos, handler.total, nil
}
