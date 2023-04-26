package examples

import (
	"bytes"
	"context"
	"encoding/base64"
	"fmt"
	"github.com/LiveAlone/GoLibSourceAnalyse/openai/utils"
	"github.com/sashabaranov/go-openai"
	"image/png"
	"log"
	"os"
)

func GenerateImage() {
	c, err := utils.NewClient(true)
	if err != nil {
		log.Fatalf("Client creation error: %v\n", err)
	}

	ctx := context.Background()

	//Example image as base64
	reqBase64 := openai.ImageRequest{
		Prompt:         "松子",
		Size:           openai.CreateImageSize512x512,
		ResponseFormat: openai.CreateImageResponseFormatB64JSON,
		N:              1,
	}

	respBase64, err := c.CreateImage(ctx, reqBase64)
	if err != nil {
		fmt.Printf("Image creation error: %v\n", err)
		return
	}

	imgBytes, err := base64.StdEncoding.DecodeString(respBase64.Data[0].B64JSON)
	if err != nil {
		fmt.Printf("Base64 decode error: %v\n", err)
		return
	}

	r := bytes.NewReader(imgBytes)
	imgData, err := png.Decode(r)
	if err != nil {
		fmt.Printf("PNG decode error: %v\n", err)
		return
	}

	file, err := os.Create("example.png")
	if err != nil {
		fmt.Printf("File creation error: %v\n", err)
		return
	}
	defer file.Close()

	if err := png.Encode(file, imgData); err != nil {
		fmt.Printf("PNG encode error: %v\n", err)
		return
	}

	fmt.Println("The image was saved as example.png")
}

func GenerateImageURL() {
	c, err := utils.NewClient(true)
	if err != nil {
		log.Fatalf("Client creation error: %v\n", err)
	}

	ctx := context.Background()
	// Sample image by link
	reqUrl := openai.ImageRequest{
		Prompt:         "美女",
		Size:           openai.CreateImageSize512x512,
		ResponseFormat: openai.CreateImageResponseFormatURL,
		N:              1,
	}

	respUrl, err := c.CreateImage(ctx, reqUrl)
	if err != nil {
		fmt.Printf("Image creation error: %v\n", err)
		return
	}
	fmt.Println(respUrl.Data[0].URL)
}
