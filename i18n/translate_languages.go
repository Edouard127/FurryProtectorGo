//go:build generate
// +build generate

package main

import (
	"encoding/json"
	"fmt"
	"github.com/Edouard127/FurryProtectorGo/i18n/locales/en"
	"github.com/joho/godotenv"
	"net/http"
	"net/url"
	"os"
	"text/template"
)

var apiKey string

const (
	api         = "https://api-free.deepl.com/v2/translate"
	defaultLang = "en"
	tmpl        = `package %s

var Map = map[string]string{
{{- range $key, $value := . }}
	"{{ $key }}": "{{ $value }}",
{{- end }}
}
`
)

func doParsing() []map[string]string {
	dir, err := os.ReadDir("locales")
	if err != nil {
		panic(err)
	}

	m := make([]map[string]string, len(dir))

	for i, file := range dir {
		if file.Name() == defaultLang {
			m[i] = en.Map
			continue
		}

		m[i] = doRequest(file.Name(), en.Map)
	}

	return m
}

func doRequest(lang string, data map[string]string) map[string]string {
	m := make(map[string]string, len(data))

	for k, v := range data {
		req, err := http.NewRequest("POST", fmt.Sprintf("%s?text=%s&target_lang=%s", api, url.QueryEscape(v), lang), nil)
		if err != nil {
			panic(err)
		}

		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		req.Header.Set("Authorization", "DeepL-Auth-Key "+apiKey)

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			panic(err)
		}

		defer resp.Body.Close()

		var result struct {
			Translations []struct {
				Text string `json:"text"`
			} `json:"translations"`
		}

		if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
			panic(err)
		}

		fmt.Println(result, lang)

		m[k] = result.Translations[0].Text
	}

	return m
}

//go:generate go run $GOFILE
func main() {
	godotenv.Load("../.env")
	apiKey = os.Getenv("DEEPL_API_KEY")

	var langA []string
	dir, _ := os.ReadDir("locales")
	for _, file := range dir {
		langA = append(langA, file.Name())
	}

	for i, lang := range doParsing() {
		tmpl, err := template.New("langs").Parse(fmt.Sprintf(tmpl, langA[i]))
		if err != nil {
			panic(err)
		}

		f, err := os.Create(fmt.Sprintf("locales/%s/%s.go", langA[i], langA[i]))
		if err != nil {
			panic(err)
		}
		defer f.Close()

		if err := tmpl.Execute(f, lang); err != nil {
			panic(err)
		}
	}

	fmt.Println("Done")
}
