package command

import (
	"context"
	"time"

	"github.com/angryronald/guestlist/internal/guest/domain/service/guest"
	"github.com/angryronald/guestlist/internal/guest/public"
)

// GuestLeavesCommand encapsulate process for guest leaves in Command
type GuestLeavesCommand struct {
	service guest.ServiceInterface
}

// NewGuestArrivedCommand build an Command for guest leaves
func NewGuestLeavesCommand(
	service guest.ServiceInterface,
) GuestLeavesCommand {
	return GuestLeavesCommand{
		service: service,
	}
}

func (c GuestLeavesCommand) Execute(ctx context.Context, request interface{}) error {
	payload := request.(public.GuestLeavesRequest)
	guest, err := c.service.GetGuestByName(ctx, payload.Name)
	if err != nil {
		return err
	}

	now := time.Now().UTC()
	guest.TimeLeaved = &now

	guest, err = c.service.UpdateGuest(ctx, guest)
	if err != nil {
		return err
	}

	err = c.service.DeleteGuest(ctx, guest.ID)
	if err != nil {
		return err
	}

	return nil
}
