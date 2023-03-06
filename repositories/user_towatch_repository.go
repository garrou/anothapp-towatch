package repositories

import (
	"anothapp_towatch/models"
	"database/sql"
	"fmt"
	"strings"
)

type userToWatchRepository struct {
	db *sql.DB
}

type UserToWatchRepository interface {
	UpdateShowsToWatchByUserId(userId string, shows []models.ShowDto)
	DeleteToWatchByUserId(userId string)
}

func NewUserToWatchRepository(db *sql.DB) UserToWatchRepository {
	return &userToWatchRepository{db}
}

func (u *userToWatchRepository) DeleteToWatchByUserId(userId string) {
	queryStmt := `
		DELETE FROM users_towatch
		WHERE user_id = $1
	`

	if _, err := u.db.Query(queryStmt, userId); err != nil {
		panic(err.Error())
	}
}

func (u *userToWatchRepository) UpdateShowsToWatchByUserId(userId string, shows []models.ShowDto) {
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

	if _, insertErr := u.db.Query(queryToWatch); insertErr != nil {
		fmt.Println(insertErr.Error())
	}

	if _, updateErr := u.db.Query(queryUpdateContinue, userId); updateErr != nil {
		fmt.Println(updateErr.Error())
	}
}
