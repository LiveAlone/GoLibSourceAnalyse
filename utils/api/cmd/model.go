package cmd

import (
	"fmt"
	"github.com/LiveAlone/GoLibSourceAnalyse/utils/domain"
	"github.com/LiveAlone/GoLibSourceAnalyse/utils/domain/config"
	"github.com/LiveAlone/GoLibSourceAnalyse/utils/manager/db"
	"log"
	"strings"

	"github.com/spf13/cobra"

	"github.com/LiveAlone/GoLibSourceAnalyse/utils/util"
)

var targetPath string

func NewModelCmd(configLoader *config.Loader, gen *db.SchemaInformationGen) *cobra.Command {
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

			tableCode, dataCode, err := gen.Gen(db.Url, db.DataBase, tbs)
			if err != nil {
				log.Fatalf("db model generate error, err :%v", err)
			}

			for tableName, code := range tableCode {
				fileName := domain.ToSnakeLower(strings.TrimPrefix(tableName, "tbl"))

				// model
				dir := fmt.Sprintf("%s/models", targetPath)
				err := util.CreateDirIfNotExists(dir)
				if err != nil {
					log.Fatalf("create dir error, err :%v", err)
				}
				err = util.WriteFile(fmt.Sprintf("%s/%s.go", dir, fileName), []byte(code))
				if err != nil {
					log.Fatalf("tb file write error, err :%v", err)
				}
				fmt.Println("数据表Model 生成完成: ", tableName)
			}

			for tableName, code := range dataCode {

				fileName := domain.ToSnakeLower(strings.TrimPrefix(tableName, "tbl"))
				dataDir := fmt.Sprintf("%s/data", targetPath)
				err = util.CreateDirIfNotExists(dataDir)
				if err != nil {
					log.Fatalf("create dir error, err :%v", err)
				}

				err = util.WriteFile(fmt.Sprintf("%s/%s.go", dataDir, fileName), []byte(code))
				if err != nil {
					log.Fatalf("tb file write error, err :%v", err)
				}
				fmt.Println("数据表Data 生成完成: ", tableName)
			}
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
