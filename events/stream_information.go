package events

import (
	"fmt"
	"strconv"
	"strings"
)

type StreamInformation struct {
	// Title ID - The id of the title that this StreamInformation relates to
	TitleIndex    int
	StreamIndex   int
	AttributeId   int
	MessageCodeId int
	Value         string
}

func (mg StreamInformation) outputMarker() {}

func parseStreamInfo(input string) (*StreamInformation, error) {
	var streamInfo StreamInformation

	trimmed, found := strings.CutPrefix(input, streamInfoPrefix)
	if !found {
		return nil, ErrPrefixNotFound
	}

	split := strings.Split(trimmed, delimiter)
	if len(split) < 5 {
		return nil, ErrNotEnoughValues
	}

	titleIndex, err := strconv.Atoi(split[0])
	if err != nil {
		return nil, fmt.Errorf("could not parse '%s' into int. %w", split[1], err)
	}
	streamInfo.TitleIndex = titleIndex
	streamIndex, err := strconv.Atoi(split[1])
	if err != nil {
		return nil, fmt.Errorf("could not parse '%s' into int. %w", split[1], err)
	}
	streamInfo.StreamIndex = streamIndex
	streamType, err := strconv.Atoi(split[2])
	if err != nil {
		return nil, fmt.Errorf("could not parse '%s' into int. %w", split[2], err)
	}
	streamInfo.AttributeId = streamType
	messageCode, err := strconv.Atoi(split[3])
	if err != nil {
		return nil, fmt.Errorf("could not parse '%s' into int. %w", split[2], err)
	}
	streamInfo.MessageCodeId = messageCode
	if split[4] == "( Lossless conversion )" {
		streamInfo.Value = "Lossless"
	} else {
		streamInfo.Value = split[4]
	}

	return &streamInfo, nil
}