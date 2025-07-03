package streamReader

import (
	"bufio"
	"io"
)


func ReadStream(reader io.Reader) <-chan string {
	c := make(chan string)
	go func() {
		defer close(c)
		scanner := bufio.NewScanner(reader)
		for scanner.Scan() {
			c <- scanner.Text()
		}
	}()
	return c
}
