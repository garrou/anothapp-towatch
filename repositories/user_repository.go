package repositories

import (
	"anothapp_towatch/database"
	"database/sql"
)

func GetUsers() *sql.Rows {
	queryStmt := `SELECT id FROM users`
	rows, err := database.Db.Query(queryStmt)

	if err != nil {
		panic(err.Error())
	}
	return rows
}
