package examples

import (
	"context"
	"fmt"
	"github.com/LiveAlone/GoLibSourceAnalyse/openai/utils"
	"github.com/sashabaranov/go-openai"
	"log"
	"os"
)

func AudioCaptions() {
	c, err := utils.NewClient(true)
	if err != nil {
		log.Fatalf("NewClient error: %v\n", err)
	}

	req := openai.AudioRequest{
		Model:    openai.Whisper1,
		FilePath: os.Args[1],
		//Format:   openai.AudioResponseFormatSRT,
	}
	resp, err := c.CreateTranscription(context.Background(), req)
	if err != nil {
		fmt.Printf("Transcription error: %v\n", err)
		return
	}
	f, err := os.Create(os.Args[1] + ".srt")
	if err != nil {
		fmt.Printf("Could not open file: %v\n", err)
		return
	}
	defer f.Close()
	if _, err := f.WriteString(resp.Text); err != nil {
		fmt.Printf("Error writing to file: %v\n", err)
		return
	}
}

func AudioToText() {
	c, err := utils.NewClient(true)
	if err != nil {
		log.Fatalf("NewClient error: %v\n", err)
	}
	ctx := context.Background()

	req := openai.AudioRequest{
		Model:    openai.Whisper1,
		FilePath: "recording.mp3",
	}
	resp, err := c.CreateTranscription(ctx, req)
	if err != nil {
		fmt.Printf("Transcription error: %v\n", err)
		return
	}
	fmt.Println(resp.Text)
}
