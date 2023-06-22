package client

import (
	"github.com/Edouard127/FurryProtectorGo/client/database"
	"github.com/Edouard127/FurryProtectorGo/client/events"
	"github.com/Edouard127/FurryProtectorGo/client/interaction"
	"github.com/Edouard127/FurryProtectorGo/registers"
	"github.com/bwmarrin/discordgo"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/collectors"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.uber.org/zap"
	"log"
	"net/http"
	"os"
)

func NewClient(logger *zap.Logger, token string) (client *discordgo.Session, err error) {
	client, err = discordgo.New("Bot " + token)
	if err != nil {
		return nil, err
	}

	return doPreInit(logger, client), nil
}

func doPreInit(logger *zap.Logger, client *discordgo.Session) *discordgo.Session {
	db := database.NewDatabase(logger.With(zap.String("module", "database")), os.Getenv("MONGO_URI"), os.Getenv("DATABASE_NAME"), "config", "users", "verification_cache")
	registry := prometheus.NewRegistry()

	go doPrometheus(registry)
	doEvents(logger, client, registry, db)
	doCommands(logger, registry, db)

	endpoint := discordgo.EndpointApplicationGlobalCommands(os.Getenv("APP_ID"))

	for _, command := range registers.InteractionCommands.Runners {
		_, err := client.RequestWithBucketID("POST", endpoint, command, endpoint)
		if err != nil {
			command.GetLogger().Error("Error while registering command", zap.Error(err))
		}
	}

	return client
}

func doEvents(logger *zap.Logger, client *discordgo.Session, registry *prometheus.Registry, db *database.Database) {
	client.AddHandler(events.NewReadyEvent(logger.With(zap.String("module", "events"), zap.String("event", "ready")), client, registry, db).Run)
	client.AddHandler(events.NewInteractionCreateEvent(logger.With(zap.String("module", "events"), zap.String("event", "interaction_create")), client, registry, db).Run)
	client.AddHandler(events.NewMessageCreateEvent(logger.With(zap.String("module", "events"), zap.String("event", "message_create")), client, registry, db).Run)
	client.AddHandler(events.NewMessageDeleteEvent(logger.With(zap.String("module", "events"), zap.String("event", "message_delete")), client, registry, db).Run)
	client.AddHandler(events.NewMessageUpdateEvent(logger.With(zap.String("module", "events"), zap.String("event", "message_update")), client, registry, db).Run)
}

func doCommands(logger *zap.Logger, registry *prometheus.Registry, db *database.Database) {
	registers.InteractionCommands.Register(interaction.NewBotInfo(logger.With(zap.String("module", "general"), zap.String("command", "info")), db))
}

func doPrometheus(registry *prometheus.Registry) {
	registry.MustRegister(collectors.NewGoCollector(),
		collectors.NewProcessCollector(collectors.ProcessCollectorOpts{}))

	http.Handle(
		"/metrics", promhttp.HandlerFor(
			registry,
			promhttp.HandlerOpts{
				EnableOpenMetrics: true,
			}),
	)

	log.Fatalln(http.ListenAndServe(":8080", nil))
}
