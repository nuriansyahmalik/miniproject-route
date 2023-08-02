package menu

import (
	"github.com/evermos/boilerplate-go/configs"
	"github.com/gofrs/uuid"
)

type MenuService interface {
	ResolveByID(id uuid.UUID) (menu Menu, err error)
}

type MenuServiceImpl struct {
	MenuRepository MenuRepository
	Config         *configs.Config
}

func ProvideMenuServiceImpl(menuRepository MenuRepository, config *configs.Config) *MenuServiceImpl {
	return &MenuServiceImpl{MenuRepository: menuRepository, Config: config}
}

func (m *MenuServiceImpl) ResolveByID(id uuid.UUID) (menu Menu, err error) {
	menu, err = m.MenuRepository.ResolveByID(id)
	if err != nil {
		return
	}
	return
}
