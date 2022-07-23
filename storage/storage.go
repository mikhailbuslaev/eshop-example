package storage

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	"github.com/mikhailbuslaev/eshop/card"
	"github.com/mikhailbuslaev/eshop/config"
)

type Storage struct {
	Db       *sqlx.DB
	DbConfig *DbConfig
}

type DbConfig struct {
	Driver   string
	Host     string
	Port     string
	User     string
	Password string
	DbName   string
	SSLmode  string
}

func New() *Storage {
	s := &Storage{}
	config.Parse(s.DbConfig, "config/dbconfig.yaml")
	dbCreds := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=%s ",
		s.DbConfig.User, s.DbConfig.Password, s.DbConfig.Host,
		s.DbConfig.Port, s.DbConfig.DbName, s.DbConfig.SSLmode)

	var err error
	s.Db, err = sqlx.Connect(s.DbConfig.Driver, dbCreds)
	if err != nil {
		log.Fatalln("connection to db failed")
	}
	return s
}

func (s *Storage) GetCards() (*[]card.Card, error) {
	cards := []card.Card{}
	err := s.Db.Select(&cards, "SELECT * FROM cards LIMIT 20")
	if err != nil {
		return nil, err
	}
	return &cards, nil
}

func (s *Storage) GetCard(id string) (*card.Card, error) {
	card := card.Card{}
	err := s.Db.Get(&card, "SELECT * FROM cards WHERE id=$1", id)
	if err != nil {
		return nil, err
	}
	return &card, nil
}

func (s *Storage) PutCard(card *card.Card) error {
	_, err := s.Db.NamedExec("INSERT INTO cards (id, title, price, description_, count_, picture_path) VALUES (:id, :title, :price, :description_, :count_, :picture_path)", card)
	return err
}

func (s *Storage) RemoveCard(id string) error {
	_, err := s.Db.NamedExec("DELETE FROM cards WHERE id=$1 ", id)
	return err
}
