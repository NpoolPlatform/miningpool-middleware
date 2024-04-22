package gooduser

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/NpoolPlatform/miningpool-middleware/pkg/db/ent/gooduser"
)

//nolint:gocognit
func (h *Handler) baseKeys() (map[string]string, error) {
	vals := make(map[string]string)
	if h.ID != nil {
		strBytes, err := json.Marshal(*h.ID)
		if err != nil {
			return vals, err
		}
		vals[gooduser.FieldID] = string(strBytes)
	}
	if h.EntID != nil {
		strBytes, err := json.Marshal(*h.EntID)
		if err != nil {
			return vals, err
		}
		vals[gooduser.FieldEntID] = string(strBytes)
	}
	if h.MiningpoolType != nil {
		strBytes, err := json.Marshal(h.MiningpoolType.String())
		if err != nil {
			return vals, err
		}
		vals[gooduser.FieldMiningpoolType] = string(strBytes)
	}
	if h.Name != nil {
		strBytes, err := json.Marshal(*h.Name)
		if err != nil {
			return vals, err
		}
		vals[gooduser.FieldName] = string(strBytes)
	}
	if h.RootUserID != nil {
		strBytes, err := json.Marshal(*h.RootUserID)
		if err != nil {
			return vals, err
		}
		vals[gooduser.FieldRootUserID] = string(strBytes)
	}
	if h.CoinType != nil {
		strBytes, err := json.Marshal(h.CoinType.String())
		if err != nil {
			return vals, err
		}
		vals[gooduser.FieldCoinType] = string(strBytes)
	}
	if h.HashRate != nil {
		strBytes, err := json.Marshal(*h.HashRate)
		if err != nil {
			return vals, err
		}
		vals[gooduser.FieldHashRate] = string(strBytes)
	}
	if h.ReadPageLink != nil {
		strBytes, err := json.Marshal(*h.ReadPageLink)
		if err != nil {
			return vals, err
		}
		vals[gooduser.FieldReadPageLink] = string(strBytes)
	}
	if h.RevenueType != nil {
		strBytes, err := json.Marshal(h.RevenueType.String())
		if err != nil {
			return vals, err
		}
		vals[gooduser.FieldRevenueType] = string(strBytes)
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
		vals[gooduser.FieldID] = string(strBytes)
	}
	if h.EntID != nil {
		strBytes, err := json.Marshal(*h.EntID)
		if err != nil {
			return vals, err
		}
		vals[gooduser.FieldEntID] = string(strBytes)
	}
	return vals, nil
}

//nolint:gocognit
func (h *Handler) genCreateSQL() (string, error) {
	vals, err := h.baseKeys()
	if err != nil {
		return "", err
	}
	delete(vals, gooduser.FieldID)

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
		*h.Name,
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
	delete(vals, gooduser.FieldID)
	delete(vals, gooduser.FieldEntID)

	if len(vals) == 0 {
		return "", fmt.Errorf("update nothing")
	}

	now := uint32(time.Now().Unix())
	vals[gooduser.FieldUpdatedAt] = fmt.Sprintf("%v", now)

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
	if v, ok := vals[gooduser.FieldMiningpoolType]; ok {
		subQKeys = append(subQKeys, fmt.Sprintf("tmp_table.%v=%v", gooduser.FieldMiningpoolType, v))
	}
	if v, ok := vals[gooduser.FieldName]; ok {
		subQKeys = append(subQKeys, fmt.Sprintf("tmp_table.%v=%v", gooduser.FieldName, v))
	}

	sql := fmt.Sprintf("update good_users set %v where %v and deleted_at=0 and  not exists (select 1 from(select * from good_users as tmp_table where %v and tmp_table.deleted_at=0 limit 1) as tmp);",
		strings.Join(keys, ","),
		strings.Join(idKeys, " AND "),
		strings.Join(subQKeys, " AND "),
	)

	return sql, nil
}
