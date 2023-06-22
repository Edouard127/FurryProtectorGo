package events

import (
	"github.com/Edouard127/FurryProtectorGo/client/database"
	"github.com/bwmarrin/discordgo"
	"github.com/prometheus/client_golang/prometheus"
	"go.uber.org/zap"
)

type MessageUpdateEvent struct {
	*zap.Logger
	*discordgo.Session
	*prometheus.Registry
	*database.Database
	messageCounter *prometheus.CounterVec
}

func NewMessageUpdateEvent(logger *zap.Logger, client *discordgo.Session, registry *prometheus.Registry, db *database.Database) *MessageUpdateEvent {
	mCounter := prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "discord_messages_update_number",
		Help: "The number of messages updated by guild by channel per user",
	}, []string{"guild", "channel", "user"})

	registry.MustRegister(mCounter)

	return &MessageUpdateEvent{logger, client, registry, db, mCounter}
}

func (m *MessageUpdateEvent) Run(_ *discordgo.Session, message *discordgo.MessageUpdate) {
	m.messageCounter.With(prometheus.Labels{"guild": message.GuildID, "channel": message.ChannelID, "user": message.Author.ID}).Inc()
}
