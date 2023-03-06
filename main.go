package main

import (
	"anothapp_towatch/database"
	"anothapp_towatch/helpers"
	"anothapp_towatch/repositories"
	"fmt"
)

var (
	db                    = database.Open()
	userRepository        = repositories.NewUserRepository(db)
	userShowRepository    = repositories.NewUserShowRepository(db)
	userToWatchRepository = repositories.NewUserToWatchRepository(db)
	watchHelper           = helpers.NewWatchHelper()
)

func main() {
	defer database.Close(db)

	usersRows := userRepository.GetUsers()
	defer usersRows.Close()

	var id string

	for usersRows.Next() {
		usersRows.Scan(&id)

		fmt.Printf("Update shows to watch for user : %s\n", id)

		showsRows := userShowRepository.GetShowsToContinueByUserId(id)
		defer showsRows.Close()

		shows := watchHelper.CheckShowsToWatch(showsRows)
		userToWatchRepository.DeleteToWatchByUserId(id)
		userToWatchRepository.UpdateShowsToWatchByUserId(id, shows)
	}
}
