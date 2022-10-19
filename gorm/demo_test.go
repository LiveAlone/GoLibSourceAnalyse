package gorm

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"testing"
	"time"
)

const DBMonitorTaskName = "tblDBMonitorTask"

type DBMonitorTask struct {
	Id        int       `gorm:"column:id" json:"id"` // 自增id
	DBId      int       `gorm:"column:db_id" json:"DBId"`
	DBName    string    `gorm:"column:db_name" json:"DBName"`
	Schedule  int       `gorm:"column:schedule" json:"schedule"`
	Tables    string    `gorm:"column:tables" json:"tables"`
	Action    int       `gorm:"column:action" json:"action"`
	Status    int8      `gorm:"column:status" json:"status"`
	CreatedAt time.Time `gorm:"column:created_at" json:"CreatedAt"` //  创建时间
	UpdatedAt time.Time `gorm:"column:updated_at" json:"UpdatedAt"` //  更新时间
}

const DBMonitorRecordName = "tblDBMonitorRecord"

type DBMonitorRecord struct {
	Id        int64     `gorm:"column:id" json:"id"` // 自增id
	TaskId    int       `gorm:"column:task_id" json:"TaskId"`
	TableName string    `gorm:"column:table_name" json:"TableName"`
	Schema    string    `gorm:"column:schema" json:"Schema"`
	DiffSql   string    `gorm:"column:diff_sql" json:"DiffSql"`
	Record    string    `gorm:"column:record" json:"Record"`
	CreatedAt time.Time `gorm:"column:created_at" json:"CreatedAt"` //  创建时间
	UpdatedAt time.Time `gorm:"column:updated_at" json:"UpdatedAt"` //  更新时间
}

// const hxxMisQa = "homework:homework@tcp(10.112.36.52:6060)/hxx_mis_qa?parseTime=true"
const hxxMis = "homework:homework@tcp(10.112.36.52:6060)/hxx_mis?parseTime=true"

func TestDbAction(t *testing.T) {
	db, err := buildDb(hxxMis)
	if err != nil {
		fmt.Println(err)
		return
	}

	// create table
	//db.Table(DBMonitorTaskName).Create(&DBMonitorTask{
	//	DBId:     1,
	//	DBName:   "test",
	//	Schedule: 20,
	//	Tables:   "*",
	//	Action:   1,
	//	Status:   1,
	//})

	// read
	var task DBMonitorTask
	//db.Table(DBMonitorTaskName).First(&task)
	db.Table(DBMonitorTaskName).First(&task, 2)
	fmt.Println(task.CreatedAt.Local())

	//update
	//db.Table(DBMonitorTaskName).Model(&task).Update("db_name", "hello")
	//db.Table(DBMonitorTaskName).Model(&task).Updates(&DBMonitorTask{DBName: "hello3", Status: 4})
	//db.Table(DBMonitorTaskName).Model(&task).Updates(map[string]interface{}{"DBName": "world3", "status": 6})

	//delete
	db.Table(DBMonitorTaskName).Delete(&task)
}

func buildDb(url string) (*gorm.DB, error) {
	db, err := gorm.Open(mysql.Open(url), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}
