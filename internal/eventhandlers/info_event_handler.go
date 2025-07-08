package eventhandlers

import (
	"io"

	"github.com/DrWalrus1/gomakemkv"
	"github.com/DrWalrus1/gomakemkv/events"
)

func MakeMkvInfoEventHandler(reader io.Reader) (standardEventsChannel chan events.MakeMkvOutput, discInfoEventChannel chan gomakemkv.MkvDiscInfo, disconnectChannel chan bool) {
	standardEventsChannel = make(chan events.MakeMkvOutput)
	discInfoEventChannel = make(chan gomakemkv.MkvDiscInfo)
	disconnectChannel = make(chan bool)

	go func() {
		c := events.ParseEventStream(reader)
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
					discInfoEventChannel <- gomakemkv.MakeMkveventsIntoMakeMkvDiscInfo(discInfoEvents)
				}
				break
			}
		}
		disconnectChannel <- true
		close(standardEventsChannel)
		close(discInfoEventChannel)
		close(disconnectChannel)
	}()
	return standardEventsChannel, discInfoEventChannel, disconnectChannel
}
