package examples

import (
	"context"
	"errors"
	"fmt"
	"github.com/LiveAlone/GoLibSourceAnalyse/openai/utils"
	"github.com/sashabaranov/go-openai"
	"io"
	"log"
)

func Chat3AdaStream() {
	c := openai.NewClient("your token")
	ctx := context.Background()

	req := openai.CompletionRequest{
		Model:     openai.GPT3Ada,
		MaxTokens: 5,
		Prompt:    "Lorem ipsum",
		Stream:    true,
	}
	stream, err := c.CreateCompletionStream(ctx, req)
	if err != nil {
		fmt.Printf("CompletionStream error: %v\n", err)
		return
	}
	defer stream.Close()

	for {
		response, err := stream.Recv()
		if errors.Is(err, io.EOF) {
			fmt.Println("Stream finished")
			return
		}

		if err != nil {
			fmt.Printf("Stream error: %v\n", err)
			return
		}

		fmt.Printf("Stream response: %v\n", response)
	}
}

func Chat3Ada() {
	c, err := utils.NewClient(true)
	if err != nil {
		log.Fatal(err)
	}
	ctx := context.Background()

	req := openai.CompletionRequest{
		Model:     openai.GPT3Ada,
		MaxTokens: 5,
		Prompt:    "你好",
	}
	resp, err := c.CreateCompletion(ctx, req)
	if err != nil {
		fmt.Printf("Completion error: %v\n", err)
		return
	}
	fmt.Println(resp.Choices[0].Text)
}

func Chat35TurboStream() {
	c, err := utils.NewClient(true)
	if err != nil {
		log.Fatal(err)
	}
	ctx := context.Background()

	req := openai.ChatCompletionRequest{
		Model:     openai.GPT3Dot5Turbo,
		MaxTokens: 200,
		Messages: []openai.ChatCompletionMessage{
			{
				Role:    openai.ChatMessageRoleUser,
				Content: "我是谁？我从哪里来？我要到哪里去？",
			},
		},
		Stream: true,
	}
	stream, err := c.CreateChatCompletionStream(ctx, req)
	if err != nil {
		fmt.Printf("ChatCompletionStream error: %v\n", err)
		return
	}
	defer stream.Close()

	fmt.Printf("Stream response: ")
	for {
		response, err := stream.Recv()
		if errors.Is(err, io.EOF) {
			fmt.Println("\nStream finished")
			return
		}

		if err != nil {
			fmt.Printf("\nStream error: %v\n", err)
			return
		}

		fmt.Printf(response.Choices[0].Delta.Content)
	}
}

func Chat35Turbo() {
	client, err := utils.NewClient(true)
	if err != nil {
		log.Fatalf("NewClient error: %v\n", err)
	}
	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: "你好，你是谁？",
				},
			},
		},
	)

	if err != nil {
		fmt.Printf("ChatCompletion error: %v\n", err)
		return
	}

	fmt.Println(resp.Choices[0].Message.Content)
}

func TextDavinci() {
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
