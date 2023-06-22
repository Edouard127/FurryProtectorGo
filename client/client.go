package client

import (
	"github.com/Edouard127/FurryProtectorGo/client/interaction/general"
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

type Client struct {
	*zap.Logger
	*discordgo.Session
	InteractionCommands *registers.RunnerRegister[discordgo.InteractionCreate]
}

func NewClient(logger *zap.Logger, token string) (c *Client, err error) {
	session, err := discordgo.New("Bot " + token)
	if err != nil {
		return nil, err
	}

	return doPreInit(logger, &Client{
		Logger:              logger,
		Session:             session,
		InteractionCommands: registers.NewRegister[discordgo.InteractionCreate](),
	}), nil
}

func doPreInit(logger *zap.Logger, client *Client) *Client {
	registry := prometheus.NewRegistry()

	go doPrometheus(registry)
	doEvents(logger, client, registry)
	doCommands(logger, client, registry)

	endpoint := discordgo.EndpointApplicationGlobalCommands(os.Getenv("APP_ID"))

	for _, command := range client.InteractionCommands.Runners {
		_, err := client.RequestWithBucketID("POST", endpoint, command, endpoint)
		if err != nil {
			command.GetLogger().Error("Error while registering command", zap.Error(err))
		}
	}

	return client
}

func doEvents(logger *zap.Logger, client *Client, registry *prometheus.Registry) {
	client.AddHandler(NewReadyEvent(logger.With(zap.String("module", "events"), zap.String("event", "ready")), client, registry).Run)
	client.AddHandler(NewInteractionCreateEvent(logger.With(zap.String("module", "events"), zap.String("event", "interaction_create")), client, registry).Run)
	client.AddHandler(NewMessageCreateEvent(logger.With(zap.String("module", "events"), zap.String("event", "message_create")), client, registry).Run)
}

func doCommands(logger *zap.Logger, client *Client, registry *prometheus.Registry) {
	client.InteractionCommands.Register(general.NewBotInfo(logger.With(zap.String("module", "general"), zap.String("command", "info"))))
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

func (c *Client) UserCount() int {
	var count int

	for _, guild := range c.State.Guilds {
		count += guild.MemberCount
	}

	return count
}
