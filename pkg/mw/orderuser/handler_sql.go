package orderuser

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/NpoolPlatform/miningpool-middleware/pkg/db/ent/orderuser"
)

//nolint:gocognit
func (h *Handler) genCreateSQL() (string, error) {
	vals := make(map[string]string)
	if h.EntID != nil {
		strBytes, err := json.Marshal(*h.EntID)
		if err != nil {
			return "", err
		}
		vals[orderuser.FieldEntID] = string(strBytes)
	}
	if h.AppID != nil {
		strBytes, err := json.Marshal(*h.AppID)
		if err != nil {
			return "", err
		}
		vals[orderuser.FieldAppID] = string(strBytes)
	}
	if h.UserID != nil {
		strBytes, err := json.Marshal(*h.UserID)
		if err != nil {
			return "", err
		}
		vals[orderuser.FieldUserID] = string(strBytes)
	}
	if h.RootUserID != nil {
		strBytes, err := json.Marshal(*h.RootUserID)
		if err != nil {
			return "", err
		}
		vals[orderuser.FieldRootUserID] = string(strBytes)
	}
	if h.GoodUserID != nil {
		strBytes, err := json.Marshal(*h.GoodUserID)
		if err != nil {
			return "", err
		}
		vals[orderuser.FieldGoodUserID] = string(strBytes)
	}
	if h.Name != nil {
		strBytes, err := json.Marshal(*h.Name)
		if err != nil {
			return "", err
		}
		vals[orderuser.FieldName] = string(strBytes)
	}
	if h.MiningpoolType != nil {
		strBytes, err := json.Marshal(h.MiningpoolType.String())
		if err != nil {
			return "", err
		}
		vals[orderuser.FieldMiningpoolType] = string(strBytes)
	}
	if h.CoinType != nil {
		strBytes, err := json.Marshal(h.CoinType.String())
		if err != nil {
			return "", err
		}
		vals[orderuser.FieldCoinType] = string(strBytes)
	}
	if h.Proportion != nil {
		strBytes, err := json.Marshal(*h.Proportion)
		if err != nil {
			return "", err
		}
		vals[orderuser.FieldProportion] = string(strBytes)
	}
	if h.RevenueAddress != nil {
		strBytes, err := json.Marshal(*h.RevenueAddress)
		if err != nil {
			return "", err
		}
		vals[orderuser.FieldRevenueAddress] = string(strBytes)
	}
	if h.ReadPageLink != nil {
		strBytes, err := json.Marshal(*h.ReadPageLink)
		if err != nil {
			return "", err
		}
		vals[orderuser.FieldReadPageLink] = string(strBytes)
	}
	if h.AutoPay != nil {
		strBytes, err := json.Marshal(*h.AutoPay)
		if err != nil {
			return "", err
		}
		vals[orderuser.FieldAutoPay] = string(strBytes)
	}

	now := uint32(time.Now().Unix())
	vals[orderuser.FieldCreatedAt] = fmt.Sprintf("%v", now)
	vals[orderuser.FieldUpdatedAt] = fmt.Sprintf("%v", now)
	vals[orderuser.FieldDeletedAt] = fmt.Sprintf("%v", 0)

	keys := []string{}
	selectVals := []string{}
	for k, v := range vals {
		keys = append(keys, k)
		selectVals = append(selectVals, fmt.Sprintf("%v as %v", v, k))
	}

	sql := fmt.Sprintf("insert into order_users (%v) select * from (select %v) as tmp where not exists (select * from order_users where miningpool_type='%v' and name='%v' and  deleted_at=0);",
		strings.Join(keys, ","),
		strings.Join(selectVals, ","),
		h.MiningpoolType.String(),
		h.Name,
	)

	fmt.Println(sql)

	return sql, nil
}
