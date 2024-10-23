// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/chainloop-dev/chainloop/app/controlplane/pkg/data/ent/apitoken"
	"github.com/chainloop-dev/chainloop/app/controlplane/pkg/data/ent/organization"
	"github.com/google/uuid"
)

// APITokenCreate is the builder for creating a APIToken entity.
type APITokenCreate struct {
	config
	mutation *APITokenMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetName sets the "name" field.
func (atc *APITokenCreate) SetName(s string) *APITokenCreate {
	atc.mutation.SetName(s)
	return atc
}

// SetDescription sets the "description" field.
func (atc *APITokenCreate) SetDescription(s string) *APITokenCreate {
	atc.mutation.SetDescription(s)
	return atc
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (atc *APITokenCreate) SetNillableDescription(s *string) *APITokenCreate {
	if s != nil {
		atc.SetDescription(*s)
	}
	return atc
}

// SetCreatedAt sets the "created_at" field.
func (atc *APITokenCreate) SetCreatedAt(t time.Time) *APITokenCreate {
	atc.mutation.SetCreatedAt(t)
	return atc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (atc *APITokenCreate) SetNillableCreatedAt(t *time.Time) *APITokenCreate {
	if t != nil {
		atc.SetCreatedAt(*t)
	}
	return atc
}

// SetExpiresAt sets the "expires_at" field.
func (atc *APITokenCreate) SetExpiresAt(t time.Time) *APITokenCreate {
	atc.mutation.SetExpiresAt(t)
	return atc
}

// SetNillableExpiresAt sets the "expires_at" field if the given value is not nil.
func (atc *APITokenCreate) SetNillableExpiresAt(t *time.Time) *APITokenCreate {
	if t != nil {
		atc.SetExpiresAt(*t)
	}
	return atc
}

// SetRevokedAt sets the "revoked_at" field.
func (atc *APITokenCreate) SetRevokedAt(t time.Time) *APITokenCreate {
	atc.mutation.SetRevokedAt(t)
	return atc
}

// SetNillableRevokedAt sets the "revoked_at" field if the given value is not nil.
func (atc *APITokenCreate) SetNillableRevokedAt(t *time.Time) *APITokenCreate {
	if t != nil {
		atc.SetRevokedAt(*t)
	}
	return atc
}

// SetOrganizationID sets the "organization_id" field.
func (atc *APITokenCreate) SetOrganizationID(u uuid.UUID) *APITokenCreate {
	atc.mutation.SetOrganizationID(u)
	return atc
}

// SetID sets the "id" field.
func (atc *APITokenCreate) SetID(u uuid.UUID) *APITokenCreate {
	atc.mutation.SetID(u)
	return atc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (atc *APITokenCreate) SetNillableID(u *uuid.UUID) *APITokenCreate {
	if u != nil {
		atc.SetID(*u)
	}
	return atc
}

// SetOrganization sets the "organization" edge to the Organization entity.
func (atc *APITokenCreate) SetOrganization(o *Organization) *APITokenCreate {
	return atc.SetOrganizationID(o.ID)
}

// Mutation returns the APITokenMutation object of the builder.
func (atc *APITokenCreate) Mutation() *APITokenMutation {
	return atc.mutation
}

// Save creates the APIToken in the database.
func (atc *APITokenCreate) Save(ctx context.Context) (*APIToken, error) {
	atc.defaults()
	return withHooks(ctx, atc.sqlSave, atc.mutation, atc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (atc *APITokenCreate) SaveX(ctx context.Context) *APIToken {
	v, err := atc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (atc *APITokenCreate) Exec(ctx context.Context) error {
	_, err := atc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (atc *APITokenCreate) ExecX(ctx context.Context) {
	if err := atc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (atc *APITokenCreate) defaults() {
	if _, ok := atc.mutation.CreatedAt(); !ok {
		v := apitoken.DefaultCreatedAt()
		atc.mutation.SetCreatedAt(v)
	}
	if _, ok := atc.mutation.ID(); !ok {
		v := apitoken.DefaultID()
		atc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (atc *APITokenCreate) check() error {
	if _, ok := atc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "APIToken.name"`)}
	}
	if _, ok := atc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "APIToken.created_at"`)}
	}
	if _, ok := atc.mutation.OrganizationID(); !ok {
		return &ValidationError{Name: "organization_id", err: errors.New(`ent: missing required field "APIToken.organization_id"`)}
	}
	if len(atc.mutation.OrganizationIDs()) == 0 {
		return &ValidationError{Name: "organization", err: errors.New(`ent: missing required edge "APIToken.organization"`)}
	}
	return nil
}

func (atc *APITokenCreate) sqlSave(ctx context.Context) (*APIToken, error) {
	if err := atc.check(); err != nil {
		return nil, err
	}
	_node, _spec := atc.createSpec()
	if err := sqlgraph.CreateNode(ctx, atc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(*uuid.UUID); ok {
			_node.ID = *id
		} else if err := _node.ID.Scan(_spec.ID.Value); err != nil {
			return nil, err
		}
	}
	atc.mutation.id = &_node.ID
	atc.mutation.done = true
	return _node, nil
}

func (atc *APITokenCreate) createSpec() (*APIToken, *sqlgraph.CreateSpec) {
	var (
		_node = &APIToken{config: atc.config}
		_spec = sqlgraph.NewCreateSpec(apitoken.Table, sqlgraph.NewFieldSpec(apitoken.FieldID, field.TypeUUID))
	)
	_spec.OnConflict = atc.conflict
	if id, ok := atc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := atc.mutation.Name(); ok {
		_spec.SetField(apitoken.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := atc.mutation.Description(); ok {
		_spec.SetField(apitoken.FieldDescription, field.TypeString, value)
		_node.Description = value
	}
	if value, ok := atc.mutation.CreatedAt(); ok {
		_spec.SetField(apitoken.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := atc.mutation.ExpiresAt(); ok {
		_spec.SetField(apitoken.FieldExpiresAt, field.TypeTime, value)
		_node.ExpiresAt = value
	}
	if value, ok := atc.mutation.RevokedAt(); ok {
		_spec.SetField(apitoken.FieldRevokedAt, field.TypeTime, value)
		_node.RevokedAt = value
	}
	if nodes := atc.mutation.OrganizationIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   apitoken.OrganizationTable,
			Columns: []string{apitoken.OrganizationColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(organization.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.OrganizationID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.APIToken.Create().
//		SetName(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.APITokenUpsert) {
//			SetName(v+v).
//		}).
//		Exec(ctx)
func (atc *APITokenCreate) OnConflict(opts ...sql.ConflictOption) *APITokenUpsertOne {
	atc.conflict = opts
	return &APITokenUpsertOne{
		create: atc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.APIToken.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (atc *APITokenCreate) OnConflictColumns(columns ...string) *APITokenUpsertOne {
	atc.conflict = append(atc.conflict, sql.ConflictColumns(columns...))
	return &APITokenUpsertOne{
		create: atc,
	}
}

type (
	// APITokenUpsertOne is the builder for "upsert"-ing
	//  one APIToken node.
	APITokenUpsertOne struct {
		create *APITokenCreate
	}

	// APITokenUpsert is the "OnConflict" setter.
	APITokenUpsert struct {
		*sql.UpdateSet
	}
)

// SetDescription sets the "description" field.
func (u *APITokenUpsert) SetDescription(v string) *APITokenUpsert {
	u.Set(apitoken.FieldDescription, v)
	return u
}

// UpdateDescription sets the "description" field to the value that was provided on create.
func (u *APITokenUpsert) UpdateDescription() *APITokenUpsert {
	u.SetExcluded(apitoken.FieldDescription)
	return u
}

// ClearDescription clears the value of the "description" field.
func (u *APITokenUpsert) ClearDescription() *APITokenUpsert {
	u.SetNull(apitoken.FieldDescription)
	return u
}

// SetRevokedAt sets the "revoked_at" field.
func (u *APITokenUpsert) SetRevokedAt(v time.Time) *APITokenUpsert {
	u.Set(apitoken.FieldRevokedAt, v)
	return u
}

// UpdateRevokedAt sets the "revoked_at" field to the value that was provided on create.
func (u *APITokenUpsert) UpdateRevokedAt() *APITokenUpsert {
	u.SetExcluded(apitoken.FieldRevokedAt)
	return u
}

// ClearRevokedAt clears the value of the "revoked_at" field.
func (u *APITokenUpsert) ClearRevokedAt() *APITokenUpsert {
	u.SetNull(apitoken.FieldRevokedAt)
	return u
}

// SetOrganizationID sets the "organization_id" field.
func (u *APITokenUpsert) SetOrganizationID(v uuid.UUID) *APITokenUpsert {
	u.Set(apitoken.FieldOrganizationID, v)
	return u
}

// UpdateOrganizationID sets the "organization_id" field to the value that was provided on create.
func (u *APITokenUpsert) UpdateOrganizationID() *APITokenUpsert {
	u.SetExcluded(apitoken.FieldOrganizationID)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.APIToken.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(apitoken.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *APITokenUpsertOne) UpdateNewValues() *APITokenUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(apitoken.FieldID)
		}
		if _, exists := u.create.mutation.Name(); exists {
			s.SetIgnore(apitoken.FieldName)
		}
		if _, exists := u.create.mutation.CreatedAt(); exists {
			s.SetIgnore(apitoken.FieldCreatedAt)
		}
		if _, exists := u.create.mutation.ExpiresAt(); exists {
			s.SetIgnore(apitoken.FieldExpiresAt)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.APIToken.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *APITokenUpsertOne) Ignore() *APITokenUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *APITokenUpsertOne) DoNothing() *APITokenUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the APITokenCreate.OnConflict
// documentation for more info.
func (u *APITokenUpsertOne) Update(set func(*APITokenUpsert)) *APITokenUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&APITokenUpsert{UpdateSet: update})
	}))
	return u
}

// SetDescription sets the "description" field.
func (u *APITokenUpsertOne) SetDescription(v string) *APITokenUpsertOne {
	return u.Update(func(s *APITokenUpsert) {
		s.SetDescription(v)
	})
}

// UpdateDescription sets the "description" field to the value that was provided on create.
func (u *APITokenUpsertOne) UpdateDescription() *APITokenUpsertOne {
	return u.Update(func(s *APITokenUpsert) {
		s.UpdateDescription()
	})
}

// ClearDescription clears the value of the "description" field.
func (u *APITokenUpsertOne) ClearDescription() *APITokenUpsertOne {
	return u.Update(func(s *APITokenUpsert) {
		s.ClearDescription()
	})
}

// SetRevokedAt sets the "revoked_at" field.
func (u *APITokenUpsertOne) SetRevokedAt(v time.Time) *APITokenUpsertOne {
	return u.Update(func(s *APITokenUpsert) {
		s.SetRevokedAt(v)
	})
}

// UpdateRevokedAt sets the "revoked_at" field to the value that was provided on create.
func (u *APITokenUpsertOne) UpdateRevokedAt() *APITokenUpsertOne {
	return u.Update(func(s *APITokenUpsert) {
		s.UpdateRevokedAt()
	})
}

// ClearRevokedAt clears the value of the "revoked_at" field.
func (u *APITokenUpsertOne) ClearRevokedAt() *APITokenUpsertOne {
	return u.Update(func(s *APITokenUpsert) {
		s.ClearRevokedAt()
	})
}

// SetOrganizationID sets the "organization_id" field.
func (u *APITokenUpsertOne) SetOrganizationID(v uuid.UUID) *APITokenUpsertOne {
	return u.Update(func(s *APITokenUpsert) {
		s.SetOrganizationID(v)
	})
}

// UpdateOrganizationID sets the "organization_id" field to the value that was provided on create.
func (u *APITokenUpsertOne) UpdateOrganizationID() *APITokenUpsertOne {
	return u.Update(func(s *APITokenUpsert) {
		s.UpdateOrganizationID()
	})
}

// Exec executes the query.
func (u *APITokenUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for APITokenCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *APITokenUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *APITokenUpsertOne) ID(ctx context.Context) (id uuid.UUID, err error) {
	if u.create.driver.Dialect() == dialect.MySQL {
		// In case of "ON CONFLICT", there is no way to get back non-numeric ID
		// fields from the database since MySQL does not support the RETURNING clause.
		return id, errors.New("ent: APITokenUpsertOne.ID is not supported by MySQL driver. Use APITokenUpsertOne.Exec instead")
	}
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *APITokenUpsertOne) IDX(ctx context.Context) uuid.UUID {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// APITokenCreateBulk is the builder for creating many APIToken entities in bulk.
type APITokenCreateBulk struct {
	config
	err      error
	builders []*APITokenCreate
	conflict []sql.ConflictOption
}

// Save creates the APIToken entities in the database.
func (atcb *APITokenCreateBulk) Save(ctx context.Context) ([]*APIToken, error) {
	if atcb.err != nil {
		return nil, atcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(atcb.builders))
	nodes := make([]*APIToken, len(atcb.builders))
	mutators := make([]Mutator, len(atcb.builders))
	for i := range atcb.builders {
		func(i int, root context.Context) {
			builder := atcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*APITokenMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, atcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = atcb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, atcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, atcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (atcb *APITokenCreateBulk) SaveX(ctx context.Context) []*APIToken {
	v, err := atcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (atcb *APITokenCreateBulk) Exec(ctx context.Context) error {
	_, err := atcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (atcb *APITokenCreateBulk) ExecX(ctx context.Context) {
	if err := atcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.APIToken.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.APITokenUpsert) {
//			SetName(v+v).
//		}).
//		Exec(ctx)
func (atcb *APITokenCreateBulk) OnConflict(opts ...sql.ConflictOption) *APITokenUpsertBulk {
	atcb.conflict = opts
	return &APITokenUpsertBulk{
		create: atcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.APIToken.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (atcb *APITokenCreateBulk) OnConflictColumns(columns ...string) *APITokenUpsertBulk {
	atcb.conflict = append(atcb.conflict, sql.ConflictColumns(columns...))
	return &APITokenUpsertBulk{
		create: atcb,
	}
}

// APITokenUpsertBulk is the builder for "upsert"-ing
// a bulk of APIToken nodes.
type APITokenUpsertBulk struct {
	create *APITokenCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.APIToken.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(apitoken.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *APITokenUpsertBulk) UpdateNewValues() *APITokenUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(apitoken.FieldID)
			}
			if _, exists := b.mutation.Name(); exists {
				s.SetIgnore(apitoken.FieldName)
			}
			if _, exists := b.mutation.CreatedAt(); exists {
				s.SetIgnore(apitoken.FieldCreatedAt)
			}
			if _, exists := b.mutation.ExpiresAt(); exists {
				s.SetIgnore(apitoken.FieldExpiresAt)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.APIToken.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *APITokenUpsertBulk) Ignore() *APITokenUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *APITokenUpsertBulk) DoNothing() *APITokenUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the APITokenCreateBulk.OnConflict
// documentation for more info.
func (u *APITokenUpsertBulk) Update(set func(*APITokenUpsert)) *APITokenUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&APITokenUpsert{UpdateSet: update})
	}))
	return u
}

// SetDescription sets the "description" field.
func (u *APITokenUpsertBulk) SetDescription(v string) *APITokenUpsertBulk {
	return u.Update(func(s *APITokenUpsert) {
		s.SetDescription(v)
	})
}

// UpdateDescription sets the "description" field to the value that was provided on create.
func (u *APITokenUpsertBulk) UpdateDescription() *APITokenUpsertBulk {
	return u.Update(func(s *APITokenUpsert) {
		s.UpdateDescription()
	})
}

// ClearDescription clears the value of the "description" field.
func (u *APITokenUpsertBulk) ClearDescription() *APITokenUpsertBulk {
	return u.Update(func(s *APITokenUpsert) {
		s.ClearDescription()
	})
}

// SetRevokedAt sets the "revoked_at" field.
func (u *APITokenUpsertBulk) SetRevokedAt(v time.Time) *APITokenUpsertBulk {
	return u.Update(func(s *APITokenUpsert) {
		s.SetRevokedAt(v)
	})
}

// UpdateRevokedAt sets the "revoked_at" field to the value that was provided on create.
func (u *APITokenUpsertBulk) UpdateRevokedAt() *APITokenUpsertBulk {
	return u.Update(func(s *APITokenUpsert) {
		s.UpdateRevokedAt()
	})
}

// ClearRevokedAt clears the value of the "revoked_at" field.
func (u *APITokenUpsertBulk) ClearRevokedAt() *APITokenUpsertBulk {
	return u.Update(func(s *APITokenUpsert) {
		s.ClearRevokedAt()
	})
}

// SetOrganizationID sets the "organization_id" field.
func (u *APITokenUpsertBulk) SetOrganizationID(v uuid.UUID) *APITokenUpsertBulk {
	return u.Update(func(s *APITokenUpsert) {
		s.SetOrganizationID(v)
	})
}

// UpdateOrganizationID sets the "organization_id" field to the value that was provided on create.
func (u *APITokenUpsertBulk) UpdateOrganizationID() *APITokenUpsertBulk {
	return u.Update(func(s *APITokenUpsert) {
		s.UpdateOrganizationID()
	})
}

// Exec executes the query.
func (u *APITokenUpsertBulk) Exec(ctx context.Context) error {
	if u.create.err != nil {
		return u.create.err
	}
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the APITokenCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for APITokenCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *APITokenUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
