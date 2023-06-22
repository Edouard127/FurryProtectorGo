package i18n

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"testing"
)

func TestLocales(t *testing.T) {
	fmt.Println(TranslateRaw("Click on the button to add questions. % remaining.", discordgo.EnglishGB, 223))
	fmt.Println(TranslateRaw("Click on the button to add questions. % remaining.", discordgo.Russian, 223))
	fmt.Println(TranslateRaw("Click on the button to add questions. % remaining.", discordgo.Ukrainian, 223))
	fmt.Println(TranslateRaw("Click on the button to add questions. % remaining.", discordgo.ChineseCN, 223))
	fmt.Println(TranslateRaw("Click on the button to add questions. % remaining.", discordgo.Japanese, 223))
	fmt.Println(TranslateRaw("Click on the button to add questions. % remaining.", discordgo.French, 223))
	fmt.Println(TranslateRaw("Click on the button to add questions. % remaining.", discordgo.Dutch, 223))
}
