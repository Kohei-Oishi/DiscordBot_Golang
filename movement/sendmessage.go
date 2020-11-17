package movement

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
)

func SendMessage(s *discordgo.Session, channelID string, msg string, nickname string)  {
	_, err := s.ChannelMessageSend(channelID, msg)

	fmt.Println("喋りやがりましたぜ by " + nickname)
	if err != nil {
		fmt.Println("Error sending message: ", err)
	}
}
