package yapi

import (
	"github.com/LiveAlone/GoLibSourceAnalyse/utils/manager/model"
)

type ApiClient struct{}

func NewApiClient() *ApiClient {
	return &ApiClient{}
}

func (y *ApiClient) QueryHttpProjectInfo(token string, apiIdList string) (*model.HttpProject, error) {
	yapiProjectInfo, err := y.QueryYapiProjectInfo(token, apiIdList)
	if err != nil {
		return nil, err
	}
	return DetailToBasicModel(yapiProjectInfo), nil
}
