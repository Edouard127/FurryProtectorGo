package i8n

import (
	"fmt"
	"testing"
)

func TestLocales(t *testing.T) {
	fmt.Println(TranslateRaw("Click on the button to add questions. % remaining.", "bg", 223))
	fmt.Println(TranslateRaw("Click on the button to add questions. % remaining.", "ru", 223))
	fmt.Println(TranslateRaw("Click on the button to add questions. % remaining.", "uk", 223))
	fmt.Println(TranslateRaw("Click on the button to add questions. % remaining.", "zh", 223))
	fmt.Println(TranslateRaw("Click on the button to add questions. % remaining.", "ja", 223))
	fmt.Println(TranslateRaw("Click on the button to add questions. % remaining.", "fr", 223))
	fmt.Println(TranslateRaw("Click on the button to add questions. % remaining.", "de", 223))
}
