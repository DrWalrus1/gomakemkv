package codes

import "errors"

const (
	ap_DskFsFlagDvdFilesPresent    int = 1
	ap_DskFsFlagHdvdFilesPresent   int = 2
	ap_DskFsFlagBlurayFilesPresent int = 4
	ap_DskFsFlagAacsFilesPresent   int = 8
	ap_DskFsFlagBdsvmFilesPresent  int = 16
)

var ErrUnknownDiskFileFlag = errors.New("unknown disk file flag")

var diskFileFlagsDescriptions = map[int]string{
	ap_DskFsFlagDvdFilesPresent:    "DVD files present on disk",
	ap_DskFsFlagHdvdFilesPresent:   "HD DVD files present on disk",
	ap_DskFsFlagBlurayFilesPresent: "Blu-ray files present on disk",
	ap_DskFsFlagAacsFilesPresent:   "Aacs files present on disk",
	ap_DskFsFlagBdsvmFilesPresent:  "Blu-ray disc movie folder files present on disk",
}

func GetDiskFileFlagDescription(id int) (string, error) {
	if desc, ok := diskFileFlagsDescriptions[id]; ok {
		return desc, nil
	}
	return "", ErrUnknownDiskFileFlag
}
