package botmanager

import "github.com/kohinigeee/DiscordBotTemplate/slashapi"

type SlashCommand struct {
	Command slashapi.SlashCommandJson
	Handler DiscorBotdHandler
}

type InteractCommand struct {
	Name    string
	Handler DiscorBotdHandler
}

type DiscordModalCommand struct {
	Name    string
	Handler DiscorBotdHandler
}

var (
	SlashCommands        []SlashCommand
	InteractCommands     []InteractCommand
	DiscordModalCommands []DiscordModalCommand
)

func init() {

	SlashCommands = []SlashCommand{
		{
			Command: slashapi.SlashCommandJson{
				Name:        "ping",
				Description: "ping pong",
			},
			Handler: pingPong,
		},
	}

	InteractCommands = []InteractCommand{}

	DiscordModalCommands = []DiscordModalCommand{}

}

func InitialSlashCommands() []SlashCommand {
	return SlashCommands
}

func InitialInteractCommands() []InteractCommand {
	return InteractCommands
}

func InitialDiscordModalCommands() []DiscordModalCommand {
	return DiscordModalCommands
}
