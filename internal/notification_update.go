// Code generated by ent, DO NOT EDIT.

package internal

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/pulkitbhardwaj/matrix/internal/notification"
	"github.com/pulkitbhardwaj/matrix/internal/predicate"
)

// NotificationUpdate is the builder for updating Notification entities.
type NotificationUpdate struct {
	config
	hooks    []Hook
	mutation *NotificationMutation
}

// Where appends a list predicates to the NotificationUpdate builder.
func (nu *NotificationUpdate) Where(ps ...predicate.Notification) *NotificationUpdate {
	nu.mutation.Where(ps...)
	return nu
}

// SetUpdatedAt sets the "updated_at" field.
func (nu *NotificationUpdate) SetUpdatedAt(t time.Time) *NotificationUpdate {
	nu.mutation.SetUpdatedAt(t)
	return nu
}

// Mutation returns the NotificationMutation object of the builder.
func (nu *NotificationUpdate) Mutation() *NotificationMutation {
	return nu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (nu *NotificationUpdate) Save(ctx context.Context) (int, error) {
	nu.defaults()
	return withHooks(ctx, nu.sqlSave, nu.mutation, nu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (nu *NotificationUpdate) SaveX(ctx context.Context) int {
	affected, err := nu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (nu *NotificationUpdate) Exec(ctx context.Context) error {
	_, err := nu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (nu *NotificationUpdate) ExecX(ctx context.Context) {
	if err := nu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (nu *NotificationUpdate) defaults() {
	if _, ok := nu.mutation.UpdatedAt(); !ok {
		v := notification.UpdateDefaultUpdatedAt()
		nu.mutation.SetUpdatedAt(v)
	}
}

func (nu *NotificationUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(notification.Table, notification.Columns, sqlgraph.NewFieldSpec(notification.FieldID, field.TypeUUID))
	if ps := nu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := nu.mutation.UpdatedAt(); ok {
		_spec.SetField(notification.FieldUpdatedAt, field.TypeTime, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, nu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{notification.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	nu.mutation.done = true
	return n, nil
}

// NotificationUpdateOne is the builder for updating a single Notification entity.
type NotificationUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *NotificationMutation
}

// SetUpdatedAt sets the "updated_at" field.
func (nuo *NotificationUpdateOne) SetUpdatedAt(t time.Time) *NotificationUpdateOne {
	nuo.mutation.SetUpdatedAt(t)
	return nuo
}

// Mutation returns the NotificationMutation object of the builder.
func (nuo *NotificationUpdateOne) Mutation() *NotificationMutation {
	return nuo.mutation
}

// Where appends a list predicates to the NotificationUpdate builder.
func (nuo *NotificationUpdateOne) Where(ps ...predicate.Notification) *NotificationUpdateOne {
	nuo.mutation.Where(ps...)
	return nuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (nuo *NotificationUpdateOne) Select(field string, fields ...string) *NotificationUpdateOne {
	nuo.fields = append([]string{field}, fields...)
	return nuo
}

// Save executes the query and returns the updated Notification entity.
func (nuo *NotificationUpdateOne) Save(ctx context.Context) (*Notification, error) {
	nuo.defaults()
	return withHooks(ctx, nuo.sqlSave, nuo.mutation, nuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (nuo *NotificationUpdateOne) SaveX(ctx context.Context) *Notification {
	node, err := nuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (nuo *NotificationUpdateOne) Exec(ctx context.Context) error {
	_, err := nuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (nuo *NotificationUpdateOne) ExecX(ctx context.Context) {
	if err := nuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (nuo *NotificationUpdateOne) defaults() {
	if _, ok := nuo.mutation.UpdatedAt(); !ok {
		v := notification.UpdateDefaultUpdatedAt()
		nuo.mutation.SetUpdatedAt(v)
	}
}

func (nuo *NotificationUpdateOne) sqlSave(ctx context.Context) (_node *Notification, err error) {
	_spec := sqlgraph.NewUpdateSpec(notification.Table, notification.Columns, sqlgraph.NewFieldSpec(notification.FieldID, field.TypeUUID))
	id, ok := nuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`internal: missing "Notification.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := nuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, notification.FieldID)
		for _, f := range fields {
			if !notification.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("internal: invalid field %q for query", f)}
			}
			if f != notification.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := nuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := nuo.mutation.UpdatedAt(); ok {
		_spec.SetField(notification.FieldUpdatedAt, field.TypeTime, value)
	}
	_node = &Notification{config: nuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, nuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{notification.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	nuo.mutation.done = true
	return _node, nil
}
