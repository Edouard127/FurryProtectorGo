package events

import (
	"github.com/Edouard127/FurryProtectorGo/client/database"
	"github.com/Edouard127/FurryProtectorGo/client/exporter"
	"github.com/bwmarrin/discordgo"
	"github.com/prometheus/client_golang/prometheus"
	"go.uber.org/zap"
)

func NewMessageUpdateEvent(logger *zap.Logger, db *database.Database) func(*discordgo.Session, *discordgo.MessageUpdate) {
	return func(session *discordgo.Session, ctx *discordgo.MessageUpdate) {
		if ctx.BeforeUpdate == nil || ctx.Content == ctx.BeforeUpdate.Content {
			return
		}

		exporter.MessageUpdateCounter.With(prometheus.Labels{"guild": ctx.GuildID, "channel": ctx.ChannelID, "user": ctx.Author.ID}).Inc()
	}
}
