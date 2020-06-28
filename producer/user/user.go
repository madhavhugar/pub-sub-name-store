package user

import (
	"encoding/json"
	"fmt"
	file "producer/interfaces/file"
	msgbroker "producer/interfaces/messagebroker"
	utils "producer/utils"
	"strings"

	"github.com/streadway/amqp"
)

// Info to represent user information
type Info struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

// GetUsers does TODO
func GetUsers(filepath string) ([]Info, error) {
	reader, err := file.Readline(filepath)
	utils.HandleError(err, "error in main reading file")
	if err != nil {
		return nil, err
	}

	var users []Info
	for line := range reader {
		info := strings.Split(line, ",")
		user := Info{
			Name:  info[0],
			Email: info[1],
		}
		fmt.Println("name:", user.Name, "|| email:", user.Email)
		users = append(users, user)
	}
	return users, nil
}

// PublishUserInfo does TODO
func PublishUserInfo(user *Info) error {
	body, err := json.Marshal(user)
	// utils.HandleError(err, "error while marshalling user")
	if err != nil {
		return err
	}

	c, err := msgbroker.GetChannel()
	// utils.HandleError(err, "error while creating message broker channel")
	if err != nil {
		return err
	}

	pubParams := msgbroker.PublishParameters{
		// TODO: exchange and key as env vars
		Exchange: "users",
		Key:      "users.info",
		Msg: amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		},
	}
	c.Publish(&pubParams)
	return nil
}
