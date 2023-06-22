package interaction

type SlashInteractionBooleanOption struct {
	*SlashInteractionOption
}

func NewSlashInteractionBooleanOption(name, description string) *SlashInteractionBooleanOption {
	return &SlashInteractionBooleanOption{NewSlashInteractionOption(Boolean, name, description)}
}
