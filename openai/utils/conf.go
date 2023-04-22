package utils

import (
	"gopkg.in/yaml.v3"
	"log"
	"os"
)

const SecretConfPath = "conf/secret.yml"

var SecretConf Conf

func init() {
	confContent, err := os.ReadFile(SecretConfPath)
	if err != nil {
		log.Fatalf("read conf faile error, paht:%v, cause %v", SecretConfPath, err)
	}
	err = yaml.Unmarshal(confContent, &SecretConf)
	if err != nil {
		log.Fatalf("init conf yaml analyse error, content:%s, cause:%v", confContent, err)
	}
}

type Conf struct {
	Token string `yaml:"token"`
	Proxy string `yaml:"proxy"`
}
