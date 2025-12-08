package bot

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)

// Run initializes and starts the bot

func Run() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	init_database()
	riot_api_key := init_riot_api()
	_ = riot_api_key

	discord_token := os.Getenv("DISCORD_BOT_TOKEN")
	dg, err := discordgo.New("Bot " + discord_token)
	if err != nil {
		log.Fatal("Error creating Discord session:", err)
	}

	dg.Identify.Intents = discordgo.IntentsGuildMessages |
		discordgo.IntentsMessageContent

	err = dg.Open()
	if err != nil {
		log.Fatal("Error opening Discord session:", err)
	}
	defer dg.Close()

	for _, v := range commands() {
		_, err := dg.ApplicationCommandCreate(dg.State.User.ID, "", v)
		if err != nil {
			log.Fatalf("Cannot create '%v' command: %v", v.Name, err)
		}
	}

	log.Println("Bot is now running. Press CTRL+C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc
}
