package pool

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	basetypes "github.com/NpoolPlatform/message/npool/basetypes/miningpool/v1"
	"github.com/NpoolPlatform/miningpool-middleware/pkg/db/ent/pool"
)

type sqlHandler struct {
	*Handler
	BondMiningpoolType *basetypes.MiningpoolType
	bondVals           map[string]string
	baseVals           map[string]string
	idVals             map[string]string
}

func (h *Handler) newSQLHandler() *sqlHandler {
	return &sqlHandler{
		Handler:  h,
		bondVals: make(map[string]string),
		baseVals: make(map[string]string),
		idVals:   make(map[string]string),
	}
}

//nolint:gocognit
func (h *sqlHandler) baseKeys() error {
	if h.ID != nil {
		strBytes, err := json.Marshal(*h.ID)
		if err != nil {
			return err
		}
		h.baseVals[pool.FieldID] = string(strBytes)
	}
	if h.EntID != nil {
		strBytes, err := json.Marshal(*h.EntID)
		if err != nil {
			return err
		}
		h.baseVals[pool.FieldEntID] = string(strBytes)
	}
	if h.MiningpoolType != nil {
		strBytes, err := json.Marshal(h.MiningpoolType.String())
		if err != nil {
			return err
		}
		h.baseVals[pool.FieldMiningpoolType] = string(strBytes)
		h.BondMiningpoolType = h.MiningpoolType
	}
	if h.Name != nil {
		strBytes, err := json.Marshal(*h.Name)
		if err != nil {
			return err
		}
		h.baseVals[pool.FieldName] = string(strBytes)
	}
	if h.Site != nil {
		strBytes, err := json.Marshal(*h.Site)
		if err != nil {
			return err
		}
		h.baseVals[pool.FieldSite] = string(strBytes)
	}
	if h.Logo != nil {
		strBytes, err := json.Marshal(*h.Logo)
		if err != nil {
			return err
		}
		h.baseVals[pool.FieldLogo] = string(strBytes)
	}
	if h.Description != nil {
		strBytes, err := json.Marshal(*h.Description)
		if err != nil {
			return err
		}
		h.baseVals[pool.FieldDescription] = string(strBytes)
	}

	if h.BondMiningpoolType == nil {
		return fmt.Errorf("please give miningpooltype")
	}
	strBytes, err := json.Marshal(h.BondMiningpoolType.String())
	if err != nil {
		return err
	}
	h.bondVals[pool.FieldMiningpoolType] = string(strBytes)
	return nil
}

func (h *sqlHandler) idKeys() error {
	if h.ID != nil {
		strBytes, err := json.Marshal(*h.ID)
		if err != nil {
			return err
		}
		h.idVals[pool.FieldID] = string(strBytes)
	}
	if h.EntID != nil {
		strBytes, err := json.Marshal(*h.EntID)
		if err != nil {
			return err
		}
		h.idVals[pool.FieldEntID] = string(strBytes)
	}
	return nil
}

//nolint:gocognit
func (h *sqlHandler) genCreateSQL() (string, error) {
	err := h.baseKeys()
	if err != nil {
		return "", err
	}
	delete(h.baseVals, pool.FieldID)

	now := uint32(time.Now().Unix())
	h.baseVals[pool.FieldCreatedAt] = fmt.Sprintf("%v", now)
	h.baseVals[pool.FieldUpdatedAt] = fmt.Sprintf("%v", now)
	h.baseVals[pool.FieldDeletedAt] = fmt.Sprintf("%v", 0)

	keys := []string{}
	selectVals := []string{}
	bondVals := []string{}

	for k, v := range h.baseVals {
		keys = append(keys, k)
		selectVals = append(selectVals, fmt.Sprintf("%v as %v", v, k))
	}

	for k, v := range h.bondVals {
		bondVals = append(bondVals, fmt.Sprintf("%v=%v", k, v))
	}

	sql := fmt.Sprintf("insert into %v (%v) select * from (select %v) as tmp where not exists (select * from %v where %v and deleted_at=0);",
		pool.Table,
		strings.Join(keys, ","),
		strings.Join(selectVals, ","),
		pool.Table,
		strings.Join(bondVals, " AND "),
	)

	return sql, nil
}

//nolint:gocognit
func (h *sqlHandler) genUpdateSQL() (string, error) {
	// get normal feilds
	err := h.baseKeys()
	if err != nil {
		return "", err
	}
	delete(h.baseVals, pool.FieldID)
	delete(h.baseVals, pool.FieldEntID)

	if len(h.baseVals) == 0 {
		return "", fmt.Errorf("update nothing")
	}

	now := uint32(time.Now().Unix())
	h.baseVals[pool.FieldUpdatedAt] = fmt.Sprintf("%v", now)

	keys := []string{}
	for k, v := range h.baseVals {
		keys = append(keys, fmt.Sprintf("%v=%v", k, v))
	}

	err = h.idKeys()
	if err != nil {
		return "", err
	}
	if len(h.idVals) == 0 {
		return "", fmt.Errorf("have neither id and ent_id")
	}

	// get id and ent_id feilds
	idKeys := []string{}
	// get sub query feilds
	bondVals := []string{}
	for k, v := range h.idVals {
		idKeys = append(idKeys, fmt.Sprintf("%v=%v", k, v))
		bondVals = append(bondVals, fmt.Sprintf("tmp_table.%v!=%v", k, v))
	}

	for k, v := range h.bondVals {
		bondVals = append(bondVals, fmt.Sprintf("tmp_table.%v=%v", k, v))
	}

	sql := fmt.Sprintf("update %v set %v where %v and deleted_at=0 and  not exists (select 1 from(select * from %v as tmp_table where %v and tmp_table.deleted_at=0 limit 1) as tmp);",
		pool.Table,
		strings.Join(keys, ","),
		strings.Join(idKeys, " AND "),
		pool.Table,
		strings.Join(bondVals, " AND "),
	)
	return sql, nil
}
