package helpers

import (
	"anothapp_towatch/models"
	"database/sql"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func CheckShowsToWatch(rows *sql.Rows) []models.ShowDto {
	var id, number int
	var showInfos models.ShowInfos
	var showsToWatch []models.ShowDto
	apiKey := os.Getenv("BETASERIES_KEY")

	for rows.Next() {
		rows.Scan(&id, &number)

		body := HttpGet(fmt.Sprintf("https://api.betaseries.com/shows/display?id=%d&key=%s", id, apiKey))

		if err := json.Unmarshal(body, &showInfos); err != nil {
			panic(err.Error())
		}

		nb, convErr := strconv.Atoi(showInfos.Show.Seasons)

		if convErr != nil {
			panic(convErr.Error())
		}
		showsToWatch = append(showsToWatch, models.ShowDto{
			Id:        id,
			NbToWatch: nb - number,
			Ended:     strings.Contains(showInfos.Show.Status, "Ended"),
		})
	}
	return showsToWatch
}
