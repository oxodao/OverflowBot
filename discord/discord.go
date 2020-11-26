package discord

import (
	"github.com/bwmarrin/discordgo"
	"github.com/oxodao/overflow-bot/log"
	"github.com/oxodao/overflow-bot/services"
)

var (
	Commands []Command
)

func RegisterCommands() {
	Commands = []Command{}

	Commands = append(Commands, &HelpCommand{&Commands})
	Commands = append(Commands, &CoursCommand{})
	Commands = append(Commands, &SoundCommand{})
}

func RegisterCustomCommands(prv *services.Provider) error {
	cmds, err := SelectCustomCommands(prv)
	if err != nil {
		return err
	}

	for _, c := range cmds {
		Commands = append(Commands, c)
	}

	return nil
}

func Initialize(prv *services.Provider) error {

	if prv.Discord != nil {
		prv.Discord.Close()
	}

	discord, err := discordgo.New("Bot " + prv.Config.Discord.Token)
	if err != nil {
		return err
	}

	RegisterCommands()

	go func() {
		for {
			err = RegisterCustomCommands(prv)
			if err != nil {
				log.Error(err)
			} else {
				log.Normal("Custom commands reloaded successfully")
			}

			<-prv.ReloadCommands
		}
	}()

	prv.Discord = discord
	discord.AddHandler(CommandHandler(prv))
	discord.AddHandler(VoiceChangeHandler(prv))

	return discord.Open()
}
