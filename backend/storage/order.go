package storage

import (
	"github.com/mikhailbuslaev/eshop/model"
	"github.com/jmoiron/sqlx/types"
	"encoding/json"
)

func (s *Storage) GetOrders(userId string, quantity, start int) (*[]model.Order, error) {
	orders := make([]model.Order, 0)
	jsons := make([]string, 0)
	err := s.Db.Select(&jsons, `SELECT data FROM orders WHERE user_id=$1 OFFSET $2 FETCH FIRST $3 ROWS ONLY`, userId, start, quantity)
	if err != nil {
		return nil, err
	}

	for i := range jsons {
		order := model.Order{}

		buf := []byte(jsons[i])
		err = json.Unmarshal(buf, &order)
		if err != nil {
			return nil, err
		}
		orders = append(orders, order)
	}
	return &orders, nil
}

func (s *Storage) GetOrder(userId, id string) (*model.Order, error) {
	order := model.Order{}
	order.Id = id
	var jsonString string
	err := s.Db.Get(&jsonString, `SELECT data FROM orders WHERE id=$1 AND user_id=$2 LIMIT 1`, id, userId)
	if err != nil {
		return nil, err
	}

	buf := []byte(jsonString)
	err = json.Unmarshal(buf, &order)
	if err != nil {
		return nil, err
	}
	return &order, nil
}

func (s *Storage) AddOrder(order *model.Order) error {
	j := types.JSONText(order.JsonData)

	jsonVal, err := j.Value()
	if err != nil {
		return err
	}
	_, err = s.Db.Exec(`INSERT INTO orders VALUES ($1, $2, $3)`, order.Id, jsonVal, order.UserId)
	return err
}

func (s *Storage) DeleteOrder(userId, id string) error {
	_, err := s.Db.Exec(`DELETE FROM orders WHERE id=$1 AND user_id=$2`, id, userId)
	return err
}

func (s *Storage) GetOrdersQuantity(userId string) (int, error) {
	var quantity int
	err := s.Db.Get(&quantity, `SELECT COUNT(id) FROM orders WHERE user_id=$1`, userId)
	if err != nil {
		return -1, err
	}
	return quantity, nil
}