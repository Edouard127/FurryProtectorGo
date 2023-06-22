package data

type EmbedType uint

const (
	RichEmbed EmbedType = iota + 1
	ImageEmbed
	VideoEmbed
	GIFVEmbed
	ArticleEmbed
	LinkEmbed
	AutoModerationEmbed
)

func (i EmbedType) String() string {
	return [...]string{"rich", "image", "video", "gifv", "article", "link", "auto_moderation_message"}[i-1]
}

const (
	EmbedDefault    = 0x000000
	EmbedWhite      = 0xffffff
	EmbedAqua       = 0x1abc9c
	EmbedGreen      = 0x2ecc71
	EmbedBlue       = 0x3498db
	EmbedYellow     = 0xffff00
	EmbedPurple     = 0x9b59b6
	EmbedGold       = 0xf1c40f
	EmbedOrange     = 0xe67e22
	EmbedRed        = 0xe74c3c
	EmbedGrey       = 0x95a5a6
	EmbedDarkAqua   = 0x11806a
	EmbedDarkGreen  = 0x1f8b4c
	EmbedDarkBlue   = 0x206694
	EmbedDarkPurple = 0x71368a
	EmbedDarkGold   = 0xc27c0e
	EmbedDarkOrange = 0xa84300
	EmbedDarkRed    = 0x992d22
	EmbedDarkGrey   = 0x979c9f
	EmbedBlurple    = 0x7289da
	EmbedDark       = 0x2c2f33
)
