package main

import (
	"context"
	"fmt"
	"log"

	"github.com/ollama/ollama/api"
)

// Ref:
// https://github.com/ollama/ollama/blob/main/examples/go-generate/main.go
// https://github.com/ollama/ollama/blob/main/examples/go-chat/main.go
// https://github.com/ollama/ollama/blob/main/examples/go-generate-streaming/main.go
// https://github.com/ollama/ollama/blob/main/examples/go-multimodal/main.go
func main() {
	client, err := api.ClientFromEnvironment()
	if err != nil {
		log.Fatal(err)
	}

	req := &api.GenerateRequest{
		Model:  "gemma2:2b",
		Prompt: "Qual a historia do jogo Genshin Impact?",

		// set streaming to false
		Stream: new(bool),
	}

	ctx := context.Background()
	respFunc := func(resp api.GenerateResponse) error {
		// Only print the response here; GenerateResponse has a number of other
		// interesting fields you want to examine.
		fmt.Println(resp.Response)
		return nil
	}

	err = client.Generate(ctx, req, respFunc)
	if err != nil {
		log.Fatal(err)
	}
}
