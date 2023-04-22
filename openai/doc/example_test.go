package doc

import (
	"context"
	"fmt"
	"github.com/sashabaranov/go-openai"
	"net/http"
	"net/url"
	"testing"
)

const Token = "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"

func TestExample(t *testing.T) {
	defaultConfig := openai.DefaultConfig(Token)
	// 设置代理服务器地址
	proxyUrl, err := url.Parse("http://127.0.0.1:7890")
	if err != nil {
		panic(err)
	}

	// 创建Transport，设置代理
	transport := &http.Transport{
		Proxy: http.ProxyURL(proxyUrl),
	}

	defaultConfig.HTTPClient = &http.Client{
		Transport: transport,
	}
	c := openai.NewClientWithConfig(defaultConfig)

	ctx := context.Background()

	resp, err := c.CreateCompletion(ctx, openai.CompletionRequest{
		Model:       openai.GPT3TextDavinci003,
		Prompt:      "Lorem ipsum",
		Temperature: 0.6,
	})
	if err != nil {
		fmt.Printf("Completion error: %v\n", err)
		return
	}
	fmt.Println(resp.Choices[0].Text)
}
