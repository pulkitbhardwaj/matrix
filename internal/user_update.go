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
	"github.com/pulkitbhardwaj/matrix/internal/predicate"
	"github.com/pulkitbhardwaj/matrix/internal/user"
)

// UserUpdate is the builder for updating User entities.
type UserUpdate struct {
	config
	hooks    []Hook
	mutation *UserMutation
}

// Where appends a list predicates to the UserUpdate builder.
func (uu *UserUpdate) Where(ps ...predicate.User) *UserUpdate {
	uu.mutation.Where(ps...)
	return uu
}

// SetUpdatedAt sets the "updated_at" field.
func (uu *UserUpdate) SetUpdatedAt(t time.Time) *UserUpdate {
	uu.mutation.SetUpdatedAt(t)
	return uu
}

// SetFirstName sets the "first_name" field.
func (uu *UserUpdate) SetFirstName(s string) *UserUpdate {
	uu.mutation.SetFirstName(s)
	return uu
}

// SetLastName sets the "last_name" field.
func (uu *UserUpdate) SetLastName(s string) *UserUpdate {
	uu.mutation.SetLastName(s)
	return uu
}

// SetEmailAddress sets the "email_address" field.
func (uu *UserUpdate) SetEmailAddress(s string) *UserUpdate {
	uu.mutation.SetEmailAddress(s)
	return uu
}

// SetUserName sets the "user_name" field.
func (uu *UserUpdate) SetUserName(s string) *UserUpdate {
	uu.mutation.SetUserName(s)
	return uu
}

// SetBio sets the "bio" field.
func (uu *UserUpdate) SetBio(s string) *UserUpdate {
	uu.mutation.SetBio(s)
	return uu
}

// SetNillableBio sets the "bio" field if the given value is not nil.
func (uu *UserUpdate) SetNillableBio(s *string) *UserUpdate {
	if s != nil {
		uu.SetBio(*s)
	}
	return uu
}

// ClearBio clears the value of the "bio" field.
func (uu *UserUpdate) ClearBio() *UserUpdate {
	uu.mutation.ClearBio()
	return uu
}

// AddFollowerIDs adds the "followers" edge to the User entity by IDs.
func (uu *UserUpdate) AddFollowerIDs(ids ...uuid.UUID) *UserUpdate {
	uu.mutation.AddFollowerIDs(ids...)
	return uu
}

// AddFollowers adds the "followers" edges to the User entity.
func (uu *UserUpdate) AddFollowers(u ...*User) *UserUpdate {
	ids := make([]uuid.UUID, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return uu.AddFollowerIDs(ids...)
}

// AddFollowingIDs adds the "following" edge to the User entity by IDs.
func (uu *UserUpdate) AddFollowingIDs(ids ...uuid.UUID) *UserUpdate {
	uu.mutation.AddFollowingIDs(ids...)
	return uu
}

// AddFollowing adds the "following" edges to the User entity.
func (uu *UserUpdate) AddFollowing(u ...*User) *UserUpdate {
	ids := make([]uuid.UUID, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return uu.AddFollowingIDs(ids...)
}

// Mutation returns the UserMutation object of the builder.
func (uu *UserUpdate) Mutation() *UserMutation {
	return uu.mutation
}

// ClearFollowers clears all "followers" edges to the User entity.
func (uu *UserUpdate) ClearFollowers() *UserUpdate {
	uu.mutation.ClearFollowers()
	return uu
}

// RemoveFollowerIDs removes the "followers" edge to User entities by IDs.
func (uu *UserUpdate) RemoveFollowerIDs(ids ...uuid.UUID) *UserUpdate {
	uu.mutation.RemoveFollowerIDs(ids...)
	return uu
}

// RemoveFollowers removes "followers" edges to User entities.
func (uu *UserUpdate) RemoveFollowers(u ...*User) *UserUpdate {
	ids := make([]uuid.UUID, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return uu.RemoveFollowerIDs(ids...)
}

// ClearFollowing clears all "following" edges to the User entity.
func (uu *UserUpdate) ClearFollowing() *UserUpdate {
	uu.mutation.ClearFollowing()
	return uu
}

// RemoveFollowingIDs removes the "following" edge to User entities by IDs.
func (uu *UserUpdate) RemoveFollowingIDs(ids ...uuid.UUID) *UserUpdate {
	uu.mutation.RemoveFollowingIDs(ids...)
	return uu
}

// RemoveFollowing removes "following" edges to User entities.
func (uu *UserUpdate) RemoveFollowing(u ...*User) *UserUpdate {
	ids := make([]uuid.UUID, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return uu.RemoveFollowingIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (uu *UserUpdate) Save(ctx context.Context) (int, error) {
	uu.defaults()
	return withHooks(ctx, uu.sqlSave, uu.mutation, uu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (uu *UserUpdate) SaveX(ctx context.Context) int {
	affected, err := uu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (uu *UserUpdate) Exec(ctx context.Context) error {
	_, err := uu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uu *UserUpdate) ExecX(ctx context.Context) {
	if err := uu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (uu *UserUpdate) defaults() {
	if _, ok := uu.mutation.UpdatedAt(); !ok {
		v := user.UpdateDefaultUpdatedAt()
		uu.mutation.SetUpdatedAt(v)
	}
}

func (uu *UserUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(user.Table, user.Columns, sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID))
	if ps := uu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := uu.mutation.UpdatedAt(); ok {
		_spec.SetField(user.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := uu.mutation.FirstName(); ok {
		_spec.SetField(user.FieldFirstName, field.TypeString, value)
	}
	if value, ok := uu.mutation.LastName(); ok {
		_spec.SetField(user.FieldLastName, field.TypeString, value)
	}
	if value, ok := uu.mutation.EmailAddress(); ok {
		_spec.SetField(user.FieldEmailAddress, field.TypeString, value)
	}
	if value, ok := uu.mutation.UserName(); ok {
		_spec.SetField(user.FieldUserName, field.TypeString, value)
	}
	if value, ok := uu.mutation.Bio(); ok {
		_spec.SetField(user.FieldBio, field.TypeString, value)
	}
	if uu.mutation.BioCleared() {
		_spec.ClearField(user.FieldBio, field.TypeString)
	}
	if uu.mutation.FollowersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   user.FollowersTable,
			Columns: user.FollowersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.RemovedFollowersIDs(); len(nodes) > 0 && !uu.mutation.FollowersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   user.FollowersTable,
			Columns: user.FollowersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.FollowersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   user.FollowersTable,
			Columns: user.FollowersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if uu.mutation.FollowingCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   user.FollowingTable,
			Columns: user.FollowingPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.RemovedFollowingIDs(); len(nodes) > 0 && !uu.mutation.FollowingCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   user.FollowingTable,
			Columns: user.FollowingPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.FollowingIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   user.FollowingTable,
			Columns: user.FollowingPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, uu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{user.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	uu.mutation.done = true
	return n, nil
}

// UserUpdateOne is the builder for updating a single User entity.
type UserUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *UserMutation
}

// SetUpdatedAt sets the "updated_at" field.
func (uuo *UserUpdateOne) SetUpdatedAt(t time.Time) *UserUpdateOne {
	uuo.mutation.SetUpdatedAt(t)
	return uuo
}

// SetFirstName sets the "first_name" field.
func (uuo *UserUpdateOne) SetFirstName(s string) *UserUpdateOne {
	uuo.mutation.SetFirstName(s)
	return uuo
}

// SetLastName sets the "last_name" field.
func (uuo *UserUpdateOne) SetLastName(s string) *UserUpdateOne {
	uuo.mutation.SetLastName(s)
	return uuo
}

// SetEmailAddress sets the "email_address" field.
func (uuo *UserUpdateOne) SetEmailAddress(s string) *UserUpdateOne {
	uuo.mutation.SetEmailAddress(s)
	return uuo
}

// SetUserName sets the "user_name" field.
func (uuo *UserUpdateOne) SetUserName(s string) *UserUpdateOne {
	uuo.mutation.SetUserName(s)
	return uuo
}

// SetBio sets the "bio" field.
func (uuo *UserUpdateOne) SetBio(s string) *UserUpdateOne {
	uuo.mutation.SetBio(s)
	return uuo
}

// SetNillableBio sets the "bio" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableBio(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetBio(*s)
	}
	return uuo
}

// ClearBio clears the value of the "bio" field.
func (uuo *UserUpdateOne) ClearBio() *UserUpdateOne {
	uuo.mutation.ClearBio()
	return uuo
}

// AddFollowerIDs adds the "followers" edge to the User entity by IDs.
func (uuo *UserUpdateOne) AddFollowerIDs(ids ...uuid.UUID) *UserUpdateOne {
	uuo.mutation.AddFollowerIDs(ids...)
	return uuo
}

// AddFollowers adds the "followers" edges to the User entity.
func (uuo *UserUpdateOne) AddFollowers(u ...*User) *UserUpdateOne {
	ids := make([]uuid.UUID, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return uuo.AddFollowerIDs(ids...)
}

// AddFollowingIDs adds the "following" edge to the User entity by IDs.
func (uuo *UserUpdateOne) AddFollowingIDs(ids ...uuid.UUID) *UserUpdateOne {
	uuo.mutation.AddFollowingIDs(ids...)
	return uuo
}

// AddFollowing adds the "following" edges to the User entity.
func (uuo *UserUpdateOne) AddFollowing(u ...*User) *UserUpdateOne {
	ids := make([]uuid.UUID, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return uuo.AddFollowingIDs(ids...)
}

// Mutation returns the UserMutation object of the builder.
func (uuo *UserUpdateOne) Mutation() *UserMutation {
	return uuo.mutation
}

// ClearFollowers clears all "followers" edges to the User entity.
func (uuo *UserUpdateOne) ClearFollowers() *UserUpdateOne {
	uuo.mutation.ClearFollowers()
	return uuo
}

// RemoveFollowerIDs removes the "followers" edge to User entities by IDs.
func (uuo *UserUpdateOne) RemoveFollowerIDs(ids ...uuid.UUID) *UserUpdateOne {
	uuo.mutation.RemoveFollowerIDs(ids...)
	return uuo
}

// RemoveFollowers removes "followers" edges to User entities.
func (uuo *UserUpdateOne) RemoveFollowers(u ...*User) *UserUpdateOne {
	ids := make([]uuid.UUID, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return uuo.RemoveFollowerIDs(ids...)
}

// ClearFollowing clears all "following" edges to the User entity.
func (uuo *UserUpdateOne) ClearFollowing() *UserUpdateOne {
	uuo.mutation.ClearFollowing()
	return uuo
}

// RemoveFollowingIDs removes the "following" edge to User entities by IDs.
func (uuo *UserUpdateOne) RemoveFollowingIDs(ids ...uuid.UUID) *UserUpdateOne {
	uuo.mutation.RemoveFollowingIDs(ids...)
	return uuo
}

// RemoveFollowing removes "following" edges to User entities.
func (uuo *UserUpdateOne) RemoveFollowing(u ...*User) *UserUpdateOne {
	ids := make([]uuid.UUID, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return uuo.RemoveFollowingIDs(ids...)
}

// Where appends a list predicates to the UserUpdate builder.
func (uuo *UserUpdateOne) Where(ps ...predicate.User) *UserUpdateOne {
	uuo.mutation.Where(ps...)
	return uuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (uuo *UserUpdateOne) Select(field string, fields ...string) *UserUpdateOne {
	uuo.fields = append([]string{field}, fields...)
	return uuo
}

// Save executes the query and returns the updated User entity.
func (uuo *UserUpdateOne) Save(ctx context.Context) (*User, error) {
	uuo.defaults()
	return withHooks(ctx, uuo.sqlSave, uuo.mutation, uuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (uuo *UserUpdateOne) SaveX(ctx context.Context) *User {
	node, err := uuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (uuo *UserUpdateOne) Exec(ctx context.Context) error {
	_, err := uuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uuo *UserUpdateOne) ExecX(ctx context.Context) {
	if err := uuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (uuo *UserUpdateOne) defaults() {
	if _, ok := uuo.mutation.UpdatedAt(); !ok {
		v := user.UpdateDefaultUpdatedAt()
		uuo.mutation.SetUpdatedAt(v)
	}
}

func (uuo *UserUpdateOne) sqlSave(ctx context.Context) (_node *User, err error) {
	_spec := sqlgraph.NewUpdateSpec(user.Table, user.Columns, sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID))
	id, ok := uuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`internal: missing "User.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := uuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, user.FieldID)
		for _, f := range fields {
			if !user.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("internal: invalid field %q for query", f)}
			}
			if f != user.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := uuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := uuo.mutation.UpdatedAt(); ok {
		_spec.SetField(user.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := uuo.mutation.FirstName(); ok {
		_spec.SetField(user.FieldFirstName, field.TypeString, value)
	}
	if value, ok := uuo.mutation.LastName(); ok {
		_spec.SetField(user.FieldLastName, field.TypeString, value)
	}
	if value, ok := uuo.mutation.EmailAddress(); ok {
		_spec.SetField(user.FieldEmailAddress, field.TypeString, value)
	}
	if value, ok := uuo.mutation.UserName(); ok {
		_spec.SetField(user.FieldUserName, field.TypeString, value)
	}
	if value, ok := uuo.mutation.Bio(); ok {
		_spec.SetField(user.FieldBio, field.TypeString, value)
	}
	if uuo.mutation.BioCleared() {
		_spec.ClearField(user.FieldBio, field.TypeString)
	}
	if uuo.mutation.FollowersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   user.FollowersTable,
			Columns: user.FollowersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.RemovedFollowersIDs(); len(nodes) > 0 && !uuo.mutation.FollowersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   user.FollowersTable,
			Columns: user.FollowersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.FollowersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   user.FollowersTable,
			Columns: user.FollowersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if uuo.mutation.FollowingCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   user.FollowingTable,
			Columns: user.FollowingPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.RemovedFollowingIDs(); len(nodes) > 0 && !uuo.mutation.FollowingCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   user.FollowingTable,
			Columns: user.FollowingPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.FollowingIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   user.FollowingTable,
			Columns: user.FollowingPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &User{config: uuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, uuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{user.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	uuo.mutation.done = true
	return _node, nil
}
