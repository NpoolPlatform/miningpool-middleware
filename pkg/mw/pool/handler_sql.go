package pool

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/NpoolPlatform/miningpool-middleware/pkg/db/ent/pool"
)

//nolint:gocognit
func (h *Handler) genCreateSQL() (string, error) {
	vals := make(map[string]string)
	if h.EntID != nil {
		strBytes, err := json.Marshal(*h.EntID)
		if err != nil {
			return "", err
		}
		vals[pool.FieldEntID] = string(strBytes)
	}
	if h.MiningpoolType != nil {
		strBytes, err := json.Marshal(h.MiningpoolType.String())
		if err != nil {
			return "", err
		}
		vals[pool.FieldMiningpoolType] = string(strBytes)
	}
	if h.Name != nil {
		strBytes, err := json.Marshal(*h.Name)
		if err != nil {
			return "", err
		}
		vals[pool.FieldName] = string(strBytes)
	}
	if h.Site != nil {
		strBytes, err := json.Marshal(*h.Site)
		if err != nil {
			return "", err
		}
		vals[pool.FieldSite] = string(strBytes)
	}
	if h.Description != nil {
		strBytes, err := json.Marshal(*h.Description)
		if err != nil {
			return "", err
		}
		vals[pool.FieldDescription] = string(strBytes)
	}

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
