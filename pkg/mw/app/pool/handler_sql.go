package apppool

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/NpoolPlatform/miningpool-middleware/pkg/db/ent/apppool"
)

func (h *Handler) genCreateSQL() (string, error) {
	vals := make(map[string]string)
	if h.EntID != nil {
		strBytes, err := json.Marshal(*h.EntID)
		if err != nil {
			return "", err
		}
		vals[apppool.FieldEntID] = string(strBytes)
	}
	if h.AppID != nil {
		strBytes, err := json.Marshal(*h.AppID)
		if err != nil {
			return "", err
		}
		vals[apppool.FieldAppID] = string(strBytes)
	}
	if h.PoolID != nil {
		strBytes, err := json.Marshal(*h.PoolID)
		if err != nil {
			return "", err
		}
		vals[apppool.FieldPoolID] = string(strBytes)
	}

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
