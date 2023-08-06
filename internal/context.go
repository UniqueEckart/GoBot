package internal

import (
	"github.com/bwmarrin/discordgo"
)

type Context interface {
	GetSession() *discordgo.Session
	GetMessage() *discordgo.Message
	GetGuild() *discordgo.Guild
	GetChannel() *discordgo.Channel
	GetAllGuildChanneles() []*discordgo.Channel
	Replay(message string) *discordgo.Message
	GetArgs() []string
}

type context struct {
	session *discordgo.Session
	message *discordgo.Message
	guild   *discordgo.Guild
	channel *discordgo.Channel
	args    []string
	Handler *CommandHandler
}

func (ctx *context) GetSession() *discordgo.Session {
	return ctx.session
}

func (ctx *context) GetGuild() *discordgo.Guild {
	return ctx.guild
}

func (ctx *context) GetAllGuildChanneles() []*discordgo.Channel {
	channels, err := ctx.session.GuildChannels(ctx.guild.ID)
	if err != nil {
		panic(err)
	}
	return channels
}

func (ctx *context) GetMessage() *discordgo.Message {
	return ctx.message
}

func (ctx *context) GetChannel() *discordgo.Channel {
	return ctx.channel
}

func (ctx *context) GetArgs() []string {
	return ctx.args
}

func (ctx *context) Replay(message string) *discordgo.Message {
	reply, err := ctx.session.ChannelMessageSend(ctx.channel.ID, message)
	if err != nil {
		panic("Could not send Message")
	}
	return reply
}
