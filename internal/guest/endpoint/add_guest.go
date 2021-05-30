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

func AddGuest(application guest.Application) endpoint.Endpoint {
	return func(ctx context.Context, req interface{}) (res interface{}, err error) {
		payload := req.(*public.AddGuestRequest)

		err = database.RunInTransaction(ctx, config.DB(), func(ctx context.Context) (err error) {
			if payload == nil {
				err = libError.New(err, http.StatusBadRequest, global.ErrInvalidRequest)
			}

			res, err = application.Commands.AddGuest.Execute(ctx, *payload)
			return err
		})
		if err == constant.ErrAlreadyExist {
			err = libError.New(err, http.StatusUnprocessableEntity, err.Error())
		}

		return res, err
	}
}
