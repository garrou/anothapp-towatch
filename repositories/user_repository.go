package repositories

import (
	"database/sql"
)

type userRepository struct {
	db *sql.DB
}

type UserRepository interface {
	GetUsers() *sql.Rows
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{db}
}

func (u *userRepository) GetUsers() *sql.Rows {
	queryStmt := `SELECT id FROM users`
	rows, err := u.db.Query(queryStmt)

	if err != nil {
		panic(err.Error())
	}
	return rows
}
