package registers

import (
	"github.com/Edouard127/FurryProtectorGo/core/builder/interaction"
	"github.com/bwmarrin/discordgo"
)

var (
	InteractionComponents = NewRegister[discordgo.InteractionCreate]()
	InteractionCommands   = NewRegister[discordgo.InteractionCreate]()
	InteractionModals     = NewRegister[discordgo.InteractionCreate]()
)

type RunnerRegister[T any] struct {
	Runners map[string]interaction.Runner[T]
}

func NewRegister[T any]() *RunnerRegister[T] {
	return &RunnerRegister[T]{
		Runners: make(map[string]interaction.Runner[T]),
	}
}

func (i *RunnerRegister[T]) Register(index string, runner interaction.Runner[T]) {
	i.Runners[index] = runner
}

func (i *RunnerRegister[T]) Get(index string) interaction.Runner[T] {
	return i.Runners[index]
}
