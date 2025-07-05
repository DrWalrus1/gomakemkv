package events

import (
	"strconv"
	"strings"
)

type MessageOutput struct {
	Code           string
	Flags          string
	ParameterCount int
	RawMessage     string
	FormatMessage  string
	MessageParams  []string
}

func (mg MessageOutput) outputMarker() {}

func parseMessageOutput(input string) (*MessageOutput, error) {
	var parsedMessage MessageOutput

	trimmed, _ := strings.CutPrefix(input, messageOutputPrefix)

	split := strings.Split(trimmed, delimiter)
	if len(split) < 5 {
		return nil, ErrNotEnoughValues
	}
	parsedMessage.Code = split[0]
	parsedMessage.Flags = split[1]
	parsedParamCount, err := strconv.Atoi(split[2])
	if err != nil {
		return nil, ErrNotEnoughValues
	}
	parsedMessage.ParameterCount = parsedParamCount
	parsedMessage.RawMessage = split[3]
	parsedMessage.FormatMessage = split[4]
	const splitOffset = 5
	ifThereAreParamsAfterOffset := func() bool {
		return len(split) > splitOffset
	}
	doParamsExist := func() bool {
		return parsedMessage.ParameterCount > 0
	}
	paramsDoNotExceedSplitBounds := func() bool {
		return parsedMessage.ParameterCount+splitOffset-1 < len(split)
	}
	if ifThereAreParamsAfterOffset() && doParamsExist() && paramsDoNotExceedSplitBounds() {
		parsedMessage.MessageParams = make([]string, parsedMessage.ParameterCount)
		for i := range parsedMessage.ParameterCount {
			parsedMessage.MessageParams[i] = split[i+splitOffset]
		}
	}

	return &parsedMessage, nil
}