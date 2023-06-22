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
	return [...]string{"CreateInstantInvite", "KickMembers", "BanMembers", "Administrator", "ManageChannels", "ManageGuild", "AddReactions", "ViewAuditLog", "PrioritySpeaker", "Stream", "ViewChannel", "SendMessages", "SendTTSMessages", "ManageMessages", "EmbedLinks", "AttachFiles", "ReadMessageHistory", "MentionEveryone", "UseExternalEmojis", "ViewGuildInsights", "Connect", "Speak", "MuteMembers", "DeafenMembers", "MoveMembers", "UseVAD", "ChangeNickname", "ManageNicknames", "ManageRoles", "ManageWebhooks", "ManageGuildExpressions", "UseApplicationCommands", "RequestToSpeak", "ManageEvents", "ManageThreads", "CreatePublicThreads", "CreatePrivateThreads", "UseExternalStickers", "SendMessagesInThreads", "StartEmbeddedActivities", "ModerateMembers", "ViewCreatorMonetization", "UseSoundboard", "ExternalSounds", "SendVoiceMessages"}[u-1]
}

func (u *UserPermission) AddPermissions(permission ...UserPermission) {
	for _, p := range permission {
		*u |= p
	}
}

func (u *UserPermission) HasPermissions(permission UserPermission) bool {
	return *u&permission == permission
}
