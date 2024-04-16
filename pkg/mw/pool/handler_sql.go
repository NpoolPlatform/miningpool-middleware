package pool

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/NpoolPlatform/miningpool-middleware/pkg/db/ent/pool"
)

func (h *Handler) baseKeys() (map[string]string, error) {
	vals := make(map[string]string)
	if h.ID != nil {
		strBytes, err := json.Marshal(*h.ID)
		if err != nil {
			return vals, err
		}
		vals[pool.FieldID] = string(strBytes)
	}
	if h.EntID != nil {
		strBytes, err := json.Marshal(*h.EntID)
		if err != nil {
			return vals, err
		}
		vals[pool.FieldEntID] = string(strBytes)
	}
	if h.MiningpoolType != nil {
		strBytes, err := json.Marshal(h.MiningpoolType.String())
		if err != nil {
			return vals, err
		}
		vals[pool.FieldMiningpoolType] = string(strBytes)
	}
	if h.Name != nil {
		strBytes, err := json.Marshal(*h.Name)
		if err != nil {
			return vals, err
		}
		vals[pool.FieldName] = string(strBytes)
	}
	if h.Site != nil {
		strBytes, err := json.Marshal(*h.Site)
		if err != nil {
			return vals, err
		}
		vals[pool.FieldSite] = string(strBytes)
	}
	if h.Description != nil {
		strBytes, err := json.Marshal(*h.Description)
		if err != nil {
			return vals, err
		}
		vals[pool.FieldDescription] = string(strBytes)
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
		vals[pool.FieldID] = string(strBytes)
	}
	if h.EntID != nil {
		strBytes, err := json.Marshal(*h.EntID)
		if err != nil {
			return vals, err
		}
		vals[pool.FieldEntID] = string(strBytes)
	}
	return vals, nil
}

//nolint:gocognit
func (h *Handler) genCreateSQL() (string, error) {
	vals, err := h.baseKeys()
	if err != nil {
		return "", err
	}
	delete(vals, pool.FieldID)

	now := uint32(time.Now().Unix())
	vals[pool.FieldCreatedAt] = fmt.Sprintf("%v", now)
	vals[pool.FieldUpdatedAt] = fmt.Sprintf("%v", now)
	vals[pool.FieldDeletedAt] = fmt.Sprintf("%v", 0)

	keys := []string{}
	selectVals := []string{}

	for k, v := range vals {
		keys = append(keys, k)
		selectVals = append(selectVals, fmt.Sprintf("%v as %v", v, k))
	}

	sql := fmt.Sprintf("insert into pools (%v) select * from (select %v) as tmp where not exists (select * from pools where miningpool_type='%v' and deleted_at=0);",
		strings.Join(keys, ","),
		strings.Join(selectVals, ","),
		h.MiningpoolType.String(),
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
	delete(vals, pool.FieldID)
	delete(vals, pool.FieldEntID)
	now := uint32(time.Now().Unix())
	vals[pool.FieldUpdatedAt] = fmt.Sprintf("%v", now)

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
	if v, ok := vals[pool.FieldMiningpoolType]; ok {
		subQKeys = append(subQKeys, fmt.Sprintf("tmp_table.%v!=%v", pool.FieldMiningpoolType, v))
	}

	sql := fmt.Sprintf("update pools set %v where %v and deleted_at=0 and  not exists (select 1 from(select * from pools as tmp_table where %v and tmp_table.deleted_at=0 limit 1) as tmp);",
		strings.Join(keys, ","),
		strings.Join(idKeys, " AND "),
		strings.Join(subQKeys, " AND "),
	)
	return sql, nil
}
