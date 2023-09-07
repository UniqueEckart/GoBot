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

func init() {
	flag.StringVar(&configPath, "c", "./config.json", "Allows to set the Config path")
	flag.BoolVar(&debug, "debug", false, "Activates the Debug Logs of the Bot")
	flag.Parse()
}

func main() {

	internal.Bot_Config, _ = internal.ParseConfigFromJSONFile(configPath)

	discord, err := discordgo.New("Bot " + internal.Bot_Config.Token)
	if err != nil {
		fmt.Println("[ERROR] Error creating discord session", err)
		return
	}

	discord.Identify.Intents = discordgo.IntentsAll
	web.Init()
	registerEvents(discord)
	registerCommands(discord, internal.Bot_Config.Prefix)

	if err = discord.Open(); err != nil {
		panic(err)
	}

	fmt.Println("[INFO] Bot is now running.  Press CTRL-C to exit")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	discord.Close()
}

func registerEvents(s *discordgo.Session) {
	s.AddHandler(events.NewJoinHandler().Handler)
	s.AddHandler(events.NewReadyHandler().Handler)
	s.AddHandler(events.NewLeaveHandler().Handler)
}

func registerCommands(s *discordgo.Session, prefix string) {
	cmdHandler := internal.NewCommandHandler(prefix, internal.Bot_Config)
	cmdHandler.OnError = func(err error, ctx internal.Context) {
		fmt.Printf("Executing of Comman failed: %s", err.Error())
	}

	cmdHandler.RegisterCommand(&commands.CmdBackup{})
	cmdHandler.RegisterMiddelware(&internal.MWPermissions{})

	s.AddHandler(cmdHandler.HandleMessage)
}
