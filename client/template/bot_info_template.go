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
			embed.NewEmbedField(i18n.Translate("RAMUsage", locale), fmt.Sprintf("```%s```", ram)),
		)
}

func getHostInfo() (cpus, ram string) {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	return strconv.Itoa(runtime.NumCPU()), formatRam(m.Alloc)
}

// formatRam formats the ram usage in KB
func formatRam(ram uint64) string {
	const unit = 1000
	if ram < unit {
		return fmt.Sprintf("%d B", ram)
	}
	div, exp := int64(unit), 0
	for n := ram / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.1f %cB",
		float64(ram)/float64(div), "kMGTPE"[exp])
}
