package events

import (
	"github.com/Edouard127/FurryProtectorGo/client/database"
	"github.com/bwmarrin/discordgo"
	"github.com/prometheus/client_golang/prometheus"
	"go.uber.org/zap"
)

type MessageDeleteEvent struct {
	*zap.Logger
	*discordgo.Session
	*prometheus.Registry
	*database.Database
	messageCounter *prometheus.CounterVec
}

func NewMessageDeleteEvent(logger *zap.Logger, client *discordgo.Session, registry *prometheus.Registry, db *database.Database) *MessageDeleteEvent {
	mCounter := prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "discord_messages_delete_number",
		Help: "The number of messages deleted by guild by channel",
	}, []string{"guild", "channel"})

	registry.MustRegister(mCounter)

	return &MessageDeleteEvent{logger, client, registry, db, mCounter}
}

func (m *MessageDeleteEvent) Run(_ *discordgo.Session, message *discordgo.MessageDelete) {
	m.messageCounter.With(prometheus.Labels{"guild": message.GuildID, "channel": message.ChannelID}).Inc()
}
