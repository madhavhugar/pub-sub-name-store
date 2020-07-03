package main

import (
	"fmt"
	"time"

	user "github.com/madhavhugar/pub-sub-name-store/user"
	utils "github.com/madhavhugar/pub-sub-name-store/utils"
)

func main() {
	// Wait for rabbitMQ to startup
	time.Sleep(10 * time.Second)

	setupInfrastructure()

	users, err := user.GetUsers("./data/data_example.csv")
	utils.HandleError(err, "error while fetching users")

	for i := 0; i < len(users); i++ {
		fmt.Println("name:", users[i].Name, "|| email:", users[i].Email)
		err = user.PublishUserInfo(&users[i])
		utils.HandleError(err, "error while publishing user info")
	}
}
