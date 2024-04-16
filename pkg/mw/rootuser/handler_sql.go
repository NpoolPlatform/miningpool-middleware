package rootuser

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/NpoolPlatform/miningpool-middleware/pkg/db/ent/rootuser"
)

//nolint:gocognit
func (h *Handler) baseKeys() (map[string]string, error) {
	vals := make(map[string]string)
	if h.ID != nil {
		strBytes, err := json.Marshal(*h.ID)
		if err != nil {
			return vals, err
		}
		vals[rootuser.FieldID] = string(strBytes)
	}
	if h.EntID != nil {
		strBytes, err := json.Marshal(*h.EntID)
		if err != nil {
			return vals, err
		}
		vals[rootuser.FieldEntID] = string(strBytes)
	}
	if h.MiningpoolType != nil {
		strBytes, err := json.Marshal(h.MiningpoolType.String())
		if err != nil {
			return vals, err
		}
		vals[rootuser.FieldMiningpoolType] = string(strBytes)
	}
	if h.Name != nil {
		strBytes, err := json.Marshal(*h.Name)
		if err != nil {
			return vals, err
		}
		vals[rootuser.FieldName] = string(strBytes)
	}
	if h.Email != nil {
		strBytes, err := json.Marshal(*h.Email)
		if err != nil {
			return vals, err
		}
		vals[rootuser.FieldEmail] = string(strBytes)
	}
	if h.AuthToken != nil {
		strBytes, err := json.Marshal(*h.AuthToken)
		if err != nil {
			return vals, err
		}
		vals[rootuser.FieldAuthToken] = string(strBytes)
	}
	if h.AuthTokenSalt != nil {
		strBytes, err := json.Marshal(*h.AuthTokenSalt)
		if err != nil {
			return vals, err
		}
		vals[rootuser.FieldAuthTokenSalt] = string(strBytes)
	}
	if h.Authed != nil {
		strBytes, err := json.Marshal(*h.Authed)
		if err != nil {
			return vals, err
		}
		vals[rootuser.FieldAuthed] = string(strBytes)
	}
	if h.Remark != nil {
		strBytes, err := json.Marshal(*h.Remark)
		if err != nil {
			return vals, err
		}
		vals[rootuser.FieldRemark] = string(strBytes)
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
		vals[rootuser.FieldID] = string(strBytes)
	}
	if h.EntID != nil {
		strBytes, err := json.Marshal(*h.EntID)
		if err != nil {
			return vals, err
		}
		vals[rootuser.FieldEntID] = string(strBytes)
	}
	return vals, nil
}

func (h *Handler) genCreateSQL() (string, error) {
	vals, err := h.baseKeys()
	if err != nil {
		return "", err
	}
	delete(vals, rootuser.FieldID)

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
		*h.Name,
		*h.Email,
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
	delete(vals, rootuser.FieldID)
	delete(vals, rootuser.FieldEntID)
	now := uint32(time.Now().Unix())
	vals[rootuser.FieldUpdatedAt] = fmt.Sprintf("%v", now)

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
	if v, ok := vals[rootuser.FieldMiningpoolType]; ok {
		subQKeys = append(subQKeys, fmt.Sprintf("tmp_table.%v=%v", rootuser.FieldMiningpoolType, v))
	}
	if v, ok := vals[rootuser.FieldName]; ok {
		subQKeys = append(subQKeys, fmt.Sprintf("tmp_table.%v=%v", rootuser.FieldName, v))
	}
	if v, ok := vals[rootuser.FieldEmail]; ok {
		subQKeys = append(subQKeys, fmt.Sprintf("tmp_table.%v=%v", rootuser.FieldEmail, v))
	}

	sql := fmt.Sprintf("update root_users set %v where %v and deleted_at=0 and  not exists (select 1 from(select * from root_users as tmp_table where %v and tmp_table.deleted_at=0 limit 1) as tmp);",
		strings.Join(keys, ","),
		strings.Join(idKeys, " AND "),
		strings.Join(subQKeys, " AND "),
	)

	return sql, nil
}
