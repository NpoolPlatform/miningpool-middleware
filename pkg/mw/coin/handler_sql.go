package coin

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/NpoolPlatform/miningpool-middleware/pkg/db/ent/coin"
)

//nolint:gocognit
func (h *Handler) baseKeys() (map[string]string, error) {
	vals := make(map[string]string)
	if h.ID != nil {
		strBytes, err := json.Marshal(*h.ID)
		if err != nil {
			return vals, err
		}
		vals[coin.FieldID] = string(strBytes)
	}
	if h.EntID != nil {
		strBytes, err := json.Marshal(*h.EntID)
		if err != nil {
			return vals, err
		}
		vals[coin.FieldEntID] = string(strBytes)
	}
	if h.MiningpoolType != nil {
		strBytes, err := json.Marshal(h.MiningpoolType.String())
		if err != nil {
			return vals, err
		}
		vals[coin.FieldMiningpoolType] = string(strBytes)
	}
	if h.CoinType != nil {
		fmt.Println(*h.CoinType)
		strBytes, err := json.Marshal(h.CoinType.String())
		if err != nil {
			return vals, err
		}
		vals[coin.FieldCoinType] = string(strBytes)
	}
	if h.RevenueTypes != nil {
		revenueTypes := []string{}
		for _, v := range *h.RevenueTypes {
			revenueTypes = append(revenueTypes, v.String())
		}
		strBytes, err := json.Marshal(revenueTypes)
		if err != nil {
			return vals, err
		}
		vals[coin.FieldRevenueTypes] = fmt.Sprintf("'%v'", string(strBytes))
	}
	if h.FeeRate != nil {
		strBytes, err := json.Marshal(*h.FeeRate)
		if err != nil {
			return vals, err
		}
		vals[coin.FieldFeeRate] = string(strBytes)
	}
	if h.FixedRevenueAble != nil {
		strBytes, err := json.Marshal(*h.FixedRevenueAble)
		if err != nil {
			return vals, err
		}
		vals[coin.FieldFixedRevenueAble] = string(strBytes)
	}
	if h.Threshold != nil {
		strBytes, err := json.Marshal(*h.Threshold)
		if err != nil {
			return vals, err
		}
		vals[coin.FieldThreshold] = string(strBytes)
	}
	if h.Remark != nil {
		strBytes, err := json.Marshal(*h.Remark)
		if err != nil {
			return vals, err
		}
		vals[coin.FieldRemark] = string(strBytes)
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
		vals[coin.FieldID] = string(strBytes)
	}
	if h.EntID != nil {
		strBytes, err := json.Marshal(*h.EntID)
		if err != nil {
			return vals, err
		}
		vals[coin.FieldEntID] = string(strBytes)
	}
	return vals, nil
}

//nolint:gocognit
func (h *Handler) genCreateSQL() (string, error) {
	vals, err := h.baseKeys()
	if err != nil {
		return "", err
	}
	delete(vals, coin.FieldID)

	now := uint32(time.Now().Unix())
	vals[coin.FieldCreatedAt] = fmt.Sprintf("%v", now)
	vals[coin.FieldUpdatedAt] = fmt.Sprintf("%v", now)
	vals[coin.FieldDeletedAt] = fmt.Sprintf("%v", 0)

	keys := []string{}
	selectVals := []string{}

	for k, v := range vals {
		keys = append(keys, k)
		selectVals = append(selectVals, fmt.Sprintf("%v as %v", v, k))
	}

	sql := fmt.Sprintf("insert into coins (%v) select * from (select %v) as tmp where not exists (select * from coins where miningpool_type='%v' and coin_type='%v' and deleted_at=0);",
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
	delete(vals, coin.FieldID)
	delete(vals, coin.FieldEntID)
	now := uint32(time.Now().Unix())
	vals[coin.FieldUpdatedAt] = fmt.Sprintf("%v", now)

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
	if v, ok := vals[coin.FieldMiningpoolType]; ok {
		subQKeys = append(subQKeys, fmt.Sprintf("tmp_table.%v=%v", coin.FieldMiningpoolType, v))
	}
	if v, ok := vals[coin.FieldCoinType]; ok {
		subQKeys = append(subQKeys, fmt.Sprintf("tmp_table.%v=%v", coin.FieldCoinType, v))
	}

	sql := fmt.Sprintf("update coins set %v where %v and deleted_at=0 and  not exists (select 1 from(select * from coins as tmp_table where %v and tmp_table.deleted_at=0 limit 1) as tmp);",
		strings.Join(keys, ","),
		strings.Join(idKeys, " AND "),
		strings.Join(subQKeys, " AND "),
	)
	return sql, nil
}
