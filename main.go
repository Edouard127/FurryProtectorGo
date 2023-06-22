package main

import (
	"github.com/Edouard127/FurryProtectorGo/client"
	"github.com/joho/godotenv"
	"go.uber.org/zap"
	"os"
)

func main() {
	godotenv.Load()

	logger, _ := zap.NewDevelopment()

	bot, err := client.NewClient(logger.With(zap.String("module", "client")), os.Getenv("TOKEN"))
	if err != nil {
		logger.Panic("Error while creating bot", zap.Error(err))
	}

	err = bot.Open()
	if err != nil {
		logger.Panic("Error while opening bot", zap.Error(err))
	}

	for {
	}
}
