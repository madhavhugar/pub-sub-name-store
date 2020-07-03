package file

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadline(t *testing.T) {
	ch, err := Readline("../../data/data_example.csv")
	var lines []string
	for line := range ch {
		lines = append(lines, line)
	}
	got := lines[0]
	wanted := "Kathy Hunter,abc.def0@bing.com"
	assert.Equal(t, wanted, got)
	assert.Equal(t, nil, err)
}
