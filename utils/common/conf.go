package common

import (
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

const ConfPath = "conf/config.yaml"

var GlobalConf Conf

// Conf 配置环境上下文
type Conf struct {
	DbTypeMap     map[string]string `yaml:"db_type_map"` // DB
	GoNullableMap map[string]string `yaml:"go_nullable_map"`
	DebugMode     bool              // debug 模式输出信息
	ApiTypeMap    map[string]string `yaml:"api_type_map"`
}

func InitConf() {
	confContent, err := os.ReadFile(ConfPath)
	if err != nil {
		log.Fatalf("read conf faile error, paht:%v, cause %v", ConfPath, err)
	}
	err = yaml.Unmarshal(confContent, &GlobalConf)
	if err != nil {
		log.Fatalf("init conf yaml analyse error, content:%s, cause:%v", confContent, err)
	}
}
