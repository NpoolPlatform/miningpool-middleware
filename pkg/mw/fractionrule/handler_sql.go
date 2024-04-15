package fractionrule

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/NpoolPlatform/miningpool-middleware/pkg/db/ent/fractionrule"
)

//nolint:gocognit
func (h *Handler) genCreateSQL() (string, error) {
	vals := make(map[string]string)
	if h.EntID != nil {
		strBytes, err := json.Marshal(*h.EntID)
		if err != nil {
			return "", err
		}
		vals[fractionrule.FieldEntID] = string(strBytes)
	}
	if h.MiningpoolType != nil {
		strBytes, err := json.Marshal(h.MiningpoolType.String())
		if err != nil {
			return "", err
		}
		vals[fractionrule.FieldMiningpoolType] = string(strBytes)
	}
	if h.CoinType != nil {
		strBytes, err := json.Marshal(h.CoinType.String())
		if err != nil {
			return "", err
		}
		vals[fractionrule.FieldCoinType] = string(strBytes)
	}
	if h.WithdrawInterval != nil {
		strBytes, err := json.Marshal(*h.WithdrawInterval)
		if err != nil {
			return "", err
		}
		vals[fractionrule.FieldWithdrawInterval] = string(strBytes)
	}
	if h.MinAmount != nil {
		strBytes, err := json.Marshal(*h.MinAmount)
		if err != nil {
			return "", err
		}
		vals[fractionrule.FieldMinAmount] = string(strBytes)
	}
	if h.WithdrawRate != nil {
		strBytes, err := json.Marshal(*h.WithdrawRate)
		if err != nil {
			return "", err
		}
		vals[fractionrule.FieldWithdrawRate] = string(strBytes)
	}

	now := uint32(time.Now().Unix())
	vals[fractionrule.FieldCreatedAt] = fmt.Sprintf("%v", now)
	vals[fractionrule.FieldUpdatedAt] = fmt.Sprintf("%v", now)
	vals[fractionrule.FieldDeletedAt] = fmt.Sprintf("%v", 0)

	keys := []string{}
	selectVals := []string{}
	for k, v := range vals {
		keys = append(keys, k)
		selectVals = append(selectVals, fmt.Sprintf("%v as %v", v, k))
	}

	sql := fmt.Sprintf("insert into fraction_rules (%v) select * from (select %v) as tmp where not exists (select * from fraction_rules where miningpool_type='%v' coin_type='%v' and deleted_at=0);",
		strings.Join(keys, ","),
		strings.Join(selectVals, ","),
		h.MiningpoolType.String(),
		h.CoinType.String(),
	)

	return sql, nil
}
