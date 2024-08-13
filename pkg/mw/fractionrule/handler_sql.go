package fractionrule

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/NpoolPlatform/go-service-framework/pkg/wlog"
	"github.com/NpoolPlatform/miningpool-middleware/pkg/db/ent/fractionrule"
	"github.com/google/uuid"
)

type sqlHandler struct {
	*Handler
	BondPoolCoinTypeID *uuid.UUID
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
			return wlog.WrapError(err)
		}
		h.baseVals[fractionrule.FieldID] = string(strBytes)
	}
	if h.EntID != nil {
		strBytes, err := json.Marshal(*h.EntID)
		if err != nil {
			return wlog.WrapError(err)
		}
		h.baseVals[fractionrule.FieldEntID] = string(strBytes)
	}
	if h.PoolCoinTypeID != nil {
		strBytes, err := json.Marshal(h.PoolCoinTypeID.String())
		if err != nil {
			return wlog.WrapError(err)
		}
		h.baseVals[fractionrule.FieldPoolCoinTypeID] = string(strBytes)
		h.BondPoolCoinTypeID = h.PoolCoinTypeID
	}
	if h.WithdrawInterval != nil {
		strBytes, err := json.Marshal(*h.WithdrawInterval)
		if err != nil {
			return wlog.WrapError(err)
		}
		h.baseVals[fractionrule.FieldWithdrawInterval] = string(strBytes)
	}
	if h.MinAmount != nil {
		strBytes, err := json.Marshal(*h.MinAmount)
		if err != nil {
			return wlog.WrapError(err)
		}
		h.baseVals[fractionrule.FieldMinAmount] = string(strBytes)
	}
	if h.PayoutThreshold != nil {
		strBytes, err := json.Marshal(*h.PayoutThreshold)
		if err != nil {
			return wlog.WrapError(err)
		}
		h.baseVals[fractionrule.FieldPayoutThreshold] = string(strBytes)
	}
	if h.WithdrawRate != nil {
		strBytes, err := json.Marshal(*h.WithdrawRate)
		if err != nil {
			return wlog.WrapError(err)
		}
		h.baseVals[fractionrule.FieldWithdrawRate] = string(strBytes)
	}

	if h.BondPoolCoinTypeID == nil {
		return wlog.Errorf("please give poolcointypeid")
	}
	strBytes, err := json.Marshal(h.BondPoolCoinTypeID.String())
	if err != nil {
		return wlog.WrapError(err)
	}
	h.bondVals[fractionrule.FieldPoolCoinTypeID] = string(strBytes)

	return nil
}

func (h *sqlHandler) idKeys() error {
	if h.ID != nil {
		strBytes, err := json.Marshal(*h.ID)
		if err != nil {
			return wlog.WrapError(err)
		}
		h.idVals[fractionrule.FieldID] = string(strBytes)
	}
	if h.EntID != nil {
		strBytes, err := json.Marshal(*h.EntID)
		if err != nil {
			return wlog.WrapError(err)
		}
		h.idVals[fractionrule.FieldEntID] = string(strBytes)
	}
	return nil
}

//nolint:gocognit
func (h *sqlHandler) genCreateSQL() (string, error) {
	err := h.baseKeys()
	if err != nil {
		return "", wlog.WrapError(err)
	}
	delete(h.baseVals, fractionrule.FieldID)

	now := uint32(time.Now().Unix())
	h.baseVals[fractionrule.FieldCreatedAt] = fmt.Sprintf("%v", now)
	h.baseVals[fractionrule.FieldUpdatedAt] = fmt.Sprintf("%v", now)
	h.baseVals[fractionrule.FieldDeletedAt] = fmt.Sprintf("%v", 0)

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
		fractionrule.Table,
		strings.Join(keys, ","),
		strings.Join(selectVals, ","),
		fractionrule.Table,
		strings.Join(bondVals, " AND "),
	)

	return sql, nil
}

//nolint:gocognit
func (h *sqlHandler) genUpdateSQL() (string, error) {
	// get normal feilds
	err := h.baseKeys()
	if err != nil {
		return "", wlog.WrapError(err)
	}

	delete(h.baseVals, fractionrule.FieldID)
	delete(h.baseVals, fractionrule.FieldEntID)

	if len(h.baseVals) == 0 {
		return "", wlog.Errorf("update nothing")
	}

	now := uint32(time.Now().Unix())
	h.baseVals[fractionrule.FieldUpdatedAt] = fmt.Sprintf("%v", now)

	keys := []string{}
	for k, v := range h.baseVals {
		keys = append(keys, fmt.Sprintf("%v=%v", k, v))
	}

	err = h.idKeys()
	if err != nil {
		return "", wlog.WrapError(err)
	}
	if len(h.idVals) == 0 {
		return "", wlog.Errorf("have neither id and ent_id")
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
		fractionrule.Table,
		strings.Join(keys, ","),
		strings.Join(idKeys, " AND "),
		fractionrule.Table,
		strings.Join(bondVals, " AND "),
	)
	return sql, nil
}
