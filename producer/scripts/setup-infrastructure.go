package main

import (
	"fmt"
	msgbroker "producer/interfaces/messagebroker"
	utils "producer/utils"
)

func main() {
	c, err := msgbroker.GetChannel()
	utils.HandleError(err, "Error creating channel")

	// Declare Topic Exchange
	eParams := msgbroker.ExchangeParameters{
		Name:    "users",
		Kind:    "topic",
		Durable: true,
		Args:    nil,
	}
	err = c.ExchangeDeclare(&eParams)
	utils.HandleError(err, "Error declaring exchange")

	// Declare Queue
	qParams := msgbroker.QueueParameters{
		// TODO: queue name as env param
		Name:    "users-info-queue",
		Durable: true,
		Args:    nil,
	}
	q, err := c.QueueDeclare(&qParams)
	utils.HandleError(err, "Error declaring queue")
	fmt.Println("Queue name: ", q.Name)

	// Bind the Queue to the Exchange with a key `users.info`
	pbParams := msgbroker.BindingParameters{
		Name:     q.Name,
		Key:      "users.info",
		Exchange: eParams.Name,
	}
	err = c.QueueBind(&pbParams)
	utils.HandleError(err, "Error binding queue to exchange")
}
