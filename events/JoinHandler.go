package events

import (
	"fmt"

	"bot/internal"

	"github.com/bwmarrin/discordgo"
)

type JoinHandler struct{}

func NewJoinHandler() *JoinHandler {
	return &JoinHandler{}
}

func (h *JoinHandler) Handler(s *discordgo.Session, e *discordgo.GuildMemberAdd) {
	s.GuildMemberRoleAdd(e.GuildID, e.User.ID, "820680792363237379")
	format := fmt.Sprintf("Wilkommen <@%s>! Ließ dir am besten erstmal die Regeln durch. Dannach kannst du in <#1128726602776326184> deine Rollen personalisieren!", e.User.ID)
	s.ChannelMessageSend(internal.Bot_Config.WelcomeChannel, format)
}
