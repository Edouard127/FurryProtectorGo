package client

import (
	"github.com/bwmarrin/discordgo"
	"go.uber.org/zap"
)

type InteractionCreateEvent struct {
	*zap.Logger
	*Client
}

func NewInteractionCreateEvent(logger *zap.Logger, c *Client) *InteractionCreateEvent {
	return &InteractionCreateEvent{logger, c}
}

func (i *InteractionCreateEvent) Run(session *discordgo.Session, ctx *discordgo.InteractionCreate) {
	i.Info("Interaction received", zap.String("name", ctx.ApplicationCommandData().Name))
	i.InteractionCommands.Get(ctx.ApplicationCommandData().Name).Run(session, ctx)
}
