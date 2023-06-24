package events

import (
	"github.com/Edouard127/FurryProtectorGo/client/database"
	"github.com/Edouard127/FurryProtectorGo/client/exporter"
	"github.com/bwmarrin/discordgo"
	"github.com/prometheus/client_golang/prometheus"
	"go.uber.org/zap"
)

type MessageUpdateEvent struct {
	*zap.Logger
	*discordgo.Session
	*database.Database
}

func NewMessageUpdateEvent(logger *zap.Logger, client *discordgo.Session, db *database.Database) *MessageUpdateEvent {
	return &MessageUpdateEvent{logger, client, db}
}

func (m *MessageUpdateEvent) Run(_ *discordgo.Session, message *discordgo.MessageUpdate) {
	if message.BeforeUpdate == nil || message.Content == message.BeforeUpdate.Content {
		return
	}

	exporter.MessageUpdateCounter.With(prometheus.Labels{"guild": message.GuildID, "channel": message.ChannelID, "user": message.Author.ID}).Inc()
}
