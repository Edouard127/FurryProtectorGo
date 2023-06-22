package interaction

type SlashInteractionUserOption struct {
	*SlashInteractionOption
}

func NewSlashInteractionUserOption(name, description string) *SlashInteractionUserOption {
	return &SlashInteractionUserOption{NewSlashInteractionOption(User, name, description)}
}
