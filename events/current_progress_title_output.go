package events

import "strings"

type CurrentProgressTitleOutput struct {
	Code string
	ID   string
	Name string
}

func (mg CurrentProgressTitleOutput) outputMarker() {}

func parseCurrentProgressTitleOutput(input string) (*CurrentProgressTitleOutput, error) {
	var currentProgressTitleOutput CurrentProgressTitleOutput

	trimmed, _ := strings.CutPrefix(input, currentProgressTitlePrefix)

	split := strings.Split(trimmed, delimiter)
	if len(split) < 3 {
		return nil, ErrNotEnoughValues
	}

	currentProgressTitleOutput.Code = split[0]
	currentProgressTitleOutput.ID = split[1]
	currentProgressTitleOutput.Name = split[2]

	return &currentProgressTitleOutput, nil
}