package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/LiveAlone/GoLibSourceAnalyse/utils/domain"
	"github.com/LiveAlone/GoLibSourceAnalyse/utils/domain/config"
	"github.com/LiveAlone/GoLibSourceAnalyse/utils/domain/mysql"
	"github.com/LiveAlone/GoLibSourceAnalyse/utils/manager/model"
	"log"
	"strings"

	"github.com/spf13/cobra"

	"github.com/LiveAlone/GoLibSourceAnalyse/utils/util"
)

var targetPath string

func NewModelCmd(configLoader *config.Loader, gen *model.SchemaInformationGen) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "model",
		Short: "Dao持久化层生成代码",
		Long:  "dest持久化生成地址，db.yaml 配置文件",
		Run: func(cmd *cobra.Command, args []string) {
			var sqlModelConfig SqlModelConfig
			err := configLoader.LoadConfigToEntity(fmt.Sprintf("%s/%s", targetPath, "db.yaml"), &sqlModelConfig)
			if err != nil {
				log.Fatalf("db.yaml load error, err :%v", err)
			}

			// 数据表生成
			db := sqlModelConfig.Db
			tbs := strings.Split(db.Tables, ",")

			tableCode, err := gen.Gen(db.Url, db.DataBase, tbs)
			if err != nil {
				log.Fatalf("db model generate error, err :%v", err)
			}
			for tableName, code := range tableCode {
				fileName := domain.ToSnakeLower(strings.TrimPrefix(tableName, "tbl"))
				err = util.WriteFile(fmt.Sprintf("%s/%s.go", targetPath, fileName), []byte(code))
				if err != nil {
					log.Fatalf("tb file write error, err :%v", err)
				}
				fmt.Println("数据表Model 生成完成: ", tableName)
			}
			//for _, tb := range tbs {
			//	code := GenerateFromTable(db.Url, db.DataBase, tb)
			//	fileName := domain.ToSnakeLower(strings.TrimPrefix(tb, "tbl"))
			//	err = util.WriteFile(fmt.Sprintf("%s/%s.go", targetPath, fileName), []byte(code))
			//	if err != nil {
			//		log.Fatalf("tb file write error, err :%v", err)
			//	}
			//	fmt.Println("数据表Model 生成完成: ", tb)
			//}
		},
	}
	cmd.Flags().StringVarP(&targetPath, "dest", "d", "", "文件生成目标地址")
	return cmd
}

// SqlModelConfig 模型配置文件
type SqlModelConfig struct {
	Db *DbConfig `yaml:"db"`
}

type DbConfig struct {
	Url      string `yaml:"url"`
	DataBase string `yaml:"dataBase"`
	Tables   string `yaml:"tables"`
}

func GenerateFromTable(url, dbName, tableName string) string {

	infomationClient, err := mysql.NewTableSchemaAnalyser(url)
	if err != nil {
		log.Fatalf("db information create fail, err %v", err)
	}

	columns, err := infomationClient.QueryColumns(dbName, tableName)
	if err != nil {
		log.Fatalf("db struct columns query fail, err %v", err)
	}
	table, err := infomationClient.QueryTable(dbName, tableName)
	if err != nil {
		log.Fatalf("db table info gain error, err:%v", err)
	}

	// 构建数据转换列表
	cols := make([]*ModelField, len(columns))
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

		if (fieldType == "string" || fieldType == "sql.NullString") && strings.Contains(column.ColumnComment, "json") {
			fieldType = "datatypes.JSON"
		}

		cols[i] = &ModelField{
			ColumnName: column.ColumnName,
			FieldType:  fieldType,
			Comment:    column.ColumnComment,
		}
	}

	// 构建数据表结构
	ms := ModelStruct{
		TableName: tableName,
		BeanName:  domain.ToCamelCaseFistLarge(strings.TrimPrefix(tableName, "tbl")),
		Columns:   cols,
		Comment:   table.TableComment,
	}

	ds, _ := json.Marshal(ms)
	fmt.Println(string(ds))

	return util.GenerateFromTemplate("basic", ms, map[string]any{
		"ToCamelCaseFistLarge": domain.ToCamelCaseFistLarge,
		"ToCamelCaseFistLower": domain.ToCamelCaseFistLower,
	})
}

// ModelStruct 模型对象
type ModelStruct struct {
	TableName string
	BeanName  string
	Columns   []*ModelField
	Comment   string // todo yqj 数据表注释
}

// ModelField 实体对象类型
type ModelField struct {
	ColumnName string // db 字段名称
	FieldType  string // 结构体数据类型
	Comment    string // 字段评论
}
