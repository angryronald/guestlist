package query

import (
	"context"

	"github.com/angryronald/guestlist/internal/guest/domain/service/guest"
	"github.com/angryronald/guestlist/internal/guest/public"
)

// ListArrivedGuestsQuery encapsulate process for list arrived guests in Query
type ListArrivedGuestsQuery struct {
	service guest.ServiceInterface
}

// NewListArrivedGuestsQuery build an Query for list arrived guests
func NewListArrivedGuestsQuery(
	service guest.ServiceInterface,
) ListArrivedGuestsQuery {
	return ListArrivedGuestsQuery{
		service: service,
	}
}

func (q ListArrivedGuestsQuery) Execute(ctx context.Context) (*public.ListGuestsArrivedResponse, error) {
	guests, err := q.service.ListGuests(ctx, true)
	if err != nil {
		return nil, err
	}

	result := &public.ListGuestsArrivedResponse{}
	for _, guest := range guests {
		result.Guests = append(result.Guests, public.GuestArrivedShown{
			Name:               guest.Name,
			AccompanyingGuests: guest.ActualAccompanyingGuests,
			TimeArrived:        guest.TimeArrived,
		})
	}

	return result, nil
}
