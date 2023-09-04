package main

import (
	"bot/commands"
	"bot/events"
	"bot/internal"
	"bot/web"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

var configPath string
var debug bool
var Config *internal.Config

func init() {
	flag.StringVar(&configPath, "c", "./config.json", "The Path of the Config to use. Default: ./config.json")
	flag.BoolVar(&debug, "debug", false, "Debugging")
}

func main() {

	Config, _ = internal.ParseConfigFromJSONFile(configPath)

	discord, err := discordgo.New("Bot " + Config.Token)
	if err != nil {
		internal.Log("Faild to create discord session!", 1)
		return
	}

	discord.Identify.Intents = discordgo.IntentsAll
	web.Init()
	registerEvents(discord)
	registerCommands(discord, Config.Prefix)

	if err = discord.Open(); err != nil {
		internal.Log(err.Error(), 1)
	}
	internal.Log("Bot is now running. Press CTRL-C to exit", 0)
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	discord.Close()
}

func registerEvents(s *discordgo.Session) {
	s.AddHandler(events.NewUserUpdateLog().Handler)
	//	s.AddHandler(events.NewJoinHandler().Handler)
	s.AddHandler(events.NewReadyHandler().Handler)
	s.AddHandler(events.NewLeaveHandler().Handler)
}

func registerCommands(s *discordgo.Session, prefix string) {
	cmdHandler := internal.NewCommandHandler(prefix, Config)
	cmdHandler.OnError = func(err error, ctx internal.Context) {
		fmt.Printf("Executing of Comman failed: %s", err.Error())
	}

	cmdHandler.RegisterCommand(&commands.CmdBackup{})
	cmdHandler.RegisterMiddelware(&internal.MWPermissions{})

	s.AddHandler(cmdHandler.HandleMessage)
}
