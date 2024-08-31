package cmd

import (
	"log/slog"
	"os"

	"github.com/joho/godotenv"
	"github.com/kohinigeee/DiscordBotTemplate/mylogger"
)

var (
	discordToken string
	appId        string
	logger       *mylogger.LoggerItem
)

func init() {
	logger = mylogger.GetLogger("Cmd Global")

	err := godotenv.Load()
	if err != nil {
		logger.Warn("Error loading .env file")
	}

	discordToken = "Bot " + os.Getenv("DISCORD_TOKEN")
	appId = os.Getenv("APPLICATION_ID")

	logger.Info("Imported env values", slog.String("DISCORD_TOKEN", discordToken), slog.String("APPLICATION_ID", appId))
}
