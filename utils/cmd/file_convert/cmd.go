package file_convert

import (
	"log"

	"github.com/LiveAlone/GoLibSourceAnalyse/utils/util"
	"github.com/spf13/cobra"
)

var fromFile, destFile string

var FileConvert = &cobra.Command{
	Use:   "convert",
	Short: "文件行转换",
	Long:  "用户需求进行文件行格式化",
	Run: func(cmd *cobra.Command, args []string) {
		lines, err := util.ReadFileLines(fromFile)
		if err != nil {
			log.Fatal(err)
			return
		}

		destLines := make([]string, 0, len(lines))
		for i, line := range lines {
			destLine, exit := lineConvert(i, line)
			destLines = append(destLines, destLine)
			if exit {
				break
			}
		}
		err = util.WriteFileLines(destFile, destLines)
		if err != nil {
			log.Fatal(err)
			return
		}
	},
}

func lineConvert(i int, line string) (string, bool) {
	//return RoleConvert(line)
	//return OrgConvert(line)
	//return UserConvert(line)
	return RelationConvert(line)
}

func init() {
	FileConvert.Flags().StringVarP(&fromFile, "from", "f", "from.text", "文件来源")
	FileConvert.Flags().StringVarP(&destFile, "to", "t", "to.text", "文件输出地址")
}
