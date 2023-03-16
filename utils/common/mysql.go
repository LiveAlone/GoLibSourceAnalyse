package common

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type TableInfoColumn struct {
	ColumnName    string `gorm:"column:COLUMN_NAME" json:"column_name"`
	DataType      string `gorm:"column:DATA_TYPE" json:"data_type"`
	ColumnKey     string `gorm:"column:COLUMN_KEY" json:"column_key"`
	IsNullable    string `gorm:"column:IS_NULLABLE" json:"is_nullable"`
	ColumnType    string `gorm:"column:COLUMN_TYPE" json:"column_type"`
	ColumnComment string `gorm:"column:COLUMN_COMMENT" json:"column_comment"`
}

type TableInfo struct {
	TableSchema  string `gorm:"column:TABLE_SCHEMA"`
	TableName    string `gorm:"column:TABLE_NAME"`
	TableComment string `gorm:"column:TABLE_COMMENT"`
}

// QueryTable 查询数据表基础信息
func (entity *DBClient) QueryTable(databaseName, table string) (tb *TableInfo, err error) {
	err = entity.Query("TABLES", &tb, func(db *gorm.DB) *gorm.DB {
		return db.Where("TABLE_SCHEMA = ? and TABLE_NAME = ?", databaseName, table)
	})
	if err != nil {
		return nil, err
	}
	return tb, nil
}

func (entity *DBClient) QueryColumns(databaseName, table string) (rs []*TableInfoColumn, err error) {
	err = entity.Query("COLUMNS", &rs, func(db *gorm.DB) *gorm.DB {
		return db.Where("TABLE_SCHEMA = ? and TABLE_NAME = ? order by ORDINAL_POSITION", databaseName, table)
	})
	if err != nil {
		return nil, err
	}
	return rs, nil
}

type DBClient struct {
	db *gorm.DB
}

func NewDBClient(dbUrl string) (*DBClient, error) {
	db, err := gorm.Open(mysql.Open(dbUrl), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return &DBClient{
		db: db,
	}, nil
}

func (entity *DBClient) Query(table string, result interface{}, where ...func(db *gorm.DB) *gorm.DB) error {
	err := entity.db.Table(table).Scopes(where...).Find(result).Error
	return err
}
