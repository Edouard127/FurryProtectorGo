package i18n

import (
	"fmt"
	"github.com/Edouard127/FurryProtectorGo/i18n/locales/bg"
	"github.com/Edouard127/FurryProtectorGo/i18n/locales/cs"
	"github.com/Edouard127/FurryProtectorGo/i18n/locales/da"
	"github.com/Edouard127/FurryProtectorGo/i18n/locales/de"
	"github.com/Edouard127/FurryProtectorGo/i18n/locales/el"
	"github.com/Edouard127/FurryProtectorGo/i18n/locales/en"
	"github.com/Edouard127/FurryProtectorGo/i18n/locales/es"
	"github.com/Edouard127/FurryProtectorGo/i18n/locales/fi"
	"github.com/Edouard127/FurryProtectorGo/i18n/locales/fr"
	"github.com/Edouard127/FurryProtectorGo/i18n/locales/hu"
	"github.com/Edouard127/FurryProtectorGo/i18n/locales/it"
	"github.com/Edouard127/FurryProtectorGo/i18n/locales/ja"
	"github.com/Edouard127/FurryProtectorGo/i18n/locales/ko"
	"github.com/Edouard127/FurryProtectorGo/i18n/locales/lt"
	"github.com/Edouard127/FurryProtectorGo/i18n/locales/pl"
	"github.com/Edouard127/FurryProtectorGo/i18n/locales/pt"
	"github.com/Edouard127/FurryProtectorGo/i18n/locales/ro"
	"github.com/Edouard127/FurryProtectorGo/i18n/locales/ru"
	"github.com/Edouard127/FurryProtectorGo/i18n/locales/uk"
	"github.com/Edouard127/FurryProtectorGo/i18n/locales/zh"
	. "github.com/bwmarrin/discordgo"
	"strings"
)

func init() {
	AddLocalizer(EnglishUS, &en.Map)
	AddLocalizer(EnglishGB, &en.Map)
	AddLocalizer(Bulgarian, &bg.Map)
	AddLocalizer(ChineseCN, &zh.Map)
	AddLocalizer(ChineseTW, &zh.Map)
	AddLocalizer(Croatian, &en.Map)
	AddLocalizer(Czech, &cs.Map)
	AddLocalizer(Danish, &da.Map)
	AddLocalizer(Dutch, &de.Map)
	AddLocalizer(Finnish, &fi.Map)
	AddLocalizer(French, &fr.Map)
	AddLocalizer(German, &de.Map)
	AddLocalizer(Greek, &el.Map)
	AddLocalizer(Hindi, &en.Map)
	AddLocalizer(Hungarian, &hu.Map)
	AddLocalizer(Italian, &it.Map)
	AddLocalizer(Japanese, &ja.Map)
	AddLocalizer(Korean, &ko.Map)
	AddLocalizer(Lithuanian, &lt.Map)
	AddLocalizer(Norwegian, &en.Map)
	AddLocalizer(Polish, &pl.Map)
	AddLocalizer(PortugueseBR, &pt.Map)
	AddLocalizer(Romanian, &ro.Map)
	AddLocalizer(Russian, &ru.Map)
	AddLocalizer(SpanishES, &es.Map)
	AddLocalizer(Swedish, &en.Map)
	AddLocalizer(Thai, &en.Map)
	AddLocalizer(Turkish, &en.Map)
	AddLocalizer(Ukrainian, &uk.Map)
	AddLocalizer(Vietnamese, &en.Map)
}

var localizers = make(map[string]*map[string]string)

func AddLocalizer(lang Locale, locale *map[string]string) {
	localizers[lang.String()] = locale
}

func TranslateRaw(text string, to Locale, data ...any) string {
	return filli18nData(replaceInvalid(localizers[to.String()], findKey(en.Map, text)), data)
}

func Translate(id string, to Locale, data ...any) string {
	return filli18nData(replaceInvalid(localizers[to.String()], id), data)
}

func replaceInvalid(m *map[string]string, id string) string {
	if (*m)[id] == "" {
		return en.Map[id]
	}
	return (*m)[id]
}

func findKey(m map[string]string, value string) string {
	for k, v := range m {
		if v == value {
			return k
		}
	}
	return ""
}

func filli18nData(text string, data []any) string {
	if len(data) == 0 {
		return text
	}

	for i := 0; i < len(data); i++ {
		text = strings.Replace(text, "%", fmt.Sprintf("%v", data[i]), -1)
	}

	return text
}
