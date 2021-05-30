package guest

import (
	"context"

	"github.com/angryronald/guestlist/internal/guest"
	"github.com/google/uuid"
)

func (s *Service) DeleteGuest(ctx context.Context, guestID uuid.UUID) error {
	guestRepo, err := s.repository.FindByID(ctx, guestID)
	if err != nil {
		return nil
	}

	if guestRepo == nil {
		return guest.ErrNotFound
	}

	return s.repository.Delete(ctx, guestRepo)
}
