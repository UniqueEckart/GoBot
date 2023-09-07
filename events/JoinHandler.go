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
<<<<<<< HEAD
	s.GuildMemberRoleAdd(e.GuildID, e.User.ID, "820680792363237379")
	format := fmt.Sprintf("Wilkommen <@%s>! Ließ dir am besten erstmal die Regeln durch. Dannach kannst du in <#1128726602776326184> deine Rollen personalisieren!", e.User.ID)
	s.ChannelMessageSend(internal.Bot_Config.WelcomeChannel, format)
=======
	var joinembed discordgo.MessageEmbed
	format := fmt.Sprintf("Wilkommen <@%s>!", e.User.ID)
	joinembed.Title = "Wilkommen"
	joinembed.Description = "Ließ dir am besten erstmal die Regeln durch. Dannach kannst du in <#1128726602776326184> deine Rollen personalisieren!"
	joinembed.Author = &discordgo.MessageEmbedAuthor{Name: e.User.Username, IconURL: e.User.AvatarURL("")}
	s.ChannelMessageSend("1128726602776326185", format)
>>>>>>> 97e918001166bca123c923229e307395197f3469
}
