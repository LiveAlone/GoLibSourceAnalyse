package appfx

import "github.com/LiveAlone/GoLibSourceAnalyse/utils/common"

func AppConstruct() []interface{} {
	depConstruct := []interface{}{
		common.NewGlobalConfig, // 全局配置
		UtilsLogger,            // 全局日志
	}

	// 命令行
	depConstruct = append(depConstruct, SubCmdConstructList()...)
	depConstruct = append(depConstruct, CommandProvider)

	return depConstruct
}
