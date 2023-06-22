package client

import (
	"github.com/Edouard127/FurryProtectorGo/client/interaction/general"
	"github.com/Edouard127/FurryProtectorGo/registers"
	"github.com/bwmarrin/discordgo"
	"go.uber.org/zap"
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
	doEvents(logger, client)
	doCommands(logger, client)

	endpoint := discordgo.EndpointApplicationGlobalCommands(os.Getenv("APP_ID"))

	for _, command := range client.InteractionCommands.Runners {
		_, err := client.RequestWithBucketID("POST", endpoint, command, endpoint)
		if err != nil {
			command.GetLogger().Error("Error while registering command", zap.Error(err))
		}
	}

	return client
}

func doEvents(logger *zap.Logger, client *Client) {
	client.AddHandler(NewReadyEvent(logger.With(zap.String("module", "events"), zap.String("event", "ready")), client).Run)
	client.AddHandler(NewInteractionCreateEvent(logger.With(zap.String("module", "events"), zap.String("event", "interaction_create")), client).Run)
}

func doCommands(logger *zap.Logger, client *Client) {
	client.InteractionCommands.Register(general.NewBotInfo(logger.With(zap.String("module", "general"), zap.String("command", "info"))))
}

func (c *Client) Users() []*discordgo.User {
	var users []*discordgo.User
	for _, guild := range c.State.Guilds {
		for _, member := range guild.Members {
			users = append(users, member.User)
		}
	}
	return users
}
