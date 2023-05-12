package api

import (
	"github.com/LiveAlone/GoLibSourceAnalyse/utils/common"
	"github.com/LiveAlone/GoLibSourceAnalyse/utils/domain/config"
	"log"
)

func ConvertProjectApisDtoDesc(apiList []*HttpApi) (rs []*DtoStructDesc) {
	for _, api := range apiList {
		rs = append(rs, ConvertBodyDescToDtoDesc(api.Prefix, api.ReqBodyDesc)...)
		rs = append(rs, ConvertBodyDescToDtoDesc(api.Prefix, api.ResBodyDesc)...)
	}
	return
}

func ConvertBodyDescToDtoDesc(prefix string, desc *BodyDesc) (rs []*DtoStructDesc) {
	if desc.Type != "object" {
		log.Fatalf("dto desc convert error, body:%v", desc)
	}

	fields := make([]*DtoFieldDesc, 0)
	for _, property := range desc.Properties {
		if property.Type == "object" {
			loopDesc := ConvertBodyDescToDtoDesc(prefix, property)
			rs = append(rs, loopDesc...)
			fields = append(fields, &DtoFieldDesc{
				Name:     common.ToCamelCaseFistLarge(property.Name),
				Type:     toStructName(prefix, property.Name),
				Example:  property.Example,
				Desc:     property.Desc,
				Required: property.Required,
				Array:    property.Array,
			})
		} else {
			fields = append(fields, &DtoFieldDesc{
				Name:     common.ToCamelCaseFistLarge(property.Name),
				Type:     toStructType(property.Type),
				Example:  property.Example,
				Desc:     property.Desc,
				Required: property.Required,
				Array:    property.Array,
			})
		}
	}

	rs = append(rs, &DtoStructDesc{
		Name:         toStructName(prefix, desc.Name),
		Example:      desc.Example,
		Desc:         desc.Desc,
		DtoFieldDesc: fields,
	})
	return
}

func toStructName(prefix string, name string) string {
	return prefix + common.ToCamelCaseFistLarge(name)
}

func toStructType(fromType string) string {
	toType, ok := config.GlobalConf.ApiTypeMap[fromType]
	if !ok {
		log.Fatalf("api from type not found :%v", fromType)
	}
	return toType
}
