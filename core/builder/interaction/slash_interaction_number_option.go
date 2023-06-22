package interaction

import (
	"encoding/json"
	"fmt"
)

type SlashInteractionNumberOption struct {
	*SlashInteractionOption
	MinValue int `json:"min_value"`
	MaxValue int `json:"max_value"`
}

func NewSlashInteractionNumberOption(name, description string) *SlashInteractionNumberOption {
	return &SlashInteractionNumberOption{NewSlashInteractionOption(Number, name, description), 0, 0}
}

func (s *SlashInteractionNumberOption) SetMinValue(i int) *SlashInteractionNumberOption {
	s.MinValue = i
	return s
}

func (s *SlashInteractionNumberOption) SetMaxValue(i int) *SlashInteractionNumberOption {
	s.MaxValue = i
	return s
}

func (s *SlashInteractionNumberOption) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`{"type":%d,"name":"%s","description":"%s","required":%t,"min_value":%d,"max_value":%d}`, s.Type, s.Name, s.Description, s.Required, s.MinValue, s.MaxValue)), nil
}

func (s *SlashInteractionNumberOption) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, s)
}
