package collector

import (
	"github.com/bwmarrin/discordgo"
	"go.uber.org/zap"
	"time"
)

type ComponentCollector struct {
	*zap.Logger
	timeout   int
	maximum   int
	stopTime  int64
	collected []*discordgo.InteractionCreate
	Callback  chan *discordgo.InteractionCreate   // sent at every collection
	Stop      chan []*discordgo.InteractionCreate // sent at the end of the collection
}

func NewComponentCollector(logger *zap.Logger, timeout, maximum int) *ComponentCollector {
	if maximum <= 0 {
		maximum = 1 // we cannot collect forever :3
	}
	return &ComponentCollector{Logger: logger, timeout: timeout, maximum: maximum, stopTime: time.Now().UnixMilli() + int64(timeout), collected: make([]*discordgo.InteractionCreate, 0), Callback: make(chan *discordgo.InteractionCreate, maximum), Stop: make(chan []*discordgo.InteractionCreate)}
}

func (c *ComponentCollector) GetLogger() *zap.Logger {
	return c.Logger
}

func (c *ComponentCollector) Run(session *discordgo.Session, ctx *discordgo.InteractionCreate) error {
	if c.stopTime < time.Now().UnixMilli() {
		return nil
	}

	c.stopTime += int64(c.timeout)

	c.collected = append(c.collected, ctx)
	c.Callback <- ctx

	if len(c.collected) >= c.maximum {
		c.Stop <- c.collected
	}

	return nil
}
