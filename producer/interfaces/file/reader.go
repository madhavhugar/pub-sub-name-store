package file

import (
	"bufio"
	"os"

	"github.com/madhavhugar/pub-sub-name-store/utils"
)

// Readline returns a channel populated with file's content
func Readline(url string) (<-chan string, error) {
	f, err := os.Open(url)
	utils.HandleError(err, "Unable to open file")
	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(f)
	err = scanner.Err()
	utils.HandleError(err, "Scanner error while buffering file")
	if err != nil {
		return nil, err
	}

	ch := make(chan string)
	go func() {
		for scanner.Scan() {
			ch <- scanner.Text()
		}
		close(ch)
	}()

	return ch, nil
}
