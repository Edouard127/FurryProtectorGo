package data

type CommandType int

const (
	ChatInput CommandType = iota + 1
	User
	Message
)

func (c CommandType) String() string {
	return [...]string{"ChatInput", "User", "Message"}[c-1]
}
