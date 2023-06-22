package events

import (
	"github.com/bwmarrin/discordgo"
	"github.com/prometheus/client_golang/prometheus"
	"go.uber.org/zap"
)

type MessageUpdateEvent struct {
	*zap.Logger
	*discordgo.Session
	*prometheus.Registry
	messageCounter *prometheus.CounterVec
}

func NewMessageUpdateEvent(logger *zap.Logger, client *discordgo.Session, registry *prometheus.Registry) *MessageUpdateEvent {
	mCounter := prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "discord_messages_update_number",
		Help: "The number of messages updated by guild by channel per user",
	}, []string{"guild", "channel", "user"})

	registry.MustRegister(mCounter)

	return &MessageUpdateEvent{logger, client, registry, mCounter}
}

func (m *MessageUpdateEvent) Run(_ *discordgo.Session, message *discordgo.MessageUpdate) {
	m.messageCounter.With(prometheus.Labels{"guild": message.GuildID, "channel": message.ChannelID, "user": message.Author.ID}).Inc()
}
