package events

import (
	"fmt"
	"github.com/Edouard127/FurryProtectorGo/client/database"
	"github.com/bwmarrin/discordgo"
	"github.com/prometheus/client_golang/prometheus"
	"go.uber.org/zap"
)

type ReadyEvent struct {
	*zap.Logger
	*discordgo.Session
	*prometheus.Registry
	*database.Database
}

func NewReadyEvent(logger *zap.Logger, client *discordgo.Session, registry *prometheus.Registry, db *database.Database) *ReadyEvent {
	return &ReadyEvent{logger, client, registry, db}
}

func (r *ReadyEvent) Run(session *discordgo.Session, ready *discordgo.Ready) {
	r.Info(fmt.Sprintf("Logged in as %s", r.State.User.String()))
	err := r.UpdateStatusComplex(discordgo.UpdateStatusData{
		Activities: []*discordgo.Activity{
			{
				Name: fmt.Sprintf("%d guilds", len(r.State.Guilds)),
				Type: discordgo.ActivityTypeWatching,
			},
		},
		Status: string(discordgo.StatusDoNotDisturb),
	})
	if err != nil {
		r.Error("Error while updating status", zap.Error(err))
	}
}
