package db

import (
	"encoding/json"
	"fmt"
	"github.com/LiveAlone/GoLibSourceAnalyse/utils/domain"
	"github.com/LiveAlone/GoLibSourceAnalyse/utils/domain/config"
	"github.com/LiveAlone/GoLibSourceAnalyse/utils/domain/mysql"
	"github.com/LiveAlone/GoLibSourceAnalyse/utils/domain/template"
	"log"
	"strings"
)

// SchemaInformationGen 基于SchemaInformation生成代码
type SchemaInformationGen struct {
	tempGenerator *template.Generator
}

func NewSchemaInformationGen(generator *template.Generator) *SchemaInformationGen {
	return &SchemaInformationGen{
		tempGenerator: generator,
	}
}

func (s *SchemaInformationGen) Gen(url string, db string, tableList []string) (rs map[string]string, err error) {
	if len(tableList) == 0 {
		return nil, nil
	}
	analyser, err := mysql.NewTableSchemaAnalyser(url)
	if err != nil {
		return nil, err
	}

	rs = make(map[string]string)
	for _, tableName := range tableList {
		ms, err := convertModelStruct(db, tableName, analyser)
		if err != nil {
			return nil, err
		}

		// 生成代码
		ds, _ := json.Marshal(ms)
		fmt.Println(string(ds))

		data, err := s.tempGenerator.GenerateTemplateContent(ms, map[string]any{
			"ToCamelCaseFistLarge": domain.ToCamelCaseFistLarge,
			"ToCamelCaseFistLower": domain.ToCamelCaseFistLower,
		})
		if err != nil {
			return nil, err
		}
		rs[ms.TableName] = data
	}
	return
}

func convertModelStruct(dbName, tableName string, analyser *mysql.TableSchemaAnalyser) (*template.ModelStruct, error) {
	columns, err := analyser.QueryColumns(dbName, tableName)
	if err != nil {
		return nil, err
	}
	table, err := analyser.QueryTable(dbName, tableName)
	if err != nil {
		return nil, err
	}

	// 构建数据转换列表
	cols := make([]*template.ModelField, len(columns))
	for i, column := range columns {
		fieldType, ok := config.GlobalConf.DbTypeMap[column.DataType]
		if !ok {
			log.Fatalf("data type not found, db:%s, table:%s, type:%s", dbName, table, column.DataType)
		}

		if column.IsNullable == "YES" {
			toFieldType, ok := config.GlobalConf.GoNullableMap[fieldType]
			if !ok {
				log.Fatalf("go nullable type not found, db:%s, table:%s, go_type:%s nullable tyle:%v", dbName, table, column.DataType, toFieldType)
			}
			fieldType = toFieldType
		}

		// json 类型datatypes.JSON 转换
		if (fieldType == "string" || fieldType == "sql.NullString") && strings.Contains(column.ColumnComment, "json") {
			fieldType = "datatypes.JSON"
		}

		cols[i] = &template.ModelField{
			ColumnName: column.ColumnName,
			FieldType:  fieldType,
			Comment:    column.ColumnComment,
		}
	}

	return &template.ModelStruct{
		TableName: tableName,
		BeanName:  domain.ToCamelCaseFistLarge(strings.TrimPrefix(tableName, "tbl")),
		Columns:   cols,
		Comment:   table.TableComment,
	}, nil
}
