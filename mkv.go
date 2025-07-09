// Package provides parsing capabilities for the makemkv console program.
// Can handle makemkvcon mkv, backup, reg and info
package gomakemkv

import (
	"bufio"
	"fmt"
	"io"

	"github.com/DrWalrus1/gomakemkv/events"
)

func parseEventStream(reader io.Reader) <-chan events.MakeMkvOutput {
	c := make(chan events.MakeMkvOutput)
	go func() {
		scanner := bufio.NewScanner(reader)
		for scanner.Scan() {
			output, err := events.Parse(scanner.Text())
			if err != nil {
				fmt.Println(err.Error())
				continue
			}
			c <- output
		}
		close(c)
	}()
	return c
}

func ParseMakeMkvLogs(reader io.Reader) <-chan events.MakeMkvOutput {
	return parseEventStream(reader)
}

func ParseMakeMkvInfoCommandLogs(reader io.Reader) (<-chan events.MakeMkvOutput, <-chan MkvDiscInfo) {
	standardEventsChannel := make(chan events.MakeMkvOutput)
	discInfoEventChannel := make(chan MkvDiscInfo)
	go func() {
		c := parseEventStream(reader)
		var discInfoEvents []events.MakeMkvOutput
		for {
			if i, ok := <-c; ok {
				if standardEvent, ok := i.(*events.TitleInformation); ok {
					discInfoEvents = append(discInfoEvents, standardEvent)
				} else if standardEvent, ok := i.(*events.DiscInformation); ok {
					discInfoEvents = append(discInfoEvents, standardEvent)
				} else if standardEvent, ok := i.(*events.StreamInformation); ok {
					discInfoEvents = append(discInfoEvents, standardEvent)
				} else {
					standardEventsChannel <- i
				}
			} else {
				if len(discInfoEvents) > 0 {
					discInfoEventChannel <- MakeMkvEventsIntoMakeMkvDiscInfo(discInfoEvents)
				}
				break
			}
		}
		close(discInfoEventChannel)
		close(standardEventsChannel)
	}()
	return standardEventsChannel, discInfoEventChannel
}
