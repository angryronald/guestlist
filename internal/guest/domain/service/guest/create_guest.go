package guest

import (
	"context"

	"github.com/angryronald/guestlist/internal/guest/domain"
	"github.com/angryronald/guestlist/internal/guest/public"
)

func (s *Service) CreateGuest(ctx context.Context, guest *public.Guest) (*public.Guest, error) {
	guestDomain := &domain.Guest{}
	guestDomain.FromPublicModel(guest)

	guestRepo, err := s.repository.Insert(ctx, guestDomain.ToRepositoryModel())
	if err != nil {
		return nil, err
	}

	guestDomain.FromRepositoryModel(guestRepo)

	return guestDomain.ToPublicModel(), nil
}
