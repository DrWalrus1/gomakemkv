package events

import (
	"fmt"
	"strconv"
	"strings"
)

type TitleInformation struct {
	TitleIndex    int
	AttributeId   int
	MessageCodeId int
	Value         string
}

func (mg TitleInformation) outputMarker() {}

func parseTitleInfo(input string) (*TitleInformation, error) {
	var titleInfo TitleInformation

	trimmed, _ := strings.CutPrefix(input, titleInfoPrefix)

	split := strings.Split(trimmed, delimiter)
	if len(split) < 4 {
		return nil, ErrNotEnoughValues
	}

	titleIndex, err := strconv.Atoi(split[0])
	if err != nil {
		return nil, fmt.Errorf("could not parse '%s' into int. %w", split[0], err)
	}
	titleInfo.TitleIndex = titleIndex
	attributeId, err := strconv.Atoi(split[1])
	if err != nil {
		return nil, fmt.Errorf("could not parse '%s' into int. %w", split[1], err)
	}
	titleInfo.AttributeId = attributeId
	messageCode, err := strconv.Atoi(split[2])
	if err != nil {
		return nil, fmt.Errorf("could not parse '%s' into int. %w", split[2], err)
	}
	titleInfo.MessageCodeId = messageCode
	titleInfo.Value = split[3]

	return &titleInfo, nil
}
