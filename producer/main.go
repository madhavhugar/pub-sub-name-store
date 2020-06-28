package main

import (
	"fmt"
	user "producer/user"
	utils "producer/utils"
)

func main() {
	users, err := user.GetUsers("./data/data_example.csv")
	utils.HandleError(err, "error while fetching users")

	for i := 0; i < len(users); i++ {
		fmt.Println("name:", users[i].Name, "|| email:", users[i].Email)
		err = user.PublishUserInfo(&users[i])
		utils.HandleError(err, "error while publishing user info")
	}
}
