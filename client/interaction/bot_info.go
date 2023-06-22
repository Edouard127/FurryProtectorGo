package interaction

import (
	"github.com/Edouard127/FurryProtectorGo/client/database"
	"github.com/Edouard127/FurryProtectorGo/client/template"
	"github.com/Edouard127/FurryProtectorGo/core/builder/components/embed"
	"github.com/Edouard127/FurryProtectorGo/core/builder/interaction"
	"github.com/bwmarrin/discordgo"
	"go.uber.org/zap"
)

type BotInfo struct {
	*zap.Logger
	*interaction.SlashInteractionBuilder
	interaction.Runner[discordgo.InteractionCreate]
}

func NewBotInfo(logger *zap.Logger, db *database.Database) (string, *BotInfo) {
	return "info", &BotInfo{Logger: logger, SlashInteractionBuilder: interaction.NewSlashInteractionBuilder("info", "Display the current information about the bot")}
}

func (b *BotInfo) GetLogger() *zap.Logger {
	return b.Logger
}

func (b *BotInfo) Run(client *discordgo.Session, ctx *discordgo.InteractionCreate) {
	client.InteractionRespond(ctx.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Embeds: []*discordgo.MessageEmbed{
				template.BotInfoTemplate(client, *ctx.GuildLocale).
					SetFooter(embed.NewEmbedFooter("Requested by " + ctx.Member.User.Username).SetIconURL(ctx.Member.AvatarURL("32"))).Build(),
			},
		},
	})
}
