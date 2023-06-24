package events

import (
	"github.com/Edouard127/FurryProtectorGo/client/database"
	"github.com/Edouard127/FurryProtectorGo/client/exporter"
	"github.com/bwmarrin/discordgo"
	"github.com/prometheus/client_golang/prometheus"
	"go.uber.org/zap"
)

type MemberEventJoin struct {
	*zap.Logger
	*discordgo.Session
	*database.Database
}

func NewMemberJoinEvent(logger *zap.Logger, client *discordgo.Session, db *database.Database) *MemberEventJoin {
	return &MemberEventJoin{logger, client, db}
}

func (m *MemberEventJoin) Run(_ *discordgo.Session, event *discordgo.GuildMemberAdd) {
	exporter.MemberGauge.With(prometheus.Labels{"guild": event.GuildID}).Inc()
	exporter.MemberJoinCounter.With(prometheus.Labels{"guild": event.GuildID}).Inc()
}
