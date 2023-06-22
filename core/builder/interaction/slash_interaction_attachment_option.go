package interaction

type SlashInteractionAttachmentOption struct {
	*SlashInteractionOption
}

func NewSlashInteractionAttachmentOption(name, description string) *SlashInteractionAttachmentOption {
	return &SlashInteractionAttachmentOption{NewSlashInteractionOption(Attachment, name, description)}
}
