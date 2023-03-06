package repositories

import (
	"database/sql"
)

type userShowRepository struct {
	db *sql.DB
}

type UserShowRepository interface {
	GetShowsToContinueByUserId(userId string) *sql.Rows
}

func NewUserShowRepository(db *sql.DB) UserShowRepository {
	return &userShowRepository{db}
}

func (u *userShowRepository) GetShowsToContinueByUserId(userId string) *sql.Rows {
	queryStmt := `
		SELECT shows.id AS id, MAX(number) AS number
        FROM users_shows
        JOIN users_seasons ON users_seasons.show_id = users_shows.show_id
        JOIN shows ON users_shows.show_id = shows.id
        WHERE users_shows.user_id = $1 AND users_shows.continue = TRUE
        GROUP BY shows.id, title, poster
	`

	rows, err := u.db.Query(queryStmt, userId)

	if err != nil {
		panic(err.Error())
	}
	return rows
}
