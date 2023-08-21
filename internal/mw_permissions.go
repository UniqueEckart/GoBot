package internal

import (
	"github.com/bwmarrin/discordgo"
)

type MWPermissions struct{}

func (mw *MWPermissions) Exec(ctx *context, cmd Command) (next bool, err error) {
	if !cmd.AdminRequired() {
		next = true
		return
	}
	defer func() {
		if !next && err == nil {
			_, err = ctx.GetSession().ChannelMessageSend(ctx.GetMessage().ChannelID,
				"Du hast nicht die rechte diesen Command auszuführen!")

		}
	}()

	guild := ctx.GetGuild()

	if guild.OwnerID == ctx.GetMessage().Author.ID {
		next = true
		return
	}
	roleMap := make(map[string]*discordgo.Role)
	for _, role := range guild.Roles {
		roleMap[role.ID] = role
	}

	for _, rID := range ctx.GetMessage().Member.Roles {
		if role, ok := roleMap[rID]; ok && role.Permissions&cmd.RequiredUserPermissions() > 0 {
			next = true
			break
		}
	}
	return
}
