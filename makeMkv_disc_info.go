package gomakemkv

import "github.com/DrWalrus1/gomakemkv/events"
import "github.com/DrWalrus1/gomakemkv/events/codes"

type MakeMkvDiscInfo struct {
	Properties map[string]MakeMkvValue `json:"properties"`
	Titles     map[int]MakeMkvTitle    `json:"titles"`
}

func (mkvDiscInfo *MakeMkvDiscInfo) addDiscInfoVerbose(discInfo events.DiscInformation) {
	desc, _ := codes.GetItemAttributeDescription(discInfo.ID)
	messagecode, err := codes.GetAppConstantDescription(discInfo.MessageCodeId)
	if err != nil {
		messagecode = ""
	}
	mkvValue := MakeMkvValue{
		MessageCodeValue: messagecode,
		Value:            discInfo.Value,
	}
	mkvDiscInfo.Properties[desc] = mkvValue
}

func (discInfo *MakeMkvDiscInfo) addTitleInformationVerbose(titleInfo events.TitleInformation) {
	if _, ok := discInfo.Titles[titleInfo.TitleIndex]; !ok {
		discInfo.Titles[titleInfo.TitleIndex] = MakeMkvTitle{
			Properties: make(map[string]MakeMkvValue),
			Streams:    make(map[int]map[string]MakeMkvValue),
		}
	}
	desc, _ := codes.GetItemAttributeDescription(titleInfo.AttributeId)
	messagecode, err := codes.GetAppConstantDescription(titleInfo.MessageCodeId)
	if err != nil {
		messagecode = ""
	}
	mkvValue := MakeMkvValue{
		MessageCodeValue: messagecode,
		Value:            titleInfo.Value,
	}
	discInfo.Titles[titleInfo.TitleIndex].Properties[desc] = mkvValue
}

func (discInfo *MakeMkvDiscInfo) addStreamInformationVerbose(streamInfo events.StreamInformation) {
	if streamInfo.Value == "" {
		return
	}
	if _, ok := discInfo.Titles[streamInfo.TitleIndex]; !ok {
		discInfo.Titles[streamInfo.TitleIndex] = MakeMkvTitle{
			Properties: make(map[string]MakeMkvValue),
			Streams:    make(map[int]map[string]MakeMkvValue),
		}
	}
	// if there are streams but not this one in particular
	if _, ok := discInfo.Titles[streamInfo.TitleIndex].Streams[streamInfo.StreamIndex]; !ok {
		discInfo.Titles[streamInfo.TitleIndex].Streams[streamInfo.StreamIndex] = make(map[string]MakeMkvValue)
	}
	desc, _ := codes.GetItemAttributeDescription(streamInfo.AttributeId)
	messagecode, err := codes.GetAppConstantDescription(streamInfo.MessageCodeId)
	if err != nil {
		messagecode = ""
	}
	mkvValue := MakeMkvValue{
		MessageCodeValue: messagecode,
		Value:            streamInfo.Value,
	}
	discInfo.Titles[streamInfo.TitleIndex].Streams[streamInfo.StreamIndex][desc] = mkvValue
}

func MakeMkveventsIntoMakeMkvDiscInfoVerbose(makemkvevents []events.MakeMkvOutput) MakeMkvDiscInfo {
	mkvDiscInfo := MakeMkvDiscInfo{
		Properties: make(map[string]MakeMkvValue),
		Titles:     make(map[int]MakeMkvTitle),
	}
	for _, x := range makemkvevents {
		if i, ok := x.(*events.DiscInformation); ok {
			mkvDiscInfo.addDiscInfoVerbose(*i)
		} else if i, ok := x.(*events.TitleInformation); ok {
			mkvDiscInfo.addTitleInformationVerbose(*i)
		} else if i, ok := x.(*events.StreamInformation); ok {
			mkvDiscInfo.addStreamInformationVerbose(*i)
		}
	}
	return mkvDiscInfo
}
