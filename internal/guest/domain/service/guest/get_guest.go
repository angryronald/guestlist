package guest

import (
	"context"

	"github.com/angryronald/guestlist/internal/guest"
	"github.com/angryronald/guestlist/internal/guest/domain"
	"github.com/angryronald/guestlist/internal/guest/public"
	"github.com/google/uuid"
)

func (s *Service) GetGuest(ctx context.Context, guestID uuid.UUID) (*public.Guest, error) {
	guestRepo, err := s.repository.FindByID(ctx, guestID)
	if err != nil {
		return nil, err
	}

	if guestRepo == nil {
		return nil, guest.ErrNotFound
	}

	guestDomain := &domain.Guest{}
	guestDomain.FromRepositoryModel(guestRepo)

	return guestDomain.ToPublicModel(), nil
}
