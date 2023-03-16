package common

import (
	"encoding/json"
	"errors"
	"github.com/LiveAlone/GoLibSourceAnalyse/utils/util"
	jsoniter "github.com/json-iterator/go"
)

type BasicResponse struct {
	ErrCode int                 `json:"errcode"`
	ErrMsg  string              `json:"errmsg"`
	Data    jsoniter.RawMessage `json:"data"`
}

// GetWithErrorCodeResp 封装查询
func GetWithErrorCodeResp(requestUrl string, params map[string]string, target interface{}) (err error) {
	baseResponse := new(BasicResponse)
	err = util.Get(requestUrl, params, baseResponse)
	if err != nil {
		return err
	}
	if baseResponse.ErrCode != 0 {
		return errors.New("httpCall" + baseResponse.ErrMsg)
	}
	return json.Unmarshal(baseResponse.Data, target)
}
