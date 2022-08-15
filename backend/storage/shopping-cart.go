package storage

import (
	"github.com/mikhailbuslaev/eshop/model"
	"github.com/jmoiron/sqlx/types"
	"encoding/json"
)

func (s *Storage) GetShoppingCart(id string) (*model.ShoppingCart, error) {
	var stringBuf string
	cart := model.ShoppingCart{}
	cart.Items = make([]model.CartItem, 0)
	err := s.Db.Get(&stringBuf, `SELECT items FROM shopping_carts WHERE id=$1 LIMIT 1`, id)
	if err != nil {
		return nil, err
	}

	buf := []byte(stringBuf)
	err = json.Unmarshal(buf, &cart)
	if err != nil {
		return nil, err
	}
	cart.Id = id
	cart.JsonItems = buf
	return &cart, nil
}

func (s *Storage) InitShoppingCart(id string) error {
	_, err := s.Db.Exec(`INSERT INTO shopping_carts (id) VALUES ($1)`, id)
	return err
}

func (s *Storage) DeleteShoppingCart(id string) error {
	_, err := s.Db.Exec(`DELETE FROM shopping_carts WHERE id=$1`, id)
	return err
}

func (s *Storage) UpdateShoppingCart(cart *model.ShoppingCart) error {
	j := types.JSONText(cart.JsonItems)

	jsonVal, err := j.Value()
	if err != nil {
		return err
	}
	_, err = s.Db.Exec(`UPDATE shopping_carts SET items=$1 WHERE id = '1'`, jsonVal)
	return err
}

func (s *Storage) ClearShoppingCart(id string) error {
	_, err := s.Db.Exec(`UPDATE shopping_carts SET items=$1 WHERE id = $2`,`{"items":[]}`, id)
	return err
}
