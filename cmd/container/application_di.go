package container

import (
	guest "github.com/angryronald/guestlist/internal/guest/application"
)

type ApplicationServiceIoC struct {
	guest guest.Application
}

func NewApplicationServiceIoC(dsIoc DomainServiceIoC, rIoc RepositoryIoC) ApplicationServiceIoC {
	return ApplicationServiceIoC{
		guest: guest.New(dsIoc.Guest()),
	}
}

func (ioc ApplicationServiceIoC) Guest() guest.Application {
	return ioc.guest
}
