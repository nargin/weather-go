package bot

import (
	"github.com/bwmarrin/discordgo"
)

func commands() []*discordgo.ApplicationCommand {
	commands := []*discordgo.ApplicationCommand{
		{
			Name:        "659ping",
			Description: "Replies with Pong!",
		},
		{
			Name:        "stats",
			Description: "Get player statistics",
		},
	}

	return commands
}
