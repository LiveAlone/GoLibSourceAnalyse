package main

import (
	"context"
	"fmt"
	"github.com/LiveAlone/GoLibSourceAnalyse/openai/utils"
	"github.com/sashabaranov/go-openai"
	"log"
)

func main() {
	c, err := utils.NewClient(true)
	if err != nil {
		log.Fatal(err)
	}
	resp, err := c.CreateCompletion(context.Background(), openai.CompletionRequest{
		Model:  openai.GPT3TextDavinci003,
		Prompt: "给我的狗启个名字",
	})
	if err != nil {
		fmt.Printf("Completion error: %v\n", err)
		return
	}
	fmt.Println(resp.Choices[0].Text)
}
