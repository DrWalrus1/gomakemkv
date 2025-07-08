package codes

import "errors"

const (
	ap_Unknown               uint = 999
	ap_Unknown2              uint = 255
	ap_DriveStateNoDrive     uint = 256
	ap_DriveStateUnmounting  uint = 257
	ap_DriveStateEmptyClosed uint = 0
	ap_DriveStateEmptyOpen   uint = 1
	ap_DriveStateInserted    uint = 2
	ap_DriveStateLoading     uint = 3
)

var ErrUnknownDriveState = errors.New("unknown drive state")

var driveStateDescriptions = map[uint]string{
	ap_DriveStateNoDrive:     "No Drive detected",
	ap_DriveStateUnmounting:  "Drive is unmounting",
	ap_DriveStateEmptyClosed: "Drive is empty and closed",
	ap_DriveStateEmptyOpen:   "Drive is empty and open",
	ap_DriveStateInserted:    "Drive has disc inserted",
	ap_DriveStateLoading:     "Drive is loading",
	ap_Unknown:               "Unknown",
	ap_Unknown2:              "Unknown",
}

func GetDriveStateDescription(id uint) (string, error) {
	if desc, ok := driveStateDescriptions[id]; ok {
		return desc, nil
	}
	return "", ErrUnknownDriveState
}
