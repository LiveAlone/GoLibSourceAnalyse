package ssh

import (
	"fmt"
	"testing"
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

func TestGainTeacher(t *testing.T) {
	db, err := buildDb(hxxSchool)
	if err != nil {
		fmt.Println(err)
		return
	}

	rs := make([]SubjectTeacher, 0)
	tx := db.Table("tblSubjectTeacher").Where("school_id = ? ", 2000252463).Scan(&rs)
	if tx.Error != nil {
		fmt.Println(tx.Error)
		return
	}
	fmt.Println(len(rs))
	for i, r := range rs {
		fmt.Println(i, r)
	}
}
