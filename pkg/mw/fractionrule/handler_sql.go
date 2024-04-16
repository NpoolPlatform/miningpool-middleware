package fractionrule

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/NpoolPlatform/miningpool-middleware/pkg/db/ent/fractionrule"
)

//nolint:gocognit
func (h *Handler) baseKeys() (map[string]string, error) {
	vals := make(map[string]string)
	if h.ID != nil {
		strBytes, err := json.Marshal(*h.ID)
		if err != nil {
			return vals, err
		}
		vals[fractionrule.FieldID] = string(strBytes)
	}
	if h.EntID != nil {
		strBytes, err := json.Marshal(*h.EntID)
		if err != nil {
			return vals, err
		}
		vals[fractionrule.FieldEntID] = string(strBytes)
	}
	if h.MiningpoolType != nil {
		strBytes, err := json.Marshal(h.MiningpoolType.String())
		if err != nil {
			return vals, err
		}
		vals[fractionrule.FieldMiningpoolType] = string(strBytes)
	}
	if h.CoinType != nil {
		strBytes, err := json.Marshal(h.CoinType.String())
		if err != nil {
			return vals, err
		}
		vals[fractionrule.FieldCoinType] = string(strBytes)
	}
	if h.WithdrawInterval != nil {
		strBytes, err := json.Marshal(*h.WithdrawInterval)
		if err != nil {
			return vals, err
		}
		vals[fractionrule.FieldWithdrawInterval] = string(strBytes)
	}
	if h.MinAmount != nil {
		strBytes, err := json.Marshal(*h.MinAmount)
		if err != nil {
			return vals, err
		}
		vals[fractionrule.FieldMinAmount] = string(strBytes)
	}
	if h.WithdrawRate != nil {
		strBytes, err := json.Marshal(*h.WithdrawRate)
		if err != nil {
			return vals, err
		}
		vals[fractionrule.FieldWithdrawRate] = string(strBytes)
	}
	return vals, nil
}

func (h *Handler) idKeys() (map[string]string, error) {
	vals := make(map[string]string)
	if h.ID != nil {
		strBytes, err := json.Marshal(*h.ID)
		if err != nil {
			return vals, err
		}
		vals[fractionrule.FieldID] = string(strBytes)
	}
	if h.EntID != nil {
		strBytes, err := json.Marshal(*h.EntID)
		if err != nil {
			return vals, err
		}
		vals[fractionrule.FieldEntID] = string(strBytes)
	}
	return vals, nil
}

//nolint:gocognit
func (h *Handler) genCreateSQL() (string, error) {
	vals, err := h.baseKeys()
	if err != nil {
		return "", err
	}
	delete(vals, fractionrule.FieldID)

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

	sql := fmt.Sprintf("insert into fraction_rules (%v) select * from (select %v) as tmp where not exists (select * from fraction_rules where miningpool_type='%v' and coin_type='%v' and deleted_at=0);",
		strings.Join(keys, ","),
		strings.Join(selectVals, ","),
		h.MiningpoolType.String(),
		h.CoinType.String(),
	)

	return sql, nil
}

//nolint:gocognit
func (h *Handler) genUpdateSQL() (string, error) {
	// get normal feilds
	vals, err := h.baseKeys()
	if err != nil {
		return "", err
	}
	delete(vals, fractionrule.FieldID)
	delete(vals, fractionrule.FieldEntID)
	now := uint32(time.Now().Unix())
	vals[fractionrule.FieldUpdatedAt] = fmt.Sprintf("%v", now)

	keys := []string{}
	for k, v := range vals {
		keys = append(keys, fmt.Sprintf("%v=%v", k, v))
	}

	idVals, err := h.idKeys()
	if err != nil {
		return "", err
	}
	if len(idVals) == 0 {
		return "", fmt.Errorf("have neither id and ent_id")
	}

	// get id and ent_id feilds
	idKeys := []string{}
	// get sub query feilds
	subQKeys := []string{}
	for k, v := range idVals {
		idKeys = append(idKeys, fmt.Sprintf("%v=%v", k, v))
		subQKeys = append(subQKeys, fmt.Sprintf("tmp_table.%v!=%v", k, v))
	}
	if v, ok := vals[fractionrule.FieldMiningpoolType]; ok {
		subQKeys = append(subQKeys, fmt.Sprintf("tmp_table.%v=%v", fractionrule.FieldMiningpoolType, v))
	}
	if v, ok := vals[fractionrule.FieldCoinType]; ok {
		subQKeys = append(subQKeys, fmt.Sprintf("tmp_table.%v=%v", fractionrule.FieldCoinType, v))
	}

	sql := fmt.Sprintf("update fraction_rules set %v where %v and deleted_at=0 and  not exists (select 1 from(select * from fraction_rules as tmp_table where %v and tmp_table.deleted_at=0 limit 1) as tmp);",
		strings.Join(keys, ","),
		strings.Join(idKeys, " AND "),
		strings.Join(subQKeys, " AND "),
	)
	return sql, nil
}
