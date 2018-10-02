package main

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

var (
	botID         string
	commandPrefix string
)

func main() {
	discord, err := discordgo.New("Bot NDkzMTUxOTM3MzgyODQyMzkw.Dog0YA.0yPdlzJ3DeoQbzJ3i1FfKVivkP8")
	errPanic("Error creating discord session", err)

	user, err := discord.User("@me")
	errPanic("Error getting bot user object", err)

	botID = user.ID

	discord.AddHandler(commandHandler)
	discord.AddHandlerOnce(statusHandler)

	err = discord.Open()
	errPanic("Could not connect to discord.", err)

	defer discord.Close()

	commandPrefix = "$"

	// hack to keep that running in background
	<-make(chan struct{})
}

func commandHandler(discord *discordgo.Session, message *discordgo.MessageCreate) {
	user := message.Author
	if user.ID == botID || user.Bot {
		return
	}

	content := message.Content

	if content[0] == commandPrefix[0] {
		if content == commandPrefix+"help" {
			sendHelp(discord, message)
			return
		}

		discord.ChannelMessageSend(message.ChannelID, fmt.Sprintf("WAT? I don't know what to do with: %s", content))
	}
}

func sendHelp(discord *discordgo.Session, message *discordgo.MessageCreate) {
	discord.ChannelMessageSend(message.ChannelID, "Here's some help: get a job.")
}

func statusHandler(discord *discordgo.Session, message *discordgo.MessageCreate) {
	err := discord.UpdateStatus(0, "Play some ShadowCord")
	if err != nil {
		fmt.Printf("Error attempting to set my status.\n")
	}

	servers := discord.State.Guilds
	fmt.Printf("ShadowCord has started on %d servers.\n", len(servers))
}

func errPanic(msg string, err error) {
	if err != nil {
		fmt.Printf("%s: %+v", msg, err)
		panic(err)
	}
}
