package common

import (
	"go.uber.org/zap"
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

const ConfPath = "conf/config.yaml"

var GlobalConf Conf

// NewGlobalConfig 初始化全局配置
func NewGlobalConfig() (*Conf, error) {
	confContent, err := os.ReadFile(ConfPath)
	if err != nil {
		log.Fatal("read conf file error", zap.String("path", ConfPath), zap.Error(err))
		return nil, err
	}
	err = yaml.Unmarshal(confContent, &GlobalConf)
	if err != nil {
		log.Fatal("init conf yaml analyse error", zap.String("content", string(confContent)), zap.Error(err))
		return nil, err
	}
	return &GlobalConf, nil
}

// Conf 配置环境上下文
type Conf struct {
	DbTypeMap     map[string]string `yaml:"db_type_map"` // DB
	GoNullableMap map[string]string `yaml:"go_nullable_map"`
	DebugMode     bool              // debug 模式输出信息
	ApiTypeMap    map[string]string `yaml:"api_type_map"`
}
