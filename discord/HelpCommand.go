package discord

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
	"github.com/oxodao/overflow-bot/services"
)

type HelpCommand struct {
	Commands *[]Command
}

func (c HelpCommand) Process(prv *services.Provider, dmsg *discordgo.MessageCreate, args string) {
	message := "OverflowBot [v.%v] by %v \n"

	for _, c := range *c.Commands {
		message += fmt.Sprintf("\t- %s: %s\n", c.CommandName(), c.HelpMessage())
	}

	prv.Discord.ChannelMessageSend(dmsg.ChannelID, fmt.Sprintf(message, prv.Software.Version, prv.Software.Author))
}

func (c HelpCommand) CommandName() string {
	return "help"
}

func (c HelpCommand) HelpMessage() string {
	return "Shows this message"
}
