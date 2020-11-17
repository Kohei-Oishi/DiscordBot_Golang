package movement

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
)

//メッセージを受信した時の、声の初めと終わりにPrintされるようだ
func OnVoiceReceived(vc *discordgo.VoiceConnection, vs *discordgo.VoiceSpeakingUpdate) {
	fmt.Println("しゃべったあああああ")
}
