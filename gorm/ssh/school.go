package ssh

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// ssh proxy tunnel
// ssh -L 8888:172.16.1.173:3306 root@123.156.228.100
const hxxSchool = "bawu_mysql:bawu#123@tcp(127.0.0.1:8888)/hxx_school?charset=utf8mb4&parseTime=True&loc=Local"

func buildDb(url string) (*gorm.DB, error) {
	db, err := gorm.Open(mysql.Open(url), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}
