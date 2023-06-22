package interaction

import (
	"encoding/json"
	"fmt"
)

type SlashInteractionIntegerOption struct {
	*SlashInteractionOption
	MinValue int `json:"min_value"`
	MaxValue int `json:"max_value"`
}

func NewSlashInteractionIntegerOption(name, description string) *SlashInteractionIntegerOption {
	return &SlashInteractionIntegerOption{NewSlashInteractionOption(Integer, name, description), 0, 0}
}

func (s *SlashInteractionIntegerOption) SetMinValue(i int) *SlashInteractionIntegerOption {
	s.MinValue = i
	return s
}

func (s *SlashInteractionIntegerOption) SetMaxValue(i int) *SlashInteractionIntegerOption {
	s.MaxValue = i
	return s
}

func (s *SlashInteractionIntegerOption) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`{"type":%d,"name":"%s","description":"%s","required":%t,"min_value":%d,"max_value":%d}`, s.Type, s.Name, s.Description, s.Required, s.MinValue, s.MaxValue)), nil
}

func (s *SlashInteractionIntegerOption) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, s)
}
