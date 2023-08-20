package commands

import (
	"bot/internal"
	"encoding/csv"
	"fmt"
	"os"
	"strconv"

	"github.com/bwmarrin/discordgo"
)

type CmdBackup struct{}

func (c *CmdBackup) Invokes() []string {
	return []string{"backup"}
}
func (c *CmdBackup) Descreption() string {
	return "Save and Load Channels of a Server like a backup just for discord Servers"
}
func (c *CmdBackup) HelpSyntax() string {
	return "Syntax: backup <subcommand>\nAvaible Options: save, load"
}

func (c *CmdBackup) AdminRequired() bool {
	return true
}

func (c *CmdBackup) HasSubCommands() bool {
	return true
}

func (c *CmdBackup) Exec(ctx internal.Context) (err error) {
	if ctx.GetArgs()[0] == "save" {
		serializeChannelPermissions(ctx.GetAllGuildChanneles())
		ctx.Replay("Alles Channels gesichert!")
	}

	if ctx.GetArgs()[0] == "no3i2nr948329ß43tn943ßt" {
		loadChannelPermissions(ctx.GetSession(), ctx.GetAllGuildChanneles())
		ctx.Replay("Alle Channel wiederhergestellt!")
	}
	return
}

func serializeChannelPermissions(c []*discordgo.Channel) {
	os.Mkdir("./BACKUP", os.ModePerm)
	for _, channel := range c {
		if channel.Type == discordgo.ChannelTypeGuildCategory {
			println("Cat")
			continue
		}
		name := fmt.Sprintf("BACKUP/%s.csv", channel.Name)
		file, err := os.Create(name)
		for _, perms := range channel.PermissionOverwrites {
			if err != nil {
				println("Error creating file", err)
				return
			}
			fmt.Fprintf(file, "%s,%b,%b,%b\n", perms.ID, perms.Type, perms.Allow, perms.Deny)
		}
		file.Close()
	}
}

func readCsvFile(filepath string) []*discordgo.PermissionOverwrite {
	file, err := os.Open(filepath)
	defer file.Close()
	if err != nil {
		fmt.Println("Unable to read file "+filepath, err)
		panic("Panic")
	}
	records, err := csv.NewReader(file).ReadAll()
	if err != nil {
		fmt.Println("Failed to parse CSV "+filepath, err)
		panic("")
	}

	var cPerm []*discordgo.PermissionOverwrite

	for _, row := range records {
		var temp discordgo.PermissionOverwrite
		temp.ID = row[0]
		temp.Type = 0
		temp.Allow, _ = strconv.ParseInt(row[1], 2, 64)
		temp.Deny, _ = strconv.ParseInt(row[2], 2, 64)
		cPerm = append(cPerm, &temp)
	}
	return cPerm
}

func loadChannelPermissions(s *discordgo.Session, c []*discordgo.Channel) {
	items, _ := os.ReadDir("./BACKUP")
	for _, item := range items {
		channel_permissions := readCsvFile("BACKUP/" + item.Name())
		for _, channel := range c {
			if channel.Name == item.Name() {
				var temp discordgo.ChannelEdit
				temp.PermissionOverwrites = channel_permissions
				s.ChannelEdit(channel.ID, &temp)
			}
		}
	}
}
