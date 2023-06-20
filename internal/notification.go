// Code generated by ent, DO NOT EDIT.

package internal

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
	"github.com/pulkitbhardwaj/matrix/internal/notification"
)

// Notification is the model entity for the Notification schema.
type Notification struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt    time.Time `json:"updated_at,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Notification) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case notification.FieldCreatedAt, notification.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		case notification.FieldID:
			values[i] = new(uuid.UUID)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Notification fields.
func (n *Notification) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case notification.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				n.ID = *value
			}
		case notification.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				n.CreatedAt = value.Time
			}
		case notification.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				n.UpdatedAt = value.Time
			}
		default:
			n.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Notification.
// This includes values selected through modifiers, order, etc.
func (n *Notification) Value(name string) (ent.Value, error) {
	return n.selectValues.Get(name)
}

// Update returns a builder for updating this Notification.
// Note that you need to call Notification.Unwrap() before calling this method if this Notification
// was returned from a transaction, and the transaction was committed or rolled back.
func (n *Notification) Update() *NotificationUpdateOne {
	return NewNotificationClient(n.config).UpdateOne(n)
}

// Unwrap unwraps the Notification entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (n *Notification) Unwrap() *Notification {
	_tx, ok := n.config.driver.(*txDriver)
	if !ok {
		panic("internal: Notification is not a transactional entity")
	}
	n.config.driver = _tx.drv
	return n
}

// String implements the fmt.Stringer.
func (n *Notification) String() string {
	var builder strings.Builder
	builder.WriteString("Notification(")
	builder.WriteString(fmt.Sprintf("id=%v, ", n.ID))
	builder.WriteString("created_at=")
	builder.WriteString(n.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(n.UpdatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Notifications is a parsable slice of Notification.
type Notifications []*Notification