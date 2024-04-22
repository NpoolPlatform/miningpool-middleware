package orderuser

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/NpoolPlatform/miningpool-middleware/pkg/db/ent/orderuser"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type sqlHandler struct {
	*Handler
	baseVals map[string]string
	idVals   map[string]string
}

func newSQLHandler(h *Handler) *sqlHandler {
	return &sqlHandler{
		Handler:  h,
		baseVals: make(map[string]string),
		idVals:   make(map[string]string),
	}
}

//nolint:gocognit
func (h *sqlHandler) baseKeysDefault() error {
	if h.EntID == nil {
		entID, err := uuid.NewUUID()
		if err != nil {
			return err
		}
		h.EntID = &entID
	}
	if h.Proportion == nil {
		h.Proportion = &decimal.Decimal{}
	}
	return nil
}

//nolint:gocognit
func (h *sqlHandler) baseKeysFiled() error {
	if h.ID != nil {
		strBytes, err := json.Marshal(*h.ID)
		if err != nil {
			return err
		}
		h.baseVals[orderuser.FieldID] = string(strBytes)
	}
	if h.EntID != nil {
		strBytes, err := json.Marshal(*h.EntID)
		if err != nil {
			return err
		}
		h.baseVals[orderuser.FieldEntID] = string(strBytes)
	}
	if h.AppID != nil {
		strBytes, err := json.Marshal(*h.AppID)
		if err != nil {
			return err
		}
		h.baseVals[orderuser.FieldAppID] = string(strBytes)
	}
	if h.UserID != nil {
		strBytes, err := json.Marshal(*h.UserID)
		if err != nil {
			return err
		}
		h.baseVals[orderuser.FieldUserID] = string(strBytes)
	}
	if h.RootUserID != nil {
		strBytes, err := json.Marshal(*h.RootUserID)
		if err != nil {
			return err
		}
		h.baseVals[orderuser.FieldRootUserID] = string(strBytes)
	}
	if h.GoodUserID != nil {
		strBytes, err := json.Marshal(*h.GoodUserID)
		if err != nil {
			return err
		}
		h.baseVals[orderuser.FieldGoodUserID] = string(strBytes)
	}
	if h.Name != nil {
		strBytes, err := json.Marshal(*h.Name)
		if err != nil {
			return err
		}
		h.baseVals[orderuser.FieldName] = string(strBytes)
	}
	if h.MiningpoolType != nil {
		strBytes, err := json.Marshal(h.MiningpoolType.String())
		if err != nil {
			return err
		}
		h.baseVals[orderuser.FieldMiningpoolType] = string(strBytes)
	}
	if h.CoinType != nil {
		strBytes, err := json.Marshal(h.CoinType.String())
		if err != nil {
			return err
		}
		h.baseVals[orderuser.FieldCoinType] = string(strBytes)
	}
	if h.Proportion != nil {
		strBytes, err := json.Marshal(*h.Proportion)
		if err != nil {
			return err
		}
		h.baseVals[orderuser.FieldProportion] = string(strBytes)
	}
	if h.RevenueAddress != nil {
		strBytes, err := json.Marshal(*h.RevenueAddress)
		if err != nil {
			return err
		}
		h.baseVals[orderuser.FieldRevenueAddress] = string(strBytes)
	}
	if h.ReadPageLink != nil {
		strBytes, err := json.Marshal(*h.ReadPageLink)
		if err != nil {
			return err
		}
		h.baseVals[orderuser.FieldReadPageLink] = string(strBytes)
	}
	if h.AutoPay != nil {
		strBytes, err := json.Marshal(*h.AutoPay)
		if err != nil {
			return err
		}
		h.baseVals[orderuser.FieldAutoPay] = string(strBytes)
	}
	return nil
}

func (h *sqlHandler) idKeys() error {
	if h.ID != nil {
		strBytes, err := json.Marshal(*h.ID)
		if err != nil {
			return err
		}
		h.idVals[orderuser.FieldID] = string(strBytes)
	}
	if h.EntID != nil {
		strBytes, err := json.Marshal(*h.EntID)
		if err != nil {
			return err
		}
		h.idVals[orderuser.FieldEntID] = string(strBytes)
	}
	return nil
}

//nolint:gocognit
func (h *sqlHandler) genCreateSQL() (string, error) {
	err := h.baseKeysDefault()
	if err != nil {
		return "", err
	}
	err = h.baseKeysFiled()
	if err != nil {
		return "", err
	}
	delete(h.baseVals, orderuser.FieldID)

	now := uint32(time.Now().Unix())
	h.baseVals[orderuser.FieldCreatedAt] = fmt.Sprintf("%v", now)
	h.baseVals[orderuser.FieldUpdatedAt] = fmt.Sprintf("%v", now)
	h.baseVals[orderuser.FieldDeletedAt] = fmt.Sprintf("%v", 0)

	keys := []string{}
	selectVals := []string{}

	for k, v := range h.baseVals {
		keys = append(keys, k)
		selectVals = append(selectVals, fmt.Sprintf("%v as %v", v, k))
	}

	sql := fmt.Sprintf("insert into %v (%v) select * from (select %v) as tmp where not exists (select * from %v where miningpool_type='%v' and name='%v' and  deleted_at=0);",
		orderuser.Table,
		strings.Join(keys, ","),
		strings.Join(selectVals, ","),
		orderuser.Table,
		h.MiningpoolType.String(),
		*h.Name,
	)

	return sql, nil
}

//nolint:gocognit
func (h *sqlHandler) genUpdateSQL() (string, error) {
	// get normal feilds
	err := h.baseKeysFiled()
	if err != nil {
		return "", err
	}
	delete(h.baseVals, orderuser.FieldID)
	delete(h.baseVals, orderuser.FieldEntID)

	if len(h.baseVals) == 0 {
		return "", fmt.Errorf("update nothing")
	}

	now := uint32(time.Now().Unix())
	h.baseVals[orderuser.FieldUpdatedAt] = fmt.Sprintf("%v", now)

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
	subQKeys := []string{}
	for k, v := range h.idVals {
		idKeys = append(idKeys, fmt.Sprintf("%v=%v", k, v))
		subQKeys = append(subQKeys, fmt.Sprintf("tmp_table.%v!=%v", k, v))
	}
	if v, ok := h.baseVals[orderuser.FieldMiningpoolType]; ok {
		subQKeys = append(subQKeys, fmt.Sprintf("tmp_table.%v=%v", orderuser.FieldMiningpoolType, v))
	}
	if v, ok := h.baseVals[orderuser.FieldName]; ok {
		subQKeys = append(subQKeys, fmt.Sprintf("tmp_table.%v=%v", orderuser.FieldName, v))
	}

	sql := fmt.Sprintf("update order_users set %v where %v and deleted_at=0 and  not exists (select 1 from(select * from order_users as tmp_table where %v and tmp_table.deleted_at=0 limit 1) as tmp);",
		strings.Join(keys, ","),
		strings.Join(idKeys, " AND "),
		strings.Join(subQKeys, " AND "),
	)
	return sql, nil
}
