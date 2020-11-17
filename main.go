package main

import (
	"Go-discord/movement"
	"fmt"
	"github.com/bwmarrin/discordgo"
)

var(
	Token = "" // Bot作成時のTokenを入れる
	BotName = "" // BotNameを入れる
	stopBot = make(chan bool)
)

func main() {
	//Discordのセッションを作成
	//"Bot"という接頭辞がないと401 unauthorizedエラーが起きます
	discord, err := discordgo.New("Bot " + Token)
	if err != nil {
		fmt.Println("Error logging in")
		fmt.Println(err)
	}

	discord.AddHandler(movement.OnMessageCreate) //全てのWSAPIイベントが発生した時のイベントハンドラを追加
	// websocketを開いてlistening開始
	err = discord.Open()
	if err != nil {
		fmt.Println(err)
	}
	defer discord.Close()

	fmt.Println("Listening...")
	<-stopBot //プログラムが終了しないようロック
	return
}
