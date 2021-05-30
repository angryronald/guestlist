package guest

import (
	"context"

	"github.com/angryronald/guestlist/internal/guest/infrastructure/repository"
	"github.com/angryronald/guestlist/internal/guest/public"
	"github.com/google/uuid"
)

// ServiceInterface represents the guest service interface
type ServiceInterface interface {
	ListGuests(ctx context.Context, isArrivedOnly bool) ([]*public.Guest, error)
	GetGuest(ctx context.Context, guestID uuid.UUID) (*public.Guest, error)
	GetGuestByName(ctx context.Context, name string) (*public.Guest, error)
	CreateGuest(ctx context.Context, guest *public.Guest) (*public.Guest, error)
	UpdateGuest(ctx context.Context, guest *public.Guest) (*public.Guest, error)
	DeleteGuest(ctx context.Context, guestID uuid.UUID) error
	GetAvailableSpace(ctx context.Context) (int, error)
}

// Service is the domain logic implementation of guest service interface
type Service struct {
	repository repository.GuestRepository
}

// NewService creates a new guest domain service
func NewService(
	repository repository.GuestRepository,
) ServiceInterface {
	return &Service{
		repository: repository,
	}
}
