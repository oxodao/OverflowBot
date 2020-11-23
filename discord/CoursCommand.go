package discord

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
	"github.com/oxodao/overflow-bot/services"
)

type CoursCommand struct{}

func (c CoursCommand) Process(prv *services.Provider, dmsg *discordgo.MessageCreate, args string) {
	cours, err := SelectCours(prv)
	if err != nil {
		prv.Discord.ChannelMessageSend(dmsg.ChannelID, "Une erreur est survenue en récupérant les cours: "+err.Error())
		return
	}

	message := ""
	if len(cours) == 2 {
		message = fmt.Sprintf("**Cours précédent: **\n%v\n**Cours suivant:**\n%v", formatCours(cours[0]), formatCours(cours[1]))
	} else if len(cours) == 1 {
		message = fmt.Sprintf("**Cours précédent: **\n%v", formatCours(cours[0]))
	}

	prv.Discord.ChannelMessageSend(dmsg.ChannelID, message)
}

func (c CoursCommand) CommandName() string {
	return "cours"
}

func (c CoursCommand) HelpMessage() string {
	return "Liste le cours précédent et le suivant"
}

func formatCours(c Cours) string {
	return "*" + c.Name + "*\n\tLe " + c.Date.Format("02/01/2006") + " de " + c.Date.Format("15:05") + " à " + c.End.Format("15:05")
}
