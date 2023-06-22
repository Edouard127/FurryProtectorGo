package interaction

type SlashInteractionOptionType int

const (
	SubCommand SlashInteractionOptionType = iota + 1
	SubCommandGroup
	String
	Integer
	Boolean
	User
	Channel
	Role
	Mentionable
	Number
	Attachment
)

func (s SlashInteractionOptionType) String() string {
	return [...]string{"SubCommand", "SubCommandGroup", "String", "Integer", "Boolean", "User", "Channel", "Role", "Mentionable", "Number", "Attachment"}[s-1]
}
