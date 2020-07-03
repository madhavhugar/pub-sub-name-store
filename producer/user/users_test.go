package user

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetUsers(t *testing.T) {
	users, err := GetUsers("../data/data_example.csv")
	wanted := Info{Name: "Kathy Hunter", Email: "abc.def0@bing.com"}
	got := users[0]
	assert.Equal(t, wanted, got)
	assert.Equal(t, nil, err)
}
