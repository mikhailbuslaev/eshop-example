package storage

import (
	"github.com/mikhailbuslaev/eshop/model"
)

func (s *Storage) GetCards(quantity, start int) (*[]model.Card, error) {
	cards := []model.Card{}
	err := s.Db.Select(&cards, `SELECT * FROM cards OFFSET $1 FETCH FIRST $2 ROWS ONLY`, start, quantity)
	if err != nil {
		return nil, err
	}
	return &cards, nil
}

func (s *Storage) GetCardsByCategory(quantity, start int, category string) (*[]model.Card, error) {
	cards := []model.Card{}
	err := s.Db.Select(&cards, `SELECT * FROM cards WHERE category=$1 OFFSET $2 FETCH FIRST $3 ROWS ONLY`, category, start, quantity)
	if err != nil {
		return nil, err
	}
	return &cards, nil
}

func (s *Storage) GetCard(id string) (*model.Card, error) {
	card := model.Card{}
	err := s.Db.Get(&card, `SELECT * FROM cards WHERE id=$1 LIMIT 1`, id)
	if err != nil {
		return nil, err
	}
	return &card, nil
}

func (s *Storage) AddCard(card *model.Card) error {
	_, err := s.Db.NamedExec(`INSERT INTO cards (id, title, price, description, count, picturepath, category) VALUES (:id, :title, :price, :description, :count, :picturepath, :category)`, card)
	return err
}

func (s *Storage) DeleteCard(id string) error {
	_, err := s.Db.Exec(`DELETE FROM cards WHERE id=$1`, id)
	return err
}

func (s *Storage) UpdateCard(card *model.Card) error {
	_, err := s.Db.NamedExec(`UPDATE cards SET id=:id, title=:title, price=:price, description=:description, count=:count, category=:category WHERE id = :id`, card)
	return err
}

func (s *Storage) GetCardsQuantity() (int, error) {
	var quantity int
	err := s.Db.Get(&quantity, `SELECT COUNT(id) FROM cards`)
	if err != nil {
		return -1, err
	}
	return quantity, nil
}

func (s *Storage) GetCardsQuantityByCategory(category string) (int, error) {
	var quantity int
	err := s.Db.Get(&quantity, `SELECT COUNT(id) FROM cards WHERE category=$1`, category)
	if err != nil {
		return -1, err
	}
	return quantity, nil
}
