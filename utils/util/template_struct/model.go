package template_struct

// ModelStruct 模型对象
type ModelStruct struct {
	TableName string
	BeanName  string
	Columns   []*Column
}

// Column 字段类型
type Column struct {
	ColumnName string
	DbType     string // 数据类型
	StructType string // 结构体类型
	IsPrimary  bool
	NotNull    bool
	Comment    string
}
