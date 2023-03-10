package util

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"

	jsoniter "github.com/json-iterator/go"
)

type BasicResponse struct {
	ErrCode int                 `json:"errcode"`
	ErrMsg  string              `json:"errmsg"`
	Data    jsoniter.RawMessage `json:"data"`
}

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

	baseResponse := new(BasicResponse)
	err = jsoniter.Unmarshal(buf.Bytes(), baseResponse)
	if err != nil || baseResponse.ErrCode != 0 {
		return errors.New("httpCall" + baseResponse.ErrMsg)
	}

	return json.Unmarshal(baseResponse.Data, target)
}
