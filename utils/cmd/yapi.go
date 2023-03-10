package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var yapiProject string
var yapiAllApi bool
var api string

var YapiCmd = &cobra.Command{
	Use:   "yapi",
	Short: "网关层协议生成",
	Run: func(cmd *cobra.Command, args []string) {
		if yapiAllApi {
			// 批量获取接口信息
		}

		fmt.Println(yapiProject, yapiAllApi, api)
		// 获取配置生成接口
		// 序列化生成construct
		// 生成配置文件，调用函数
	},
}

func init() {
	YapiCmd.Flags().StringVarP(&yapiProject, "project", "p", "", "输入需要生成项目")
	YapiCmd.Flags().BoolVarP(&yapiAllApi, "full", "f", false, "是否全量接口同步")
	YapiCmd.Flags().StringVarP(&api, "api", "a", "", "输入单个接口列表")
}

func FillProjectInfo() error {
	// 查询接口
	return nil
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
	Path         string `json:"path"`
	Title        string `json:"title"`
	ReqBodyType  string `json:"req_body_type"`
	ReqBodyOther string `json:"req_body_other"`
	ResBodyType  string `json:"res_body_type"`
	ResBody      string `json:"res_body"`
}
