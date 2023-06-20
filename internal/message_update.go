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
	"github.com/google/uuid"
	"github.com/pulkitbhardwaj/matrix/internal/message"
	"github.com/pulkitbhardwaj/matrix/internal/predicate"
)

// MessageUpdate is the builder for updating Message entities.
type MessageUpdate struct {
	config
	hooks    []Hook
	mutation *MessageMutation
}

// Where appends a list predicates to the MessageUpdate builder.
func (mu *MessageUpdate) Where(ps ...predicate.Message) *MessageUpdate {
	mu.mutation.Where(ps...)
	return mu
}

// SetUpdatedAt sets the "updated_at" field.
func (mu *MessageUpdate) SetUpdatedAt(t time.Time) *MessageUpdate {
	mu.mutation.SetUpdatedAt(t)
	return mu
}

// SetBody sets the "body" field.
func (mu *MessageUpdate) SetBody(s string) *MessageUpdate {
	mu.mutation.SetBody(s)
	return mu
}

// AddFromIDs adds the "from" edge to the Message entity by IDs.
func (mu *MessageUpdate) AddFromIDs(ids ...uuid.UUID) *MessageUpdate {
	mu.mutation.AddFromIDs(ids...)
	return mu
}

// AddFrom adds the "from" edges to the Message entity.
func (mu *MessageUpdate) AddFrom(m ...*Message) *MessageUpdate {
	ids := make([]uuid.UUID, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return mu.AddFromIDs(ids...)
}

// AddToIDs adds the "to" edge to the Message entity by IDs.
func (mu *MessageUpdate) AddToIDs(ids ...uuid.UUID) *MessageUpdate {
	mu.mutation.AddToIDs(ids...)
	return mu
}

// AddTo adds the "to" edges to the Message entity.
func (mu *MessageUpdate) AddTo(m ...*Message) *MessageUpdate {
	ids := make([]uuid.UUID, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return mu.AddToIDs(ids...)
}

// Mutation returns the MessageMutation object of the builder.
func (mu *MessageUpdate) Mutation() *MessageMutation {
	return mu.mutation
}

// ClearFrom clears all "from" edges to the Message entity.
func (mu *MessageUpdate) ClearFrom() *MessageUpdate {
	mu.mutation.ClearFrom()
	return mu
}

// RemoveFromIDs removes the "from" edge to Message entities by IDs.
func (mu *MessageUpdate) RemoveFromIDs(ids ...uuid.UUID) *MessageUpdate {
	mu.mutation.RemoveFromIDs(ids...)
	return mu
}

// RemoveFrom removes "from" edges to Message entities.
func (mu *MessageUpdate) RemoveFrom(m ...*Message) *MessageUpdate {
	ids := make([]uuid.UUID, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return mu.RemoveFromIDs(ids...)
}

// ClearTo clears all "to" edges to the Message entity.
func (mu *MessageUpdate) ClearTo() *MessageUpdate {
	mu.mutation.ClearTo()
	return mu
}

// RemoveToIDs removes the "to" edge to Message entities by IDs.
func (mu *MessageUpdate) RemoveToIDs(ids ...uuid.UUID) *MessageUpdate {
	mu.mutation.RemoveToIDs(ids...)
	return mu
}

// RemoveTo removes "to" edges to Message entities.
func (mu *MessageUpdate) RemoveTo(m ...*Message) *MessageUpdate {
	ids := make([]uuid.UUID, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return mu.RemoveToIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (mu *MessageUpdate) Save(ctx context.Context) (int, error) {
	mu.defaults()
	return withHooks(ctx, mu.sqlSave, mu.mutation, mu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (mu *MessageUpdate) SaveX(ctx context.Context) int {
	affected, err := mu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (mu *MessageUpdate) Exec(ctx context.Context) error {
	_, err := mu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mu *MessageUpdate) ExecX(ctx context.Context) {
	if err := mu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (mu *MessageUpdate) defaults() {
	if _, ok := mu.mutation.UpdatedAt(); !ok {
		v := message.UpdateDefaultUpdatedAt()
		mu.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (mu *MessageUpdate) check() error {
	if v, ok := mu.mutation.Body(); ok {
		if err := message.BodyValidator(v); err != nil {
			return &ValidationError{Name: "body", err: fmt.Errorf(`internal: validator failed for field "Message.body": %w`, err)}
		}
	}
	return nil
}

func (mu *MessageUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := mu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(message.Table, message.Columns, sqlgraph.NewFieldSpec(message.FieldID, field.TypeUUID))
	if ps := mu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := mu.mutation.UpdatedAt(); ok {
		_spec.SetField(message.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := mu.mutation.Body(); ok {
		_spec.SetField(message.FieldBody, field.TypeString, value)
	}
	if mu.mutation.FromCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   message.FromTable,
			Columns: message.FromPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(message.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := mu.mutation.RemovedFromIDs(); len(nodes) > 0 && !mu.mutation.FromCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   message.FromTable,
			Columns: message.FromPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(message.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := mu.mutation.FromIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   message.FromTable,
			Columns: message.FromPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(message.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if mu.mutation.ToCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   message.ToTable,
			Columns: message.ToPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(message.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := mu.mutation.RemovedToIDs(); len(nodes) > 0 && !mu.mutation.ToCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   message.ToTable,
			Columns: message.ToPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(message.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := mu.mutation.ToIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   message.ToTable,
			Columns: message.ToPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(message.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, mu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{message.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	mu.mutation.done = true
	return n, nil
}

// MessageUpdateOne is the builder for updating a single Message entity.
type MessageUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *MessageMutation
}

// SetUpdatedAt sets the "updated_at" field.
func (muo *MessageUpdateOne) SetUpdatedAt(t time.Time) *MessageUpdateOne {
	muo.mutation.SetUpdatedAt(t)
	return muo
}

// SetBody sets the "body" field.
func (muo *MessageUpdateOne) SetBody(s string) *MessageUpdateOne {
	muo.mutation.SetBody(s)
	return muo
}

// AddFromIDs adds the "from" edge to the Message entity by IDs.
func (muo *MessageUpdateOne) AddFromIDs(ids ...uuid.UUID) *MessageUpdateOne {
	muo.mutation.AddFromIDs(ids...)
	return muo
}

// AddFrom adds the "from" edges to the Message entity.
func (muo *MessageUpdateOne) AddFrom(m ...*Message) *MessageUpdateOne {
	ids := make([]uuid.UUID, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return muo.AddFromIDs(ids...)
}

// AddToIDs adds the "to" edge to the Message entity by IDs.
func (muo *MessageUpdateOne) AddToIDs(ids ...uuid.UUID) *MessageUpdateOne {
	muo.mutation.AddToIDs(ids...)
	return muo
}

// AddTo adds the "to" edges to the Message entity.
func (muo *MessageUpdateOne) AddTo(m ...*Message) *MessageUpdateOne {
	ids := make([]uuid.UUID, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return muo.AddToIDs(ids...)
}

// Mutation returns the MessageMutation object of the builder.
func (muo *MessageUpdateOne) Mutation() *MessageMutation {
	return muo.mutation
}

// ClearFrom clears all "from" edges to the Message entity.
func (muo *MessageUpdateOne) ClearFrom() *MessageUpdateOne {
	muo.mutation.ClearFrom()
	return muo
}

// RemoveFromIDs removes the "from" edge to Message entities by IDs.
func (muo *MessageUpdateOne) RemoveFromIDs(ids ...uuid.UUID) *MessageUpdateOne {
	muo.mutation.RemoveFromIDs(ids...)
	return muo
}

// RemoveFrom removes "from" edges to Message entities.
func (muo *MessageUpdateOne) RemoveFrom(m ...*Message) *MessageUpdateOne {
	ids := make([]uuid.UUID, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return muo.RemoveFromIDs(ids...)
}

// ClearTo clears all "to" edges to the Message entity.
func (muo *MessageUpdateOne) ClearTo() *MessageUpdateOne {
	muo.mutation.ClearTo()
	return muo
}

// RemoveToIDs removes the "to" edge to Message entities by IDs.
func (muo *MessageUpdateOne) RemoveToIDs(ids ...uuid.UUID) *MessageUpdateOne {
	muo.mutation.RemoveToIDs(ids...)
	return muo
}

// RemoveTo removes "to" edges to Message entities.
func (muo *MessageUpdateOne) RemoveTo(m ...*Message) *MessageUpdateOne {
	ids := make([]uuid.UUID, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return muo.RemoveToIDs(ids...)
}

// Where appends a list predicates to the MessageUpdate builder.
func (muo *MessageUpdateOne) Where(ps ...predicate.Message) *MessageUpdateOne {
	muo.mutation.Where(ps...)
	return muo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (muo *MessageUpdateOne) Select(field string, fields ...string) *MessageUpdateOne {
	muo.fields = append([]string{field}, fields...)
	return muo
}

// Save executes the query and returns the updated Message entity.
func (muo *MessageUpdateOne) Save(ctx context.Context) (*Message, error) {
	muo.defaults()
	return withHooks(ctx, muo.sqlSave, muo.mutation, muo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (muo *MessageUpdateOne) SaveX(ctx context.Context) *Message {
	node, err := muo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (muo *MessageUpdateOne) Exec(ctx context.Context) error {
	_, err := muo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (muo *MessageUpdateOne) ExecX(ctx context.Context) {
	if err := muo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (muo *MessageUpdateOne) defaults() {
	if _, ok := muo.mutation.UpdatedAt(); !ok {
		v := message.UpdateDefaultUpdatedAt()
		muo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (muo *MessageUpdateOne) check() error {
	if v, ok := muo.mutation.Body(); ok {
		if err := message.BodyValidator(v); err != nil {
			return &ValidationError{Name: "body", err: fmt.Errorf(`internal: validator failed for field "Message.body": %w`, err)}
		}
	}
	return nil
}

func (muo *MessageUpdateOne) sqlSave(ctx context.Context) (_node *Message, err error) {
	if err := muo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(message.Table, message.Columns, sqlgraph.NewFieldSpec(message.FieldID, field.TypeUUID))
	id, ok := muo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`internal: missing "Message.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := muo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, message.FieldID)
		for _, f := range fields {
			if !message.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("internal: invalid field %q for query", f)}
			}
			if f != message.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := muo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := muo.mutation.UpdatedAt(); ok {
		_spec.SetField(message.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := muo.mutation.Body(); ok {
		_spec.SetField(message.FieldBody, field.TypeString, value)
	}
	if muo.mutation.FromCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   message.FromTable,
			Columns: message.FromPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(message.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := muo.mutation.RemovedFromIDs(); len(nodes) > 0 && !muo.mutation.FromCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   message.FromTable,
			Columns: message.FromPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(message.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := muo.mutation.FromIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   message.FromTable,
			Columns: message.FromPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(message.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if muo.mutation.ToCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   message.ToTable,
			Columns: message.ToPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(message.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := muo.mutation.RemovedToIDs(); len(nodes) > 0 && !muo.mutation.ToCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   message.ToTable,
			Columns: message.ToPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(message.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := muo.mutation.ToIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   message.ToTable,
			Columns: message.ToPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(message.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Message{config: muo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, muo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{message.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	muo.mutation.done = true
	return _node, nil
}