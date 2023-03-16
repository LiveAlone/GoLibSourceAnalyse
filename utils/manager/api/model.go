// Package api api 接口模型
package api

type HttpApi struct {
	Schema           string
	Path             string
	Method           string
	Prefix           string // 接口标识
	Description      string
	RequestBodyType  string
	ResponseBodyType string
}

// BodyDesc 结构体描述信息
type BodyDesc struct {
	Name       string // 字段名称
	Type       string // 基础类型, 结构体类型, todo swagger format
	Example    string
	Required   bool
	Array      bool
	Properties []*BodyDesc
}
