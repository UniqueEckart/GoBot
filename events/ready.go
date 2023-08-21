package events

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

type ReadyHandler struct{}

func NewReadyHandler() *ReadyHandler {
	return &ReadyHandler{}
}

func (h *ReadyHandler) Handler(s *discordgo.Session, e *discordgo.Ready) {
	s.UpdateGameStatus(1, "mit den Dweebis.")
	fmt.Println("[INFO] Bot seesion is ready!")
	fmt.Printf("[LOG] Logged in as %s\n", e.User.String())
}
