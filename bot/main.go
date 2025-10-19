package bot

import (
	"context"
	"log"
	"net/http"
	"os"

	_ "github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/v2/mongo"
	_ "go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func init_database() {
	mongo_uri := os.Getenv("MONGODB_URI")

	client, err := mongo.Connect(options.Client().ApplyURI(mongo_uri))
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(context.TODO())

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Connected to MongoDB!")
}

func init_riot_api() string {
	riot_api_key := os.Getenv("RIOT_API_KEY")

	// Initialize Riot API client here
	resp, err := http.Get(
		"https://euw1.api.riotgames.com/lol/platform/v3/champion-rotations" +
			"?api_key=" + riot_api_key)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Fatalf("Riot API request failed with status: %s", resp.Status)
	}

	log.Println("Riot API initialized successfully!")
	return riot_api_key
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	init_database()
	riot_api_key := init_riot_api()
	_ = riot_api_key

	discord_token := os.Getenv("DISCORD_BOT_TOKEN")
	_ = discord_token

}
