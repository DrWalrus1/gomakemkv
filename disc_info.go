package gomakemkv

import (
	"strconv"

	"github.com/DrWalrus1/gomakemkv/events"
)

type DiscInfo struct {
	Name     string        `json:"name"`
	Language string        `json:"language"`
	Type     string        `json:"type"`
	Titles   map[int]Title `json:"titles"`
}

func (di DiscInfo) GetTypeName() string {
	return "DiscInfoComplete"
}

func NewDisc() DiscInfo {
	return DiscInfo{
		Titles: make(map[int]Title),
	}
}

func (di *DiscInfo) UpdateDiscInfo(info events.DiscInformation) {
	switch info.ID {
	case events.Name:
		di.Name = info.Value
	case events.MetadataLanguageName:
		di.Language = info.Value
	case events.Type:
		di.Type = info.Value
	}
}

func (di *DiscInfo) UpsertDiscTitleMetadata(info events.TitleInformation) {
	title, exists := di.Titles[info.TitleIndex]
	if !exists {
		title = NewTitle(strconv.Itoa(info.TitleIndex))
	}
	title.UpdateTitle(info)
	di.Titles[info.TitleIndex] = title
}

func (di *DiscInfo) UpsertTitleStreamMetadata(info events.StreamInformation) {
	title, exists := di.Titles[info.TitleIndex]
	if !exists {
		title = NewTitle(strconv.Itoa(info.TitleIndex))
	}
	title.UpsertStreamData(info)
	di.Titles[info.TitleIndex] = title
}

type Title struct {
	Id             string                `json:"id"`
	Name           string                `json:"name"`
	Size           string                `json:"size"`
	SizeInBytes    string                `json:"sizeInBytes"`
	Duration       string                `json:"duration"`
	Language       string                `json:"language"`
	Chapters       string                `json:"chapters"`
	OutputFileName string                `json:"outputFileName"`
	VideoTracks    map[int]VideoTrack    `json:"video"`
	AudioTracks    map[int]AudioTrack    `json:"audio"`
	SubtitleTracks map[int]SubtitleTrack `json:"subtitles"`
}

func NewTitle(id string) Title {
	return Title{
		Id:             id,
		VideoTracks:    make(map[int]VideoTrack),
		AudioTracks:    make(map[int]AudioTrack),
		SubtitleTracks: make(map[int]SubtitleTrack),
	}
}

func (t *Title) UpdateTitle(info events.TitleInformation) {
	switch info.AttributeId {
	case events.Name:
		t.Name = info.Value
	case events.DiskSize:
		t.Size = info.Value
	case events.DiskSizeBytes:
		t.SizeInBytes = info.Value
	case events.Duration:
		t.Duration = info.Value
	case events.MetadataLanguageName:
		t.Language = info.Value
	case events.ChapterCount:
		t.Chapters = info.Value
	case events.OutputFileName:
		t.OutputFileName = info.Value
	}
}

func (t *Title) UpsertStreamData(info events.StreamInformation) {
	// Create new stream if the type is detected
	if info.AttributeId == events.Type {
		switch info.Value {
		case "Video":
			t.VideoTracks[info.StreamIndex] = VideoTrack{
				Type: info.Value,
			}
		case "Audio":
			t.AudioTracks[info.StreamIndex] = AudioTrack{
				Type: info.Value,
			}
		default:
			t.SubtitleTracks[info.StreamIndex] = SubtitleTrack{
				Type: info.Value,
			}
		}
	}
	videoTrack, ok := t.VideoTracks[info.StreamIndex]
	if ok {
		videoTrack.UpdateVideoTrack(info)
		t.VideoTracks[info.StreamIndex] = videoTrack
		return
	}
	audioTrack, ok := t.AudioTracks[info.StreamIndex]
	if ok {
		audioTrack.UpdateAudioTrack(info)
		t.AudioTracks[info.StreamIndex] = audioTrack
		return
	}

	subtitleTrack, ok := t.SubtitleTracks[info.StreamIndex]
	if ok {
		subtitleTrack.UpdateSubtitleTrack(info)
		t.SubtitleTracks[info.StreamIndex] = subtitleTrack
		return
	}
	panic("Attempted to parse out of order stream information")
	// TODO: CONSIDER making the array a queue. if we don't find the type first skip for now and re-enqueue

}

type VideoTrack struct {
	Type           string `json:"type"`
	Framerate      string `json:"framerate"`
	VideoSize      string `json:"videoSize"`
	Codec          string `json:"codec"`
	Language       string `json:"language"`
	ConversionType string `json:"conversionType"`
}

func (vt *VideoTrack) UpdateVideoTrack(info events.StreamInformation) {
	switch info.AttributeId {
	case events.VideoFrameRate:
		vt.Framerate = info.Value
	case events.VideoSize:
		vt.VideoSize = info.Value
	case events.CodecShort:
		vt.Codec = info.Value
	case events.MetadataLanguageName:
		vt.Language = info.Value
	case events.OutputConversionType:
		vt.ConversionType = info.Value
	}
}

type AudioTrack struct {
	Type           string `json:"type"`
	Name           string `json:"name"`
	Language       string `json:"language"`
	Bitrate        string `json:"bitrate"`
	SampleRate     string `json:"sampleRate"`
	SampleSize     string `json:"sampleSize"`
	ChannelNumbers string `json:"channelNumbers"`
	ConversionType string `json:"conversionType"`
}

func (vt *AudioTrack) UpdateAudioTrack(info events.StreamInformation) {
	switch info.AttributeId {
	case events.Name:
		vt.Name = info.Value
	case events.MetadataLanguageName:
		vt.Language = info.Value
	case events.Bitrate:
		vt.Bitrate = info.Value
	case events.AudioSampleRate:
		vt.SampleRate = info.Value
	case events.AudioSampleSize:
		vt.SampleSize = info.Value
	case events.AudioChannelsCount:
		vt.ChannelNumbers = info.Value
	case events.OutputConversionType:
		vt.ConversionType = info.Value
	}
}

type SubtitleTrack struct {
	Type           string `json:"type"`
	Language       string `json:"language"`
	Codec          string `json:"codec"`
	ConversionType string `json:"conversionType"`
}

func (vt *SubtitleTrack) UpdateSubtitleTrack(info events.StreamInformation) {
	switch info.AttributeId {
	case events.MetadataLanguageName:
		vt.Language = info.Value
	case events.CodecShort:
		vt.Codec = info.Value
	case events.OutputConversionType:
		vt.ConversionType = info.Value
	}

}

func MakeMkveventsIntoMakeMkvDiscInfo(makemkvevents []events.MakeMkvOutput) DiscInfo {
	mkvDiscInfo := NewDisc()
	for _, x := range makemkvevents {
		if i, ok := x.(*events.DiscInformation); ok {
			mkvDiscInfo.UpdateDiscInfo(*i)
		} else if i, ok := x.(*events.TitleInformation); ok {
			mkvDiscInfo.UpsertDiscTitleMetadata(*i)
		} else if i, ok := x.(*events.StreamInformation); ok {
			mkvDiscInfo.UpsertTitleStreamMetadata(*i)
		}
	}
	return mkvDiscInfo
}
