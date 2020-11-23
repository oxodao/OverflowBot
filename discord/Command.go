package discord

import (
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/oxodao/overflow-bot/services"
)

type Command interface {
	Process(*services.Provider, *discordgo.MessageCreate, string)
	CommandName() string
	HelpMessage() string
}

func CommandHandler(prv *services.Provider) func(*discordgo.Session, *discordgo.MessageCreate) {
	return func(s *discordgo.Session, m *discordgo.MessageCreate) {
		// @TODO Permissions

		if !strings.HasPrefix(m.Content, "!") {
			return
		}

		cmd := strings.ToLower(strings.Trim(m.Content[1:], " "))
		args := ""

		if strings.Index(cmd, " ") > 0 {
			args = cmd[strings.Index(cmd, " ")+1:]
			cmd = cmd[:strings.Index(cmd, " ")]
		}

		var command Command
		for _, c := range Commands {
			if c.CommandName() == cmd {
				command = c
			}
		}

		if command != nil {
			command.Process(prv, m, args)
		} else {
			s.ChannelMessageSend(m.ChannelID, "Commande inexistante!")
		}
	}
}
