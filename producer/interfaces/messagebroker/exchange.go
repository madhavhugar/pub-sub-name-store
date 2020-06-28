package messagebroker

import "github.com/streadway/amqp"

// ExchangeParameters restricts and defines the parameters
// that can be passed when declaring a exchange
type ExchangeParameters struct {
	Name    string
	Kind    string
	Durable bool
	Args    amqp.Table
}

// ExchangeDeclare declares an exchange on the message broker
func (a AmqpChannel) ExchangeDeclare(params *ExchangeParameters) error {
	return a.channel.ExchangeDeclare(
		params.Name,
		params.Kind,
		params.Durable,
		false, // autoDelete set to false
		false, // internal set to false
		false, // noWait set to false
		params.Args,
	)
}
