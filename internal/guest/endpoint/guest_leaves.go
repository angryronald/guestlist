package endpoint

import (
	"context"
	"net/http"

	"github.com/angryronald/guestlist/config"
	"github.com/angryronald/guestlist/internal/global"
	constant "github.com/angryronald/guestlist/internal/guest"
	guest "github.com/angryronald/guestlist/internal/guest/application"
	"github.com/angryronald/guestlist/internal/guest/public"
	"github.com/angryronald/guestlist/lib/database"
	libError "github.com/angryronald/guestlist/lib/error"
	"github.com/go-kit/kit/endpoint"
)

func GuestLeaves(application guest.Application) endpoint.Endpoint {
	return func(ctx context.Context, req interface{}) (res interface{}, err error) {
		payload := req.(*public.GuestLeavesRequest)

		err = database.RunInTransaction(ctx, config.DB(), func(ctx context.Context) (err error) {
			if payload == nil {
				err = libError.New(err, http.StatusBadRequest, global.ErrInvalidRequest)
			}

			err = application.Commands.GuestLeaves.Execute(ctx, *payload)
			return err
		})
		if err == constant.ErrNotFound {
			err = libError.New(err, http.StatusNotFound, err.Error())
		}

		return nil, err
	}
}
