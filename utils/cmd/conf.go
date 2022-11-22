package cmd

import (
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

const ConfPath = "conf/conf.yaml"

var GlobalConf Conf

// Conf 配置环境上下文
type Conf struct {
	DbTypeMap map[string]string
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
