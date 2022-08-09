package storage

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/mikhailbuslaev/eshop/config"
)

type Storage struct {
	Db       *sqlx.DB
	DbConfig DbConfig
}

type DbConfig struct {
	DriverName string
	Host       string
	Port       string
	User       string
	Password   string
	DbName     string
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
		log.Fatalln("connection to db failed:" + err.Error())
	}
	return s
}
