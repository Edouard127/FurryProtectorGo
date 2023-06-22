package template

import (
	"fmt"
	"github.com/Edouard127/FurryProtectorGo/core/builder/components/embed"
	"github.com/Edouard127/FurryProtectorGo/i18n"
	"github.com/bwmarrin/discordgo"
	"runtime"
	"strconv"
)

var BotInfoTemplate = func(session *discordgo.Session, locale discordgo.Locale) *embed.Embed {
	cpus, ram := getHostInfo()
	var users int
	for _, guild := range session.State.Guilds {
		users += guild.MemberCount
	}
	return embed.NewEmbedBuilder().
		SetTitle(i18n.Translate("BotInfo", locale)).
		AddField(
			embed.NewEmbedField(i18n.Translate("Guilds", locale), fmt.Sprintf("```%d```", len(session.State.Guilds))),
			embed.NewEmbedField(i18n.Translate("Users", locale), fmt.Sprintf("```%d```", users)),
			embed.NewEmbedField(i18n.Translate("CPUAmount", locale), fmt.Sprintf("```%s```", cpus)),
			embed.NewEmbedField(i18n.Translate("RAMUsage", locale), fmt.Sprintf("```%s MB```", ram)),
		)
}

func getHostInfo() (cpus, ram string) {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	return strconv.Itoa(runtime.NumCPU()), strconv.Itoa(int(m.Sys / 1024 / 1024))
}
