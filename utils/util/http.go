package util

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

func Get(requestUrl string, params map[string]string, res interface{}) error {
	var values url.Values
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
	return json.Unmarshal(buf.Bytes(), res)
}
