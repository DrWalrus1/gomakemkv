package streamReader_test

import (
	"fmt"
	"io"
	"testing"

	"github.com/DrWalrus1/gomakemkv/internal/streamReader"
	"github.com/go-playground/assert/v2"
)

func simulateMakeMkvProgressOutput(t *testing.T) io.Reader {
	reader, writer := io.Pipe()

	go func() {
		defer writer.Close()

		maxPercentage := 20
		totalPercentage := 10
		for currentPercentage := 0; currentPercentage <= totalPercentage; currentPercentage++ {
			line := fmt.Sprintf("PRGV:%d,%d,%d\n", currentPercentage, totalPercentage, maxPercentage)
			_, err := writer.Write([]byte(line))
			if err != nil {
				t.Errorf("Error writing to pipe: %v\n", err)
				return
			}
			// Use this if you want to slow down the stream
			// time.Sleep(1 * time.Second)
		}
	}()
	return reader
}


func TestReadStream(t *testing.T) {
	mockOutput := simulateMakeMkvProgressOutput(t)
	c := streamReader.ReadStream(mockOutput)
	actualLineCount := 0
	for range c {
		actualLineCount++
	}
	assert.Equal(t, 11, actualLineCount)

}
