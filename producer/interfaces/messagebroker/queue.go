package messagebroker

import "github.com/streadway/amqp"

// QueueParameters restricts and defines the parameters
// that can be passed when declaring a queue
type QueueParameters struct {
	Name    string
	Durable bool
	Args    amqp.Table
}

// QueueDeclare declares a Queue on rabbitMQ
func (a AmqpChannel) QueueDeclare(params *QueueParameters) (amqp.Queue, error) {
	return a.channel.QueueDeclare(
		params.Name,
		params.Durable,
		false, // We don't want to auto-delete queue
		false, // We don't need exclusive queues
		false,
		params.Args,
	)
}
