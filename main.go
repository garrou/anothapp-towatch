package main

import (
	"anothapp_towatch/database"
	"anothapp_towatch/helpers"
	"anothapp_towatch/repositories"
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
	/*
		var id string

			for usersRows.Next() {
				usersRows.Scan(&id)

				fmt.Printf("Update shows to watch for user : %s\n", id)
	*/
	showsRows := userShowRepository.GetShowsToContinueByUserId("107385984010527617533")
	defer showsRows.Close()

	shows := watchHelper.CheckShowsToWatch(showsRows)
	userToWatchRepository.DeleteToWatchByUserId("107385984010527617533")
	userToWatchRepository.UpdateShowsToWatchByUserId("107385984010527617533", shows)
	//}
}
