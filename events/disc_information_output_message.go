package events

import (
	"strconv"
	"strings"
)

type DiscInformationOutputMessage struct {
	TitleCount int
}

func (mg DiscInformationOutputMessage) outputMarker() {}

func parseDiscInformationOutputMessage(input string) (*DiscInformationOutputMessage, error) {
	var discInformationOutput DiscInformationOutputMessage

	trimmed, _ := strings.CutPrefix(input, discInfoOutputPrefix)

	titleCount, err := strconv.Atoi(trimmed)
	if err != nil {
		return nil, err
	}
	discInformationOutput.TitleCount = titleCount
	return &discInformationOutput, nil
}