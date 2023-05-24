package main

import (
	"fmt"

	"github.com/bregydoc/gtranslate"
)

func main() {
	text := "Hello World"
	translated, err := gtranslate.TranslateWithParams(
		text,
		gtranslate.TranslationParams{
			From: "en",
			To:   "pt",
		},
	)
	if err != nil {
		panic(err)
	}

	fmt.Printf("en: %s | pt: %s \n", text, translated)
}
