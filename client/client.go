package client

import (
	"github.com/Edouard127/FurryProtectorGo/client/database"
	"github.com/Edouard127/FurryProtectorGo/client/events"
	"github.com/Edouard127/FurryProtectorGo/client/exporter"
	"github.com/Edouard127/FurryProtectorGo/client/interaction/commands"
	"github.com/Edouard127/FurryProtectorGo/core/builder/interaction"
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

func NewClient(logger *zap.Logger, token string) {
	client, err := discordgo.New("Bot " + token)
	if err != nil {
		logger.Panic("Error while creating bot", zap.Error(err))
	}

	client.Identify.Intents = discordgo.IntentsAll

	StartClient(logger, doPreInit(logger, client), exporter.DoRegister(prometheus.NewRegistry()))
}

func StartClient(logger *zap.Logger, client *discordgo.Session, registry *prometheus.Registry) {
	err := client.Open()
	if err != nil {
		logger.Panic("Error while opening bot", zap.Error(err))
	}

	doPrometheus(registry) // blocking call
}

func doPreInit(logger *zap.Logger, client *discordgo.Session) *discordgo.Session {
	db := database.NewDatabase(logger.With(zap.String("module", "database")), os.Getenv("MONGO_URI"), os.Getenv("DATABASE_NAME"), "config", "users", "verification_cache")

	doEvents(logger, client, db)
	doCommands(logger, client, db)

	return client
}

func doEvents(logger *zap.Logger, client *discordgo.Session, db *database.Database) {
	client.AddHandler(events.NewInteractionCreateEvent(logger.With(zap.String("module", "events"), zap.String("event", "interaction_create")), db))
	client.AddHandler(events.NewMemberJoinEvent(logger.With(zap.String("module", "events"), zap.String("event", "member_join")), db))
	client.AddHandler(events.NewMemberDeleteEvent(logger.With(zap.String("module", "events"), zap.String("event", "member_leave")), db))
	client.AddHandler(events.NewMessageCreateEvent(logger.With(zap.String("module", "events"), zap.String("event", "message_create")), db))
	client.AddHandler(events.NewMessageDeleteEvent(logger.With(zap.String("module", "events"), zap.String("event", "message_delete")), db))
	client.AddHandler(events.NewMessageUpdateEvent(logger.With(zap.String("module", "events"), zap.String("event", "message_update")), db))
	client.AddHandler(events.NewReadyEvent(logger.With(zap.String("module", "events"), zap.String("event", "ready")), db))
}

func commandRegister(logger *zap.Logger, client *discordgo.Session) (func(*interaction.SlashInteractionBuilder, interaction.Runner[discordgo.InteractionCreate]) (string, interaction.Runner[discordgo.InteractionCreate]), func()) {
	var builders = make([]*discordgo.ApplicationCommand, 0)

	return func(builder *interaction.SlashInteractionBuilder, runner interaction.Runner[discordgo.InteractionCreate]) (string, interaction.Runner[discordgo.InteractionCreate]) {
			builders = append(builders, builder.Build())
			return builder.Name, runner
		}, func() {
			for _, builder := range builders {
				_, err := client.ApplicationCommandCreate(os.Getenv("APP_ID"), "", builder)
				if err != nil {
					logger.Error("Error while registering command", zap.String("command", builder.Name), zap.Error(err))
				}
			}
		}
}

func doCommands(logger *zap.Logger, client *discordgo.Session, db *database.Database) {
	registerer, done := commandRegister(logger, client)

	registers.InteractionCommands.Register(registerer(commands.NewAddEmoji(logger, db)))
	registers.InteractionCommands.Register(registerer(commands.NewBotInfo(logger, db)))
	registers.InteractionCommands.Register(registerer(commands.NewSetup(logger, db)))

	done()
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
