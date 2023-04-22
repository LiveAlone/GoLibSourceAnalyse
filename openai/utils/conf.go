package utils

import (
	"gopkg.in/yaml.v3"
	"log"
	"os"
)

const ConfPath = "conf/config.yml"

var GlobalConf Conf

func init() {
	confContent, err := os.ReadFile(ConfPath)
	if err != nil {
		log.Fatalf("read conf faile error, paht:%v, cause %v", ConfPath, err)
	}
	err = yaml.Unmarshal(confContent, &GlobalConf)
	if err != nil {
		log.Fatalf("init conf yaml analyse error, content:%s, cause:%v", confContent, err)
	}
}

type Conf struct {
	OpenAi map[string]string `yaml:"openai"` // open ai
}

func GetOpenAiConfig(key string) string {
	return GlobalConf.OpenAi[key]
}
