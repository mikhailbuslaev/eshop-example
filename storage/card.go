package storage

import (
	"github.com/mikhailbuslaev/eshop/model"
)

func (s *Storage) GetCards(quantity int) (*[]model.Card, error) {
	cards := []model.Card{}
	err := s.Db.Select(&cards, `SELECT * FROM cards LIMIT $1`, quantity)
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
	_, err := s.Db.NamedExec(`INSERT INTO cards (id, title, price, description, count, picturepath) VALUES (:id, :title, :price, :description, :count, :picturepath)`, card)
	return err
}

func (s *Storage) DeleteCard(id string) error {
	_, err := s.Db.Exec(`DELETE FROM cards WHERE id=$1`, id)
	return err
}

func (s *Storage) UpdateCard(card *model.Card) error {
	_, err := s.Db.NamedExec(`UPDATE cards SET id=:id, title=:title, price=:price, description=:description, count=:count WHERE id = :id`, card)
	return err
}
