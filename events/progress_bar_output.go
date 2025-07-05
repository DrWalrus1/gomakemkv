package events

import "strings"

type ProgressBarOutput struct {
	CurrentProgress string
	TotalProgress   string
	MaxProgress     string
}

func (mg ProgressBarOutput) outputMarker() {}

func parseProgressBarOutput(input string) (*ProgressBarOutput, error) {
	var progressOutput ProgressBarOutput

	trimmed, _ := strings.CutPrefix(input, progressBarOutputPrefix)

	split := strings.Split(trimmed, delimiter)
	if len(split) < 3 {
		return nil, ErrNotEnoughValues
	}
	progressOutput.CurrentProgress = split[0]
	progressOutput.TotalProgress = split[1]
	progressOutput.MaxProgress = split[2]
	return &progressOutput, nil
}