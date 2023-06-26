package data

type UserPermission uint64

const (
	CreateInstantInvite UserPermission = 1 << iota
	KickMembers
	BanMembers
	Administrator
	ManageChannels
	ManageGuild
	AddReactions
	ViewAuditLog
	PrioritySpeaker
	Stream
	ViewChannel
	SendMessages
	SendTTSMessages
	ManageMessages
	EmbedLinks
	AttachFiles
	ReadMessageHistory
	MentionEveryone
	UseExternalEmojis
	ViewGuildInsights
	Connect
	Speak
	MuteMembers
	DeafenMembers
	MoveMembers
	UseVAD
	ChangeNickname
	ManageNicknames
	ManageRoles
	ManageWebhooks
	ManageGuildExpressions
	UseApplicationCommands
	RequestToSpeak
	ManageEvents
	ManageThreads
	CreatePublicThreads
	CreatePrivateThreads
	UseExternalStickers
	SendMessagesInThreads
	StartEmbeddedActivities
	ModerateMembers
	ViewCreatorMonetization
	UseSoundboard
	ExternalSounds
	SendVoiceMessages
)

func (u UserPermission) String() string {
	return [...]string{"CreateInstantInvite", "KickMembers", "BanMembers", "Administrator", "ManageChannels", "ManageGuild", "AddReactions", "ViewAuditLog", "PrioritySpeaker", "Stream", "ViewChannel", "SendMessages", "SendTTSMessages", "ManageMessages", "EmbedLinks", "AttachFiles", "ReadMessageHistory", "MentionEveryone", "UseExternalEmojis", "ViewGuildInsights", "Connect", "Speak", "MuteMembers", "DeafenMembers", "MoveMembers", "UseVAD", "ChangeNickname", "ManageNicknames", "ManageRoles", "ManageWebhooks", "ManageGuildExpressions", "UseApplicationCommands", "RequestToSpeak", "ManageEvents", "ManageThreads", "CreatePublicThreads", "CreatePrivateThreads", "UseExternalStickers", "SendMessagesInThreads", "StartEmbeddedActivities", "ModerateMembers", "ViewCreatorMonetization", "UseSoundboard", "ExternalSounds", "SendVoiceMessages"}[getIndex(u)]
}

func getIndex(n UserPermission) int {
	low := 0
	high := 44

	for low < high {
		mid := (low + high) / 2

		switch {
		case n < 1<<uint(mid):
			high = mid
		case n > 1<<uint(mid):
			low = mid + 1
		default:
			return mid
		}
	}

	return 0
}
