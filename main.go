package main

import (
	"anothapp_towatch/database"
	"anothapp_towatch/helpers"
	"anothapp_towatch/repositories"
	"fmt"
)

func main() {
	database.Open()
	defer database.Close()

	usersRows := repositories.GetUsers()
	defer usersRows.Close()

	var id string

	for usersRows.Next() {
		usersRows.Scan(&id)

		fmt.Printf("User : %s\n", id)

		showsRows := repositories.GetShowsToContinueByUserId(id)
		defer showsRows.Close()

		shows := helpers.CheckShowsToWatch(showsRows)
		repositories.DeleteToWatchByUserId(id)
		repositories.UpdateShowsToWatchByUserId(id, shows)
	}
}
