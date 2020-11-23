package main

import (
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/oxodao/overflow-bot/config"
	"github.com/oxodao/overflow-bot/discord"
	"github.com/oxodao/overflow-bot/services"
	"github.com/oxodao/overflow-bot/web"
)

const (
	VERSION = "1.0"
	AUTHOR  = "Oxodao"
)

func main() {
	fmt.Printf("Overflow Bot [v.%s] by %s\n", VERSION, AUTHOR)
	cfg, err := config.Load()
	if err != nil {
		panic(err)
	}

	prv, err := services.NewProvider(cfg, VERSION, AUTHOR)
	if err != nil {
		panic(err)
	}

	GETCOURS := false
	if len(os.Args) > 1 && strings.ToLower(os.Args[1]) == "import_cours" {
		GETCOURS = true
	}

	if GETCOURS {
		FetchCours(prv, "cours.csv")
		return
	}

	if len(cfg.Discord.Token) > 0 {
		err := discord.Initialize(prv)
		if err != nil {
			fmt.Println("Could not connect to discord: ", err)
		}
	}

	go web.Initialize(prv)

	fmt.Println("OverflowBot is connected.\nCTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	prv.Discord.Close()
}
