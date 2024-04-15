package coin

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/NpoolPlatform/miningpool-middleware/pkg/db/ent/coin"
)

//nolint:gocognit
func (h *Handler) genCreateSQL() (string, error) {
	vals := make(map[string]string)
	if h.EntID != nil {
		strBytes, err := json.Marshal(*h.EntID)
		if err != nil {
			return "", err
		}
		vals[coin.FieldEntID] = string(strBytes)
	}
	if h.MiningpoolType != nil {
		strBytes, err := json.Marshal(h.MiningpoolType.String())
		if err != nil {
			return "", err
		}
		vals[coin.FieldMiningpoolType] = string(strBytes)
	}
	if h.CoinType != nil {
		fmt.Println(*h.CoinType)
		strBytes, err := json.Marshal(h.CoinType.String())
		if err != nil {
			return "", err
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
			return "", err
		}
		vals[coin.FieldRevenueTypes] = fmt.Sprintf("'%v'", string(strBytes))
	}
	if h.FeeRate != nil {
		strBytes, err := json.Marshal(*h.FeeRate)
		if err != nil {
			return "", err
		}
		vals[coin.FieldFeeRate] = string(strBytes)
	}
	if h.FixedRevenueAble != nil {
		strBytes, err := json.Marshal(*h.FixedRevenueAble)
		if err != nil {
			return "", err
		}
		vals[coin.FieldFixedRevenueAble] = string(strBytes)
	}
	if h.Threshold != nil {
		strBytes, err := json.Marshal(*h.Threshold)
		if err != nil {
			return "", err
		}
		vals[coin.FieldThreshold] = string(strBytes)
	}
	if h.Remark != nil {
		strBytes, err := json.Marshal(*h.Remark)
		if err != nil {
			return "", err
		}
		vals[coin.FieldRemark] = string(strBytes)
	}

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

	fmt.Println(sql)

	return sql, nil
}
