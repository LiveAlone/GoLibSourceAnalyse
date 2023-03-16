package util

import (
	"bytes"
	"fmt"
	"net/http"
	"net/url"

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
