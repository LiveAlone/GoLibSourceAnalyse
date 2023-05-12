package mysql

import "gorm.io/gorm"

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

type TableSchemaAnalyser struct {
	DbUrl  string
	Client *DBClient
}

func NewTableSchemaAnalyser(url string) (*TableSchemaAnalyser, error) {
	client, err := NewDBClient(url)
	if err != nil {
		return nil, err
	}
	return &TableSchemaAnalyser{
		DbUrl:  url,
		Client: client,
	}, nil
}

// QueryTable 查询数据表基础信息
func (entity *TableSchemaAnalyser) QueryTable(databaseName, table string) (tb *TableInfo, err error) {
	err = entity.Client.Query("TABLES", &tb, func(db *gorm.DB) *gorm.DB {
		return db.Where("TABLE_SCHEMA = ? and TABLE_NAME = ?", databaseName, table)
	})
	if err != nil {
		return nil, err
	}
	return tb, nil
}

func (entity *TableSchemaAnalyser) QueryColumns(databaseName, table string) (rs []*TableInfoColumn, err error) {
	err = entity.Client.Query("COLUMNS", &rs, func(db *gorm.DB) *gorm.DB {
		return db.Where("TABLE_SCHEMA = ? and TABLE_NAME = ? order by ORDINAL_POSITION", databaseName, table)
	})
	if err != nil {
		return nil, err
	}
	return rs, nil
}
