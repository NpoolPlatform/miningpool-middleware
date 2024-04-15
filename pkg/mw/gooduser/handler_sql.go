package gooduser

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/NpoolPlatform/miningpool-middleware/pkg/db/ent/gooduser"
)

//nolint:gocognit
func (h *Handler) genCreateSQL() (string, error) {
	vals := make(map[string]string)
	if h.EntID != nil {
		strBytes, err := json.Marshal(*h.EntID)
		if err != nil {
			return "", err
		}
		vals[gooduser.FieldEntID] = string(strBytes)
	}
	if h.MiningpoolType != nil {
		strBytes, err := json.Marshal(h.MiningpoolType.String())
		if err != nil {
			return "", err
		}
		vals[gooduser.FieldMiningpoolType] = string(strBytes)
	}
	if h.Name != nil {
		strBytes, err := json.Marshal(*h.Name)
		if err != nil {
			return "", err
		}
		vals[gooduser.FieldName] = string(strBytes)
	}
	if h.RootUserID != nil {
		strBytes, err := json.Marshal(*h.RootUserID)
		if err != nil {
			return "", err
		}
		vals[gooduser.FieldRootUserID] = string(strBytes)
	}
	if h.CoinType != nil {
		strBytes, err := json.Marshal(h.CoinType.String())
		if err != nil {
			return "", err
		}
		vals[gooduser.FieldCoinType] = string(strBytes)
	}
	if h.HashRate != nil {
		strBytes, err := json.Marshal(*h.HashRate)
		if err != nil {
			return "", err
		}
		vals[gooduser.FieldHashRate] = string(strBytes)
	}
	if h.ReadPageLink != nil {
		strBytes, err := json.Marshal(*h.ReadPageLink)
		if err != nil {
			return "", err
		}
		vals[gooduser.FieldReadPageLink] = string(strBytes)
	}
	if h.RevenueType != nil {
		strBytes, err := json.Marshal(h.RevenueType.String())
		if err != nil {
			return "", err
		}
		vals[gooduser.FieldRevenueType] = string(strBytes)
	}

	now := uint32(time.Now().Unix())
	vals[gooduser.FieldCreatedAt] = fmt.Sprintf("%v", now)
	vals[gooduser.FieldUpdatedAt] = fmt.Sprintf("%v", now)
	vals[gooduser.FieldDeletedAt] = fmt.Sprintf("%v", 0)

	keys := []string{}
	selectVals := []string{}
	for k, v := range vals {
		keys = append(keys, k)
		selectVals = append(selectVals, fmt.Sprintf("%v as %v", v, k))
	}

	sql := fmt.Sprintf("insert into good_users (%v) select * from (select %v) as tmp where not exists (select * from good_users where miningpool_type='%v' and name='%v' and  deleted_at=0);",
		strings.Join(keys, ","),
		strings.Join(selectVals, ","),
		h.MiningpoolType.String(),
		h.Name,
	)

	fmt.Println(sql)

	return sql, nil
}
