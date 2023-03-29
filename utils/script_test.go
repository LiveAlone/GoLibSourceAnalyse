package main

import (
	"fmt"
	"github.com/LiveAlone/GoLibSourceAnalyse/utils/util"
	jsoniter "github.com/json-iterator/go"
	"log"
	"testing"
)

type PrepareItem struct {
	SchoolName  string `json:"school_name" binding:"required"`
	GradeName   string `json:"grade_name" binding:"required"`
	ClassName   string `json:"class_name" binding:"required"`
	SubjectName string `json:"subject_name" binding:"required"`
	UserName    string `json:"user_name" binding:"required"`
	Phone       string `json:"phone" binding:"required"`
}

// 生成sql脚本
func TestTeacherSJTSql(t *testing.T) {
	// 脚本执行
	excelFile := "dest/data.xlsx"
	datas, err := util.ReadExcelData(excelFile, 0)
	if err != nil {
		log.Fatalf("read excel data failed, err: %v", err)
	}

	rs := make([]PrepareItem, 0, len(datas))
	for i, data := range datas {
		if i == 0 {
			continue
		}
		if i > 10 {
			break
		}
		if len(data) < 9 || len(data[8]) == 0 {
			break
		}

		entity := PrepareItem{
			SchoolName:  data[1],
			GradeName:   data[2],
			ClassName:   data[3],
			SubjectName: data[4],
			UserName:    data[5],
			Phone:       data[7],
		}

		rs = append(rs, entity)
	}

	for _, item := range rs {
		body, err := jsoniter.Marshal(map[string][]PrepareItem{
			"item_list": {item},
		})
		if err != nil {
			log.Fatalf("json marshal failed, err: %v", err)
		}
		rs := make(map[string]interface{})
		err = util.Post("http://localhost:8080/sjt/prepare_data", string(body), &rs)
		if err != nil {
			log.Fatalf("http post failed, err: %v", err)
		}
		fmt.Printf("http post success, rs: %v", rs)
		break
	}
}
