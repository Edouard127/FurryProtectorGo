package client

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/prometheus/client_golang/prometheus"
	"go.uber.org/zap"
)

type ReadyEvent struct {
	*zap.Logger
	*Client
	*prometheus.Registry
}

func NewReadyEvent(logger *zap.Logger, client *Client, registry *prometheus.Registry) *ReadyEvent {
	return &ReadyEvent{logger, client, registry}
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
