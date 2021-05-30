package container

import (
	"github.com/angryronald/guestlist/internal/guest/domain/service/guest"
)

type DomainServiceIoC struct {
	guest guest.ServiceInterface
}

func NewDomainServiceIoC(ioc RepositoryIoC) DomainServiceIoC {
	return DomainServiceIoC{
		guest: guest.NewService(ioc.Guest()),
	}
}

func (ioc DomainServiceIoC) Guest() guest.ServiceInterface {
	return ioc.guest
}
