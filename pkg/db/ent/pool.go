// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/NpoolPlatform/miningpool-middleware/pkg/db/ent/pool"
	"github.com/google/uuid"
)

// Pool is the model entity for the Pool schema.
type Pool struct {
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
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Site holds the value of the "site" field.
	Site string `json:"site,omitempty"`
	// Description holds the value of the "description" field.
	Description string `json:"description,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Pool) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case pool.FieldID, pool.FieldCreatedAt, pool.FieldUpdatedAt, pool.FieldDeletedAt:
			values[i] = new(sql.NullInt64)
		case pool.FieldMiningpoolType, pool.FieldName, pool.FieldSite, pool.FieldDescription:
			values[i] = new(sql.NullString)
		case pool.FieldEntID:
			values[i] = new(uuid.UUID)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Pool", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Pool fields.
func (po *Pool) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case pool.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			po.ID = uint32(value.Int64)
		case pool.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				po.CreatedAt = uint32(value.Int64)
			}
		case pool.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				po.UpdatedAt = uint32(value.Int64)
			}
		case pool.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				po.DeletedAt = uint32(value.Int64)
			}
		case pool.FieldEntID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field ent_id", values[i])
			} else if value != nil {
				po.EntID = *value
			}
		case pool.FieldMiningpoolType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field miningpool_type", values[i])
			} else if value.Valid {
				po.MiningpoolType = value.String
			}
		case pool.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				po.Name = value.String
			}
		case pool.FieldSite:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field site", values[i])
			} else if value.Valid {
				po.Site = value.String
			}
		case pool.FieldDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description", values[i])
			} else if value.Valid {
				po.Description = value.String
			}
		}
	}
	return nil
}

// Update returns a builder for updating this Pool.
// Note that you need to call Pool.Unwrap() before calling this method if this Pool
// was returned from a transaction, and the transaction was committed or rolled back.
func (po *Pool) Update() *PoolUpdateOne {
	return (&PoolClient{config: po.config}).UpdateOne(po)
}

// Unwrap unwraps the Pool entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (po *Pool) Unwrap() *Pool {
	_tx, ok := po.config.driver.(*txDriver)
	if !ok {
		panic("ent: Pool is not a transactional entity")
	}
	po.config.driver = _tx.drv
	return po
}

// String implements the fmt.Stringer.
func (po *Pool) String() string {
	var builder strings.Builder
	builder.WriteString("Pool(")
	builder.WriteString(fmt.Sprintf("id=%v, ", po.ID))
	builder.WriteString("created_at=")
	builder.WriteString(fmt.Sprintf("%v", po.CreatedAt))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(fmt.Sprintf("%v", po.UpdatedAt))
	builder.WriteString(", ")
	builder.WriteString("deleted_at=")
	builder.WriteString(fmt.Sprintf("%v", po.DeletedAt))
	builder.WriteString(", ")
	builder.WriteString("ent_id=")
	builder.WriteString(fmt.Sprintf("%v", po.EntID))
	builder.WriteString(", ")
	builder.WriteString("miningpool_type=")
	builder.WriteString(po.MiningpoolType)
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(po.Name)
	builder.WriteString(", ")
	builder.WriteString("site=")
	builder.WriteString(po.Site)
	builder.WriteString(", ")
	builder.WriteString("description=")
	builder.WriteString(po.Description)
	builder.WriteByte(')')
	return builder.String()
}

// Pools is a parsable slice of Pool.
type Pools []*Pool

func (po Pools) config(cfg config) {
	for _i := range po {
		po[_i].config = cfg
	}
}