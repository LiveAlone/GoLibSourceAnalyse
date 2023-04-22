package utils

import (
	"github.com/sashabaranov/go-openai"
	"net/http"
	"net/url"
)

func NewClient(proxy bool) (*openai.Client, error) {
	defaultConfig := openai.DefaultConfig(GetOpenAiConfig("token"))

	if proxy {
		proxyUrl, err := url.Parse(GetOpenAiConfig("proxy"))
		if err != nil {
			return nil, err
		}
		defaultConfig.HTTPClient = &http.Client{
			Transport: &http.Transport{
				Proxy: http.ProxyURL(proxyUrl),
			},
		}
	}
	return openai.NewClientWithConfig(defaultConfig), nil
}
