package main

import (
	"github.com/Edouard127/FurryProtectorGo/client"
	"github.com/joho/godotenv"
	"go.uber.org/zap"
	_ "net/http/pprof"
	"os"
)

func main() {
	godotenv.Load()

	logger, _ := zap.NewDevelopment()

	client.NewClient(logger.With(zap.String("module", "client")), os.Getenv("TOKEN")) // blocking call
}
