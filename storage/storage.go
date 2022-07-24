package storage

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/mikhailbuslaev/eshop/card"
	"github.com/mikhailbuslaev/eshop/config"
)

type Storage struct {
	Db       *sqlx.DB
	DbConfig DbConfig
}

type DbConfig struct {
	DriverName   string
	Host     string
	Port     string
	User     string
	Password string
	DbName   string
}

func New() *Storage {
	s := &Storage{}
	config.Parse(&s.DbConfig, "config/dbconfig.yaml")
	dbCreds := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
	s.DbConfig.Host, s.DbConfig.Port, s.DbConfig.User,
	s.DbConfig.Password, s.DbConfig.DbName)

	var err error
	s.Db, err = sqlx.Connect(s.DbConfig.DriverName, dbCreds)
	if err != nil {
		log.Fatalln("connection to db failed:"+err.Error())
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
	_, err := s.Db.NamedExec("INSERT INTO cards (id, title, price, description, count, picturepath) VALUES (:id, :title, :price, :description, :count, :picturepath)", card)
	return err
}

func (s *Storage) RemoveCard(id string) error {
	_, err := s.Db.NamedExec("DELETE FROM cards WHERE id=$1 ", id)
	return err
}
