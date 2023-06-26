package events

import (
	"github.com/Edouard127/FurryProtectorGo/client/database"
	"github.com/Edouard127/FurryProtectorGo/client/exporter"
	"github.com/Edouard127/FurryProtectorGo/core/data"
	"github.com/bwmarrin/discordgo"
	"github.com/prometheus/client_golang/prometheus"
	"go.uber.org/zap"
)

func NewMessageCreateEvent(logger *zap.Logger, db *database.Database) func(*discordgo.Session, *discordgo.MessageCreate) {
	var users = database.NewMongoCache[data.UserData](db, "users", 5000)

	return func(session *discordgo.Session, ctx *discordgo.MessageCreate) {
		if ctx.Author.Bot {
			return
		}

		exporter.MessageCreateCounter.With(prometheus.Labels{"guild": ctx.GuildID, "channel": ctx.ChannelID, "user": ctx.Author.ID}).Inc()

		exporter.MessageTest.With(prometheus.Labels{"guild": ctx.GuildID}).Observe(1)

		users.Set(ctx.Author.ID, users.Get(ctx.Author.ID).AddMessage(data.NewUserMessage(ctx.Content, ctx.GuildID, ctx.ChannelID, ctx.ID, ctx.Timestamp.UnixMilli())))
	}
}
