package interaction

import (
	"github.com/Edouard127/FurryProtectorGo/core/data"
	"github.com/bwmarrin/discordgo"
)

type Runner[T any] func(client *discordgo.Session, ctx *T) error

type SlashInteractionBuilder discordgo.ApplicationCommand

func NewSlashInteractionBuilder(name, description string) *SlashInteractionBuilder {
	return &SlashInteractionBuilder{
		Type:        discordgo.ChatApplicationCommand,
		Name:        name,
		Description: description,
	}
}

func (b *SlashInteractionBuilder) AddOption(option ...*SlashInteractionOption) *SlashInteractionBuilder {
	for _, o := range option {
		b.Options = append(b.Options, (*discordgo.ApplicationCommandOption)(o))
	}
	return b
}

func (b *SlashInteractionBuilder) SetDefaultPermission(defaultPermission data.UserPermission) *SlashInteractionBuilder {
	f := int64(defaultPermission)
	b.DefaultMemberPermissions = &f
	return b
}

func (b *SlashInteractionBuilder) SetDMPermission(dmPermission bool) *SlashInteractionBuilder {
	b.DMPermission = &dmPermission
	return b
}

func (b *SlashInteractionBuilder) SetNsfw(nsfw bool) *SlashInteractionBuilder {
	b.NSFW = &nsfw
	return b
}

func (b *SlashInteractionBuilder) Build() *discordgo.ApplicationCommand {
	return (*discordgo.ApplicationCommand)(b)
}
