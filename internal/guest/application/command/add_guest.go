package command

import (
	"context"

	constant "github.com/angryronald/guestlist/internal/guest"
	"github.com/angryronald/guestlist/internal/guest/domain/service/guest"
	"github.com/angryronald/guestlist/internal/guest/public"
)

// AddGuestCommand encapsulate process for add guest in Command
type AddGuestCommand struct {
	service guest.ServiceInterface
}

// NewAddGuestCommand build an Command for add guest
func NewAddGuestCommand(
	service guest.ServiceInterface,
) AddGuestCommand {
	return AddGuestCommand{
		service: service,
	}
}

func (c AddGuestCommand) Execute(ctx context.Context, request interface{}) (*public.AddGuestResponse, error) {
	payload := request.(public.AddGuestRequest)
	guestExistsCheck, _ := c.service.GetGuestByName(ctx, payload.Name)
	if guestExistsCheck != nil {
		return nil, constant.ErrAlreadyExist
	}

	insertedGuest, err := c.service.CreateGuest(ctx, &public.Guest{
		Name:               payload.Name,
		Table:              payload.Table,
		AccompanyingGuests: payload.AccompanyingGuests,
	})
	if err != nil {
		return nil, err
	}

	return &public.AddGuestResponse{
		Name: insertedGuest.Name,
	}, nil
}
