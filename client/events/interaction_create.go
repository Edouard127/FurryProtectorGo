package events

import (
	"github.com/Edouard127/FurryProtectorGo/client/database"
	"github.com/Edouard127/FurryProtectorGo/client/exporter"
	"github.com/Edouard127/FurryProtectorGo/core/builder/components/embed"
	"github.com/Edouard127/FurryProtectorGo/core/builder/interaction"
	"github.com/Edouard127/FurryProtectorGo/i18n"
	"github.com/Edouard127/FurryProtectorGo/registers"
	"github.com/bwmarrin/discordgo"
	"github.com/prometheus/client_golang/prometheus"
	"go.uber.org/zap"
	"time"
)

var (
	runner interaction.Runner[discordgo.InteractionCreate]
	ok     bool
	err    error
)

func NewInteractionCreateEvent(logger *zap.Logger, db *database.Database) func(*discordgo.Session, *discordgo.InteractionCreate) {
	return func(session *discordgo.Session, ctx *discordgo.InteractionCreate) {
		logger.Info("Interaction received", zap.String("name", ctx.ApplicationCommandData().Name))
		var before = time.Now()

		switch ctx.Type {
		case discordgo.InteractionMessageComponent:
			runner, ok = registers.InteractionComponents.Get(ctx.MessageComponentData().CustomID)
			if !ok {
				session.InteractionRespond(ctx.Interaction, &discordgo.InteractionResponse{
					Type: discordgo.InteractionResponseDeferredMessageUpdate,
					Data: &discordgo.InteractionResponseData{
						Embeds: []*discordgo.MessageEmbed{
							embed.NewEmbedBuilder().
								SetTitle(i18n.Translate("UnknownComponent", *ctx.GuildLocale)).
								SetDescription(i18n.Translate("UnknownComponentDescription", *ctx.GuildLocale)).
								SetColor(0xff0000).
								SetFooter(embed.NewEmbedFooter(i18n.Translate("RequestedBy", *ctx.GuildLocale, ctx.Member.User.Username)).SetIconURL(ctx.Member.AvatarURL("256"))).
								Build(),
						},
					},
				})
				return
			}
		case discordgo.InteractionApplicationCommand:
			runner, ok = registers.InteractionCommands.Get(ctx.ApplicationCommandData().Name)
			if !ok {
				session.InteractionRespond(ctx.Interaction, &discordgo.InteractionResponse{
					Type: discordgo.InteractionResponseDeferredMessageUpdate,
					Data: &discordgo.InteractionResponseData{
						Embeds: []*discordgo.MessageEmbed{
							embed.NewEmbedBuilder().
								SetTitle(i18n.Translate("UnknownCommand", *ctx.GuildLocale)).
								SetDescription(i18n.Translate("UnknownCommandDescription", *ctx.GuildLocale)).
								SetColor(0xff0000).
								SetFooter(embed.NewEmbedFooter(i18n.Translate("RequestedBy", *ctx.GuildLocale, ctx.Member.User.Username)).SetIconURL(ctx.Member.AvatarURL("256"))).
								Build(),
						},
					},
				})
				return
			}
		case discordgo.InteractionModalSubmit:
			runner, ok = registers.InteractionModals.Get(ctx.MessageComponentData().CustomID)
			if !ok {
				session.InteractionRespond(ctx.Interaction, &discordgo.InteractionResponse{
					Type: discordgo.InteractionResponseDeferredMessageUpdate,
					Data: &discordgo.InteractionResponseData{
						Embeds: []*discordgo.MessageEmbed{
							embed.NewEmbedBuilder().
								SetTitle(i18n.Translate("UnknownModal", *ctx.GuildLocale)).
								SetDescription(i18n.Translate("UnknownModalDescription", *ctx.GuildLocale)).
								SetColor(0xff0000).
								SetFooter(embed.NewEmbedFooter(i18n.Translate("RequestedBy", *ctx.GuildLocale, ctx.Member.User.Username)).SetIconURL(ctx.Member.AvatarURL("256"))).
								Build(),
						},
					},
				})
				return
			}
		}

		err = runner(session, ctx)

		if err != nil {
			logger.Error("Error while running interaction", zap.Error(err))
		}

		exporter.InterationHist.With(prometheus.Labels{"guild": ctx.GuildID, "user": ctx.Member.User.ID, "interaction": ctx.ApplicationCommandData().Name}).Observe(time.Since(before).Seconds())
	}
}
