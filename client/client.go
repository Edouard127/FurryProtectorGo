package client

import (
	"github.com/Edouard127/FurryProtectorGo/client/database"
	"github.com/Edouard127/FurryProtectorGo/client/events"
	"github.com/Edouard127/FurryProtectorGo/client/exporter"
	"github.com/Edouard127/FurryProtectorGo/client/interaction/commands"
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

	exporter.DoRegister(registry)
	go doPrometheus(registry)
	doEvents(logger, client, db)
	doCommands(logger, db)

	endpoint := discordgo.EndpointApplicationGlobalCommands(os.Getenv("APP_ID"))

	for _, command := range registers.InteractionCommands.Runners {
		_, err := client.RequestWithBucketID("POST", endpoint, command, endpoint)
		if err != nil {
			command.GetLogger().Error("Error while registering command", zap.Error(err))
		}
	}

	return client
}

func doEvents(logger *zap.Logger, client *discordgo.Session, db *database.Database) {
	client.AddHandler(events.NewInteractionCreateEvent(logger.With(zap.String("module", "events"), zap.String("event", "interaction_create")), client, db).Run)
	client.AddHandler(events.NewMemberJoinEvent(logger.With(zap.String("module", "events"), zap.String("event", "member_join")), client, db).Run)
	client.AddHandler(events.NewMemberDeleteEvent(logger.With(zap.String("module", "events"), zap.String("event", "member_leave")), client, db).Run)
	client.AddHandler(events.NewMessageCreateEvent(logger.With(zap.String("module", "events"), zap.String("event", "message_create")), client, db).Run)
	client.AddHandler(events.NewMessageDeleteEvent(logger.With(zap.String("module", "events"), zap.String("event", "message_delete")), client, db).Run)
	client.AddHandler(events.NewMessageUpdateEvent(logger.With(zap.String("module", "events"), zap.String("event", "message_update")), client, db).Run)
	client.AddHandler(events.NewReadyEvent(logger.With(zap.String("module", "events"), zap.String("event", "ready")), client, db).Run)
}

func doCommands(logger *zap.Logger, db *database.Database) {
	registers.InteractionCommands.Register(commands.NewBotInfo(logger.With(zap.String("module", "general"), zap.String("command", "info")), db))
	registers.InteractionCommands.Register(commands.NewAddEmoji(logger.With(zap.String("module", "general"), zap.String("command", "add_emoji")), db))
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
