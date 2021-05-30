package public

import (
	"time"

	"github.com/google/uuid"
)

type Guest struct {
	ID                       uuid.UUID  `json:"id"`
	Name                     string     `json:"name"`
	Table                    int        `json:"table"`
	AccompanyingGuests       int        `json:"accompanying_guests"`
	ActualAccompanyingGuests int        `json:"actual_accompanying_guests"`
	TimeArrived              *time.Time `json:"time_arrived"`
	TimeLeaved               *time.Time `json:"time_leaved"`
}

type GuestShown struct {
	Name               string `json:"name"`
	Table              int    `json:"table"`
	AccompanyingGuests int    `json:"accompanying_guests"`
}

type GuestArrivedShown struct {
	Name               string     `json:"name"`
	AccompanyingGuests int        `json:"accompanying_guests"`
	TimeArrived        *time.Time `json:"time_arrived"`
}

type AddGuestRequest struct {
	Name               string `url_param:"name"`
	Table              int    `json:"table"`
	AccompanyingGuests int    `json:"accompanying_guests"`
}

type AddGuestResponse struct {
	Name string `json:"name"`
}

type ListGuestsResponse struct {
	Guests []GuestShown `json:"guests"`
}

type GuestArrivedRequest struct {
	Name               string `url_param:"name"`
	AccompanyingGuests int    `json:"accompanying_guests"`
}

type GuestArrivedResponse struct {
	Name string `json:"name"`
}

type ListGuestsArrivedResponse struct {
	Guests []GuestArrivedShown `json:"guests"`
}

type CountEmptySeats struct {
	SeatsEmpty int `json:"seats_empty"`
}

type GuestLeavesRequest struct {
	Name string `url_param:"name"`
}
