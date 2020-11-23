package discord

import (
	"database/sql"
	"fmt"

	"github.com/bwmarrin/dgvoice"
	"github.com/bwmarrin/discordgo"

	"github.com/oxodao/overflow-bot/services"
)

type SoundCommand struct{}

func (c SoundCommand) Process(prv *services.Provider, dmsg *discordgo.MessageCreate, args string) {
	if args == "list" {
		sounds, err := SelectSounds(prv)
		if err != nil {
			prv.Discord.ChannelMessageSend(dmsg.ChannelID, fmt.Sprintf("Error listing sounds: %v", err))
			return
		}

		msg := "Liste des sons: \n"
		for _, s := range sounds {
			msg += fmt.Sprintf("\t- %v: %v\n", s.Name, s.Description)
		}

		prv.Discord.ChannelMessageSend(dmsg.ChannelID, msg)
		return
	}

	channel, guild, err := findChannel(prv.Discord, dmsg.GuildID, dmsg.ChannelID)
	if err != nil {
		prv.Discord.ChannelMessageSend(dmsg.ChannelID, "Can't find the channel!")
		CurrentChannel = nil
		return
	}

	vs, err := findVocalChannel(guild, channel, dmsg.Author.ID)
	if err != nil {
		prv.Discord.ChannelMessageSend(dmsg.ChannelID, "Can't find the channel!")
		CurrentChannel = nil
		return
	}

	// If the bot is not in a room or is in another room, we call him there
	if CurrentChannel == nil || CurrentChannel.Channel.ChannelID != vs.ChannelID {
		vc, err := prv.Discord.ChannelVoiceJoin(vs.GuildID, vs.ChannelID, false, true)
		if err != nil {
			prv.Discord.ChannelMessageSend(dmsg.ChannelID, "Can't join the channel!")
			CurrentChannel = nil
			return
		}

		CurrentChannel = &VoiceChannelInfo{
			Channel: vc,
			Caller:  dmsg.Author.ID,
		}
	} else if CurrentChannel != nil && CurrentChannel.Caller != dmsg.Author.ID {
		// If the bot was called by someone else, we change the caller
		CurrentChannel.Caller = dmsg.Author.ID
	}

	// We're all setup, now time to find the sound requested
	sound, err := SelectSound(prv, args)
	if err != nil {
		if err == sql.ErrNoRows {
			prv.Discord.ChannelMessageSend(dmsg.ChannelID, "Ce son n'existe pas")
			return
		}

		prv.Discord.ChannelMessageSend(dmsg.ChannelID, "Une erreur est survenue: "+err.Error())
		return
	}

	dgvoice.PlayAudioFile(CurrentChannel.Channel, fmt.Sprintf("sounds/%s", sound.File), make(chan bool))
}

func (c SoundCommand) CommandName() string {
	return "sound"
}

func (c SoundCommand) HelpMessage() string {
	return "Lance un son dans le channel vocal. !sound list pour la liste des sons"
}
