package embed

import (
	"github.com/bwmarrin/discordgo"
	"time"
)

type Embed discordgo.MessageEmbed

func NewEmbedBuilder() *Embed {
	return &Embed{}
}

func (e *Embed) SetTitle(title string) *Embed {
	e.Title = title
	return e
}

func (e *Embed) SetType(embedType discordgo.EmbedType) *Embed {
	e.Type = embedType
	return e
}

func (e *Embed) SetDescription(description string) *Embed {
	e.Description = description
	return e
}

func (e *Embed) SetURL(url string) *Embed {
	e.URL = url
	return e
}

func (e *Embed) SetTimestamp(timestamp *time.Time) *Embed {
	e.Timestamp = timestamp.String()
	return e
}

func (e *Embed) SetColor(color int) *Embed {
	e.Color = color
	return e
}

func (e *Embed) SetFooter(footer *EmbedFooter) *Embed {
	e.Footer = (*discordgo.MessageEmbedFooter)(footer)
	return e
}

func (e *Embed) SetImage(image *EmbedImage) *Embed {
	e.Image = (*discordgo.MessageEmbedImage)(image)
	return e
}

func (e *Embed) SetThumbnail(thumbnail *EmbedThumbnail) *Embed {
	e.Thumbnail = (*discordgo.MessageEmbedThumbnail)(thumbnail)
	return e
}

func (e *Embed) SetVideo(video *EmbedVideo) *Embed {
	e.Video = (*discordgo.MessageEmbedVideo)(video)
	return e
}

func (e *Embed) SetProvider(provider *EmbedProvider) *Embed {
	e.Provider = (*discordgo.MessageEmbedProvider)(provider)
	return e
}

func (e *Embed) SetAuthor(author *EmbedAuthor) *Embed {
	e.Author = (*discordgo.MessageEmbedAuthor)(author)
	return e
}

func (e *Embed) AddField(field ...*EmbedField) *Embed {
	for _, f := range field {
		e.Fields = append(e.Fields, (*discordgo.MessageEmbedField)(f))
	}
	return e
}

func (e *Embed) Build() *discordgo.MessageEmbed {
	return (*discordgo.MessageEmbed)(e)
}

type EmbedThumbnail discordgo.MessageEmbedThumbnail

func NewEmbedThumbnail(url string) *EmbedThumbnail {
	return &EmbedThumbnail{URL: url}
}

func (e *EmbedThumbnail) SetProxyURL(proxyURL string) *EmbedThumbnail {
	e.ProxyURL = proxyURL
	return e
}

func (e *EmbedThumbnail) SetHeight(height int) *EmbedThumbnail {
	e.Height = height
	return e
}

func (e *EmbedThumbnail) SetWidth(width int) *EmbedThumbnail {
	e.Width = width
	return e
}

type EmbedVideo discordgo.MessageEmbedVideo

func NewEmbedVideo(url string) *EmbedVideo {
	return &EmbedVideo{URL: url}
}

func (e *EmbedVideo) SetHeight(height int) *EmbedVideo {
	e.Height = height
	return e
}

func (e *EmbedVideo) SetWidth(width int) *EmbedVideo {
	e.Width = width
	return e
}

type EmbedImage discordgo.MessageEmbedImage

func NewEmbedImage(url string) *EmbedImage {
	return &EmbedImage{URL: url}
}

func (e *EmbedImage) SetProxyURL(proxyURL string) *EmbedImage {
	e.ProxyURL = proxyURL
	return e
}

func (e *EmbedImage) SetHeight(height int) *EmbedImage {
	e.Height = height
	return e
}

func (e *EmbedImage) SetWidth(width int) *EmbedImage {
	e.Width = width
	return e
}

type EmbedProvider discordgo.MessageEmbedProvider

func NewEmbedProvider(name string, url string) *EmbedProvider {
	return &EmbedProvider{Name: name, URL: url}
}

type EmbedAuthor discordgo.MessageEmbedAuthor

func NewEmbedAuthor(name string) *EmbedAuthor {
	return &EmbedAuthor{Name: name}
}

func (e *EmbedAuthor) SetURL(url string) *EmbedAuthor {
	e.URL = url
	return e
}

func (e *EmbedAuthor) SetIconURL(iconURL string) *EmbedAuthor {
	e.IconURL = iconURL
	return e
}

func (e *EmbedAuthor) SetProxyIconURL(proxyIconURL string) *EmbedAuthor {
	e.ProxyIconURL = proxyIconURL
	return e
}

type EmbedFooter discordgo.MessageEmbedFooter

func NewEmbedFooter(text string) *EmbedFooter {
	return &EmbedFooter{Text: text}
}

func (e *EmbedFooter) SetIconURL(iconURL string) *EmbedFooter {
	e.IconURL = iconURL
	return e
}

func (e *EmbedFooter) SetProxyIconURL(proxyIconURL string) *EmbedFooter {
	e.ProxyIconURL = proxyIconURL
	return e
}

type EmbedField discordgo.MessageEmbedField

func NewEmbedField(name string, value string) *EmbedField {
	return &EmbedField{Name: name, Value: value}
}

func (e *EmbedField) SetInline(inline bool) *EmbedField {
	e.Inline = inline
	return e
}
