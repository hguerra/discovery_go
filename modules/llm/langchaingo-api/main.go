package main

import (
	"context"
	"fmt"
	"log"

	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/llms/ollama"
)

// Ref:
// https://github.com/tmc/langchaingo/blob/main/examples/json-mode-example/json_mode_example.go
func main() {
	ctx := context.Background()
	llm, err := ollama.New(ollama.WithModel("gemma2:2b"))
	if err != nil {
		log.Fatal(err)
	}
	completion, err := llms.GenerateFromSinglePrompt(ctx,
		llm,
		"Who was first man to walk on the moon? Respond in json format, include `first_man` in response keys.",
		llms.WithTemperature(0.0),
		llms.WithJSONMode(),
	)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(completion)
}
