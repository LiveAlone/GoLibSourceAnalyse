package util

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Column 数据表Schema 字段定义
type Column struct {
	ColumnName    string `gorm:"column:COLUMN_NAME" json:"column_name"`
	DataType      string `gorm:"column:DATA_TYPE" json:"data_type"`
	ColumnKey     string `gorm:"column:COLUMN_KEY" json:"column_key"`
	IsNullable    string `gorm:"column:IS_NULLABLE" json:"is_nullable"`
	ColumnType    string `gorm:"column:COLUMN_TYPE" json:"column_type"`
	ColumnComment string `gorm:"column:COLUMN_COMMENT" json:"column_comment"`
}

func QueryColumns(url, databaseName, table string) ([]*Column, error) {
	db, err := BuildDbClient(url)
	if err != nil {
		return nil, err
	}
	var columns []*Column
	tx := db.Table("COLUMNS").Where("TABLE_SCHEMA = ? and TABLE_NAME = ? order by ORDINAL_POSITION", databaseName, table).Find(&columns)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return columns, err
}

func BuildDbClient(dbUrl string) (*gorm.DB, error) {
	db, err := gorm.Open(mysql.Open(dbUrl), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}
