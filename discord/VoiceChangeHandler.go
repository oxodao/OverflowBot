package discord

import (
	"errors"

	"github.com/bwmarrin/discordgo"
	"github.com/oxodao/overflow-bot/services"
)

type VoiceChannelInfo struct {
	Channel *discordgo.VoiceConnection
	Caller  string
}

var CurrentChannel *VoiceChannelInfo = nil

func VoiceChangeHandler(prv *services.Provider) func(*discordgo.Session, *discordgo.VoiceStateUpdate) {
	return func(s *discordgo.Session, vsu *discordgo.VoiceStateUpdate) {
		// If the user who called the bot last is no longer in the channel
		if CurrentChannel != nil && vsu.UserID == CurrentChannel.Caller && vsu.ChannelID != CurrentChannel.Channel.ChannelID {
			CurrentChannel.Channel.Disconnect()
			CurrentChannel = nil
		}
	}
}

func findChannel(s *discordgo.Session, gID, chanID string) (*discordgo.Channel, *discordgo.Guild, error) {
	c, err := s.State.Channel(chanID)
	if err != nil {
		return nil, nil, err
	}

	g, err := s.State.Guild(c.GuildID)
	return c, g, err
}

func findVocalChannel(g *discordgo.Guild, c *discordgo.Channel, userID string) (*discordgo.VoiceState, error) {
	for _, vs := range g.VoiceStates {
		if vs.UserID == userID {
			return vs, nil
		}
	}

	return nil, errors.New("Can't find room")
}
