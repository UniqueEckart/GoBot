package events

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

type VoiecUpdate struct{}

func NewVoiceUpdate() *VoiecUpdate {
	return &VoiecUpdate{}
}

func (h *VoiecUpdate) Handler(s *discordgo.Session, e *discordgo.VoiceStateUpdate) {
	fmt.Printf("Hello World")
}
