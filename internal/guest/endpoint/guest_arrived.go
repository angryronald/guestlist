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

func GuestArrived(application guest.Application) endpoint.Endpoint {
	return func(ctx context.Context, req interface{}) (res interface{}, err error) {
		payload := req.(*public.GuestArrivedRequest)

		err = database.RunInTransaction(ctx, config.DB(), func(ctx context.Context) (err error) {
			if payload == nil {
				err = libError.New(err, http.StatusBadRequest, global.ErrInvalidRequest)
			}

			res, err = application.Commands.GuestArrived.Execute(ctx, *payload)
			return err
		})
		if err == constant.ErrInsufficientSpace {
			err = libError.New(err, http.StatusUnprocessableEntity, err.Error())
		}
		if err == constant.ErrAlreadyArrived {
			err = libError.New(err, http.StatusUnprocessableEntity, err.Error())
		}
		if err == constant.ErrNotFound {
			err = libError.New(err, http.StatusNotFound, err.Error())
		}

		return res, err
	}
}
