package guest

import (
	"context"

	"github.com/angryronald/guestlist/internal/guest/domain"
	"github.com/angryronald/guestlist/internal/guest/public"
)

func (s *Service) ListGuests(ctx context.Context, isArrivedOnly bool) ([]*public.Guest, error) {
	guestsRepo, err := s.repository.FindAll(ctx, isArrivedOnly)
	if err != nil {
		return nil, err
	}

	result := []*public.Guest{}
	for _, guestRepo := range guestsRepo {
		guestDomain := &domain.Guest{}
		guestDomain.FromRepositoryModel(guestRepo)
		result = append(result, guestDomain.ToPublicModel())
	}

	return result, nil
}
