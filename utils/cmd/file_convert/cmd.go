package file_convert

import "github.com/spf13/cobra"

var fromFile, destFile string

var FileConvert = &cobra.Command{
	Use:   "convert",
	Short: "文件行转换",
	Long:  "用户需求进行文件行格式化",
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func init() {
	FileConvert.Flags().StringVarP(&fromFile, "from", "f", "from.text", "文件来源")
	FileConvert.Flags().StringVarP(&destFile, "to", "t", "to.text", "文件输出地址")
}
