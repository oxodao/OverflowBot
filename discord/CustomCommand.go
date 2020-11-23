package discord

import (
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/oxodao/overflow-bot/services"
)

type CustomCommand struct {
	Name    string `db:"command_name"`
	Help    string `db:"command_help"`
	Content string `db:"command_response"`
}

func (c CustomCommand) Process(prv *services.Provider, mc *discordgo.MessageCreate, args string) {
	prv.Discord.ChannelMessageSend(mc.ChannelID, strings.ReplaceAll(c.Content, `\n`, "\n"))
}

func (c CustomCommand) CommandName() string {
	return c.Name
}

func (c CustomCommand) HelpMessage() string {
	return c.Help
}
