package interaction

type SlashInteractionChoice struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

func NewSlashInteractionChoice(name, value string) *SlashInteractionChoice {
	return &SlashInteractionChoice{name, value}
}
