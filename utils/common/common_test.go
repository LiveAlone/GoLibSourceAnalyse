package common

import (
	"fmt"
	"testing"

	jsoniter "github.com/json-iterator/go"
)

func TestYapiAnalyse(t *testing.T) {
	js := "{\"type\":\"object\",\"properties\":{\"errNo\":{\"type\":\"number\"},\"errMsg\":{\"type\":\"string\"},\"data\":{\"type\":\"object\",\"properties\":{\"total\":{\"type\":\"number\",\"description\":\"总数\"},\"list\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"roleId\":{\"type\":\"number\",\"description\":\"角色ID\"},\"roleName\":{\"type\":\"string\",\"description\":\"角色名称\"},\"roleRemark\":{\"type\":\"string\",\"description\":\"角色描述\"},\"roleType\":{\"type\":\"number\",\"description\":\"角色类型，1默认，2普通，3管理员角色\"}},\"required\":[\"roleId\",\"roleName\",\"roleRemark\",\"roleType\"]},\"description\":\"列表\"}},\"required\":[\"total\",\"list\"]}},\"required\":[\"errNo\",\"errMsg\",\"data\"]}"
	rs := ConvertToStructInfo("test", js)
	for i, r := range rs {
		data, _ := jsoniter.Marshal(r)
		fmt.Println(i, string(data))
	}
}
