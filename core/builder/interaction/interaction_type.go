package interaction

type InteractionType uint

const (
	ChatInputInteraction InteractionType = iota + 1
	UserInteraction
	MessageInteraction
)

func (i InteractionType) String() string {
	return [...]string{"ChatInputInteraction", "UserInteraction", "MessageInteraction"}[i-1]
}
