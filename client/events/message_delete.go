package events

import (
	"github.com/Edouard127/FurryProtectorGo/client/database"
	"github.com/Edouard127/FurryProtectorGo/client/exporter"
	"github.com/bwmarrin/discordgo"
	"github.com/prometheus/client_golang/prometheus"
	"go.uber.org/zap"
)

func NewMessageDeleteEvent(logger *zap.Logger, db *database.Database) func(*discordgo.Session, *discordgo.MessageDelete) {
	return func(session *discordgo.Session, ctx *discordgo.MessageDelete) {
		exporter.MessageDeleteCounter.With(prometheus.Labels{"guild": ctx.GuildID, "channel": ctx.ChannelID}).Inc()
	}
}
