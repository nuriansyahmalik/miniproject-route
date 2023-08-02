package menu

import (
	"github.com/evermos/boilerplate-go/infras"
	"github.com/evermos/boilerplate-go/shared/failure"
	"github.com/evermos/boilerplate-go/shared/logger"
	"github.com/gofrs/uuid"
)

var (
	menuQueries = struct {
		insertMenu string
		selectMenu string
	}{
		insertMenu: `
			INSERT INTO menu (
			    id,
			    name,
			    price,
			    category,
			    stock
			) VALUES (
			    :id,
			    :name,
			    :price,
		        :category,
			    :stock)`,
		selectMenu: `
			SELECT
				menu.id,
			    menu.name,
			    menu.price,
			    menu.category,
			    menu.stock
			FROM menu`,
	}
)

type MenuRepository interface {
	Create(menu Menu) (err error)
	ExistsByID(id uuid.UUID) (exists bool, err error)
	ResolveByID(id uuid.UUID) (menu Menu, err error)
}

type MenuRepositoryMysql struct {
	DB *infras.MySQLConn
}

func ProvideMenuRepositoryMySQL(db *infras.MySQLConn) *MenuRepositoryMysql {
	return &MenuRepositoryMysql{DB: db}
}

func (m *MenuRepositoryMysql) Create(menu Menu) (err error) {
	return err
}

func (m *MenuRepositoryMysql) ResolveByID(id uuid.UUID) (mennu Menu, err error) {
	err = m.DB.Read.Get(
		&mennu,
		menuQueries.selectMenu+" WHERE menu.id = ?",
		id.String())
	if err != nil {
		err = failure.NotFound("menu")
		logger.ErrorWithStack(err)
		return
	}
	return
}

func (m *MenuRepositoryMysql) ExistsByID(id uuid.UUID) (exists bool, err error) {
	err = m.DB.Read.Get(
		&exists,
		"SELECT COUNT(id) FROM menu WHERE menu.id = ?",
		id.String())
	if err != nil {
		logger.ErrorWithStack(err)
	}

	return
}
