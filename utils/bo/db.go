package bo

// ModelStruct template/model 模型
type ModelStruct struct {
	TableName string        // tblModelTable
	BeanName  string        // ModelTable
	Columns   []*ModelField // 字段列表
	Comment   string        // 数据表注释
}

// ModelField 实体对象类型
type ModelField struct {
	ColumnName string // db 字段名称
	FieldType  string // 结构体数据类型
	Comment    string // 字段评论
}

// DataStruct template/data 数据模型
type DataStruct struct {
	BeanName  string // ModelTable
	DataIndex []*DataIndex
}

type DataIndex struct {
	IndexName    string
	Unique       bool
	Fields       []*DataIndexField
	IndexComment string
}

type DataIndexField struct {
	Index      int    // 位置
	ColumnName string // db 字段名称
	FieldType  string // 结构体数据类型
}
