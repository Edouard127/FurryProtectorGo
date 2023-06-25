package events

import (
	"fmt"
	"github.com/Edouard127/FurryProtectorGo/client/database"
	"github.com/bwmarrin/discordgo"
	"go.uber.org/zap"
)

type ReadyEvent struct {
	*zap.Logger
	*discordgo.Session
	*database.Database
}

func NewReadyEvent(logger *zap.Logger, client *discordgo.Session, db *database.Database) *ReadyEvent {
	return &ReadyEvent{logger, client, db}
}

func (r *ReadyEvent) Run(session *discordgo.Session, ready *discordgo.Ready) {
	r.Info(fmt.Sprintf("Logged in as %s", r.State.User.String()))
	r.UpdateStatusComplex(discordgo.UpdateStatusData{
		Activities: []*discordgo.Activity{
			{
				Name: fmt.Sprintf("%d guilds", len(r.State.Guilds)),
				Type: discordgo.ActivityTypeWatching,
			},
		},
		Status: string(discordgo.StatusDoNotDisturb),
	})
}
