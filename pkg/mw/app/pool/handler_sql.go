package apppool

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/NpoolPlatform/miningpool-middleware/pkg/db/ent/apppool"
)

func (h *Handler) baseKeys() (map[string]string, error) {
	vals := make(map[string]string)
	if h.ID != nil {
		strBytes, err := json.Marshal(*h.ID)
		if err != nil {
			return vals, err
		}
		vals[apppool.FieldID] = string(strBytes)
	}
	if h.EntID != nil {
		strBytes, err := json.Marshal(*h.EntID)
		if err != nil {
			return vals, err
		}
		vals[apppool.FieldEntID] = string(strBytes)
	}
	if h.AppID != nil {
		strBytes, err := json.Marshal(*h.AppID)
		if err != nil {
			return vals, err
		}
		vals[apppool.FieldAppID] = string(strBytes)
	}
	if h.PoolID != nil {
		strBytes, err := json.Marshal(*h.PoolID)
		if err != nil {
			return vals, err
		}
		vals[apppool.FieldPoolID] = string(strBytes)
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
		vals[apppool.FieldID] = string(strBytes)
	}
	if h.EntID != nil {
		strBytes, err := json.Marshal(*h.EntID)
		if err != nil {
			return vals, err
		}
		vals[apppool.FieldEntID] = string(strBytes)
	}
	return vals, nil
}

//nolint:gocognit
func (h *Handler) genCreateSQL() (string, error) {
	vals, err := h.baseKeys()
	if err != nil {
		return "", err
	}
	delete(vals, apppool.FieldID)

	now := uint32(time.Now().Unix())
	vals[apppool.FieldCreatedAt] = fmt.Sprintf("%v", now)
	vals[apppool.FieldUpdatedAt] = fmt.Sprintf("%v", now)
	vals[apppool.FieldDeletedAt] = fmt.Sprintf("%v", 0)

	keys := []string{}
	selectVals := []string{}

	for k, v := range vals {
		keys = append(keys, k)
		selectVals = append(selectVals, fmt.Sprintf("%v as %v", v, k))
	}

	sql := fmt.Sprintf("insert into app_pools (%v) select * from (select %v) as tmp where not exists (select * from app_pools where app_id='%v' and pool_id='%v' and deleted_at=0);",
		strings.Join(keys, ","),
		strings.Join(selectVals, ","),
		h.AppID.String(),
		h.PoolID.String(),
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
	delete(vals, apppool.FieldID)
	delete(vals, apppool.FieldEntID)
	now := uint32(time.Now().Unix())
	vals[apppool.FieldUpdatedAt] = fmt.Sprintf("%v", now)

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
	if v, ok := vals[apppool.FieldAppID]; ok {
		subQKeys = append(subQKeys, fmt.Sprintf("tmp_table.%v!=%v", apppool.FieldAppID, v))
	}
	if v, ok := vals[apppool.FieldPoolID]; ok {
		subQKeys = append(subQKeys, fmt.Sprintf("tmp_table.%v!=%v", apppool.FieldPoolID, v))
	}

	sql := fmt.Sprintf("update app_pools set %v where %v and deleted_at=0 and  not exists (select 1 from(select * from app_pools as tmp_table where %v and tmp_table.deleted_at=0 limit 1) as tmp);",
		strings.Join(keys, ","),
		strings.Join(idKeys, " AND "),
		strings.Join(subQKeys, " AND "),
	)
	return sql, nil
}
