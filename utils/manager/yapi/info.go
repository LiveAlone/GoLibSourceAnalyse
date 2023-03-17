package yapi

import (
	"github.com/LiveAlone/GoLibSourceAnalyse/utils/common"
	"log"
	"strconv"
)

type PageApiInfo struct {
	Count int        `json:"count"`
	Total int        `json:"total"`
	List  []*ApiInfo `json:"list"`
}

func PageQueryApiInfo(token string, projectId, page, size int) *PageApiInfo {
	pageApiInfo := new(PageApiInfo)
	err := common.GetWithErrorCodeResp("https://yapi.zuoyebang.cc/api/interface/list", map[string]string{
		"token":      token,
		"project_id": strconv.Itoa(projectId),
		"page":       strconv.Itoa(page),
		"size":       strconv.Itoa(size),
	}, pageApiInfo)
	if err != nil {
		log.Fatalf("page api info err:%v", err)
	}
	return pageApiInfo
}

type ProjectInfo struct {
	ID       int    `json:"_id"`
	Name     string `json:"name"`     // brick
	Basepath string `json:"basepath"` // /brick
}

func GetProjectInfo(token string) *ProjectInfo {
	projectBaseInfo := new(ProjectInfo)
	var err error
	err = common.GetWithErrorCodeResp("https://yapi.zuoyebang.cc/api/project/get", map[string]string{
		"token": token,
	}, projectBaseInfo)
	if err != nil {
		log.Fatalf("gain basic project info error, casue:%v", err)
	}
	return projectBaseInfo
}

type ApiInfo struct {
	Id           int64           `json:"_id"`
	Method       string          `json:"method"`
	Path         string          `json:"path"`
	Title        string          `json:"title"`
	ReqQueryList []*ReqQueryItem `json:"req_query"`     // GET
	ReqBodyType  string          `json:"req_body_type"` // POST
	ReqBodyOther string          `json:"req_body_other"`
	ResBodyType  string          `json:"res_body_type"` // POST
	ResBody      string          `json:"res_body"`
}
type ReqQueryItem struct {
	Id       string `json:"_id"`
	Name     string `json:"name"`
	Example  string `json:"example"`
	Desc     string `json:"desc"`
	Required string `json:"required"`
}

func GetInterfaceApi(token, apiId string) *ApiInfo {
	apiInfo := new(ApiInfo)
	err := common.GetWithErrorCodeResp("https://yapi.zuoyebang.cc/api/interface/get", map[string]string{
		"token": token,
		"id":    apiId,
	}, apiInfo)
	if err != nil {
		log.Fatalf("single api info err:%v", err)
	}
	return apiInfo
}
