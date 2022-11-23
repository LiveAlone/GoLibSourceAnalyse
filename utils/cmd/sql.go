package cmd

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/spf13/cobra"

	"github.com/LiveAlone/GoLibSourceAnalyse/utils/util"
	"github.com/LiveAlone/GoLibSourceAnalyse/utils/util/template_struct"
)

var targetPath string

func init() {
	SqlCmd.Flags().StringVarP(&targetPath, "dest", "d", "", "文件生成目标地址")
}

var SqlCmd = &cobra.Command{
	Use:   "model",
	Short: "Dao持久化层生成代码",
	Long:  "dest持久化生成地址，db.yaml 配置文件",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("数据表Model文件转换, ", targetPath)
		content, err := os.ReadFile("dest//db.yaml")
		if err != nil {
			log.Fatalf("err %v", err)
		}
		fmt.Println(string(content))
	},
}

func GenerateFromTable(db, table string) string {
	url := "homework:homework@tcp(10.112.36.52:6060)/information_schema?charset=utf8mb4&parseTime=True&loc=Local"
	columns, err := util.QueryColumns(url, db, table)
	if err != nil {
		log.Fatalf("db struct columns query fail, err %v", err)
	}

	// 数据转换
	cols := make([]*template_struct.Column, len(columns))
	for i, column := range columns {
		structType, ok := GlobalConf.DbTypeMap[column.DataType]
		if !ok {
			log.Fatalf("data type not found, db:%s, table:%s, type:%s", db, table, column.DataType)
		}

		isPrimary := column.ColumnKey == "PRI"
		notNull := column.IsNullable == "NO"

		cols[i] = &template_struct.Column{
			ColumnName: column.ColumnName,
			DbType:     column.DataType,
			StructType: structType,
			IsPrimary:  isPrimary,
			NotNull:    notNull,
			Comment:    column.ColumnComment,
		}
	}

	ms := template_struct.ModelStruct{
		TableName: table,
		Columns:   cols,
	}

	ms.BeanName = ToCamelCaseFistLarge(strings.TrimPrefix(table, "tbl"))

	return util.GenerateFromTemplate("model", ms, map[string]any{
		"ToCamelCaseFistLarge": ToCamelCaseFistLarge,
	})
}
