package events

import (
	"github.com/Edouard127/FurryProtectorGo/client/database"
	"github.com/Edouard127/FurryProtectorGo/client/exporter"
	"github.com/bwmarrin/discordgo"
	"github.com/prometheus/client_golang/prometheus"
	"go.uber.org/zap"
)

type MemberEventDelete struct {
	*zap.Logger
	*discordgo.Session
	*database.Database
}

func NewMemberDeleteEvent(logger *zap.Logger, client *discordgo.Session, db *database.Database) *MemberEventDelete {
	return &MemberEventDelete{logger, client, db}
}

func (m *MemberEventDelete) Run(_ *discordgo.Session, event *discordgo.GuildMemberAdd) {
	exporter.MemberGauge.With(prometheus.Labels{"guild": event.GuildID}).Dec()
	exporter.MemberDeleteCounter.With(prometheus.Labels{"guild": event.GuildID}).Inc()
	m.Logger.Debug("Member deleted", zap.String("guild", event.GuildID), zap.String("member", event.Member.User.ID))
}
