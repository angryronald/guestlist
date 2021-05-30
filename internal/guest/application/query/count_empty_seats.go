package query

import (
	"context"

	"github.com/angryronald/guestlist/internal/guest/domain/service/guest"
	"github.com/angryronald/guestlist/internal/guest/public"
)

// CountEmptySeatsQuery encapsulate process for count empty seats in Query
type CountEmptySeatsQuery struct {
	service guest.ServiceInterface
}

// NewCountEmptySeatsQuery build an Query for count empty seats
func NewCountEmptySeatsQuery(
	service guest.ServiceInterface,
) CountEmptySeatsQuery {
	return CountEmptySeatsQuery{
		service: service,
	}
}

func (q CountEmptySeatsQuery) Execute(ctx context.Context) (*public.CountEmptySeats, error) {
	emptySeats, err := q.service.GetAvailableSpace(ctx)
	if err != nil {
		return nil, err
	}

	return &public.CountEmptySeats{
		SeatsEmpty: emptySeats,
	}, nil
}
