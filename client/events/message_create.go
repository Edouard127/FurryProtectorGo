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

type MessageCreateEvent struct {
	*zap.Logger
	*discordgo.Session
	*database.Database
}

func NewMessageCreateEvent(logger *zap.Logger, client *discordgo.Session, db *database.Database) *MessageCreateEvent {
	return &MessageCreateEvent{logger, client, db}
}

func (m *MessageCreateEvent) Run(_ *discordgo.Session, message *discordgo.MessageCreate) {
	if message.Author.Bot {
		return
	}

	exporter.MessageCreateCounter.With(prometheus.Labels{"guild": message.GuildID, "channel": message.ChannelID, "user": message.Author.ID}).Inc()

	exporter.MessageTest.With(prometheus.Labels{"guild": message.GuildID}).Observe(1)

	var user data.UserData
	err := m.Read("users", bson.M{"id": message.Author.ID}, &user)
	if err != nil {
		m.Log(zap.DebugLevel, "New user", zap.String("user", message.Author.ID))
		user = data.NewUserData(message.Author.ID)
	}
	m.Write("users", user.AddMessage(data.NewUserMessage(message.Content, message.GuildID, message.ChannelID, message.ID, message.Timestamp.UnixMilli())))
}
