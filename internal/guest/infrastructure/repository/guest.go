package repository

import (
	"context"
	"time"

	"github.com/google/uuid"
)

// Guest represents the repository guest object
type Guest struct {
	ID                       uuid.UUID  `json:"id" gorm:"primaryKey,not null"`
	Name                     string     `json:"name" gorm:"index:idx_guest_phone,unique,not null"`
	Table                    int        `json:"table"`
	AccompanyingGuests       int        `json:"accompanying_guests"`
	ActualAccompanyingGuests int        `json:"actual_accompanying_guests"`
	TimeArrived              *time.Time `json:"time_arrived"`
	TimeLeaved               *time.Time `json:"time_leaved"`
	CreatedAt                time.Time  `json:"created_at" gorm:"not null"`
	UpdatedAt                time.Time  `json:"updated_at" gorm:"not null"`
	DeletedAt                *time.Time `json:"deleted_at"`
}

// GuestRepository represents the guest repository interface
type GuestRepository interface {
	FindAll(ctx context.Context, isArrivedOnly bool) ([]*Guest, error)
	FindByID(ctx context.Context, guestID uuid.UUID) (*Guest, error)
	FindByName(ctx context.Context, name string) (*Guest, error)
	Insert(ctx context.Context, guest *Guest) (*Guest, error)
	Update(ctx context.Context, guest *Guest) (*Guest, error)
	Delete(ctx context.Context, guest *Guest) error
}
