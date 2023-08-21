package events

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

type UserUpdateLog struct{}

func NewUserUpdateLog() *UserUpdateLog {
	return &UserUpdateLog{}
}

func (h *UserUpdateLog) Handler(s *discordgo.Session, e *discordgo.GuildMemberUpdate) {
	var user_update discordgo.MessageEmbed
	if e.BeforeUpdate.Nick != e.Nick {
		user_id := fmt.Sprintf("ID: %s", e.User.ID)
		user_update.Title = "Username Changed"
		user_update.Author = &discordgo.MessageEmbedAuthor{Name: e.User.Username, IconURL: e.User.AvatarURL("")}
		user_update.Fields = append(user_update.Fields, &discordgo.MessageEmbedField{Name: "Before", Value: e.BeforeUpdate.Nick, Inline: false})
		user_update.Fields = append(user_update.Fields, &discordgo.MessageEmbedField{Name: "After", Value: e.Nick, Inline: false})
		user_update.Footer = &discordgo.MessageEmbedFooter{Text: user_id}
	}
	if e.BeforeUpdate.Avatar != e.Avatar {
		user_id := fmt.Sprintf("ID: %s", e.User.ID)
		user_update.Title = "Avatar changed"
		user_update.Author = &discordgo.MessageEmbedAuthor{Name: e.User.Username, IconURL: e.User.AvatarURL("")}
		user_update.Fields = append(user_update.Fields, &discordgo.MessageEmbedField{Name: "User", Value: e.User.Username, Inline: false})
		user_update.Thumbnail = &discordgo.MessageEmbedThumbnail{URL: e.AvatarURL("")}
		user_update.Footer = &discordgo.MessageEmbedFooter{Text: user_id}
	}
	s.ChannelMessageSendEmbed("1143244214189170759", &user_update)
}
