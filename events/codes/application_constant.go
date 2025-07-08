package codes

import "errors"

const (
	appDumpDonePartial                        int = 5004
	appDumpDone                               int = 5005
	appInitFailed                             int = 5009
	appAskFolderCreate                        int = 5013
	appFolderInvalid                          int = 5016
	progressAppSaveMkvFreeSpace               int = 5033
	protDemoKeyExpired                        int = 5021
	appKeytypeInvalid                         int = 5095
	appEvalTimeNever                          int = 5067
	appBackupFailed                           int = 5069
	appBackupCompleted                        int = 5070
	appBackupCompletedHashfail                int = 5079
	profileNameDefault                        int = 5086
	vitemName                                 int = 5202
	vitemTimestamp                            int = 5223
	appIfaceTitle                             int = 6000
	appCaptionMsg                             int = 6001
	appAboutboxTitle                          int = 6002
	appIfaceOpenfileTitle                     int = 6003
	appSettingdlgTitle                        int = 6135
	appBackupdlgTitle                         int = 6136
	appIfaceOpenfileFilterTemplate1           int = 6007
	appIfaceOpenfileFilterTemplate2           int = 6008
	appIfaceOpenfolderTitle                   int = 6005
	appIfaceOpenfolderInfoTitle               int = 6006
	appIfaceProgressTitle                     int = 6038
	appIfaceProgressElapsedOnly               int = 6039
	appIfaceProgressElapsedEta                int = 6040
	appIfaceActOpenfilesName                  int = 6010
	appIfaceActOpenfilesSkey                  int = 6011
	appIfaceActOpenfilesStip                  int = 6012
	appIfaceActOpenfilesDvdName               int = 6024
	appIfaceActOpenfilesDvdStip               int = 6026
	appIfaceActClosediskName                  int = 6013
	appIfaceActClosediskStip                  int = 6014
	appIfaceActSetfolderName                  int = 6015
	appIfaceActSetfolderStip                  int = 6016
	appIfaceActSaveallmkvName                 int = 6017
	appIfaceActSaveallmkvStip                 int = 6018
	appIfaceActCancelName                     int = 6036
	appIfaceActCancelStip                     int = 6037
	appIfaceActStreamingName                  int = 6131
	appIfaceActStreamingStip                  int = 6132
	appIfaceActBackupName                     int = 6133
	appIfaceActBackupStip                     int = 6134
	appIfaceActQuitName                       int = 6019
	appIfaceActQuitSkey                       int = 6020
	appIfaceActQuitStip                       int = 6021
	appIfaceActAboutName                      int = 6022
	appIfaceActAboutStip                      int = 6023
	appIfaceActSettingsName                   int = 6042
	appIfaceActSettingsStip                   int = 6043
	appIfaceActHelppageName                   int = 6045
	appIfaceActHelppageStip                   int = 6046
	appIfaceActRegisterName                   int = 6047
	appIfaceActRegisterStip                   int = 6048
	appIfaceActPurchaseName                   int = 6145
	appIfaceActPurchaseStip                   int = 6146
	appIfaceActClearlogName                   int = 6110
	appIfaceActClearlogStip                   int = 6111
	appIfaceActEjectName                      int = 6052
	appIfaceActEjectStip                      int = 6053
	appIfaceActRevertName                     int = 6105
	appIfaceActRevertStip                     int = 6106
	appIfaceActNewinstanceName                int = 6107
	appIfaceActNewinstanceStip                int = 6108
	appIfaceActOpendiscDvd                    int = 6062
	appIfaceActOpendiscHddvd                  int = 6063
	appIfaceActOpendiscBray                   int = 6064
	appIfaceActOpendiscLoading                int = 6065
	appIfaceActOpendiscUnknown                int = 6099
	appIfaceActOpendiscNodisc                 int = 6109
	appIfaceActTtreeToggle                    int = 6066
	appIfaceActTtreeSelectAll                 int = 6067
	appIfaceActTtreeUnselectAll               int = 6068
	appIfaceMenuFile                          int = 6030
	appIfaceMenuView                          int = 6031
	appIfaceMenuHelp                          int = 6032
	appIfaceMenuToolbar                       int = 6034
	appIfaceMenuSettings                      int = 6044
	appIfaceMenuDrives                        int = 6035
	appIfaceCancelConfirm                     int = 6041
	appIfaceFatalComm                         int = 6050
	appIfaceFatalMem                          int = 6051
	appIfaceGuiVersion                        int = 6054
	appIfaceLatestVersion                     int = 6158
	appIfaceLicenseType                       int = 6055
	appIfaceEvalState                         int = 6056
	appIfaceEvalExpiration                    int = 6057
	appIfaceProgExpiration                    int = 6142
	appIfaceWebsiteUrl                        int = 6159
	appIfaceVideoFolderNameWin                int = 6058
	appIfaceVideoFolderNameMac                int = 6059
	appIfaceVideoFolderNameLinux              int = 6060
	appIfaceDefaultFolderName                 int = 6061
	appIfaceMainFrameInfo                     int = 6069
	appIfaceMainFrameMakeMkv                  int = 6070
	appIfaceMainFrameProfile                  int = 6180
	appIfaceMainFrameProperties               int = 6181
	appIfaceEmptyFrameInfo                    int = 6075
	appIfaceEmptyFrameSource                  int = 6071
	appIfaceEmptyFrameType                    int = 6072
	appIfaceEmptyFrameLabel                   int = 6073
	appIfaceEmptyFrameProtection              int = 6074
	appIfaceEmptyFrameDvdManual               int = 6084
	appIfaceRegisterText                      int = 6076
	appIfaceRegisterCodeIncorrect             int = 6077
	appIfaceRegisterCodeNotSaved              int = 6078
	appIfaceRegisterCodeSaved                 int = 6079
	appIfaceSettingsIoOptions                 int = 6080
	appIfaceSettingsIoAuto                    int = 6081
	appIfaceSettingsIoReadRetry               int = 6082
	appIfaceSettingsIoReadBuffer              int = 6083
	appIfaceSettingsIoNoDirectAccess          int = 6150
	appIfaceSettingsIoDarwinK2Workaround      int = 6151
	appIfaceSettingsIoSingleDrive             int = 6168
	appIfaceSettingsDvdAuto                   int = 6085
	appIfaceSettingsDvdMinLength              int = 6086
	appIfaceSettingsDvdSpRemove               int = 6087
	appIfaceSettingsAacsKeyDir                int = 6088
	appIfaceSettingsBdpMisc                   int = 6129
	appIfaceSettingsBdpDumpAlways             int = 6130
	appIfaceSettingsDestTypeNone              int = 6089
	appIfaceSettingsDestTypeAuto              int = 6090
	appIfaceSettingsDestTypeSemiauto          int = 6091
	appIfaceSettingsDestTypeCustom            int = 6092
	appIfaceSettingsDestdir                   int = 6093
	appIfaceSettingsGeneralMisc               int = 6094
	appIfaceSettingsLogDebugMsg               int = 6095
	appIfaceSettingsDataDir                   int = 6167
	appIfaceSettingsExpertMode                int = 6169
	appIfaceSettingsShowAvsync                int = 6170
	appIfaceSettingsGeneralOnlineUpdates      int = 6188
	appIfaceSettingsEnableInternetAccess      int = 6187
	appIfaceSettingsProxyServer               int = 6189
	appIfaceSettingsTabGeneral                int = 6096
	appIfaceSettingsMsgFailed                 int = 6097
	appIfaceSettingsMsgRestart                int = 6098
	appIfaceSettingsTabLanguage               int = 6152
	appIfaceSettingsLangInterface             int = 6153
	appIfaceSettingsLangPreferred             int = 6154
	appIfaceSettingsLanguageAuto              int = 6156
	appIfaceSettingsLanguageNone              int = 6157
	appIfaceSettingsTabIo                     int = 6164
	appIfaceSettingsTabStreaming              int = 6165
	appIfaceSettingsTabProt                   int = 6166
	appIfaceSettingsTabAdvanced               int = 6172
	appIfaceSettingsAdvDefaultProfile         int = 6173
	appIfaceSettingsAdvDefaultSelection       int = 6174
	appIfaceSettingsAdvExternExecPath         int = 6175
	appIfaceSettingsProtJavaPath              int = 6177
	appIfaceSettingsAdvOutputFileNameTemplate int = 6178
	appIfaceSettingsTabIntegration            int = 6190
	appIfaceSettingsIntText                   int = 6191
	appIfaceSettingsIntHdrPath                int = 6192
	appIfaceKeyText                           int = 6179
	appIfaceKeyName                           int = 6182
	appIfaceKeyType                           int = 6183
	appIfaceKeyDate                           int = 6184
	appIfaceBackupdlgTextCaption              int = 6137
	appIfaceBackupdlgText                     int = 6138
	appIfaceBackupdlgFolder                   int = 6139
	appIfaceBackupdlgOptions                  int = 6147
	appIfaceBackupdlgDecrypt                  int = 6148
	appIfaceDriveinfoLoading                  int = 6100
	appIfaceDriveinfoUnmounting               int = 6112
	appIfaceDriveinfoWait                     int = 6101
	appIfaceDriveinfoNodisc                   int = 6102
	appIfaceDriveinfoDatadisc                 int = 6103
	appIfaceDriveinfoNone                     int = 6104
	appIfaceFlagsDirectorsComments            int = 6125
	appIfaceFlagsAltDirectorsComments         int = 6126
	appIfaceFlagsSecondaryAudio               int = 6127
	appIfaceFlagsForVisuallyImpaired          int = 6128
	appIfaceFlagsCoreAudio                    int = 6143
	appIfaceFlagsForcedSubtitles              int = 6144
	appIfaceFlagsProfileSecondaryStream       int = 6171
	appIfaceIteminfoSource                    int = 6119
	appIfaceIteminfoTitle                     int = 6120
	appIfaceIteminfoTrack                     int = 6121
	appIfaceIteminfoAttachment                int = 6122
	appIfaceIteminfoChapter                   int = 6123
	appIfaceIteminfoChapters                  int = 6124
	appTtreeTitle                             int = 6200
	appTtreeVideo                             int = 6201
	appTtreeAudio                             int = 6202
	appTtreeSubpicture                        int = 6203
	appTtreeAttachment                        int = 6214
	appTtreeChapters                          int = 6215
	appTtreeChapter                           int = 6216
	appTtreeForcedSubtitles                   int = 6211
	appTtreeHdrType                           int = 6204
	appTtreeHdrDesc                           int = 6205
	dvdTypeDisk                               int = 6206
	brayTypeDisk                              int = 6209
	hddvdTypeDisk                             int = 6212
	mkvTypeFile                               int = 6213
	appTtreeChapDesc                          int = 6207
	appTtreeAngleDesc                         int = 6210
	appDvdManualTitle                         int = 6220
	appDvdManualText                          int = 6225
	appDvdTitlesCount                         int = 6221
	appDvdCountCells                          int = 6222
	appDvdCountPgc                            int = 6223
	appDvdBrokenTitleEntry                    int = 6224
	appSingleDriveTitle                       int = 6226
	appSingleDriveText                        int = 6227
	appSingleDriveAll                         int = 6228
	appSingleDriveCaption                     int = 6229
	appSiDriveinfo                            int = 6300
	appSiProfile                              int = 6301
	appSiManufacturer                         int = 6302
	appSiProduct                              int = 6303
	appSiRevision                             int = 6304
	appSiSerial                               int = 6305
	appSiFirmware                             int = 6306
	appSiFirdate                              int = 6307
	appSiBecflags                             int = 6308
	appSiHighestAacs                          int = 6309
	appSiDiscinfo                             int = 6320
	appSiNodisc                               int = 6321
	appSiDiscload                             int = 6322
	appSiCapacity                             int = 6323
	appSiDisctype                             int = 6324
	appSiDiscsize                             int = 6325
	appSiDiscrate                             int = 6326
	appSiDisclayers                           int = 6327
	appSiDisccbl                              int = 6329
	appSiDisccbl25                            int = 6330
	appSiDisccbl27                            int = 6331
	appSiDevice                               int = 6332
)

var ErrUnknownConstant = errors.New("unknown application item constant")

var appConstantsDescriptions = map[int]string{
	appDumpDonePartial:                        "Partial dump done",
	appDumpDone:                               "Dump done",
	appInitFailed:                             "Initialization failed",
	appAskFolderCreate:                        "Ask to create folder",
	appFolderInvalid:                          "Invalid folder",
	progressAppSaveMkvFreeSpace:               "Save MKV free space progress",
	protDemoKeyExpired:                        "Demo key expired",
	appKeytypeInvalid:                         "Invalid key type",
	appEvalTimeNever:                          "Evaluation time never",
	appBackupFailed:                           "Backup failed",
	appBackupCompleted:                        "Backup completed",
	appBackupCompletedHashfail:                "Backup completed with hash failure",
	profileNameDefault:                        "Default profile name",
	vitemName:                                 "Virtual item name",
	vitemTimestamp:                            "Virtual item timestamp",
	appIfaceTitle:                             "Interface title",
	appCaptionMsg:                             "Caption message",
	appAboutboxTitle:                          "About box title",
	appIfaceOpenfileTitle:                     "Open file title",
	appSettingdlgTitle:                        "Settings dialog title",
	appBackupdlgTitle:                         "Backup dialog title",
	appIfaceOpenfileFilterTemplate1:           "Open file filter template 1",
	appIfaceOpenfileFilterTemplate2:           "Open file filter template 2",
	appIfaceOpenfolderTitle:                   "Open folder title",
	appIfaceOpenfolderInfoTitle:               "Open folder info title",
	appIfaceProgressTitle:                     "Progress title",
	appIfaceProgressElapsedOnly:               "Progress elapsed only",
	appIfaceProgressElapsedEta:                "Progress elapsed and ETA",
	appIfaceActOpenfilesName:                  "Open files action name",
	appIfaceActOpenfilesSkey:                  "Open files shortcut key",
	appIfaceActOpenfilesStip:                  "Open files tooltip",
	appIfaceActOpenfilesDvdName:               "Open DVD files action name",
	appIfaceActOpenfilesDvdStip:               "Open DVD files tooltip",
	appIfaceActClosediskName:                  "Close disk action name",
	appIfaceActClosediskStip:                  "Close disk tooltip",
	appIfaceActSetfolderName:                  "Set folder action name",
	appIfaceActSetfolderStip:                  "Set folder tooltip",
	appIfaceActSaveallmkvName:                 "Save all MKV action name",
	appIfaceActSaveallmkvStip:                 "Save all MKV tooltip",
	appIfaceActCancelName:                     "Cancel action name",
	appIfaceActCancelStip:                     "Cancel tooltip",
	appIfaceActStreamingName:                  "Streaming action name",
	appIfaceActStreamingStip:                  "Streaming tooltip",
	appIfaceActBackupName:                     "Backup action name",
	appIfaceActBackupStip:                     "Backup tooltip",
	appIfaceActQuitName:                       "Quit action name",
	appIfaceActQuitSkey:                       "Quit shortcut key",
	appIfaceActQuitStip:                       "Quit tooltip",
	appIfaceActAboutName:                      "About action name",
	appIfaceActAboutStip:                      "About tooltip",
	appIfaceActSettingsName:                   "Settings action name",
	appIfaceActSettingsStip:                   "Settings tooltip",
	appIfaceActHelppageName:                   "Help page action name",
	appIfaceActHelppageStip:                   "Help page tooltip",
	appIfaceActRegisterName:                   "Register action name",
	appIfaceActRegisterStip:                   "Register tooltip",
	appIfaceActPurchaseName:                   "Purchase action name",
	appIfaceActPurchaseStip:                   "Purchase tooltip",
	appIfaceActClearlogName:                   "Clear log action name",
	appIfaceActClearlogStip:                   "Clear log tooltip",
	appIfaceActEjectName:                      "Eject action name",
	appIfaceActEjectStip:                      "Eject tooltip",
	appIfaceActRevertName:                     "Revert action name",
	appIfaceActRevertStip:                     "Revert tooltip",
	appIfaceActNewinstanceName:                "New instance action name",
	appIfaceActNewinstanceStip:                "New instance tooltip",
	appIfaceActOpendiscDvd:                    "Open DVD disc action",
	appIfaceActOpendiscHddvd:                  "Open HD DVD disc action",
	appIfaceActOpendiscBray:                   "Open Blu-ray disc action",
	appIfaceActOpendiscLoading:                "Open disc loading action",
	appIfaceActOpendiscUnknown:                "Open unknown disc action",
	appIfaceActOpendiscNodisc:                 "Open no disc action",
	appIfaceActTtreeToggle:                    "Toggle tree action",
	appIfaceActTtreeSelectAll:                 "Select all tree action",
	appIfaceActTtreeUnselectAll:               "Unselect all tree action",
	appIfaceMenuFile:                          "File menu",
	appIfaceMenuView:                          "View menu",
	appIfaceMenuHelp:                          "Help menu",
	appIfaceMenuToolbar:                       "Toolbar menu",
	appIfaceMenuSettings:                      "Settings menu",
	appIfaceMenuDrives:                        "Drives menu",
	appIfaceCancelConfirm:                     "Cancel confirmation",
	appIfaceFatalComm:                         "Fatal communication error",
	appIfaceFatalMem:                          "Fatal memory error",
	appIfaceGuiVersion:                        "GUI version",
	appIfaceLatestVersion:                     "Latest version",
	appIfaceLicenseType:                       "License type",
	appIfaceEvalState:                         "Evaluation state",
	appIfaceEvalExpiration:                    "Evaluation expiration",
	appIfaceProgExpiration:                    "Program expiration",
	appIfaceWebsiteUrl:                        "Website URL",
	appIfaceVideoFolderNameWin:                "Video folder name (Windows)",
	appIfaceVideoFolderNameMac:                "Video folder name (Mac)",
	appIfaceVideoFolderNameLinux:              "Video folder name (Linux)",
	appIfaceDefaultFolderName:                 "Default folder name",
	appIfaceMainFrameInfo:                     "Main frame info",
	appIfaceMainFrameMakeMkv:                  "Main frame MakeMKV",
	appIfaceMainFrameProfile:                  "Main frame profile",
	appIfaceMainFrameProperties:               "Main frame properties",
	appIfaceEmptyFrameInfo:                    "Empty frame info",
	appIfaceEmptyFrameSource:                  "Empty frame source",
	appIfaceEmptyFrameType:                    "Empty frame type",
	appIfaceEmptyFrameLabel:                   "Empty frame label",
	appIfaceEmptyFrameProtection:              "Empty frame protection",
	appIfaceEmptyFrameDvdManual:               "Empty frame DVD manual",
	appIfaceRegisterText:                      "Register text",
	appIfaceRegisterCodeIncorrect:             "Register code incorrect",
	appIfaceRegisterCodeNotSaved:              "Register code not saved",
	appIfaceRegisterCodeSaved:                 "Register code saved",
	appIfaceSettingsIoOptions:                 "Settings IO options",
	appIfaceSettingsIoAuto:                    "Settings IO auto",
	appIfaceSettingsIoReadRetry:               "Settings IO read retry",
	appIfaceSettingsIoReadBuffer:              "Settings IO read buffer",
	appIfaceSettingsIoNoDirectAccess:          "Settings IO no direct access",
	appIfaceSettingsIoDarwinK2Workaround:      "Settings IO Darwin K2 workaround",
	appIfaceSettingsIoSingleDrive:             "Settings IO single drive",
	appIfaceSettingsDvdAuto:                   "Settings DVD auto",
	appIfaceSettingsDvdMinLength:              "Settings DVD minimum length",
	appIfaceSettingsDvdSpRemove:               "Settings DVD SP remove",
	appIfaceSettingsAacsKeyDir:                "Settings AACS key directory",
	appIfaceSettingsBdpMisc:                   "Settings BDP miscellaneous",
	appIfaceSettingsBdpDumpAlways:             "Settings BDP dump always",
	appIfaceSettingsDestTypeNone:              "Settings destination type none",
	appIfaceSettingsDestTypeAuto:              "Settings destination type auto",
	appIfaceSettingsDestTypeSemiauto:          "Settings destination type semi-auto",
	appIfaceSettingsDestTypeCustom:            "Settings destination type custom",
	appIfaceSettingsDestdir:                   "Settings destination directory",
	appIfaceSettingsGeneralMisc:               "Settings general miscellaneous",
	appIfaceSettingsLogDebugMsg:               "Settings log debug messages",
	appIfaceSettingsDataDir:                   "Settings data directory",
	appIfaceSettingsExpertMode:                "Settings expert mode",
	appIfaceSettingsShowAvsync:                "Settings show AV sync",
	appIfaceSettingsGeneralOnlineUpdates:      "Settings general online updates",
	appIfaceSettingsEnableInternetAccess:      "Settings enable internet access",
	appIfaceSettingsProxyServer:               "Settings proxy server",
	appIfaceSettingsTabGeneral:                "Settings tab general",
	appIfaceSettingsMsgFailed:                 "Settings message failed",
	appIfaceSettingsMsgRestart:                "Settings message restart",
	appIfaceSettingsTabLanguage:               "Settings tab language",
	appIfaceSettingsLangInterface:             "Settings language interface",
	appIfaceSettingsLangPreferred:             "Settings language preferred",
	appIfaceSettingsLanguageAuto:              "Settings language auto",
	appIfaceSettingsLanguageNone:              "Settings language none",
	appIfaceSettingsTabIo:                     "Settings tab IO",
	appIfaceSettingsTabStreaming:              "Settings tab streaming",
	appIfaceSettingsTabProt:                   "Settings tab protection",
	appIfaceSettingsTabAdvanced:               "Settings tab advanced",
	appIfaceSettingsAdvDefaultProfile:         "Settings advanced default profile",
	appIfaceSettingsAdvDefaultSelection:       "Settings advanced default selection",
	appIfaceSettingsAdvExternExecPath:         "Settings advanced external execution path",
	appIfaceSettingsProtJavaPath:              "Settings protection Java path",
	appIfaceSettingsAdvOutputFileNameTemplate: "Settings advanced output file name template",
	appIfaceSettingsTabIntegration:            "Settings tab integration",
	appIfaceSettingsIntText:                   "Settings integration text",
	appIfaceSettingsIntHdrPath:                "Settings integration HDR path",
	appIfaceKeyText:                           "Key text",
	appIfaceKeyName:                           "Key name",
	appIfaceKeyType:                           "Key type",
	appIfaceKeyDate:                           "Key date",
	appIfaceBackupdlgTextCaption:              "Backup dialog text caption",
	appIfaceBackupdlgText:                     "Backup dialog text",
	appIfaceBackupdlgFolder:                   "Backup dialog folder",
	appIfaceBackupdlgOptions:                  "Backup dialog options",
	appIfaceBackupdlgDecrypt:                  "Backup dialog decrypt",
	appIfaceDriveinfoLoading:                  "Drive info loading",
	appIfaceDriveinfoUnmounting:               "Drive info unmounting",
	appIfaceDriveinfoWait:                     "Drive info wait",
	appIfaceDriveinfoNodisc:                   "Drive info no disc",
	appIfaceDriveinfoDatadisc:                 "Drive info data disc",
	appIfaceDriveinfoNone:                     "Drive info none",
	appIfaceFlagsDirectorsComments:            "Flags director's comments",
	appIfaceFlagsAltDirectorsComments:         "Flags alternate director's comments",
	appIfaceFlagsSecondaryAudio:               "Flags secondary audio",
	appIfaceFlagsForVisuallyImpaired:          "Flags for visually impaired",
	appIfaceFlagsCoreAudio:                    "Flags core audio",
	appIfaceFlagsForcedSubtitles:              "Flags forced subtitles",
	appIfaceFlagsProfileSecondaryStream:       "Flags profile secondary stream",
	appIfaceIteminfoSource:                    "Item info source",
	appIfaceIteminfoTitle:                     "Item info title",
	appIfaceIteminfoTrack:                     "Item info track",
	appIfaceIteminfoAttachment:                "Item info attachment",
	appIfaceIteminfoChapter:                   "Item info chapter",
	appIfaceIteminfoChapters:                  "Item info chapters",
	appTtreeTitle:                             "Tree title",
	appTtreeVideo:                             "Tree video",
	appTtreeAudio:                             "Tree audio",
	appTtreeSubpicture:                        "Tree subpicture",
	appTtreeAttachment:                        "Tree attachment",
	appTtreeChapters:                          "Tree chapters",
	appTtreeChapter:                           "Tree chapter",
	appTtreeForcedSubtitles:                   "Tree forced subtitles",
	appTtreeHdrType:                           "Tree HDR type",
	appTtreeHdrDesc:                           "Tree HDR description",
	dvdTypeDisk:                               "DVD type disk",
	brayTypeDisk:                              "Blu-ray type disk",
	hddvdTypeDisk:                             "HD DVD type disk",
	mkvTypeFile:                               "MKV type file",
	appTtreeChapDesc:                          "Tree chapter description",
	appTtreeAngleDesc:                         "Tree angle description",
	appDvdManualTitle:                         "DVD manual title",
	appDvdManualText:                          "DVD manual text",
	appDvdTitlesCount:                         "DVD titles count",
	appDvdCountCells:                          "DVD count cells",
	appDvdCountPgc:                            "DVD count PGC",
	appDvdBrokenTitleEntry:                    "DVD broken title entry",
	appSingleDriveTitle:                       "Single drive title",
	appSingleDriveText:                        "Single drive text",
	appSingleDriveAll:                         "Single drive all",
	appSingleDriveCaption:                     "Single drive caption",
	appSiDriveinfo:                            "SI drive info",
	appSiProfile:                              "SI profile",
	appSiManufacturer:                         "SI manufacturer",
	appSiProduct:                              "SI product",
	appSiRevision:                             "SI revision",
	appSiSerial:                               "SI serial",
	appSiFirmware:                             "SI firmware",
	appSiFirdate:                              "SI firmware date",
	appSiBecflags:                             "SI BEC flags",
	appSiHighestAacs:                          "SI highest AACS",
	appSiDiscinfo:                             "SI disc info",
	appSiNodisc:                               "SI no disc",
	appSiDiscload:                             "SI disc load",
	appSiCapacity:                             "SI capacity",
	appSiDisctype:                             "SI disc type",
	appSiDiscsize:                             "SI disc size",
	appSiDiscrate:                             "SI disc rate",
	appSiDisclayers:                           "SI disc layers",
	appSiDisccbl:                              "SI disc CBL",
	appSiDisccbl25:                            "SI disc CBL 25",
	appSiDisccbl27:                            "SI disc CBL 27",
	appSiDevice:                               "SI device",
}

func GetAppConstantDescription(id int) (string, error) {
	if desc, ok := appConstantsDescriptions[id]; ok {
		return desc, nil
	}
	return "", ErrUnknownConstant
}
