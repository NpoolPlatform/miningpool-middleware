// Code generated by ent, DO NOT EDIT.

package ent

import (
	"encoding/json"
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/NpoolPlatform/miningpool-middleware/pkg/db/ent/coin"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

// Coin is the model entity for the Coin schema.
type Coin struct {
	config `json:"-"`
	// ID of the ent.
	ID uint32 `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt uint32 `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt uint32 `json:"updated_at,omitempty"`
	// DeletedAt holds the value of the "deleted_at" field.
	DeletedAt uint32 `json:"deleted_at,omitempty"`
	// EntID holds the value of the "ent_id" field.
	EntID uuid.UUID `json:"ent_id,omitempty"`
	// MiningpoolType holds the value of the "miningpool_type" field.
	MiningpoolType string `json:"miningpool_type,omitempty"`
	// CoinType holds the value of the "coin_type" field.
	CoinType string `json:"coin_type,omitempty"`
	// RevenueTypes holds the value of the "revenue_types" field.
	RevenueTypes []string `json:"revenue_types,omitempty"`
	// FeeRate holds the value of the "fee_rate" field.
	FeeRate decimal.Decimal `json:"fee_rate,omitempty"`
	// FixedRevenueAble holds the value of the "fixed_revenue_able" field.
	FixedRevenueAble bool `json:"fixed_revenue_able,omitempty"`
	// Remark holds the value of the "remark" field.
	Remark string `json:"remark,omitempty"`
	// Threshold holds the value of the "threshold" field.
	Threshold decimal.Decimal `json:"threshold,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Coin) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case coin.FieldRevenueTypes:
			values[i] = new([]byte)
		case coin.FieldFeeRate, coin.FieldThreshold:
			values[i] = new(decimal.Decimal)
		case coin.FieldFixedRevenueAble:
			values[i] = new(sql.NullBool)
		case coin.FieldID, coin.FieldCreatedAt, coin.FieldUpdatedAt, coin.FieldDeletedAt:
			values[i] = new(sql.NullInt64)
		case coin.FieldMiningpoolType, coin.FieldCoinType, coin.FieldRemark:
			values[i] = new(sql.NullString)
		case coin.FieldEntID:
			values[i] = new(uuid.UUID)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Coin", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Coin fields.
func (c *Coin) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case coin.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			c.ID = uint32(value.Int64)
		case coin.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				c.CreatedAt = uint32(value.Int64)
			}
		case coin.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				c.UpdatedAt = uint32(value.Int64)
			}
		case coin.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				c.DeletedAt = uint32(value.Int64)
			}
		case coin.FieldEntID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field ent_id", values[i])
			} else if value != nil {
				c.EntID = *value
			}
		case coin.FieldMiningpoolType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field miningpool_type", values[i])
			} else if value.Valid {
				c.MiningpoolType = value.String
			}
		case coin.FieldCoinType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field coin_type", values[i])
			} else if value.Valid {
				c.CoinType = value.String
			}
		case coin.FieldRevenueTypes:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field revenue_types", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &c.RevenueTypes); err != nil {
					return fmt.Errorf("unmarshal field revenue_types: %w", err)
				}
			}
		case coin.FieldFeeRate:
			if value, ok := values[i].(*decimal.Decimal); !ok {
				return fmt.Errorf("unexpected type %T for field fee_rate", values[i])
			} else if value != nil {
				c.FeeRate = *value
			}
		case coin.FieldFixedRevenueAble:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field fixed_revenue_able", values[i])
			} else if value.Valid {
				c.FixedRevenueAble = value.Bool
			}
		case coin.FieldRemark:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field remark", values[i])
			} else if value.Valid {
				c.Remark = value.String
			}
		case coin.FieldThreshold:
			if value, ok := values[i].(*decimal.Decimal); !ok {
				return fmt.Errorf("unexpected type %T for field threshold", values[i])
			} else if value != nil {
				c.Threshold = *value
			}
		}
	}
	return nil
}

// Update returns a builder for updating this Coin.
// Note that you need to call Coin.Unwrap() before calling this method if this Coin
// was returned from a transaction, and the transaction was committed or rolled back.
func (c *Coin) Update() *CoinUpdateOne {
	return (&CoinClient{config: c.config}).UpdateOne(c)
}

// Unwrap unwraps the Coin entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (c *Coin) Unwrap() *Coin {
	_tx, ok := c.config.driver.(*txDriver)
	if !ok {
		panic("ent: Coin is not a transactional entity")
	}
	c.config.driver = _tx.drv
	return c
}

// String implements the fmt.Stringer.
func (c *Coin) String() string {
	var builder strings.Builder
	builder.WriteString("Coin(")
	builder.WriteString(fmt.Sprintf("id=%v, ", c.ID))
	builder.WriteString("created_at=")
	builder.WriteString(fmt.Sprintf("%v", c.CreatedAt))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(fmt.Sprintf("%v", c.UpdatedAt))
	builder.WriteString(", ")
	builder.WriteString("deleted_at=")
	builder.WriteString(fmt.Sprintf("%v", c.DeletedAt))
	builder.WriteString(", ")
	builder.WriteString("ent_id=")
	builder.WriteString(fmt.Sprintf("%v", c.EntID))
	builder.WriteString(", ")
	builder.WriteString("miningpool_type=")
	builder.WriteString(c.MiningpoolType)
	builder.WriteString(", ")
	builder.WriteString("coin_type=")
	builder.WriteString(c.CoinType)
	builder.WriteString(", ")
	builder.WriteString("revenue_types=")
	builder.WriteString(fmt.Sprintf("%v", c.RevenueTypes))
	builder.WriteString(", ")
	builder.WriteString("fee_rate=")
	builder.WriteString(fmt.Sprintf("%v", c.FeeRate))
	builder.WriteString(", ")
	builder.WriteString("fixed_revenue_able=")
	builder.WriteString(fmt.Sprintf("%v", c.FixedRevenueAble))
	builder.WriteString(", ")
	builder.WriteString("remark=")
	builder.WriteString(c.Remark)
	builder.WriteString(", ")
	builder.WriteString("threshold=")
	builder.WriteString(fmt.Sprintf("%v", c.Threshold))
	builder.WriteByte(')')
	return builder.String()
}

// Coins is a parsable slice of Coin.
type Coins []*Coin

func (c Coins) config(cfg config) {
	for _i := range c {
		c[_i].config = cfg
	}
}
