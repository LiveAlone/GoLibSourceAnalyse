package api

import (
	"github.com/LiveAlone/GoLibSourceAnalyse/utils/domain"
	"github.com/LiveAlone/GoLibSourceAnalyse/utils/manager/yapi"
	"github.com/LiveAlone/GoLibSourceAnalyse/utils/util"
	"log"
	"strings"
)

func DetailToBasicModel(detail *yapi.ProjectDetailInfo) *HttpProject {
	yapiProjectInfo := detail.ProjectInfo
	res := &HttpProject{
		ID:       yapiProjectInfo.ID,
		Title:    "",
		Name:     yapiProjectInfo.Name,
		BasePath: yapiProjectInfo.Basepath,
	}

	apiList := make([]*HttpApi, 0)
	for _, yapiApi := range detail.ApiList {
		httpApi := &HttpApi{
			Schema:      "http",
			Path:        yapiApi.Path,
			Method:      yapiApi.Method,
			Prefix:      domain.ToCamelCaseFistLarge(strings.ReplaceAll(strings.Trim(yapiApi.Path, "/"), "/", "_")),
			Description: yapiApi.Title,
			ReqBodyType: yapiApi.ReqBodyType,
			ResBodyType: yapiApi.ReqBodyType,
		}
		success := analyseBodyStruct(yapiApi, httpApi)
		if !success {
			log.Printf("api analyse ignore :%v", yapiApi)
			continue
		}
		apiList = append(apiList, httpApi)
	}
	res.ApiList = apiList
	return res
}

func analyseBodyStruct(yapiApi *yapi.ApiInfo, httpApi *HttpApi) bool {
	if yapiApi.Method != "GET" && yapiApi.Method != "POST" {
		return false
	}
	if yapiApi.Method == "GET" {
		httpApi.ReqBodyDesc = yapiGetBodyToDesc("req", yapiApi.ReqQueryList)
	} else {
		if yapiApi.ReqBodyType != "json" {
			return false
		}
		httpApi.ReqBodyDesc = yapiJsonBodyToDesc("req", yapiApi.ReqBodyOther)
	}

	if yapiApi.ResBodyType != "json" {
		return false
	}
	httpApi.ResBodyDesc = yapiJsonBodyToDesc("res", yapiApi.ResBody)
	return true
}

func yapiJsonBodyToDesc(name, jsonDesc string) *BodyDesc {
	yapiWrapper, err := yapi.ConvertJsonStructWrap(jsonDesc)
	if err != nil {
		log.Fatalf("yapi json convert error, json:%v, err:%v", jsonDesc, err)
	}
	return yapiAnalyseWrapper(name, yapiWrapper, true)
}

func yapiAnalyseWrapper(name string, wrapper *yapi.StructWrapper, first bool) *BodyDesc {
	if wrapper.Type != "object" {
		log.Fatalf("analyse not object fail, wrapper:%v, name: %v", wrapper, name)
	}

	// 去除data层, yapi
	if first && wrapper.Type == "object" && len(wrapper.Properties) == 3 {
		if data, ok := wrapper.Properties["data"]; ok {
			return yapiAnalyseWrapper(name, data, false)
		}
	}

	rs := &BodyDesc{
		Name:     name,
		Type:     wrapper.Type,
		Example:  "",
		Desc:     wrapper.Description,
		Required: util.ContainsForArrayString(name, wrapper.Required),
	}
	for wn, w := range wrapper.Properties {
		if w.Type == "object" {
			item := yapiAnalyseWrapper(wn, w, false)
			rs.Properties = append(rs.Properties, item)
		} else if w.Type == "array" {
			newW := w.Items
			if newW.Type == "array" {
				// todo 数组嵌套等待支持
				log.Fatalf("api array not support")
			}

			if newW.Type == "object" {
				// 转换兑现
				item := yapiAnalyseWrapper(wn, newW, false)
				item.Array = true
				rs.Properties = append(rs.Properties, item)
			} else {
				item := &BodyDesc{
					Name:     wn,
					Type:     newW.Type,
					Desc:     newW.Description,
					Required: util.ContainsForArrayString(wn, wrapper.Required),
					Array:    true,
				}
				rs.Properties = append(rs.Properties, item)
			}
		} else {
			item := &BodyDesc{
				Name:     wn,
				Type:     w.Type,
				Desc:     w.Description,
				Required: util.ContainsForArrayString(wn, wrapper.Required),
			}
			rs.Properties = append(rs.Properties, item)
		}
	}
	return rs
}

func yapiGetBodyToDesc(name string, items []*yapi.ReqQueryItem) *BodyDesc {
	if len(items) == 0 {
		return nil
	}
	itemDescList := make([]*BodyDesc, len(items))
	for _, item := range items {
		itemDescList = append(itemDescList, &BodyDesc{
			Name:     item.Name,
			Type:     "string",
			Example:  item.Example,
			Desc:     item.Desc,
			Required: item.Required == "1",
		})
	}
	return &BodyDesc{
		Name:       name,
		Type:       "object",
		Properties: itemDescList,
	}
}
