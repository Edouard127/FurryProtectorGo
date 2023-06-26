package commands

import (
	"github.com/Edouard127/FurryProtectorGo/client/database"
	"github.com/Edouard127/FurryProtectorGo/core/builder/components/embed"
	"github.com/Edouard127/FurryProtectorGo/core/builder/interaction"
	"github.com/Edouard127/FurryProtectorGo/core/data"
	"github.com/Edouard127/FurryProtectorGo/i18n"
	"github.com/bwmarrin/discordgo"
	"go.uber.org/zap"
)

func NewAddEmoji(logger *zap.Logger, db *database.Database) (*interaction.SlashInteractionBuilder, interaction.Runner[discordgo.InteractionCreate]) {
	return interaction.NewSlashInteractionBuilder("addemoji", "Add an emoji to the server").
		AddOption(
			interaction.NewSlashInteractionStringOption("emoji", "The emoji").SetRequired(true),
			interaction.NewSlashInteractionStringOption("name", "The name of the emoji")).
		SetDefaultPermission(data.ManageGuildExpressions), runAddEmoji(logger)
}

func runAddEmoji(logger *zap.Logger) interaction.Runner[discordgo.InteractionCreate] {
	return func(client *discordgo.Session, ctx *discordgo.InteractionCreate) error {
		emoji := ctx.ApplicationCommandData().Options[0].StringValue()

		permission, _ := client.State.UserChannelPermissions(client.State.User.ID, ctx.ChannelID)
		if data.UserPermission(permission)&data.ManageGuildExpressions != data.ManageGuildExpressions {
			return client.InteractionRespond(ctx.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Embeds: []*discordgo.MessageEmbed{
						embed.NewEmbedBuilder().
							SetDescription(i18n.Translate("MissingPermission", *ctx.GuildLocale, data.ManageGuildExpressions)).
							SetFooter(embed.NewEmbedFooter(i18n.Translate("RequestedBy", *ctx.GuildLocale, ctx.Member.User.Username)).SetIconURL(ctx.Member.AvatarURL("256"))).
							Build(),
					},
				},
			})
		}

		e, ok := data.ParseEmoji(emoji)
		if !ok {
			return client.InteractionRespond(ctx.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Embeds: []*discordgo.MessageEmbed{
						embed.NewEmbedBuilder().
							SetDescription(i18n.Translate("InvalidEmoji", *ctx.GuildLocale)).
							SetFooter(embed.NewEmbedFooter(i18n.Translate("RequestedBy", *ctx.GuildLocale, ctx.Member.User.Username)).SetIconURL(ctx.Member.AvatarURL("256"))).
							Build(),
					},
				},
			})
		}

		if len(ctx.ApplicationCommandData().Options) > 1 {
			e.Name = ctx.ApplicationCommandData().Options[1].StringValue()
		}

		newEmoji, _ := client.GuildEmojiCreate(ctx.GuildID, e.API()) // Error safe since we already check the permission

		return client.InteractionRespond(ctx.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Embeds: []*discordgo.MessageEmbed{
					embed.NewEmbedBuilder().
						SetDescription(i18n.Translate("EmojiAdded", *ctx.GuildLocale, newEmoji.MessageFormat())).
						SetFooter(embed.NewEmbedFooter(i18n.Translate("RequestedBy", *ctx.GuildLocale, ctx.Member.User.Username)).SetIconURL(ctx.Member.AvatarURL("256"))).
						Build(),
				},
			},
		})
	}
}
