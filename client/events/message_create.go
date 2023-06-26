package events

import (
	"github.com/Edouard127/FurryProtectorGo/client/database"
	"github.com/Edouard127/FurryProtectorGo/client/exporter"
	"github.com/Edouard127/FurryProtectorGo/core/data"
	"github.com/bwmarrin/discordgo"
	"github.com/prometheus/client_golang/prometheus"
	"go.mongodb.org/mongo-driver/bson"
	"go.uber.org/zap"
)

func NewMessageCreateEvent(logger *zap.Logger, db *database.Database) func(*discordgo.Session, *discordgo.MessageCreate) {
	return func(session *discordgo.Session, ctx *discordgo.MessageCreate) {
		if ctx.Author.Bot {
			return
		}

		exporter.MessageCreateCounter.With(prometheus.Labels{"guild": ctx.GuildID, "channel": ctx.ChannelID, "user": ctx.Author.ID}).Inc()

		exporter.MessageTest.With(prometheus.Labels{"guild": ctx.GuildID}).Observe(1)

		var user data.UserData
		err := db.Read("users", bson.M{"id": ctx.Author.ID}, &user)
		if err != nil {
			logger.Debug("New user", zap.String("user", ctx.Author.ID))
			user = data.NewUserData(ctx.Author.ID)
		}
		db.Write("users", user.AddMessage(data.NewUserMessage(ctx.Content, ctx.GuildID, ctx.ChannelID, ctx.ID, ctx.Timestamp.UnixMilli())))
	}
}
