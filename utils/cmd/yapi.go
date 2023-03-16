package cmd

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"gopkg.in/yaml.v3"

	"github.com/LiveAlone/GoLibSourceAnalyse/utils/util"
	"github.com/spf13/cobra"
)

var yapiProject string
var yapiAllApi bool
var api string
var yapiDest string

var YapiCmd = &cobra.Command{
	Use:   "yapi",
	Short: "网关层协议生成",
	Run: func(cmd *cobra.Command, args []string) {
		content, err := os.ReadFile(fmt.Sprintf("%s/%s", yapiDest, "yapi.yaml"))
		if err != nil {
			log.Fatalf("yaml file read error %v", err)
		}
		var config YapiConfig
		err = yaml.Unmarshal(content, &config)
		if err != nil {
			log.Fatalf("yaml convert err %v", err)
		}
		token, ok := config.Token[yapiProject]
		if !ok {
			log.Fatalf("gain token fail config:%v, project:%v", config, yapiProject)
		}

		projectInfo := FillProjectInfo(token)
		for _, api := range projectInfo.ApiList {
			GenerateSingleApi(projectInfo.BaseInfo, api)
		}
	},
}

func GenerateSingleApi(base *ProjectBaseInfo, api *ProjectApiInfo) {
	if api.Method != "POST" || api.ReqBodyType != "json" || api.ResBodyType != "json" {
		fmt.Printf("api type not support ignore title:%s, url:%s \n", api.Title, api.Path)
		return
	}

	fmt.Println(base, api)

	prefix := ToCamelCaseFistLarge(strings.ReplaceAll(strings.TrimPrefix(api.Path, "/"), "/", "_"))

	// todo yqj 分析输入数据结构

	fmt.Println(prefix)

	// yapi 工具转换
	fmt.Println(api.ReqBodyOther)
	fmt.Println(api.ResBody)

	// AST 构建生成结构体
}

func init() {
	YapiCmd.Flags().StringVarP(&yapiProject, "project", "p", "", "输入需要生成项目")
	YapiCmd.Flags().BoolVarP(&yapiAllApi, "full", "f", false, "是否全量接口同步")
	YapiCmd.Flags().StringVarP(&api, "api", "a", "", "输入单个接口列表")
	YapiCmd.Flags().StringVarP(&yapiDest, "dest", "d", "", "输入目标文件路径")
}

func FillProjectInfo(token string) *ProjectInfo {
	projectBaseInfo := new(ProjectBaseInfo)
	var err error
	err = util.Get("https://yapi.zuoyebang.cc/api/project/get", map[string]string{
		"token": token,
	}, projectBaseInfo)
	if err != nil {
		log.Fatalf("gain basic project info error, casue:%v", err)
	}

	var apiIds []string
	var apiList []*ProjectApiInfo
	if yapiAllApi {
		page, size := 1, 20
		for true {
			pageApiInfo := new(PageApiInfo)
			err = util.Get("https://yapi.zuoyebang.cc/api/interface/list", map[string]string{
				"token":      token,
				"project_id": strconv.Itoa(projectBaseInfo.ID),
				"page":       strconv.Itoa(page),
				"size":       strconv.Itoa(size),
			}, pageApiInfo)
			if err != nil {
				log.Fatalf("page api info err:%v", err)
			}
			if len(pageApiInfo.List) == 0 {
				break
			}
			for _, info := range pageApiInfo.List {
				apiIds = append(apiIds, strconv.FormatInt(info.Id, 10))
			}
			page += 1
		}
	}

	if !yapiAllApi {
		apiIds = strings.Split(strings.TrimSpace(api), ",")
	}
	for _, apiId := range apiIds {
		apiInfo := new(ProjectApiInfo)
		err = util.Get("https://yapi.zuoyebang.cc/api/interface/get", map[string]string{
			"token": token,
			"id":    apiId,
		}, apiInfo)
		if err != nil {
			log.Fatalf("single api info err:%v", err)
		}
		apiList = append(apiList, apiInfo)
	}

	if len(apiList) == 0 {
		return nil
	}

	return &ProjectInfo{
		BaseInfo: projectBaseInfo,
		ApiList:  apiList,
	}
}

type ProjectInfo struct {
	BaseInfo *ProjectBaseInfo
	ApiList  []*ProjectApiInfo
}

type ProjectBaseInfo struct {
	ID       int    `json:"_id"`
	Name     string `json:"name"`     // brick
	Basepath string `json:"basepath"` // /brick
}

type ProjectApiInfo struct {
	Id           int64  `json:"_id"`
	Method       string `json:"method"`
	Path         string `json:"path"`
	Title        string `json:"title"`
	ReqBodyType  string `json:"req_body_type"`
	ReqBodyOther string `json:"req_body_other"`
	ResBodyType  string `json:"res_body_type"`
	ResBody      string `json:"res_body"`
}

type PageApiInfo struct {
	Count int               `json:"count"`
	Total int               `json:"total"`
	List  []*ProjectApiInfo `json:"list"`
}

type YapiConfig struct {
	Token map[string]string `yaml:"token"`
}
