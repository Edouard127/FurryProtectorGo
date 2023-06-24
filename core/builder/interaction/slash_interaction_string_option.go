package interaction

import (
	"encoding/json"
	"fmt"
)

type SlashInteractionStringOption struct {
	*SlashInteractionOption
	MinLength int `json:"min_length,omitempty"`
	MaxLength int `json:"max_length,omitempty"`
}

func NewSlashInteractionStringOption(name, description string) *SlashInteractionStringOption {
	return &SlashInteractionStringOption{NewSlashInteractionOption(String, name, description), 1, 255}
}

func (s *SlashInteractionStringOption) SetMinLength(i int) *SlashInteractionStringOption {
	s.MinLength = i
	return s
}

func (s *SlashInteractionStringOption) SetMaxLength(i int) *SlashInteractionStringOption {
	s.MaxLength = i
	return s
}

func (s *SlashInteractionStringOption) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`{"type":%d,"name":"%s","description":"%s","required":%t,"min_length":%d,"max_length":%d}`, s.Type, s.Name, s.Description, s.Required, s.MinLength, s.MaxLength)), nil
}

func (s *SlashInteractionStringOption) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, s)
}
