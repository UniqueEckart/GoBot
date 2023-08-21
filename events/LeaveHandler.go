package events

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

type LeaveHandler struct{}

func NewLeaveHandler() *LeaveHandler {
	return &LeaveHandler{}
}

func (h *LeaveHandler) Handler(s *discordgo.Session, e *discordgo.GuildMemberRemove) {
	format := fmt.Sprintf("<@%s>/%s hat uns leider verlassen", e.User.ID, e.User.Username)
	s.ChannelMessageSend("1136055788582998076", format)
}
