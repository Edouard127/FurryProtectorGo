package interaction

import (
	"encoding/json"
	"fmt"
)

type SlashInteraction interface {
	MarshalJSON() ([]byte, error)
	UnmarshalJSON([]byte) error
}

type SlashInteractionOption struct {
	SlashInteraction
	Type        SlashInteractionOptionType `json:"type"`
	Name        string                     `json:"name"`
	Description string                     `json:"description"`
	Required    bool                       `json:"required,omitempty"`
}

func NewSlashInteractionOption(t SlashInteractionOptionType, name, description string) *SlashInteractionOption {
	return &SlashInteractionOption{
		Type:        t,
		Name:        name,
		Description: description,
	}
}

func (s *SlashInteractionOption) SetRequired(b bool) *SlashInteractionOption {
	s.Required = b
	return s
}

func (s *SlashInteractionOption) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`{"type":%d,"name":"%s","description":"%s","required":%t}`, s.Type, s.Name, s.Description, s.Required)), nil
}

func (s *SlashInteractionOption) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, s)
}
