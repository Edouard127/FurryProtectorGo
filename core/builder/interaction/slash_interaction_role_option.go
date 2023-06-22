package interaction

type SlashInteractionRoleOption struct {
	*SlashInteractionOption
}

func NewSlashInteractionRoleOption(name, description string) *SlashInteractionRoleOption {
	return &SlashInteractionRoleOption{NewSlashInteractionOption(Role, name, description)}
}
