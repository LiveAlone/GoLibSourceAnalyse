package api

import (
	"github.com/LiveAlone/GoLibSourceAnalyse/utils/domain"
	"github.com/LiveAlone/GoLibSourceAnalyse/utils/domain/config"
	"log"
)

// SchemaGen 基于APISchema 生成code
type SchemaGen struct {
}

func NewSchemaGen() *SchemaGen {
	return &SchemaGen{}
}

// GenFromYapi 通过生成code
func (s *SchemaGen) GenFromYapi(token string, allApi bool, apiList []string) (rs map[string]string, err error) {
	return nil, err
}

func ConvertProjectApisDtoDesc(apiList []*HttpApi) (rs []*DtoStructDesc) {
	for _, api := range apiList {
		rs = append(rs, convertBodyDescToDtoDesc(api.Prefix, api.ReqBodyDesc)...)
		rs = append(rs, convertBodyDescToDtoDesc(api.Prefix, api.ResBodyDesc)...)
	}
	return
}

func convertBodyDescToDtoDesc(prefix string, desc *BodyDesc) (rs []*DtoStructDesc) {
	if desc.Type != "object" {
		log.Fatalf("dto desc convert error, body:%v", desc)
	}

	fields := make([]*DtoFieldDesc, 0)
	for _, property := range desc.Properties {
		if property.Type == "object" {
			loopDesc := convertBodyDescToDtoDesc(prefix, property)
			rs = append(rs, loopDesc...)
			fields = append(fields, &DtoFieldDesc{
				Name:     domain.ToCamelCaseFistLarge(property.Name),
				Type:     toStructName(prefix, property.Name),
				Example:  property.Example,
				Desc:     property.Desc,
				Required: property.Required,
				Array:    property.Array,
			})
		} else {
			fields = append(fields, &DtoFieldDesc{
				Name:     domain.ToCamelCaseFistLarge(property.Name),
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
	return prefix + domain.ToCamelCaseFistLarge(name)
}

func toStructType(fromType string) string {
	toType, ok := config.GlobalConf.ApiTypeMap[fromType]
	if !ok {
		log.Fatalf("api from type not found :%v", fromType)
	}
	return toType
}
