package movement

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
)

func VoiceChannelSearch(s *discordgo.Session, userID string) (string)  {
	for _, g := range s.State.Guilds{
		for _, voiceStates := range g.VoiceStates{
			fmt.Println(voiceStates.UserID)

			if userID == voiceStates.UserID {
				return voiceStates.ChannelID
			}

		}
	}

	//return DemoChannelID
	return ""
}
