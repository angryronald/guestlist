package application

import (
	"github.com/angryronald/guestlist/internal/guest/application/command"
	"github.com/angryronald/guestlist/internal/guest/application/query"
	"github.com/angryronald/guestlist/internal/guest/domain/service/guest"
)

type Commands struct {
	AddGuest     command.AddGuestCommand
	GuestArrived command.GuestArrivedCommand
	GuestLeaves  command.GuestLeavesCommand
}

type Queries struct {
	CountEmptySeats   query.CountEmptySeatsQuery
	ListArrivedGuests query.ListArrivedGuestsQuery
	ListGuests        query.ListGuestsQuery
}

type Application struct {
	Commands Commands
	Queries  Queries
}

func New(
	service guest.ServiceInterface,
) Application {
	return Application{
		Commands: Commands{
			AddGuest:     command.NewAddGuestCommand(service),
			GuestArrived: command.NewGuestArrivedCommand(service),
			GuestLeaves:  command.NewGuestLeavesCommand(service),
		},
		Queries: Queries{
			CountEmptySeats:   query.NewCountEmptySeatsQuery(service),
			ListArrivedGuests: query.NewListArrivedGuestsQuery(service),
			ListGuests:        query.NewListGuestsQuery(service),
		},
	}
}
