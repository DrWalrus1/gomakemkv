package events

import (
	"strconv"
	"strings"

	"github.com/DrWalrus1/gomakemkv/events/codes"
)

type DriveScanMessage struct {
	DriveIndex string
	Visible    string
	Enabled    string
	Flags      string
	DriveName  string
	DiscName   string
	DeviceName string
}

func (mg DriveScanMessage) outputMarker() {}

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
	driveScanVisibibleString, err := codes.GetDriveStateDescription(uint(visibleCode))
	if err != nil {
		return nil, err
	}
	driveScanMessage.Visible = driveScanVisibibleString
	enabledCode, err := strconv.Atoi(split[2])
	if err != nil {
		return nil, err
	}
	driveScanEnabledString, err := codes.GetDriveStateDescription(uint(enabledCode))
	if err != nil {
		return nil, err
	}
	driveScanMessage.Enabled = driveScanEnabledString
	driveScanMessage.Flags = split[3]
	driveScanMessage.DriveName = split[4]
	driveScanMessage.DiscName = split[5]
	driveScanMessage.DeviceName = split[6]
	return &driveScanMessage, nil
}
