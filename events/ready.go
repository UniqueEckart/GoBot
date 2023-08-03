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
	fmt.Println("Bot seesion is ready!")
	fmt.Printf("Logged ub as %s\n", e.User.String())
}
