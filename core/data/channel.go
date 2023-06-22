package data

type ChannelType uint8

const (
	GuildText ChannelType = iota
	DirectMessage
	GuildVoice
	GroupDirectMessage
	GuildCategory
	GuildNews
	AnnouncementThread
	PublicThread
	PrivateThread
	GuildStageVoice
	GuildDirectory
	GuildForum
)

func (c ChannelType) String() string {
	return [...]string{"GuildText", "DirectMessage", "GuildVoice", "GroupDirectMessage", "GuildCategory", "GuildNews", "AnnouncementThread", "PublicThread", "PrivateThread", "GuildStageVoice", "GuildDirectory", "GuildForum"}[c]
}
