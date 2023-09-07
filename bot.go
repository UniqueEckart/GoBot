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
<<<<<<< HEAD

func init() {
	flag.StringVar(&configPath, "c", "./config.json", "Allows to set the Config path")
	flag.BoolVar(&debug, "debug", false, "Activates the Debug Logs of the Bot")
	flag.Parse()
=======
var Config *internal.Config

func init() {
	flag.StringVar(&configPath, "c", "./config.json", "The Path of the Config to use. Default: ./config.json")
	flag.BoolVar(&debug, "debug", false, "Debugging")
>>>>>>> 97e918001166bca123c923229e307395197f3469
}

func main() {

<<<<<<< HEAD
	internal.Bot_Config, _ = internal.ParseConfigFromJSONFile(configPath)

	discord, err := discordgo.New("Bot " + internal.Bot_Config.Token)
=======
	Config, _ = internal.ParseConfigFromJSONFile(configPath)

	discord, err := discordgo.New("Bot " + Config.Token)
>>>>>>> 97e918001166bca123c923229e307395197f3469
	if err != nil {
		internal.Log("Faild to create discord session!", 1)
		return
	}

	discord.Identify.Intents = discordgo.IntentsAll
	web.Init()
	registerEvents(discord)
<<<<<<< HEAD
	registerCommands(discord, internal.Bot_Config.Prefix)
=======
	registerCommands(discord, Config.Prefix)
>>>>>>> 97e918001166bca123c923229e307395197f3469

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
	s.AddHandler(events.NewJoinHandler().Handler)
	s.AddHandler(events.NewReadyHandler().Handler)
	s.AddHandler(events.NewLeaveHandler().Handler)
}

func registerCommands(s *discordgo.Session, prefix string) {
<<<<<<< HEAD
	cmdHandler := internal.NewCommandHandler(prefix, internal.Bot_Config)
=======
	cmdHandler := internal.NewCommandHandler(prefix, Config)
>>>>>>> 97e918001166bca123c923229e307395197f3469
	cmdHandler.OnError = func(err error, ctx internal.Context) {
		fmt.Printf("Executing of Comman failed: %s", err.Error())
	}

	cmdHandler.RegisterCommand(&commands.CmdBackup{})
	cmdHandler.RegisterMiddelware(&internal.MWPermissions{})

	s.AddHandler(cmdHandler.HandleMessage)
}
