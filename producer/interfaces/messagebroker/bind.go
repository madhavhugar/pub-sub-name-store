package messagebroker

// BindingParameters defines params while bindings queue to exchanges
type BindingParameters struct {
	Name     string
	Key      string
	Exchange string
}

// QueueBind binds a queue to an exchange
func (c AmqpChannel) QueueBind(params *BindingParameters) error {
	return c.channel.QueueBind(
		params.Name,     // name of the queue
		params.Key,      // routing key; if empty then routing queue is the queue's name
		params.Exchange, // name of the exchange
		false,           // noWait
		nil,             // args
	)
}
