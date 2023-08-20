package main

import (
	//	"bot/commands"
	"bot/events"
	"bot/internal"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

var configPath string
var debug bool
var cfg *internal.Config

func init() {
	flag.StringVar(&configPath, "c", "./config.json", "Allows to set the Config path")
	flag.BoolVar(&debug, "debug", false, "Activates the Debug Logs of the Bot")
}

func main() {

	cfg, _ = internal.ParseConfigFromJSONFile(configPath)

	discord, err := discordgo.New("Bot " + cfg.Token)
	if err != nil {
		fmt.Println("Error creating discord session", err)
		return
	}

	discord.Identify.Intents = discordgo.IntentsAll

	registerEvents(discord)
	registerCommands(discord, cfg.Prefix)

	if err = discord.Open(); err != nil {
		panic(err)
	}

	fmt.Println("Bot is now running.  Press CTRL-C to exit")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	discord.Close()
}

func registerEvents(s *discordgo.Session) {
	s.AddHandler(events.NewReadyHandler().Handler)
	s.AddHandler(events.NewLeaveHandler().Handler)
}

func registerCommands(s *discordgo.Session, prefix string) {
	cmdHandler := internal.NewCommandHandler(prefix, cfg)
	cmdHandler.OnError = func(err error, ctx internal.Context) {
		fmt.Printf("Executing of Comman failed: %s", err.Error())
	}

	//	cmdHandler.RegisterCommand(&commands.CmdBackup{})
	cmdHandler.RegisterMiddelware(&internal.MWPermissions{})

	s.AddHandler(cmdHandler.HandleMessage)
}
