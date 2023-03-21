package util

import (
	"bytes"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	jsoniter "github.com/json-iterator/go"
)

func Get(requestUrl string, params map[string]string, target interface{}) error {
	var values url.Values = make(map[string][]string)
	for k, v := range params {
		values.Set(k, v)
	}
	resp, err := http.Get(fmt.Sprintf("%s?%s", requestUrl, values.Encode()))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	buf := new(bytes.Buffer)
	_, err = buf.ReadFrom(resp.Body)
	if err != nil {
		return err
	}
	return jsoniter.Unmarshal(buf.Bytes(), target)
}

func Post(requestUrl string, body string, target interface{}) error {
	resp, err := http.Post(requestUrl, "application/json", strings.NewReader(body))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	buf := new(bytes.Buffer)
	_, err = buf.ReadFrom(resp.Body)
	if err != nil {
		return err
	}
	return jsoniter.Unmarshal(buf.Bytes(), target)
}
