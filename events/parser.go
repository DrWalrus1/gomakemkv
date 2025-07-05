package events

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"regexp"
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
