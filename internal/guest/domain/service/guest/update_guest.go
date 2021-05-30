package guest

import (
	"context"

	"github.com/angryronald/guestlist/internal/guest/domain"
	"github.com/angryronald/guestlist/internal/guest/public"
)

func (s *Service) UpdateGuest(ctx context.Context, guest *public.Guest) (*public.Guest, error) {
	guestDomain := &domain.Guest{}
	guestDomain.FromPublicModel(guest)

	guestRepo, err := s.repository.Update(ctx, guestDomain.ToRepositoryModel())
	if err != nil {
		return nil, err
	}

	guestDomain.FromRepositoryModel(guestRepo)

	return guestDomain.ToPublicModel(), nil
}
