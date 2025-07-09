package codes

import "errors"

type appItemAttributeId = int

const (
	Unknown                      appItemAttributeId = 0
	Type                         appItemAttributeId = 1
	Name                         appItemAttributeId = 2
	LangCode                     appItemAttributeId = 3
	LangName                     appItemAttributeId = 4
	CodecId                      appItemAttributeId = 5
	CodecShort                   appItemAttributeId = 6
	CodecLong                    appItemAttributeId = 7
	ChapterCount                 appItemAttributeId = 8
	Duration                     appItemAttributeId = 9
	DiskSize                     appItemAttributeId = 10
	DiskSizeBytes                appItemAttributeId = 11
	StreamTypeExtension          appItemAttributeId = 12
	Bitrate                      appItemAttributeId = 13
	AudioChannelsCount           appItemAttributeId = 14
	AngleInfo                    appItemAttributeId = 15
	SourceFileName               appItemAttributeId = 16
	AudioSampleRate              appItemAttributeId = 17
	AudioSampleSize              appItemAttributeId = 18
	VideoSize                    appItemAttributeId = 19
	VideoAspectRatio             appItemAttributeId = 20
	VideoFrameRate               appItemAttributeId = 21
	StreamFlags                  appItemAttributeId = 22
	DateTime                     appItemAttributeId = 23
	OriginalTitleId              appItemAttributeId = 24
	SegmentsCount                appItemAttributeId = 25
	SegmentsMap                  appItemAttributeId = 26
	OutputFileName               appItemAttributeId = 27
	MetadataLanguageCode         appItemAttributeId = 28
	MetadataLanguageName         appItemAttributeId = 29
	TreeInfo                     appItemAttributeId = 30
	PanelTitle                   appItemAttributeId = 31
	VolumeName                   appItemAttributeId = 32
	OrderWeight                  appItemAttributeId = 33
	OutputFormat                 appItemAttributeId = 34
	OutputFormatDescription      appItemAttributeId = 35
	SeamlessInfo                 appItemAttributeId = 36
	PanelText                    appItemAttributeId = 37
	MkvFlags                     appItemAttributeId = 38
	MkvFlagsText                 appItemAttributeId = 39
	AudioChannelLayoutName       appItemAttributeId = 40
	OutputCodecShort             appItemAttributeId = 41
	OutputConversionType         appItemAttributeId = 42
	OutputAudioSampleRate        appItemAttributeId = 43
	OutputAudioSampleSize        appItemAttributeId = 44
	OutputAudioChannelsCount     appItemAttributeId = 45
	OutputAudioChannelLayoutName appItemAttributeId = 46
	OutputAudioChannelLayout     appItemAttributeId = 47
	OutputAudioMixDescription    appItemAttributeId = 48
	Comment                      appItemAttributeId = 49
	OffsetSequenceId             appItemAttributeId = 50
	MaxValue                     appItemAttributeId = 51
)

var ErrUnknownAttribute = errors.New("unknown application item attribute")

var attributeDetailedDescription = map[appItemAttributeId]string{
	Unknown:                      "Unknown",
	Type:                         "Type",
	Name:                         "Name",
	LangCode:                     "LanguageCode",
	LangName:                     "LanguageName",
	CodecId:                      "CodecID",
	CodecShort:                   "ShortCodecName",
	CodecLong:                    "LongCodecName",
	ChapterCount:                 "NumberOfChapters",
	Duration:                     "Duration",
	DiskSize:                     "DiskSize",
	DiskSizeBytes:                "DiskSizeInBytes",
	StreamTypeExtension:          "StreamTypeExtension",
	Bitrate:                      "Bitrate",
	AudioChannelsCount:           "NumberOfAudioChannels",
	AngleInfo:                    "AngleInformation",
	SourceFileName:               "SourceFileName",
	AudioSampleRate:              "AudioSampleRate",
	AudioSampleSize:              "AudioSampleSize",
	VideoSize:                    "VideoSize",
	VideoAspectRatio:             "VideoAspectRatio",
	VideoFrameRate:               "VideoFrameRate",
	StreamFlags:                  "StreamFlags",
	DateTime:                     "DateAndTime",
	OriginalTitleId:              "OriginalTitleID",
	SegmentsCount:                "NumberofSegments",
	SegmentsMap:                  "SegmentsMap",
	OutputFileName:               "OutputFileName",
	MetadataLanguageCode:         "MetadataLanguageCode",
	MetadataLanguageName:         "MetadataLanguageName",
	TreeInfo:                     "TreeInformation",
	PanelTitle:                   "PanelTitle",
	VolumeName:                   "VolumeName",
	OrderWeight:                  "OrderWeight",
	OutputFormat:                 "OutputFormat",
	OutputFormatDescription:      "OutputFormatDescription",
	SeamlessInfo:                 "SeamlessInformation",
	PanelText:                    "PanelText",
	MkvFlags:                     "MKVFlags",
	MkvFlagsText:                 "MKVFlagsText",
	AudioChannelLayoutName:       "AudioChannelLayoutName",
	OutputCodecShort:             "eventshortCodecName",
	OutputConversionType:         "OutputConversionType",
	OutputAudioSampleRate:        "OutputAudioSampleRate",
	OutputAudioSampleSize:        "OutputAudioSampleSize",
	OutputAudioChannelsCount:     "OutputNumberOfAudioChannels",
	OutputAudioChannelLayoutName: "OutputAudioChannelLayout Name",
	OutputAudioChannelLayout:     "OutputAudioChannelLayout",
	OutputAudioMixDescription:    "OutputAudioMixDescription",
	Comment:                      "Comment",
	OffsetSequenceId:             "OffsetSequenceID",
	MaxValue:                     "MaxValue",
}

func GetItemAttributeDescription(id int) (string, error) {
	if desc, ok := attributeDetailedDescription[id]; ok {
		return desc, nil
	}
	return "", ErrUnknownAttribute
}

// CURRENTLY UNUSED CODES

// const (
// 	ap_MaxCdromDevices              int = 16
// 	ap_Progress_MaxValue            int = 65536
// 	ap_Progress_MaxLayoutItems      int = 10
// 	ap_UIMSG_BOX_MASK               int = 3854
// 	ap_UIMSG_BOXOK                  int = 260
// 	ap_UIMSG_BOXERROR               int = 516
// 	ap_UIMSG_BOXWARNING             int = 1028
// 	ap_UIMSG_BOXYESNO               int = 776
// 	ap_UIMSG_BOXYESNO_ERR           int = 1288
// 	ap_UIMSG_BOXYESNO_REG           int = 1544
// 	ap_UIMSG_YES                    int = 0
// 	ap_UIMSG_NO                     int = 1
// 	ap_UIMSG_DEBUG                  int = 32
// 	ap_UIMSG_HIDDEN                 int = 64
// 	ap_UIMSG_EVENT                  int = 128
// 	ap_UIMSG_HAVE_URL               int = 131072
// 	ap_UIMSG_VITEM_BASE             int = 5200
// 	ap_MMBD_DISC_FLAG_BUSENC        int = 2
// 	ap_MMBD_MMBD_DISC_FLAG_AACS     int = 4
// 	ap_MMBD_MMBD_DISC_FLAG_BDPLUS   int = 8
// 	ap_vastr_Name                   int = 0
// 	ap_vastr_Version                int = 1
// 	ap_vastr_Platform               int = 2
// 	ap_vastr_Build                  int = 3
// 	ap_vastr_KeyType                int = 4
// 	ap_vastr_KeyFeatures            int = 5
// 	ap_vastr_KeyExpiration          int = 6
// 	ap_vastr_EvalState              int = 7
// 	ap_vastr_ProgExpiration         int = 8
// 	ap_vastr_LatestVersion          int = 9
// 	ap_vastr_RestartRequired        int = 10
// 	ap_vastr_ExpertMode             int = 11
// 	ap_vastr_ProfileCount           int = 12
// 	ap_vastr_ProgExpired            int = 13
// 	ap_vastr_OutputFolderName       int = 14
// 	ap_vastr_OutputBaseName         int = 15
// 	ap_vastr_CurrentProfile         int = 16
// 	ap_vastr_OpenFileFilter         int = 17
// 	ap_vastr_WebSiteURL             int = 18
// 	ap_vastr_OpenDVDFileFilter      int = 19
// 	ap_vastr_DefaultSelectionString int = 20
// 	ap_vastr_DefaultOutputFileName  int = 21
// 	ap_vastr_ExternalAppItem        int = 22
// 	ap_vastr_InterfaceLanguage      int = 23
// 	ap_vastr_ProfileString          int = 24
// 	ap_vastr_KeyString              int = 25
// )
//
// const (
// 	apNotifyUpdateLayoutFlagNoTime   = 1
// 	apProgressCurrentIndexSourceName = 65280
// 	apBackupFlagDecryptVideo         = 1
// 	apOpenFlagManualMode             = 1
// 	apUpdateDrivesFlagNoScan         = 1
// 	apUpdateDrivesFlagNoSingleDrive  = 2
// )
//
// const (
// 	apAVStreamFlagDirectorsComments          = 1
// 	apAVStreamFlagAlternateDirectorsComments = 2
// 	apAVStreamFlagForVisuallyImpaired        = 4
// 	apAVStreamFlagCoreAudio                  = 256
// 	apAVStreamFlagSecondaryAudio             = 512
// 	apAVStreamFlagHasCoreAudio               = 1024
// 	apAVStreamFlagDerivedStream              = 2048
// 	apAVStreamFlagForcedSubtitles            = 4096
// 	apAVStreamFlagProfileSecondaryStream     = 16384
// 	apAVStreamFlagOffsetSequenceIdPresent    = 32768
// )
//
// const (
// 	apAPPLOCMAX = 7000
// )
