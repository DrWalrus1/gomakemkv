package events

import (
	"fmt"
	"strconv"
	"strings"
)

type DiscInformation struct {
	ID            int
	MessageCodeId int
	Value         string
}

func (mg DiscInformation) outputMarker() {}

func parseDiscInfo(input string) (*DiscInformation, error) {
	var discInfo DiscInformation

	trimmed, _ := strings.CutPrefix(input, discInfoPrefix)

	split := strings.Split(trimmed, delimiter)
	if len(split) < 3 {
		return nil, ErrNotEnoughValues
	}

	id, err := strconv.Atoi(split[0])
	if err != nil {
		return nil, fmt.Errorf("could not parse '%s' into int. %w", split[0], err)
	}
	discInfo.ID = id
	messageCodeId, err := strconv.Atoi(split[1])
	if err != nil {
		return nil, fmt.Errorf("could not parse '%s' into int. %w", split[0], err)
	}
	discInfo.MessageCodeId = messageCodeId
	discInfo.Value = split[2]

	return &discInfo, nil
}