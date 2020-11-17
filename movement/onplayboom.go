package movement

import (
	"fmt"
	"github.com/bwmarrin/dgvoice"
	"github.com/bwmarrin/discordgo"
	"io/ioutil"
)

func OnPlayBoom(s *discordgo.Session, m *discordgo.MessageCreate, nickname string)  {
	SendMessage(s, m.ChannelID, "爆弾起爆!!!!!!!", nickname)
	fmt.Println("Reading Folder : ", *BoonFolder)
	files, _ := ioutil.ReadDir(*BoonFolder)
	for _, f := range files{
		fmt.Println("PlayAudioFile: ", f.Name())
		s.UpdateStatus(0, f.Name())

		filename := fmt.Sprintf("%s/%s", *BoonFolder, f.Name())
		fmt.Println(filename)

		dgvoice.PlayAudioFile(vcsession, filename, stopBot)
	}
}
