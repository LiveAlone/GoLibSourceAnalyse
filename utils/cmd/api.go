package cmd

import (
	"fmt"
	"github.com/LiveAlone/GoLibSourceAnalyse/utils/manager/api"
	"github.com/LiveAlone/GoLibSourceAnalyse/utils/manager/yapi"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
	"log"
	"os"
)

var apiProject string
var apiAllApi bool
var apiList string
var apiDest string

var apiDestConfig *ApiConfig

var ApiCmd = &cobra.Command{
	Use:   "api",
	Short: "网关SDK生成",
	Run: func(cmd *cobra.Command, args []string) {
		// 初始化配置
		content, err := os.ReadFile(fmt.Sprintf("%s/%s", apiDest, "api.yaml"))
		if err != nil {
			log.Fatalf("yaml file read error %v", err)
		}

		apiDestConfig := new(ApiConfig)
		err = yaml.Unmarshal(content, apiDestConfig)
		if err != nil {
			log.Fatalf("yaml convert err %v", err)
		}

		generateFromApi()
	},
}

func init() {
	ApiCmd.Flags().StringVarP(&apiProject, "project", "p", "", "输入需要生成项目")
	ApiCmd.Flags().StringVarP(&apiDest, "dest", "d", "", "输入目标文件路径")
	ApiCmd.Flags().BoolVarP(&apiAllApi, "full", "f", false, "是否全量接口同步")
	ApiCmd.Flags().StringVarP(&apiList, "api", "a", "", "输入单个接口列表")
}

type ApiConfig struct {
	Token map[string]string `yaml:"token"`
}

func generateFromApi() {
	token, ok := apiDestConfig.Token[apiProject]
	if !ok {
		log.Fatalf("project fail get token, projet:%v", apiProject)
	}

	var yapiProject *yapi.ProjectDetailInfo
	if apiAllApi {
		yapiProject = yapi.QueryProjectInfo(token, "")
	} else {
		yapiProject = yapi.QueryProjectInfo(token, apiList)
	}

	httpProject := api.DetailToBasicModel(yapiProject)

	// todo yqj 渲染结构体模版，生成client 和SDK
	fmt.Println(httpProject)
}
