package gorm

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"testing"
	"time"
)

const dbUrl = "homework:homework@tcp(10.112.36.52:6060)/hxx_mis_qa?parseTime=true"

type DbRegister struct {
	ID        int    `gorm:"column:id;primaryKey;autoIncrement:true"`
	Name      string ``
	Username  string
	CreatedAt time.Time `gorm:"column:created_at"` //  创建时间
	UpdatedAt time.Time `gorm:"column:updated_at"` //  更新时间
}

func TestQuery(t *testing.T) {
	db, err := gorm.Open(mysql.Open(dbUrl), &gorm.Config{})
	if err != nil {
		t.Errorf("gain db error, %v", err)
		return
	}

	var register DbRegister
	db.Table("tblDBRegisters").First(&register)
	fmt.Println(register)
	fmt.Println(register.CreatedAt)
	fmt.Println(register.UpdatedAt)
}
