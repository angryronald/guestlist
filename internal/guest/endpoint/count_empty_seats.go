package endpoint

import (
	"context"

	guest "github.com/angryronald/guestlist/internal/guest/application"
	// httpResponse "github.com/angryronald/guestlist/lib/net/http"
	"github.com/go-kit/kit/endpoint"
)

func CountEmptySeats(application guest.Application) endpoint.Endpoint {
	return func(ctx context.Context, req interface{}) (res interface{}, err error) {
		res, err = application.Queries.CountEmptySeats.Execute(ctx)

		// I have provided the specific function to adjust the response structure
		// and adding response time in the metadata
		// return httpResponse.ResponseWithRequestTime(ctx, res, nil), err
		return res, nil
	}
}
