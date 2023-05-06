package main

import (
	"fmt"
	"github.com/LiveAlone/GoLibSourceAnalyse/utils/script/model"
	"github.com/LiveAlone/GoLibSourceAnalyse/utils/util"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"time"
)

type SubjectTeacher struct {
	Id         int64  `gorm:"primaryKey;autoIncrement" xlsx:"id"          json:"id"         form:"id"`
	SchoolId   string `gorm:"column:school_id"         xlsx:"school_id"   json:"schoolId"   form:"schoolId"`
	TeacherUid string `gorm:"column:teacher_uid"       xlsx:"teacher_uid" json:"teacherUid" form:"teacherUid"`
	SubjectId  int64  `gorm:"column:subject_id"        xlsx:"subject_id"  json:"subjectId"  form:"subjectId"`
	Status     int8   `gorm:"column:status"            xlsx:"status"      json:"status"     form:"status"` //1.正常，2.删除
	OpUid      string `gorm:"column:op_uid"            xlsx:"op_uid"      json:"opUid"      form:"opUid"`
	CreateTime int64  `gorm:"column:create_time"       xlsx:"create_time" json:"createTime" form:"createTime"`
	UpdateTime int64  `gorm:"column:update_time"       xlsx:"update_time" json:"updateTime" form:"updateTime"`
}

func main() {
	db, err := buildDb(hxxSchool)
	if err != nil {
		fmt.Println(err)
		return
	}

	schIds, err := util.ReadFileLines("dest/data.csv")
	if err != nil {
		log.Fatalf("ReadFileLines err: %v", err)
		return
	}

	// bawu : 2000252463
	var totalResult = make([]string, 0)
	for i, schId := range schIds {

		totalResult = append(totalResult, fmt.Sprintf("# schId %s \n", schId))
		rs, err := GenerateSchSql(db, schId)
		if err != nil {
			log.Fatalf("GenerateSchSql err: %v", err)
			return
		}
		totalResult = append(totalResult, rs...)
		log.Printf("i: %d, id: %s success", i, schId)
	}

	err = util.WriteFileLines("dest/local.sql", totalResult)
	if err != nil {
		log.Fatalf("WriteFileLines err: %v", err)
	}
}

func GenerateSchSql(db *gorm.DB, schId string) (rs []string, err error) {
	subjectTeacherList := make([]SubjectTeacher, 0)
	tx := db.Table("tblSubjectTeacher").Where("school_id = ? and status = 1", schId).Scan(&subjectTeacherList)
	if tx.Error != nil {
		return nil, tx.Error
	}
	if len(subjectTeacherList) == 0 {
		return nil, nil
	}

	teacherSubjectMap := make(map[string][]int64)
	for _, st := range subjectTeacherList {
		teacherSubjectMap[st.TeacherUid] = append(teacherSubjectMap[st.TeacherUid], st.SubjectId)
	}

	teacherSubjectBitMap := make(map[string]int64)
	for teacherUid, subjectIds := range teacherSubjectMap {
		teacherSubjectBitMap[teacherUid] = bitGenerate(subjectIds)
	}

	current := time.Now().Unix()
	for teacherUid, bitV := range teacherSubjectBitMap {
		rs = append(rs, fmt.Sprintf(
			"UPDATE `tblTeacher` SET `op_uid`=2350629888,`update_time`= %d,`subject`= %d WHERE `tblTeacher`.`school_id` = %s AND `tblTeacher`.`teacher_uid` = %s;\n",
			current,
			bitV,
			schId,
			teacherUid,
		))
	}
	return
}

func bitGenerate(subIds []int64) int64 {
	var ret int64
	for _, id := range subIds {
		if v, ok := model.SubjectID2Bit[id]; ok {
			ret = ret | v
		}
	}
	return ret
}

// ssh proxy tunnel
// ssh -L 8888:172.16.1.173:3306 root@123.156.228.100
const hxxSchool = "bawu_mysql:bawu#123@tcp(127.0.0.1:8888)/hxx_school?charset=utf8mb4&parseTime=True&loc=Local"

//const hxxSchool = "homework:homework@tcp(10.117.0.4:3306)/hxx_school_qa?charset=utf8mb4&parseTime=True&loc=Local"

func buildDb(url string) (*gorm.DB, error) {
	db, err := gorm.Open(mysql.Open(url), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}
