package interaction

import (
	"github.com/Edouard127/FurryProtectorGo/core/data"
	"github.com/bwmarrin/discordgo"
	"go.uber.org/zap"
	"strconv"
)

type Runner[T any] interface {
	GetLogger() *zap.Logger
	Run(client *discordgo.Session, ctx *T)
}

type SlashInteractionBuilder struct {
	Type              data.CommandType          `json:"type"`
	Name              string                    `json:"name"`
	Description       string                    `json:"description"`
	Required          bool                      `json:"required,omitempty"`
	Choices           []*SlashInteractionChoice `json:"choices,omitempty"`
	Options           []SlashInteraction        `json:"options,omitempty"`
	ChannelTypes      []data.ChannelType        `json:"channel_types,omitempty"`
	DefaultPermission string                    `json:"default_member_permission,omitempty"`
	DMPermission      bool                      `json:"dm_permissions,omitempty"`
	Nsfw              bool                      `json:"nsfw,omitempty"`
	Version           data.Snowflake            `json:"version"`
}

func NewSlashInteractionBuilder(name, description string) *SlashInteractionBuilder {
	return &SlashInteractionBuilder{
		Type:        data.ChatInput,
		Name:        name,
		Description: description,
	}
}

func (b *SlashInteractionBuilder) AddOption(option SlashInteraction) *SlashInteractionBuilder {
	b.Options = append(b.Options, option)
	return b
}

func (b *SlashInteractionBuilder) SetDefaultPermission(defaultPermission data.UserPermission) *SlashInteractionBuilder {
	b.DefaultPermission = strconv.FormatUint(uint64(defaultPermission), 10)
	return b
}

func (b *SlashInteractionBuilder) SetDMPermission(dmPermission bool) *SlashInteractionBuilder {
	b.DMPermission = dmPermission
	return b
}

func (b *SlashInteractionBuilder) SetNsfw(nsfw bool) *SlashInteractionBuilder {
	b.Nsfw = nsfw
	return b
}

func (b *SlashInteractionBuilder) SetVersion(version data.Snowflake) *SlashInteractionBuilder {
	b.Version = version
	return b
}
