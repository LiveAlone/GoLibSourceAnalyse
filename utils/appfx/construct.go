package appfx

import (
	"github.com/LiveAlone/GoLibSourceAnalyse/utils/domain/config"
	"github.com/LiveAlone/GoLibSourceAnalyse/utils/domain/template"
	"github.com/LiveAlone/GoLibSourceAnalyse/utils/manager/model"
)

func AppConstruct() []interface{} {
	depConstruct := []interface{}{
		config.NewConfigLoader, // 配置加载器
		config.NewGlobalConfig, // 全局配置
		UtilsLogger,            // 全局日志
	}

	// 支持命令行
	depConstruct = append(depConstruct, SubCmdConstructList()...)
	depConstruct = append(depConstruct, CommandProvider)

	// 模版生成器
	depConstruct = append(depConstruct, template.NewGenerator)

	// db 模型
	depConstruct = append(depConstruct, model.NewSchemaInformationGen)

	return depConstruct
}
