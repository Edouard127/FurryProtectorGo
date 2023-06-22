package client

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"go.uber.org/zap"
)

type ReadyEvent struct {
	*zap.Logger
	*Client
}

func NewReadyEvent(logger *zap.Logger, client *Client) *ReadyEvent {
	return &ReadyEvent{logger, client}
}

func (r *ReadyEvent) Run(session *discordgo.Session, ready *discordgo.Ready) {
	r.Info(fmt.Sprintf("Logged in as %s", r.State.User.String()))
}
