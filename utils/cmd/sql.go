package cmd

import (
	"fmt"
	"log"
	"os"
	"strings"

	"gopkg.in/yaml.v3"

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
		content, err := os.ReadFile(fmt.Sprintf("%s/%s", targetPath, "db.yaml"))
		if err != nil {
			log.Fatalf("yaml file read error %v", err)
		}
		var config SqlModelConfig
		err = yaml.Unmarshal(content, &config)
		if err != nil {
			log.Fatalf("yaml convert err %v", err)
		}

		db := config.Db
		tbs := strings.Split(db.Tables, ",")
		for _, tb := range tbs {
			code := GenerateFromTable(db.Url, db.DataBase, tb)
			fileName := ToSnakeLower(strings.TrimPrefix(tb, "tbl"))
			err = os.WriteFile(fmt.Sprintf("%s/%s.go", targetPath, fileName), []byte(code), 0666)
			if err != nil {
				log.Fatalf("tb file write error, err :%v", err)
			}
		}
	},
}

type SqlModelConfig struct {
	Db *DbConfig `yaml:"db"`
}

type DbConfig struct {
	Url      string `yaml:"url"`
	DataBase string `yaml:"dataBase"`
	Tables   string `yaml:"tables"`
}

func GenerateFromTable(url, db, table string) string {
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

	return util.GenerateFromTemplate("model2", ms, map[string]any{
		"ToCamelCaseFistLarge": ToCamelCaseFistLarge,
	})
}
