package data

import (
	"encoding/base64"
	"github.com/bwmarrin/discordgo"
	"io"
	"net/http"
	"regexp"
)

var EmojiRegex = regexp.MustCompile(`<?(?:(a):)?(\w{2,32}):(\d{17,19})?>?`)

type Emoji discordgo.Emoji

func NewEmoji(name, id string, animated bool) Emoji {
	return Emoji{ID: id, Name: name, Animated: animated}
}

func (e Emoji) GetLink() string {
	if e.Animated {
		return "https://cdn.discordapp.com/emojis/" + e.ID + ".gif"
	}

	return "https://cdn.discordapp.com/emojis/" + e.ID + ".png"
}

func (e Emoji) GetMention() string {
	if e.Animated {
		return "<a:" + e.Name + ":" + e.ID + ">"
	}

	return "<:" + e.Name + ":" + e.ID + ">"
}

func (e Emoji) API() *discordgo.EmojiParams {
	return &discordgo.EmojiParams{
		Name:  e.Name,
		Image: parseImage(e.GetLink(), e.Animated),
	}
}

// parseImage returns a base64 encoded string of the image data
func parseImage(url string, animated bool) string {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return ""
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return ""
	}

	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return ""
	}

	if animated {
		return "data:image/gif;base64," + base64.StdEncoding.EncodeToString(data)
	}

	return "data:image/png;base64," + base64.StdEncoding.EncodeToString(data)
}

func ParseEmoji(s string) (*Emoji, bool) {
	if !EmojiRegex.MatchString(s) {
		return nil, false
	}

	capture := EmojiRegex.FindStringSubmatch(s)
	emoji := new(Emoji)

	emoji.Name = capture[2]
	emoji.ID = capture[3]
	emoji.Animated = capture[1] == "a"

	return emoji, true
}
