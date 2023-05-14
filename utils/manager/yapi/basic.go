package yapi

import (
	jsoniter "github.com/json-iterator/go"
	"strconv"
	"strings"
)

func QueryProjectInfo(token string, apiIdList string) *ProjectDetailInfo {
	projectBaseInfo := getProjectInfo(token)

	var apiIds []string
	if len(apiIdList) > 0 {
		apiIds = strings.Split(strings.TrimSpace(apiIdList), ",")
	} else {
		page, size := 1, 20
		for {
			pageApiInfo := pageQueryApiInfo(token, projectBaseInfo.ID, page, size)
			if len(pageApiInfo.List) == 0 {
				break
			}
			for _, info := range pageApiInfo.List {
				apiIds = append(apiIds, strconv.FormatInt(info.Id, 10))
			}
			page += 1
		}
	}

	apiList := make([]*ApiInfo, 0, len(apiIds))
	for _, apiId := range apiIds {
		interfaceApiInfo := getInterfaceApi(token, apiId)
		apiList = append(apiList, interfaceApiInfo)
	}

	if len(apiList) == 0 {
		return nil
	}

	return &ProjectDetailInfo{
		ProjectInfo: projectBaseInfo,
		ApiList:     apiList,
	}
}

type ProjectDetailInfo struct {
	ProjectInfo *ProjectInfo
	ApiList     []*ApiInfo
}

func ConvertJsonStructWrap(json string) (res *StructWrapper, err error) {
	res = new(StructWrapper)
	err = jsoniter.Unmarshal([]byte(json), res)
	return
}

// StructWrapper type(object, array, 基础类型)
type StructWrapper struct {
	Type        string                    `json:"type"`
	Properties  map[string]*StructWrapper `json:"properties"`
	Required    []string                  `json:"required"`
	Items       *StructWrapper            `json:"items"`
	Description string                    `json:"description"`
}
