package commands

import (
	"fmt"
	"github.com/Edouard127/FurryProtectorGo/client/database"
	"github.com/Edouard127/FurryProtectorGo/core/builder/components/embed"
	"github.com/Edouard127/FurryProtectorGo/core/builder/interaction"
	"github.com/Edouard127/FurryProtectorGo/core/data"
	"github.com/bwmarrin/discordgo"
	"go.uber.org/zap"
)

type AddEmoji struct {
	*zap.Logger
	*interaction.SlashInteractionBuilder
}

func NewAddEmoji(logger *zap.Logger, db *database.Database) (string, *AddEmoji) {
	return "addemoji", &AddEmoji{Logger: logger, SlashInteractionBuilder: interaction.NewSlashInteractionBuilder("addemoji", "Add an emoji to the server").
		AddOption(
			interaction.NewSlashInteractionStringOption("emoji", "The emoji").SetRequired(true),
			interaction.NewSlashInteractionStringOption("name", "The name of the emoji")).
		SetDefaultPermission(data.ManageGuildExpressions)}
}

func (a *AddEmoji) GetLogger() *zap.Logger {
	return a.Logger
}

func (a *AddEmoji) Run(client *discordgo.Session, ctx *discordgo.InteractionCreate) {
	emoji := ctx.ApplicationCommandData().Options[0].StringValue()

	e, ok := data.ParseEmoji(emoji)
	if !ok {
		client.InteractionRespond(ctx.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Embeds: []*discordgo.MessageEmbed{
					embed.NewEmbedBuilder().
						SetTitle("Add an emoji").
						SetDescription("The emoji is invalid").
						SetFooter(embed.NewEmbedFooter("Requested by " + ctx.Member.User.Username).SetIconURL(ctx.Member.AvatarURL("256"))).
						Build(),
				},
			},
		})
		return
	}

	if len(ctx.ApplicationCommandData().Options) > 1 {
		e.Name = ctx.ApplicationCommandData().Options[1].StringValue()
	}

	newEmoji, err := client.GuildEmojiCreate(ctx.GuildID, e.API()) // TODO: Check permission
	if err != nil {
		a.Logger.Error("Failed to create emoji", zap.Error(err))
		client.InteractionRespond(ctx.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Embeds: []*discordgo.MessageEmbed{
					embed.NewEmbedBuilder().
						SetTitle("Add an emoji").
						SetDescription("Failed to create emoji").
						SetFooter(embed.NewEmbedFooter("Requested by " + ctx.Member.User.Username).SetIconURL(ctx.Member.AvatarURL("256"))).
						Build(),
				},
			},
		})
		return
	}

	client.InteractionRespond(ctx.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Embeds: []*discordgo.MessageEmbed{
				embed.NewEmbedBuilder().
					SetTitle("Add an emoji").
					SetDescription(fmt.Sprintf("Emoji added ! %s", newEmoji.MessageFormat())).
					SetFooter(embed.NewEmbedFooter("Requested by " + ctx.Member.User.Username).SetIconURL(ctx.Member.AvatarURL("256"))).
					Build(),
			},
		},
	})
}
