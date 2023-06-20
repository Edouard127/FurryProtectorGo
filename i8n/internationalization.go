package i8n

import (
	"fmt"
	"github.com/Edouard127/FurryProtectorGo/i8n/locales/bg"
	"github.com/Edouard127/FurryProtectorGo/i8n/locales/cs"
	"github.com/Edouard127/FurryProtectorGo/i8n/locales/da"
	"github.com/Edouard127/FurryProtectorGo/i8n/locales/de"
	"github.com/Edouard127/FurryProtectorGo/i8n/locales/el"
	"github.com/Edouard127/FurryProtectorGo/i8n/locales/en"
	"github.com/Edouard127/FurryProtectorGo/i8n/locales/fi"
	"github.com/Edouard127/FurryProtectorGo/i8n/locales/fr"
	"github.com/Edouard127/FurryProtectorGo/i8n/locales/hu"
	"github.com/Edouard127/FurryProtectorGo/i8n/locales/id"
	"github.com/Edouard127/FurryProtectorGo/i8n/locales/it"
	"github.com/Edouard127/FurryProtectorGo/i8n/locales/ja"
	"github.com/Edouard127/FurryProtectorGo/i8n/locales/ko"
	"github.com/Edouard127/FurryProtectorGo/i8n/locales/lt"
	"github.com/Edouard127/FurryProtectorGo/i8n/locales/nl"
	"github.com/Edouard127/FurryProtectorGo/i8n/locales/pl"
	"github.com/Edouard127/FurryProtectorGo/i8n/locales/pt"
	"github.com/Edouard127/FurryProtectorGo/i8n/locales/ro"
	"github.com/Edouard127/FurryProtectorGo/i8n/locales/ru"
	"github.com/Edouard127/FurryProtectorGo/i8n/locales/sv"
	"github.com/Edouard127/FurryProtectorGo/i8n/locales/uk"
	"github.com/Edouard127/FurryProtectorGo/i8n/locales/zh"
	"strings"
)

func init() {
	localizers = make(map[string]*map[string]string)
	AddLocalizer("bg", &bg.Map)
	AddLocalizer("cs", &cs.Map)
	AddLocalizer("da", &da.Map)
	AddLocalizer("de", &de.Map)
	AddLocalizer("el", &el.Map)
	AddLocalizer("en", &en.Map)
	AddLocalizer("fi", &fi.Map)
	AddLocalizer("fr", &fr.Map)
	AddLocalizer("hu", &hu.Map)
	AddLocalizer("id", &id.Map)
	AddLocalizer("it", &it.Map)
	AddLocalizer("ja", &ja.Map)
	AddLocalizer("ko", &ko.Map)
	AddLocalizer("lt", &lt.Map)
	AddLocalizer("nl", &nl.Map)
	AddLocalizer("pl", &pl.Map)
	AddLocalizer("pt", &pt.Map)
	AddLocalizer("ro", &ro.Map)
	AddLocalizer("ru", &ru.Map)
	AddLocalizer("sv", &sv.Map)
	AddLocalizer("uk", &uk.Map)
	AddLocalizer("zh", &zh.Map)
}

var localizers map[string]*map[string]string

func AddLocalizer(lang string, locale *map[string]string) {
	localizers[lang] = locale
}

func TranslateRaw(text string, to string, data ...any) string {
	return filli18nData((*localizers[to])[findKey(en.Map, text)], data)
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

	replacements := make([]string, len(data))
	for i, value := range data {
		replacements[i] = fmt.Sprintf("%v", value)
	}

	for _, replacement := range replacements {
		text = strings.Replace(text, "%", replacement, -1)
	}

	return text
}
