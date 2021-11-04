package main

import (
	"log"
	"github.com/bwmarrin/discordgo"
)

const token string = ""

var (
	BotId  string
	Prefix string = "!"
)

func main() {
	log.Printf("Starting Bot...")

	dg, err := discordgo.New("Bot " + token)

	if err != nil {
		log.Printf("%s", err.Error())
		return
	}

	u, err := dg.User("@me")

	if err != nil {
		log.Printf("%s", err.Error())
		return
	}

	BotId = u.ID
	err = dg.Open()

	if err != nil {
		log.Printf("%s", err.Error())
		return
	}

	log.Printf("Bot is ready!")

	<-make(chan struct{})
	return
}
