package container

import (
	guest "github.com/angryronald/guestlist/internal/guest/infrastructure/repository"
	guestRepository "github.com/angryronald/guestlist/internal/guest/infrastructure/repository/database"
)

type RepositoryIoC struct {
	guest guest.GuestRepository
}

func NewRepositoryIoC() RepositoryIoC {
	return RepositoryIoC{
		guest: guestRepository.NewGuestPostgres(),
	}
}

func (ioc RepositoryIoC) Guest() guest.GuestRepository {
	return ioc.guest
}
