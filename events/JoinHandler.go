package events

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

type JoinHandler struct{}

func NewJoinHandler() *JoinHandler {
	return &JoinHandler{}
}

func (h *JoinHandler) Handler(s *discordgo.Session, e *discordgo.GuildMemberAdd) {
	format := fmt.Sprintf("Wilkommen <@%s>! Ließ dir am besten erstmal die Regeln durch. Dannach kannst du in <#1128726602776326184> deine Rollen personalisieren!", e.User.ID)
	s.ChannelMessageSend("1128726602776326185", format)
}
