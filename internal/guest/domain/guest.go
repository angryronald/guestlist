package domain

import (
	"time"

	"github.com/angryronald/guestlist/internal/guest/infrastructure/repository"
	"github.com/angryronald/guestlist/internal/guest/public"
	"github.com/angryronald/guestlist/lib/encoding"
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

func (g *Guest) FromPublicModel(guestPublic interface{}) {
	_ = encoding.TransformObject(guestPublic, g)
}

func (g *Guest) ToPublicModel() *public.Guest {
	guestPublic := &public.Guest{}
	_ = encoding.TransformObject(g, guestPublic)
	return guestPublic
}

func (g *Guest) FromRepositoryModel(guestRepo interface{}) {
	_ = encoding.TransformObject(guestRepo, g)
}

func (g *Guest) ToRepositoryModel() *repository.Guest {
	guestRepo := &repository.Guest{}
	_ = encoding.TransformObject(g, guestRepo)
	return guestRepo
}
