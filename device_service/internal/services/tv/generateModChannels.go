package tv_service

import (
	tv_entity "device_service/internal/entity/tv"
	"fmt"
)

func GenerateMockChannels() []*tv_entity.Channel {
	channelNames := []string{
		"BBC News", "CNN", "Al Jazeera", "Fox News", "Sky News",
		"MSNBC", "Euronews", "France 24", "RT", "Bloomberg",
		"CNBC", "ABC News", "CBS News", "NBC News", "ITV News",
		"Channel 4 News", "RTÉ News", "Telemundo", "Univision", "CBC News",
		"Global News", "CTV News", "Newsmax", "One America News", "Cheddar",
		"Newsy", "NewsNation", "News 12", "NHK World", "Deutsche Welle",
		"TRT World", "Arirang", "CGTN", "Press TV", "India Today",
		"WION", "NDTV", "Zee News", "Aaj Tak", "Republic TV",
		"TV5Monde", "BFMTV", "EWTN", "Telesur", "ABC Australia",
		"7 News", "SBS World News", "Ten News", "Nine News", "The Weather Channel",
	}

	var channels []*tv_entity.Channel

	for i, name := range channelNames {
		channel := &tv_entity.Channel{
			ChannelName:   name,
			ChannelNumber: fmt.Sprintf("%02d", i+1),
		}
		channels = append(channels, channel)
	}

	return channels
}
