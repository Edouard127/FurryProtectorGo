package events

import (
	"github.com/Edouard127/FurryProtectorGo/client/database"
	"github.com/Edouard127/FurryProtectorGo/core/data"
	"github.com/bwmarrin/discordgo"
	"github.com/prometheus/client_golang/prometheus"
	"go.mongodb.org/mongo-driver/bson"
	"go.uber.org/zap"
)

type MessageCreateEvent struct {
	*zap.Logger
	*discordgo.Session
	*prometheus.Registry
	*database.Database
	messageCounter *prometheus.CounterVec
}

func NewMessageCreateEvent(logger *zap.Logger, client *discordgo.Session, registry *prometheus.Registry, db *database.Database) *MessageCreateEvent {
	mCounter := prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "discord_messages_number",
		Help: "The number of messages received by guild by channel per user",
	}, []string{"guild", "channel", "user"})

	registry.MustRegister(mCounter)

	return &MessageCreateEvent{logger, client, registry, db, mCounter}
}

func (m *MessageCreateEvent) Run(_ *discordgo.Session, message *discordgo.MessageCreate) {
	m.messageCounter.With(prometheus.Labels{"guild": message.GuildID, "channel": message.ChannelID, "user": message.Author.ID}).Inc()

	if message.Author.Bot {
		return
	}

	var user *data.UserData
	err := m.Read("users", bson.M{"id": message.Author.ID}, &user)
	if err != nil {
		m.Log(zap.DebugLevel, "New user", zap.String("user", message.Author.ID))
		user = data.NewUserData(message.Author.ID)
	}
	m.Write("users", user.AddMessage(data.NewUserMessage(message.Content, message.GuildID, message.ChannelID, message.ID, message.Timestamp.UnixMilli())))
}
