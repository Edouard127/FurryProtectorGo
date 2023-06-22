package events

import (
	"github.com/Edouard127/FurryProtectorGo/client/database"
	"github.com/Edouard127/FurryProtectorGo/registers"
	"github.com/bwmarrin/discordgo"
	"github.com/prometheus/client_golang/prometheus"
	"go.uber.org/zap"
)

type InteractionCreateEvent struct {
	*zap.Logger
	*discordgo.Session
	*prometheus.Registry
	*database.Database
	interactionCounter *prometheus.CounterVec
}

func NewInteractionCreateEvent(logger *zap.Logger, client *discordgo.Session, registry *prometheus.Registry, db *database.Database) *InteractionCreateEvent {
	iCounter := prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "discord_interactions_request_number",
		Help: "The number of interaction received by guild per user",
	}, []string{"guild", "user", "interaction"})

	registry.MustRegister(iCounter)

	return &InteractionCreateEvent{logger, client, registry, db, iCounter}
}

func (i *InteractionCreateEvent) Run(session *discordgo.Session, ctx *discordgo.InteractionCreate) {
	i.Info("Interaction received", zap.String("name", ctx.ApplicationCommandData().Name))
	i.interactionCounter.With(prometheus.Labels{"guild": ctx.GuildID, "user": ctx.Member.User.ID, "interaction": ctx.ApplicationCommandData().Name}).Inc()
	registers.InteractionCommands.Get(ctx.ApplicationCommandData().Name).Run(session, ctx)
}
