package movement

import (
	"Go-discord/algorithm"
	"flag"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"time"
)

var (
	DemoChannelID = "" // ChannelIDを入れる
	stopBot = make(chan bool)
	explosiontimer = make(chan bool)

	vcsession *discordgo.VoiceConnection
	BoonFolder = flag.String("b", "soundeffect", "Folder of files to play.")
)

const (
	greatjob = "おつかれさま"
	startfan = "センプウキつけて"
	botsorry = "お前うるさいらしいぞ"
	explosion = "爆発させて"
	join1, join2, join3, join4 = "こい", "きて", "来て", "来い"
	teachchannel1, teachchannel2 = "このチャンネルのことを教えて", "このチャンネルのこと教えて"
	leave1, leave2, leave3, leave4 = "死になさい", "失せろ", "バイバイ", "うせろ"
	nowtime1, nowtime2 = "今の時間を教えて", "今の時間教えて"
	minuteafter = "分後"
)

func OnMessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {

	if m.Author.Bot {
		return
	}

	nickname := m.Author.Username
	member, err := s.State.Member(m.GuildID, m.Author.ID)
	if err == nil && member.Nick != "" {
		nickname = member.Nick
	}

	fmt.Println("< " + m.Content + " by " + nickname)
	checkminute, minute := algorithm.BMCheck(m.Content, minuteafter)
	checkexplosion, _ := algorithm.BMCheck(m.Content, explosion)

	switch {
	// Botに対して、「おつかれさま」と実行したとき、Botが「ありがとう」返してくれるもの
	case m.Content == greatjob:
		SendMessage(s, m.ChannelID, "ありがとう", nickname)

	case m.Content == startfan:
		SendMessage(s, m.ChannelID, "おうよ、馬吉センプウキ起動", nickname)
		SendMessage(s, m.ChannelID, "ブォーーーーー!!!!!!!", nickname)
		SendMessage(s, m.ChannelID, "わっしょいわっしょい", nickname)

	case m.Content == botsorry:
		SendMessage(s, m.ChannelID, "誠にお騒がせして申し訳ございませんでした", nickname)
		SendMessage(s, m.ChannelID, "なんて謝ると思ったか、はっはっは～", nickname)
		SendMessage(s, m.ChannelID, "まぁ、うるさかったら抜けてください", nickname)

	// Botに対して、「爆発させて」と実行したとき、BotがonPlayBoom(爆発音をボイスチャンネル内に鳴り響かせるもの)を実行するもの
	case m.Content == explosion:
		OnPlayBoom(s, m, nickname)

	// Botに対して、「きて」・「こい」・「来て」・「来い」と実行したとき、Botが言ったユーザーのボイスチャンネル内に入ってくるもの
	case m.Content == join1 || m.Content == join2 || m.Content == join3 || m.Content == join4:
		SendMessage(s, m.ChannelID, "がってんしょうちのすけ", nickname)
		c, err := s.State.Channel(m.ChannelID)
		if err != nil {
			fmt.Println(err)
		}

		voicechannelID := VoiceChannelSearch(s, member.User.ID)

		if voicechannelID != "" {
			vcsession, _ = s.ChannelVoiceJoin(c.GuildID, voicechannelID, false, false)
			vcsession.AddHandler(OnVoiceReceived)
		}else {
			fmt.Println("なんも入ってないやないかぁい")
		}

	// Botに対して、「このチャンネルのことを教えて」と実行したとき、Botがこのチャンネル全体のことを教えてくれるもの
	case m.Content == teachchannel1 || m.Content == teachchannel2:
		c, err := s.State.Channel(m.ChannelID)
		if err != nil {
			fmt.Println(err)
		}
		guildChannnels, _ := s.GuildChannels(c.GuildID)

		var sendText string
		for _, a := range guildChannnels{
			sendText += fmt.Sprintf("%vチャンネルの%v(IDは%v)\n", a.Type, a.Name, a.ID)
		}
		SendMessage(s, m.ChannelID, sendText, nickname)

	// Botに対して、「死になさい」・「失せろ」・「うせろ」・「バイバイ」と実行したとき、Botが任意のボイスチャンネルから退出するもの
	case m.Content == leave1 || m.Content == leave2 || m.Content == leave3 || m.Content == leave4:
		SendMessage(s, m.ChannelID, "すんません", nickname)
		vcsession.Disconnect()

	// Botに対して、「今の時間を教えて」・「今の時間教えて」と実行したとき、Botが現在の時間を教えてくれるもの
	case m.Content == nowtime1 || m.Content == nowtime2:
		SendMessage(s, m.ChannelID, "ヒヒーン", nickname)
		now := time.Now()
		SendMessage(s, m.ChannelID, "今は " + now.String() + " だよ", nickname)

	// Botに対して、「分後」と「爆発させて」を含んだ文字列で実行したとき、Botが指定した分数後にボイスチャンネルで爆発させるもの
	case checkminute != 0 && checkexplosion != 0:
		SendMessage(s, m.ChannelID, "任せろ、爆弾設置!!!", nickname)
		go func() {
			time.Sleep(time.Minute * time.Duration(minute))
			OnPlayBoom(s, m, nickname)
			SendMessage(s, m.ChannelID, "ちなみにさっきの爆破は俺のおならだぜ", nickname)

			explosiontimer <- true
		}()

	default:
		SendMessage(s, m.ChannelID, "何言ってるかわかりませんねぇ", nickname)
	}

	<- explosiontimer
}

