package messagebroker

import "github.com/streadway/amqp"

// PublishParameters restricts and defines the parameters
// that can be passed when publishing a message
type PublishParameters struct {
	Exchange string
	Key      string
	Msg      amqp.Publishing
}

// Publish publishes a given message to an exchange with a provided key
func (c AmqpChannel) Publish(params *PublishParameters) error {
	return c.channel.Publish(
		params.Exchange,
		params.Key,
		false, // Mandatory: If this flag is set, the server will return an unroutable message with a Return method.
		false, // Immediate: If this flag is set, the server will return an undeliverable message with a Return method
		params.Msg,
	)
}
