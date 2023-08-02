package menu

import "encoding/json"

type Menu struct {
	ID       string `db:"id"`
	Name     string `db:"name"`
	Price    string `db:"price"`
	Category string `db:"category"`
	Stock    string `db:"stock"`
}
type MenuResponseFormat struct {
	ID       string
	Name     string
	Price    string
	Category string
	Stock    string
}

func (m Menu) MarshalJSON() ([]byte, error) {
	return json.Marshal(m.ToResponseFormat())
}

func (m Menu) ToResponseFormat() MenuResponseFormat {
	return MenuResponseFormat{
		ID:       m.ID,
		Name:     m.Name,
		Price:    m.Price,
		Category: m.Category,
		Stock:    m.Stock,
	}
}
