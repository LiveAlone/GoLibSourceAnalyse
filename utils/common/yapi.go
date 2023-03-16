package common

import (
	"github.com/LiveAlone/GoLibSourceAnalyse/utils/util"
	"log"

	jsoniter "github.com/json-iterator/go"
)

type YapiStructInfo struct {
	StructName string
	Items      []*YapiStructItem
}

type YapiStructItem struct {
	Name        string
	TypeName    string
	Description string
	Required    bool
	Array       bool
}

func ConvertToStructInfo(name, structJson string) []*YapiStructInfo {
	wrapper := new(StructWrapper)
	err := jsoniter.Unmarshal([]byte(structJson), wrapper)
	if err != nil {
		log.Fatalf("convert to struct info error, json:%v, cause:%v", structJson, err)
	}
	return analyseWrapper(name, wrapper, true)
}

func analyseWrapper(name string, wrapper *StructWrapper, first bool) (rs []*YapiStructInfo) {
	if wrapper.Type != "object" {
		log.Fatalf("analyse not object fail, wrapper:%v, name: %v", wrapper, name)
	}

	// 去除data层, yapi
	if first && wrapper.Type == "object" && len(wrapper.Properties) == 3 {
		if data, ok := wrapper.Properties["data"]; ok {
			return analyseWrapper(name, data, false)
		}
	}

	info := &YapiStructInfo{
		StructName: name,
		Items:      make([]*YapiStructItem, 0),
	}
	for wn, w := range wrapper.Properties {
		item := &YapiStructItem{
			Name:        wn,
			TypeName:    w.Type,
			Description: w.Description,
			Required:    util.ContainsForArrayString(wn, wrapper.Required),
		}
		if item.TypeName == "object" {
			curRsList := analyseWrapper(item.Name, w, false)
			rs = append(rs, curRsList...)
		} else if w.Type == "array" {
			item.Array = true
			newW := w.Items
			if newW.Type == "object" {
				curRsList := analyseWrapper(item.Name, newW, false)
				rs = append(rs, curRsList...)
			}
			item.TypeName = newW.Type
			item.Description = newW.Description
		}
		info.Items = append(info.Items, item)
	}
	rs = append(rs, info)
	return
}

// StructWrapper type(object, array, 基础类型)
type StructWrapper struct {
	Type        string                    `json:"type"`
	Properties  map[string]*StructWrapper `json:"properties"`
	Required    []string                  `json:"required"`
	Items       *StructWrapper            `json:"items"`
	Description string                    `json:"description"`
}
