package rootuser

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/NpoolPlatform/miningpool-middleware/pkg/db/ent/rootuser"
)

//nolint:gocognit
func (h *Handler) genCreateSQL() (string, error) {
	vals := make(map[string]string)
	if h.EntID != nil {
		strBytes, err := json.Marshal(*h.EntID)
		if err != nil {
			return "", err
		}
		vals[rootuser.FieldEntID] = string(strBytes)
	}
	if h.MiningpoolType != nil {
		strBytes, err := json.Marshal(h.MiningpoolType.String())
		if err != nil {
			return "", err
		}
		vals[rootuser.FieldMiningpoolType] = string(strBytes)
	}
	if h.Name != nil {
		strBytes, err := json.Marshal(*h.Name)
		if err != nil {
			return "", err
		}
		vals[rootuser.FieldName] = string(strBytes)
	}
	if h.Email != nil {
		strBytes, err := json.Marshal(*h.Email)
		if err != nil {
			return "", err
		}
		vals[rootuser.FieldEmail] = string(strBytes)
	}
	if h.AuthToken != nil {
		strBytes, err := json.Marshal(*h.AuthToken)
		if err != nil {
			return "", err
		}
		vals[rootuser.FieldAuthToken] = string(strBytes)
	}
	if h.Authed != nil {
		strBytes, err := json.Marshal(*h.Authed)
		if err != nil {
			return "", err
		}
		vals[rootuser.FieldAuthed] = string(strBytes)
	}
	if h.Remark != nil {
		strBytes, err := json.Marshal(*h.Remark)
		if err != nil {
			return "", err
		}
		vals[rootuser.FieldRemark] = string(strBytes)
	}

	now := uint32(time.Now().Unix())
	vals[rootuser.FieldCreatedAt] = fmt.Sprintf("%v", now)
	vals[rootuser.FieldUpdatedAt] = fmt.Sprintf("%v", now)
	vals[rootuser.FieldDeletedAt] = fmt.Sprintf("%v", 0)

	keys := []string{}
	selectVals := []string{}
	for k, v := range vals {
		keys = append(keys, k)
		selectVals = append(selectVals, fmt.Sprintf("%v as %v", v, k))
	}

	sql := fmt.Sprintf("insert into root_users (%v) select * from (select %v) as tmp where not exists (select * from root_users where miningpool_type='%v' and name='%v' and email='%v' and  deleted_at=0);",
		strings.Join(keys, ","),
		strings.Join(selectVals, ","),
		h.MiningpoolType.String(),
		h.Name,
		h.Email,
	)

	return sql, nil
}
