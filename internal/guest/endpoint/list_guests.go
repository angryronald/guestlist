package endpoint

import (
	"context"

	guest "github.com/angryronald/guestlist/internal/guest/application"
	"github.com/go-kit/kit/endpoint"
)

func ListGuests(application guest.Application) endpoint.Endpoint {
	return func(ctx context.Context, req interface{}) (res interface{}, err error) {
		res, err = application.Queries.ListGuests.Execute(ctx)

		return res, err
	}
}
