package domain

import (
	"time"

	"github.com/gofrs/uuid"
)

const (
	Read   EventStatus = "read"
	Unread EventStatus = "unread"
)

// ID defines model for id.
type ID = uuid.UUID

// UserID defines model for user_uuid.
type UserID = uuid.UUID

// OperationID defines model for operation_id.
type OperationID = uuid.UUID

// SystemID defines model for system_id.
type SystemID = uuid.UUID

// Time defines model for time.
type Time = time.Time

// CreatedAt Timestamp when the resource was created.
type CreatedAt time.Time

// DeletedAt Timestamp when the resource was deleted.
type DeletedAt time.Time

// UpdatedAt Timestamp when the resource was updated.
type UpdatedAt = time.Time

// EventAction defines model for event_action.
type EventAction string

// EventStatus defines model for event_state.
type EventStatus string

// EventSeverity defines model for event_severity.
type EventSeverity string

type (
	// Event defines the model for event_trackers
	Event struct {
		ID        string    `db:"id" json:"id,omitempty"`
		UserID    string    `db:"user_id" json:"user_id,omitempty"`
		SystemID  string    `db:"system_id" json:"system_id,omitempty"`
		CreatedAt CreatedAt `db:"created_at" json:"created_at,omitempty"`
		UpdatedAt UpdatedAt `db:"updated_at" json:"updated_at,omitempty"`
		DeletedAt DeletedAt `db:"deleted_at" json:"deleted_at,omitempty"`
	}
	EventEntity struct {
		ID        string
		UserID    string
		SystemID  string
		CreatedAt CreatedAt
		UpdatedAt UpdatedAt
		DeletedAt DeletedAt
	}
	EventRepository interface {
		FindEventByID(id string) (*EventEntity, error)
		FindAllEvents() ([]*EventEntity, error)
		CreateEvent(eventID *EventEntity) error
		UpdateEvent(eventID *EventEntity) error
		DeleteEvent(eventID *EventEntity) error
	}
	EventService interface {
		FindEventByID(id string) (*Event, error)
		FindEvents() ([]*Event, error)
		CreateEvent(eventID *Event) error
		UpdateEvent(eventID *Event) error
		DeleteEvent(eventID *Event) error
	}
	EventHandler interface {
		HandleEvent(eventID *Event) error
	}
	BulkEventHandler interface {
		FindAllEvents() ([]*Event, error)
		FindByType(eventType string) ([]*Event, error)
		BulkCreate(events []*Event) error
		BulkUpdate(events []*Event) error
		BulkDelete(events []*Event) error
	}
)
