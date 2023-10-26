package repositories

import (
	"anothapp_towatch/database"
	"database/sql"
)

func GetShowsToContinueByUserId(userId string) *sql.Rows {
	queryStmt := `
		SELECT shows.id AS id, MAX(number) AS number
        FROM users_shows
        JOIN users_seasons ON users_seasons.show_id = users_shows.show_id
        JOIN shows ON users_shows.show_id = shows.id
        WHERE users_shows.user_id = $1 AND users_shows.continue = TRUE
        GROUP BY shows.id, title, poster
	`

	rows, err := database.Db.Query(queryStmt, userId)

	if err != nil {
		panic(err.Error())
	}
	return rows
}
