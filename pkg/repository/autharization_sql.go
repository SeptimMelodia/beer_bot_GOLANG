package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

type Active bool

func (a *Active) Scan(value interface{}) error {
	// Преобразование значения из базы данных в тип Active
	switch v := value.(type) {
	case bool:
		*a = Active(v)
	case []byte:
		if string(v) == "\x01" {
			*a = Active(true)
		} else {
			*a = Active(false)
		}
	default:
		return fmt.Errorf("неожиданный тип значения для Active: %T", value)
	}
	return nil
}

type User struct {
	ID            int    `db:"id" json:"id"`
	Name          string `db:"name" json:"name"`
	TelegramID    int64  `db:"telegram_id" json:"telegram_id"`
	TelegramAlias string `db:"telegram_alias" json:"telegram_alias"`
	IntraName     string `db:"intra_name" json:"intra_name"`
	Active        Active `db:"active" json:"active"`
	CityID        int    `db:"city_id" json:"city_id"`
	RoleID        int    `db:"role_id" json:"role_id"`
}

type Auth struct {
	db *sqlx.DB
}

func NewAuth(db *sqlx.DB) *Auth {
	return &Auth{db: db}
}

func (r *Auth) CreateUser(user User) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (name, telegram_id, telegram_alias, intra_name, active, city_id, role_id) VALUES ($1, $2, $3, $4, $5, $6, $7)", usersTable)
	row := r.db.QueryRow(query, user.Name, user.TelegramID, user.TelegramAlias, user.IntraName, user.Active, user.CityID, user.RoleID)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (r *Auth) GetUser(telegramId string) (User, error) {
	var user User
	query := fmt.Sprintf("SELECT * FROM %s WHERE telegram_id=?", usersTable)
	err := r.db.Get(&user, query, telegramId)
	return user, err
}
