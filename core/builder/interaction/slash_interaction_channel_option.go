package interaction

type SlashInteractionChannelOption struct {
	*SlashInteractionOption
}

func NewSlashInteractionChannelOption(name, description string) *SlashInteractionChannelOption {
	return &SlashInteractionChannelOption{NewSlashInteractionOption(Channel, name, description)}
}
