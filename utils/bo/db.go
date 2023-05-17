package bo

// ModelStruct template/model 模型
type ModelStruct struct {
	TableName string
	BeanName  string
	Columns   []*ModelField
	Comment   string // todo yqj 数据表注释
}

// ModelField 实体对象类型
type ModelField struct {
	ColumnName string // db 字段名称
	FieldType  string // 结构体数据类型
	Comment    string // 字段评论
}

// IndexField 索引描述信息
type IndexField struct {
	// todo yqj
}
