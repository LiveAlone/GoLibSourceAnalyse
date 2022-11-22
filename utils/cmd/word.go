package cmd

import (
	"log"
	"strings"
	"unicode"

	"github.com/gobeam/stringy"

	"github.com/spf13/cobra"
)

const (
	ModeUpper                = iota + 1 // 全部转大写
	ModeLower                           // 全部转小写
	ModeToCamelCaseFistLarge            // 转大写驼峰
	ModeToCamelCaseFistLower            // 转小写驼峰
	ModeToSnakeLower                    // 转小写下划线
	ModeToSnakeLarge                    // 转大写下划线
)

var desc = strings.Join([]string{
	"该子命令支持各种单词格式转换，模式如下：",
	"1：全部转大写",
	"2：全部转小写",
	"3：转大写驼峰",
	"4：转小写驼峰",
	"5：转下划线小写",
	"6：转下划线大写",
}, "\n")

var str string
var mode int8

var WordCmd = &cobra.Command{
	Use:   "word",
	Short: "单词格式转换",
	Long:  desc,
	Run: func(cmd *cobra.Command, args []string) {
		var content string
		switch mode {
		case ModeUpper:
			content = ToUpper(str)
		case ModeLower:
			content = ToLower(str)
		case ModeToCamelCaseFistLarge:
			content = ToCamelCaseFistLarge(str)
		case ModeToCamelCaseFistLower:
			content = ToCamelCaseFistLower(str)
		case ModeToSnakeLower:
			content = ToSnakeLower(str)
		case ModeToSnakeLarge:
			content = ToSnakeLarge(str)
		default:
			log.Fatalf("暂不支持该转换模式，请执行 help word 查看帮助文档")
		}

		log.Printf("输出结果: %s", content)
	},
}

func init() {
	WordCmd.Flags().StringVarP(&str, "str", "s", "", "请输入单词内容")
	WordCmd.Flags().Int8VarP(&mode, "mode", "m", 0, "请输入单词转换的模式")
}

func ToUpper(s string) string {
	return strings.ToUpper(s)
}

func ToLower(s string) string {
	return strings.ToLower(s)
}

func ToCamelCaseFistLarge(s string) string {
	return stringy.New(s).CamelCase()
}

func ToCamelCaseFistLower(s string) string {
	rs := stringy.New(s).CamelCase()
	if len(rs) > 0 {
		return string(unicode.ToLower(rune(rs[0]))) + rs[1:]
	}
	return rs
}

func ToSnakeLower(s string) string {
	rs := stringy.New(s).SnakeCase().ToLower()
	return rs
}

func ToSnakeLarge(s string) string {
	rs := stringy.New(s).SnakeCase().ToUpper()
	return rs
}
