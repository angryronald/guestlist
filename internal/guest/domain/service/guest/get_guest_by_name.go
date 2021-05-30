package guest

import (
	"context"

	"github.com/angryronald/guestlist/internal/guest"
	"github.com/angryronald/guestlist/internal/guest/domain"
	"github.com/angryronald/guestlist/internal/guest/public"
)

func (s *Service) GetGuestByName(ctx context.Context, name string) (*public.Guest, error) {
	guestRepo, err := s.repository.FindByName(ctx, name)
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
