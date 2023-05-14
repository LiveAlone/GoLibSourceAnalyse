package cmd

import (
	"fmt"
	"github.com/LiveAlone/GoLibSourceAnalyse/utils/domain"
	"github.com/LiveAlone/GoLibSourceAnalyse/utils/domain/config"
	"github.com/LiveAlone/GoLibSourceAnalyse/utils/manager/api"
	"github.com/LiveAlone/GoLibSourceAnalyse/utils/manager/yapi"
	"github.com/LiveAlone/GoLibSourceAnalyse/utils/util"
	"github.com/spf13/cobra"
	"log"
)

type ApiParam struct {
	project string
	allApi  bool
	list    string
	dest    string
}

var apiParam = new(ApiParam)

func NewApiParam(configLoader *config.Loader) *cobra.Command {
	apiCmd := &cobra.Command{
		Use:   "api",
		Short: "网关SDK生成",
		Run: func(cmd *cobra.Command, args []string) {
			// 初始化配置
			apiDestConfig := new(ApiConfig)
			err := configLoader.LoadConfigToEntity(fmt.Sprintf("%s/%s", apiParam.dest, "api.yaml"), apiDestConfig)
			if err != nil {
				log.Fatalf("yaml file read error %v", err)
			}
			generateFromApi(apiDestConfig.Token[apiParam.project])
		},
	}
	apiCmd.Flags().StringVarP(&apiParam.project, "project", "p", "", "输入需要生成项目")
	apiCmd.Flags().StringVarP(&apiParam.dest, "dest", "d", "", "输入目标文件路径")
	apiCmd.Flags().BoolVarP(&apiParam.allApi, "full", "f", false, "是否全量接口同步")
	apiCmd.Flags().StringVarP(&apiParam.list, "api", "a", "", "输入单个接口列表")
	return apiCmd
}

type ApiConfig struct {
	Token map[string]string `yaml:"token"`
}

func generateFromApi(token string) {
	if len(token) == 0 {
		log.Fatalf("project fail get token, projet:%v", apiParam.project)
	}

	var yapiProject *yapi.ProjectDetailInfo
	if apiParam.allApi {
		yapiProject = yapi.QueryProjectInfo(token, "")
	} else {
		yapiProject = yapi.QueryProjectInfo(token, apiParam.list)
	}

	httpProject := api.DetailToBasicModel(yapiProject)

	// dto generate
	dtoStructs := api.ConvertProjectApisDtoDesc(httpProject.ApiList)

	var content string
	var err error

	//write dto
	content = util.GenerateFromTemplate("api/dto", map[string]any{
		"dtoList": dtoStructs,
	}, map[string]any{
		"ToCamelCaseFistLower": domain.ToCamelCaseFistLower,
	})
	err = util.WriteFile(fmt.Sprintf("%s/%s_dto.go", apiParam.dest, httpProject.Name), []byte(content))
	if err != nil {
		log.Fatalf("wirte dto file error, %v", err)
	}

	// write client
	content = util.GenerateFromTemplate("api/client", map[string]any{
		"apiList":  httpProject.ApiList,
		"basePath": httpProject.BasePath,
		"name":     domain.ToCamelCaseFistLarge(httpProject.Name),
	}, map[string]any{})
	err = util.WriteFile(fmt.Sprintf("%s/%s_api.go", apiParam.dest, httpProject.Name), []byte(content))
	if err != nil {
		log.Fatalf("wirte client file error, %v", err)
	}

	// cont service
	for _, httpApi := range httpProject.ApiList {
		content = util.GenerateFromTemplate("api/control", httpApi, map[string]any{})
		err = util.WriteFile(fmt.Sprintf("%s/%s_%s_controller.go", apiParam.dest, domain.ToSnakeLower(httpApi.Prefix), httpProject.Name), []byte(content))
		if err != nil {
			log.Fatalf("wirte file error, %v", err)
		}

		content = util.GenerateFromTemplate("api/service", httpApi, map[string]any{})
		err = util.WriteFile(fmt.Sprintf("%s/%s_%s_service.go", apiParam.dest, domain.ToSnakeLower(httpApi.Prefix), httpProject.Name), []byte(content))
		if err != nil {
			log.Fatalf("wirte file error, %v", err)
		}
	}
}
