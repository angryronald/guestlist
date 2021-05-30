package command

import (
	"context"
	"time"

	constant "github.com/angryronald/guestlist/internal/guest"
	"github.com/angryronald/guestlist/internal/guest/domain/service/guest"
	"github.com/angryronald/guestlist/internal/guest/public"
)

// GuestArrivedCommand encapsulate process for guest arrives in Command
type GuestArrivedCommand struct {
	service guest.ServiceInterface
}

// NewGuestArrivedCommand build an Command for guest arrives
func NewGuestArrivedCommand(
	service guest.ServiceInterface,
) GuestArrivedCommand {
	return GuestArrivedCommand{
		service: service,
	}
}

func (c GuestArrivedCommand) Execute(ctx context.Context, request interface{}) (*public.GuestArrivedResponse, error) {
	payload := request.(public.GuestArrivedRequest)
	guest, err := c.service.GetGuestByName(ctx, payload.Name)
	if err != nil {
		return nil, err
	}

	if guest.TimeArrived != nil {
		return nil, constant.ErrAlreadyArrived
	}

	totalAvailableSpace, err := c.service.GetAvailableSpace(ctx)
	if err != nil {
		return nil, err
	}

	isAbleToProceed := true
	if payload.AccompanyingGuests > guest.AccompanyingGuests || totalAvailableSpace < payload.AccompanyingGuests {
		isAbleToProceed = false
	}

	if !isAbleToProceed {
		return nil, constant.ErrInsufficientSpace
	}

	guest.ActualAccompanyingGuests = guest.AccompanyingGuests
	if guest.AccompanyingGuests != payload.AccompanyingGuests {
		guest.ActualAccompanyingGuests = payload.AccompanyingGuests
	}

	now := time.Now().UTC()
	guest.TimeArrived = &now

	guest, err = c.service.UpdateGuest(ctx, guest)
	if err != nil {
		return nil, err
	}

	return &public.GuestArrivedResponse{
		Name: guest.Name,
	}, nil
}
