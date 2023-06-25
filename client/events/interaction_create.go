package events

import (
	"github.com/Edouard127/FurryProtectorGo/client/database"
	"github.com/Edouard127/FurryProtectorGo/client/exporter"
	"github.com/Edouard127/FurryProtectorGo/registers"
	"github.com/bwmarrin/discordgo"
	"github.com/prometheus/client_golang/prometheus"
	"go.uber.org/zap"
	"time"
)

type InteractionCreateEvent struct {
	*zap.Logger
	*discordgo.Session
	*database.Database
}

func NewInteractionCreateEvent(logger *zap.Logger, client *discordgo.Session, db *database.Database) *InteractionCreateEvent {
	return &InteractionCreateEvent{logger, client, db}
}

func (i *InteractionCreateEvent) Run(session *discordgo.Session, ctx *discordgo.InteractionCreate) {
	i.Info("Interaction received", zap.String("name", ctx.ApplicationCommandData().Name))
	var before = time.Now()

	switch ctx.Type {
	case discordgo.InteractionMessageComponent:
		registers.InteractionComponents.Get(ctx.MessageComponentData().CustomID).Run(session, ctx)
	case discordgo.InteractionApplicationCommand:
		registers.InteractionCommands.Get(ctx.ApplicationCommandData().Name).Run(session, ctx)
	case discordgo.InteractionModalSubmit:
		registers.InteractionModals.Get(ctx.MessageComponentData().CustomID).Run(session, ctx)
	}

	exporter.InterationHist.With(prometheus.Labels{"guild": ctx.GuildID, "user": ctx.Member.User.ID, "interaction": ctx.ApplicationCommandData().Name}).Observe(time.Since(before).Seconds())
}
