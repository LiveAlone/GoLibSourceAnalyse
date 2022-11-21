package util

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func BuildDbClient(dbUrl string) (*gorm.DB, error) {
	db, err := gorm.Open(mysql.Open(dbUrl), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}
