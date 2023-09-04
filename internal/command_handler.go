package internal

import (
	"strings"

	"github.com/bwmarrin/discordgo"
)

type CommandHandler struct {
	prefix string

	cmdInstaces []Command
	cmdMap      map[string]Command
	middlewares []Middelware

	OnError func(err error, ctx Context)
}

func NewCommandHandler(prefix string, cfg *Config) *CommandHandler {
	return &CommandHandler{
		prefix:      prefix,
		cmdInstaces: make([]Command, 0),
		cmdMap:      make(map[string]Command),
		middlewares: make([]Middelware, 0),
		OnError:     func(error, Context) {},
	}
}

func (c *CommandHandler) RegisterCommand(cmd Command) {
	c.cmdInstaces = append(c.cmdInstaces, cmd)
	for _, invoke := range cmd.Invokes() {
		c.cmdMap[invoke] = cmd
	}
}

func (c *CommandHandler) RegisterMiddelware(mw Middelware) {
	c.middlewares = append(c.middlewares, mw)
}

func (c *CommandHandler) HandleMessage(s *discordgo.Session, e *discordgo.MessageCreate) {

	if e.Author.ID == s.State.User.ID || e.Author.Bot || !strings.HasPrefix(e.Content, c.prefix) {
		return
	}

	split := strings.Split(e.Content[len(c.prefix):], " ")
	if len(split) < 1 {
		return
	}

	invoke := strings.ToLower(split[0])
	args := split[1:]

	cmd, has := c.cmdMap[invoke]
	if !has || cmd == nil {
		return
	}

	ctxguild, err := s.Guild(e.GuildID)
	if err != nil {
		Log("Could not get Guild", 1)
	}

	ctxchannel, err := s.Channel(e.ChannelID)
	if err != nil {
		Log("Could not get Channel", 1)
	}

	ctx := &context{
		session: s,
		message: e.Message,
		guild:   ctxguild,
		channel: ctxchannel,
		args:    args,
		Handler: c,
	}
	if cmd.HasSubCommands() && len(ctx.GetArgs()) == 0 {
		ctx.Replay(cmd.HelpSyntax())
		return
	}

	for _, mw := range c.middlewares {
		next, err := mw.Exec(ctx, cmd)
		if err != nil {
			c.OnError(err, ctx)
			return
		}
		if !next {
			return
		}
	}

	if err := cmd.Exec(ctx); err != nil {
		c.OnError(err, ctx)
	}

}
