package events

import (
	"bot/internal"

	"github.com/bwmarrin/discordgo"
)

type ReadyHandler struct{}

func NewReadyHandler() *ReadyHandler {
	return &ReadyHandler{}
}

func (h *ReadyHandler) Handler(s *discordgo.Session, e *discordgo.Ready) {
	s.UpdateGameStatus(1, "mit den Dweebis.")
	internal.Log("Bot is ready!", 0)
	internal.LogFormat("Logged in as %s\n", 0, e.User.Username)
}
