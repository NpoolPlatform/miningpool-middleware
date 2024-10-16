// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/miningpool-middleware/pkg/db/ent/fractionwithdrawalrule"
	"github.com/NpoolPlatform/miningpool-middleware/pkg/db/ent/predicate"
)

// FractionWithdrawalRuleDelete is the builder for deleting a FractionWithdrawalRule entity.
type FractionWithdrawalRuleDelete struct {
	config
	hooks    []Hook
	mutation *FractionWithdrawalRuleMutation
}

// Where appends a list predicates to the FractionWithdrawalRuleDelete builder.
func (fwrd *FractionWithdrawalRuleDelete) Where(ps ...predicate.FractionWithdrawalRule) *FractionWithdrawalRuleDelete {
	fwrd.mutation.Where(ps...)
	return fwrd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (fwrd *FractionWithdrawalRuleDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(fwrd.hooks) == 0 {
		affected, err = fwrd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*FractionWithdrawalRuleMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			fwrd.mutation = mutation
			affected, err = fwrd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(fwrd.hooks) - 1; i >= 0; i-- {
			if fwrd.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = fwrd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, fwrd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (fwrd *FractionWithdrawalRuleDelete) ExecX(ctx context.Context) int {
	n, err := fwrd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (fwrd *FractionWithdrawalRuleDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: fractionwithdrawalrule.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint32,
				Column: fractionwithdrawalrule.FieldID,
			},
		},
	}
	if ps := fwrd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, fwrd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	return affected, err
}

// FractionWithdrawalRuleDeleteOne is the builder for deleting a single FractionWithdrawalRule entity.
type FractionWithdrawalRuleDeleteOne struct {
	fwrd *FractionWithdrawalRuleDelete
}

// Exec executes the deletion query.
func (fwrdo *FractionWithdrawalRuleDeleteOne) Exec(ctx context.Context) error {
	n, err := fwrdo.fwrd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{fractionwithdrawalrule.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (fwrdo *FractionWithdrawalRuleDeleteOne) ExecX(ctx context.Context) {
	fwrdo.fwrd.ExecX(ctx)
}