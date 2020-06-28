package messagebroker

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type channelMock interface {
	QueueBind(params *BindingParameters) error
}

type amqpWrapper interface {
	channel channelMock
}

func TestQueueBind(t *testing.T) {
	fmt.Println("testing...")
	c, err := GetChannel()
	
	assert.Equal(t, 1, err)
}
