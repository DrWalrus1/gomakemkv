package events

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"regexp"
	"strconv"
	"strings"
)

const (
	messageOutputPrefix        = "MSG:"
	driveScanMessagePrefix     = "DRV:"
	currentProgressTitlePrefix = "PRGC:"
	discInfoOutputPrefix       = "TCOUNT:"
	discInfoPrefix             = "CINFO:"
	progressBarOutputPrefix    = "PRGV:"
	streamInfoPrefix           = "SINFO:"
	titleInfoPrefix            = "TINFO:"
	totalProgressTitlePrefix   = "PRGT:"
	delimiter                  = ","
)

var ErrPrefixNotFound = errors.New("prefix did not match expected")
var ErrNotEnoughValues = errors.New("not enough values found in input")
var ErrEmptyInput = errors.New("input is empty")

func sanitiseString(input string) string {
	input = strings.TrimSpace(input)
	input = strings.ReplaceAll(input, "\"", "")
	boldOrBreakLineRegex := regexp.MustCompile("\u003c(b|/b|br)\u003e")
	input = boldOrBreakLineRegex.ReplaceAllString(input, "")
	return input
}

func ParseEventStream(reader io.Reader) <-chan MakeMkvOutput {
	c := make(chan MakeMkvOutput)
	go func() {
		scanner := bufio.NewScanner(reader)
		for scanner.Scan() {
			output, err := Parse(scanner.Text())
			if err != nil {
				fmt.Println(err.Error())
				continue
			}
			c <- output
		}
		close(c)
	}()
	return c
}

func Parse(input string) (MakeMkvOutput, error) {
	sanitised := sanitiseString(input)
	if sanitised == "" {
		return nil, ErrEmptyInput
	}

	switch {
	case strings.HasPrefix(sanitised, messageOutputPrefix):
		return parseMessageOutput(sanitised)
	case strings.HasPrefix(sanitised, driveScanMessagePrefix):
		return parseDriveScanMessage(sanitised)
	case strings.HasPrefix(sanitised, currentProgressTitlePrefix):
		return parseCurrentProgressTitleOutput(sanitised)
	case strings.HasPrefix(sanitised, discInfoOutputPrefix):
		return parseDiscInformationOutputMessage(sanitised)
	case strings.HasPrefix(sanitised, discInfoPrefix):
		return parseDiscInfo(sanitised)
	case strings.HasPrefix(sanitised, progressBarOutputPrefix):
		return parseProgressBarOutput(sanitised)
	case strings.HasPrefix(sanitised, streamInfoPrefix):
		return parseStreamInfo(sanitised)
	case strings.HasPrefix(sanitised, titleInfoPrefix):
		return parseTitleInfo(sanitised)
	case strings.HasPrefix(sanitised, totalProgressTitlePrefix):
		return parseTotalProgressTitleOutput(sanitised)
	}
	return nil, ErrPrefixNotFound
}

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

func parseDriveScanMessage(input string) (*DriveScanMessage, error) {
	var driveScanMessage DriveScanMessage

	trimmed, _ := strings.CutPrefix(input, driveScanMessagePrefix)

	split := strings.Split(trimmed, delimiter)
	if len(split) < 7 {
		return nil, ErrNotEnoughValues
	}

	driveScanMessage.DriveIndex = split[0]
	visibleCode, err := strconv.Atoi(split[1])
	if err != nil {
		return nil, err
	}
	driveScanMessage.Visible = GetDriveStateDescription(uint(visibleCode))
	enabledCode, err := strconv.Atoi(split[2])
	if err != nil {
		return nil, err
	}
	driveScanMessage.Enabled = GetDriveStateDescription(uint(enabledCode))
	driveScanMessage.Flags = split[3]
	driveScanMessage.DriveName = split[4]
	driveScanMessage.DiscName = split[5]
	driveScanMessage.DeviceName = split[6]
	return &driveScanMessage, nil
}

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
