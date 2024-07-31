package speaker_service

import (
	speaker_entity "device_service/internal/entity/speaker"
	"fmt"
)

func GenerateMockPopSongs() []*speaker_entity.Channel {
	songNames := []string{
		"Blinding Lights", "Uptown Funk", "Shape of You", "Despacito", "Happy",
		"Shake It Off", "Can't Stop the Feeling!", "Old Town Road", "Sunflower", "Levitating",
		"Rolling in the Deep", "Call Me Maybe", "Poker Face", "Just Dance", "Single Ladies",
		"Bad Guy", "Havana", "Sorry", "All About That Bass", "Roar",
		"Uptown Funk", "Shallow", "Truth Hurts", "Someone Like You", "Locked Out of Heaven",
		"Bad Romance", "Teenage Dream", "Firework", "Moves Like Jagger", "We Belong Together",
	}

	var popSongs []*speaker_entity.Channel

	for i, name := range songNames {
		song := &speaker_entity.Channel{
			ChannelName:   name,
			ChannelNumber: fmt.Sprintf("%02d", i+1),
		}
		popSongs = append(popSongs, song)
	}

	return popSongs
}
