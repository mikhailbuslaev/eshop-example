package storage

import (
	"github.com/mikhailbuslaev/eshop/model"
	"github.com/jmoiron/sqlx/types"
	"encoding/json"
)

func (s *Storage) GetOrders(userId string, quantity, start int) (*[]model.Order, error) {
	orders := make([]model.Order, 0)
	ids := make([]string, 0)
	var stringIds string
	var stringJsons string

	err := s.Db.Select(&stringIds, `SELECT id FROM orders WHERE user_id=$1 OFFSET $2 FETCH FIRST $3 ROWS ONLY`, userId, start, quantity)
	if err != nil {
		return nil, err
	}

	jsons := make([]string, 0)
	err = s.Db.Select(&stringJsons, `SELECT data FROM orders WHERE user_id=$1 OFFSET $2 FETCH FIRST $3 ROWS ONLY`, userId, start, quantity)
	if err != nil {
		return nil, err
	}

	buf := []byte(stringIds)
	err = json.Unmarshal(buf, &ids)
	if err != nil {
		return nil, err
	}

	buf = []byte(stringJsons)
	err = json.Unmarshal(buf, &jsons)
	if err != nil {
		return nil, err
	}

	for i := range ids {
		order := model.Order{}
		order.Id = ids[i]
		order.JsonData = jsons[i]
		order.UserId = userId
		orders = append(orders, order)
	}
	return &orders, nil
}

func (s *Storage) GetOrder(userId, id string) (*model.Order, error) {
	order := model.Order{}
	order.Id = id
	err := s.Db.Get(&order.JsonData, `SELECT data FROM orders WHERE id=$1 AND user_id=$2 LIMIT 1`, id, userId)
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