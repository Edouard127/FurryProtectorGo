package events

import (
	"github.com/bwmarrin/discordgo"
	"github.com/prometheus/client_golang/prometheus"
	"go.uber.org/zap"
)

type MessageCreateEvent struct {
	*zap.Logger
	*discordgo.Session
	*prometheus.Registry
	messageCounter *prometheus.CounterVec
}

func NewMessageCreateEvent(logger *zap.Logger, client *discordgo.Session, registry *prometheus.Registry) *MessageCreateEvent {
	mCounter := prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "discord_messages_number",
		Help: "The number of messages received by guild by channel per user",
	}, []string{"guild", "channel", "user"})

	registry.MustRegister(mCounter)

	return &MessageCreateEvent{logger, client, registry, mCounter}
}

func (m *MessageCreateEvent) Run(_ *discordgo.Session, message *discordgo.MessageCreate) {
	m.messageCounter.With(prometheus.Labels{"guild": message.GuildID, "channel": message.ChannelID, "user": message.Author.ID}).Inc()
}
