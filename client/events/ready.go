package events

import (
	"fmt"
	"github.com/Edouard127/FurryProtectorGo/client/database"
	"github.com/bwmarrin/discordgo"
	"go.uber.org/zap"
)

func NewReadyEvent(logger *zap.Logger, db *database.Database) func(*discordgo.Session, *discordgo.Ready) {
	return func(session *discordgo.Session, ctx *discordgo.Ready) {
		logger.Info(fmt.Sprintf("Logged in as %s", session.State.User.String()))
		session.UpdateStatusComplex(discordgo.UpdateStatusData{
			Activities: []*discordgo.Activity{
				{
					Name: fmt.Sprintf("%d guilds", len(session.State.Guilds)),
					Type: discordgo.ActivityTypeWatching,
				},
			},
			Status: string(discordgo.StatusDoNotDisturb),
		})
	}
}
