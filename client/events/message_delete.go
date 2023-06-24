package events

import (
	"github.com/Edouard127/FurryProtectorGo/client/database"
	"github.com/Edouard127/FurryProtectorGo/client/exporter"
	"github.com/bwmarrin/discordgo"
	"github.com/prometheus/client_golang/prometheus"
	"go.uber.org/zap"
)

type MessageDeleteEvent struct {
	*zap.Logger
	*discordgo.Session
	*database.Database
}

func NewMessageDeleteEvent(logger *zap.Logger, client *discordgo.Session, db *database.Database) *MessageDeleteEvent {
	return &MessageDeleteEvent{logger, client, db}
}

func (m *MessageDeleteEvent) Run(_ *discordgo.Session, message *discordgo.MessageDelete) {
	exporter.MessageDeleteCounter.With(prometheus.Labels{"guild": message.GuildID, "channel": message.ChannelID}).Inc()
}
