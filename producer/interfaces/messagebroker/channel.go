package messagebroker

import (
	"producer/utils"
	"sync"

	"github.com/streadway/amqp"
)

var connection *amqp.Connection
var once sync.Once

func connect() (*amqp.Connection, error) {
	// TODO: replace to use env var
	rabbitMQUrl := "amqp://guest:guest@localhost:5672/"
	conn, err := amqp.Dial(rabbitMQUrl)
	return conn, err
}

// AmqpChannel is the wrapper to wrap the calls to the message broker
type AmqpChannel struct {
	channel *amqp.Channel
}

// getConnection returns the singleton connection instance to rabbitMQ via AMQP
func getConnection() (*amqp.Connection, error) {
	var err error
	once.Do(func() {
		connection, err = connect()
	})
	return connection, err
}

func closeConnection() {
	c, _ := getConnection()
	defer c.Close()
}

// GetChannel returns a new channel over the rabbitMQ connection
func GetChannel() (AmqpChannel, error) {
	c, err := getConnection()
	utils.HandleError(err, "Failed to connect to rabbitmq server")
	channel, err := c.Channel()
	utils.HandleError(err, "Failed to create rabbitmq channel")
	a := AmqpChannel{channel: channel}
	return a, err
}
