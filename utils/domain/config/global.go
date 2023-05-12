package config

const ConfPath = "conf/config.yaml"

var GlobalConf Conf

// NewGlobalConfig 初始化全局配置
func NewGlobalConfig(loader *Loader) (*Conf, error) {
	err := loader.loadConfigToEntity(ConfPath, &GlobalConf)
	if err != nil {
		return nil, err
	}
	return &GlobalConf, nil
}

// Conf 配置环境上下文
type Conf struct {
	DbTypeMap     map[string]string `yaml:"db_type_map"` // DB
	GoNullableMap map[string]string `yaml:"go_nullable_map"`
	ApiTypeMap    map[string]string `yaml:"api_type_map"`
}
