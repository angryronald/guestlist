package query

import (
	"context"

	"github.com/angryronald/guestlist/internal/guest/domain/service/guest"
	"github.com/angryronald/guestlist/internal/guest/public"
)

// ListGuestsQuery encapsulate process for list guests in Query
type ListGuestsQuery struct {
	service guest.ServiceInterface
}

// NewListGuestsQuery build an Query for list guests
func NewListGuestsQuery(
	service guest.ServiceInterface,
) ListGuestsQuery {
	return ListGuestsQuery{
		service: service,
	}
}

func (q ListGuestsQuery) Execute(ctx context.Context) (*public.ListGuestsResponse, error) {
	guests, err := q.service.ListGuests(ctx, false)
	if err != nil {
		return nil, err
	}

	result := &public.ListGuestsResponse{}
	for _, guest := range guests {
		result.Guests = append(result.Guests, public.GuestShown{
			Name:               guest.Name,
			Table:              guest.Table,
			AccompanyingGuests: guest.AccompanyingGuests,
		})
	}

	return result, nil
}
