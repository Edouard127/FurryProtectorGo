package commands

import (
	"github.com/Edouard127/FurryProtectorGo/client/database"
	"github.com/Edouard127/FurryProtectorGo/core/builder/interaction"
	"github.com/bwmarrin/discordgo"
	"go.uber.org/zap"
)

func NewSetup(logger *zap.Logger, db *database.Database) (*interaction.SlashInteractionBuilder, interaction.Runner[discordgo.InteractionCreate]) {
	return interaction.NewSlashInteractionBuilder("setup", "Setup the bot"), runSetup(logger)
}

func runSetup(logger *zap.Logger) interaction.Runner[discordgo.InteractionCreate] {
	return func(client *discordgo.Session, ctx *discordgo.InteractionCreate) error {
		return client.InteractionRespond(ctx.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: "This command is not implemented yet.",
			},
		})
	}
}
