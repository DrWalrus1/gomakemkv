package eventhandlers

import (
	"io"

	"github.com/DrWalrus1/gomakemkv/events"
)

func MakeMkvMkvEventHandler(reader io.Reader) <-chan events.MakeMkvOutput {
	return events.ParseEventStream(reader)
}
