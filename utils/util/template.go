package util

import (
	"bytes"
	"log"
	"os"
	"text/template"
)

var supportTemplates = map[string]string{
	"model":  "conf/template/model.template",
	"model2": "conf/template/model2.template",
}

// GenerateFromTemplate 模版生成文本内容
func GenerateFromTemplate(templateName string, data any, funcMap template.FuncMap) string {
	filePath, ok := supportTemplates[templateName]
	if !ok {
		log.Fatalf("template file not found %s", templateName)
	}
	bc, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatalf("os read file error, %s, %v", filePath, err)
	}

	templateContent := string(bc)
	current, err := template.New("current").Funcs(funcMap).Parse(templateContent)
	if err != nil {
		log.Fatalf("template compile error, content:%s, cause:%v ", templateContent, err)
	}

	var rs bytes.Buffer
	err = current.Execute(&rs, data)
	if err != nil {
		log.Fatalf("template execute error, data:%v, tmp:%v, cause:%v", data, templateName, err)
	}
	return rs.String()
}
