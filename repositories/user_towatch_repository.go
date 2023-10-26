package repositories

import (
	"anothapp_towatch/database"
	"anothapp_towatch/models"
	"fmt"
	"strings"
)

func DeleteToWatchByUserId(userId string) {
	queryStmt := `
		DELETE FROM users_towatch
		WHERE user_id = $1
	`

	if _, err := database.Db.Query(queryStmt, userId); err != nil {
		panic(err.Error())
	}
}

func UpdateShowsToWatchByUserId(userId string, shows []models.ShowDto) {
	queryToWatch := `
		INSERT INTO users_towatch (user_id, show_id, nb)
		VALUES
	`

	queryUpdateContinue := `
		UPDATE users_shows
		SET continue = FALSE
		WHERE user_id = $1
		AND show_id IN (
	`

	for _, s := range shows {
		if s.NbToWatch == 0 && s.Ended {
			queryUpdateContinue += fmt.Sprintf("%d,", s.Id)
		} else if s.NbToWatch > 0 {
			queryToWatch += fmt.Sprintf(" (%s, %d, %d),", userId, s.Id, s.NbToWatch)
		}
	}
	queryToWatch = strings.TrimSuffix(queryToWatch, ",") + ";"
	queryUpdateContinue = strings.TrimSuffix(queryUpdateContinue, ",") + ");"

	if _, insertErr := database.Db.Query(queryToWatch); insertErr != nil {
		fmt.Println(insertErr.Error())
	}

	if _, updateErr := database.Db.Query(queryUpdateContinue, userId); updateErr != nil {
		fmt.Println(updateErr.Error())
	}
}
