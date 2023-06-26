package interaction

import (
	"github.com/bwmarrin/discordgo"
)

type SlashInteractionOption discordgo.ApplicationCommandOption

func NewSlashInteractionOption(t discordgo.ApplicationCommandOptionType, name, description string) *SlashInteractionOption {
	return &SlashInteractionOption{
		Type:        t,
		Name:        name,
		Description: description,
	}
}

func (s *SlashInteractionOption) SetChannelTypes(channelTypes ...discordgo.ChannelType) *SlashInteractionOption {
	s.ChannelTypes = channelTypes
	return s
}

func (s *SlashInteractionOption) SetRequired(b bool) *SlashInteractionOption {
	s.Required = b
	return s
}

func (s *SlashInteractionOption) AddOption(option ...*SlashInteractionOption) *SlashInteractionOption {
	for _, o := range option {
		s.Options = append(s.Options, (*discordgo.ApplicationCommandOption)(o))
	}
	return s
}

func (s *SlashInteractionOption) SetAutoComplete(b bool) *SlashInteractionOption {
	s.Autocomplete = b
	return s
}

func (s *SlashInteractionOption) AddChoice(choice ...*discordgo.ApplicationCommandOptionChoice) *SlashInteractionOption {
	s.Choices = append(s.Choices, choice...)
	return s
}

func (s *SlashInteractionOption) SetMinValue(min float64) *SlashInteractionOption {
	s.MinValue = &min
	return s
}

func (s *SlashInteractionOption) SetMaxValue(max float64) *SlashInteractionOption {
	s.MaxValue = max
	return s
}

func (s *SlashInteractionOption) SetMinLength(min int) *SlashInteractionOption {
	s.MinLength = &min
	return s
}

func (s *SlashInteractionOption) SetMaxLength(max int) *SlashInteractionOption {
	s.MaxLength = max
	return s
}

func NewSlashInteractionAttachmentOption(name, description string) *SlashInteractionOption {
	return NewSlashInteractionOption(discordgo.ApplicationCommandOptionAttachment, name, description)
}

func NewSlashInteractionBooleanOption(name, description string) *SlashInteractionOption {
	return NewSlashInteractionOption(discordgo.ApplicationCommandOptionBoolean, name, description)
}

func NewSlashInteractionChannelOption(name, description string) *SlashInteractionOption {
	return NewSlashInteractionOption(discordgo.ApplicationCommandOptionChannel, name, description)
}

func NewSlashInteractionMentionableOption(name, description string) *SlashInteractionOption {
	return NewSlashInteractionOption(discordgo.ApplicationCommandOptionMentionable, name, description)
}

func NewSlashInteractionNumberOption(name, description string) *SlashInteractionOption {
	return NewSlashInteractionOption(discordgo.ApplicationCommandOptionInteger, name, description)
}

func NewSlashInteractionRoleOption(name, description string) *SlashInteractionOption {
	return NewSlashInteractionOption(discordgo.ApplicationCommandOptionRole, name, description)
}

func NewSlashInteractionStringOption(name, description string) *SlashInteractionOption {
	return NewSlashInteractionOption(discordgo.ApplicationCommandOptionString, name, description)
}

func NewSlashInteractionUserOption(name, description string) *SlashInteractionOption {
	return NewSlashInteractionOption(discordgo.ApplicationCommandOptionUser, name, description)
}
