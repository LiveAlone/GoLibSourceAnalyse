package main

import (
	"context"
	"fmt"
	"github.com/sashabaranov/go-openai"
)

const Token = "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"

func main() {
	c := openai.NewClient(Token)
	ctx := context.Background()

	req := openai.CompletionRequest{
		Model:       openai.GPT3Ada,
		Prompt:      "Lorem ipsum",
		Temperature: 0.6,
	}
	resp, err := c.CreateCompletion(ctx, req)
	if err != nil {
		fmt.Printf("Completion error: %v\n", err)
		return
	}
	fmt.Println(resp.Choices[0].Text)
}
