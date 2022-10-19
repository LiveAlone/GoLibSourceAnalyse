package gorm

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"testing"
)

const DBMonitorTaskName = "tblDBMonitorTask"

type DBMonitorTask struct {
	Id         int    `gorm:"column:id" json:"id"` // 自增id
	DBId       int    `gorm:"column:db_id" json:"DBId"`
	DBName     string `gorm:"column:db_name" json:"DBName"`
	Schedule   int    `gorm:"column:schedule" json:"schedule"`
	Tables     string `gorm:"column:tables" json:"tables"`
	Action     int    `gorm:"column:action" json:"action"`
	Status     int8   `gorm:"column:status" json:"status"`
	CreateTime int64  `gorm:"create_time" json:"createTime"` // 创建时间
	UpdateTime int64  `gorm:"update_time" json:"updateTime"` // 更新时间
}

const DBMonitorRecordName = "tblDBMonitorRecord"

type DBMonitorRecord struct {
	Id         int64  `gorm:"column:id" json:"id"` // 自增id
	TaskId     int    `gorm:"column:task_id" json:"TaskId"`
	TableName  string `gorm:"column:table_name" json:"TableName"`
	Schema     string `gorm:"column:schema" json:"Schema"`
	DiffSql    string `gorm:"column:diff_sql" json:"DiffSql"`
	Record     string `gorm:"column:record" json:"Record"`
	CreateTime int64  `gorm:"create_time" json:"createTime"` // 创建时间
	UpdateTime int64  `gorm:"update_time" json:"updateTime"` // 更新时间
}

// const hxxMisQa = "homework:homework@tcp(10.112.36.52:6060)/hxx_mis_qa?parseTime=true"
const hxxMis = "homework:homework@tcp(10.112.36.52:6060)/hxx_mis?parseTime=true"

func TestDbAction(t *testing.T) {
	fmt.Println(hxxMis)

	db, err := buildDb(hxxMis)
	if err != nil {
		fmt.Println(err)
		return
	}

	// create table
	db.Table(DBMonitorTaskName).Create(&DBMonitorTask{
		DBId:     1,
		DBName:   "test",
		Schedule: 20,
		Tables:   "*",
		Action:   1,
		Status:   1,
	})
}

func buildDb(url string) (*gorm.DB, error) {
	db, err := gorm.Open(mysql.Open(dbUrl), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}
