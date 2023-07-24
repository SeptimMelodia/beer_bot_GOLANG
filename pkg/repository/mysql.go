package repository

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
}

const (
	usersTable        = "users"
	citiesTable       = "cities"
	eagerPeersTable   = "eager_peers"
	randomGroupsTable = "random_groups"
	reportsTable      = "reports"
	userRole          = "user_role"
)

func NewMySqlDB(conf Config) (*sqlx.DB, error) {
	db, err := sqlx.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", conf.Username, conf.Password, conf.Host, conf.Port, conf.DBName))
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil
}
