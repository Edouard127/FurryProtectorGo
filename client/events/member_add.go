package events

import (
	"github.com/Edouard127/FurryProtectorGo/client/database"
	"github.com/Edouard127/FurryProtectorGo/client/exporter"
	"github.com/bwmarrin/discordgo"
	"github.com/prometheus/client_golang/prometheus"
	"go.uber.org/zap"
)

func NewMemberJoinEvent(logger *zap.Logger, db *database.Database) func(*discordgo.Session, *discordgo.GuildMemberAdd) {
	return func(session *discordgo.Session, event *discordgo.GuildMemberAdd) {
		exporter.MemberGauge.With(prometheus.Labels{"guild": event.GuildID}).Inc()
		exporter.MemberJoinCounter.With(prometheus.Labels{"guild": event.GuildID}).Inc()
		logger.Debug("Member joined", zap.String("guild", event.GuildID), zap.String("member", event.Member.User.ID))
	}
}
