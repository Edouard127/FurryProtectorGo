package interaction

type SlashInteractionMentionableOption struct {
	*SlashInteractionOption
}

func NewSlashInteractionMentionableOption(name, description string) *SlashInteractionMentionableOption {
	return &SlashInteractionMentionableOption{NewSlashInteractionOption(Mentionable, name, description)}
}
