package events

import "strings"

type TotalProgressTitleOutput struct {
	Code string
	ID   string
	Name string
}

func (mg TotalProgressTitleOutput) outputMarker() {}

func parseTotalProgressTitleOutput(input string) (*TotalProgressTitleOutput, error) {
	var currentProgressTitleOutput TotalProgressTitleOutput

	trimmed, _ := strings.CutPrefix(input, totalProgressTitlePrefix)

	split := strings.Split(trimmed, delimiter)
	if len(split) < 3 {
		return nil, ErrNotEnoughValues
	}

	currentProgressTitleOutput.Code = split[0]
	currentProgressTitleOutput.ID = split[1]
	currentProgressTitleOutput.Name = split[2]

	return &currentProgressTitleOutput, nil

}
