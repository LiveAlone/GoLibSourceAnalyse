package template_struct

// ModelStruct 模型对象
type ModelStruct struct {
	TableName string
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

// DBTypeToStructType 数据库Golang 映射关系
var DBTypeToStructType = map[string]string{
	"int":        "int32",
	"tinyint":    "int8",
	"smallint":   "int",
	"mediumint":  "int64",
	"bigint":     "int64",
	"bit":        "int",
	"bool":       "bool",
	"enum":       "string",
	"set":        "string",
	"varchar":    "string",
	"char":       "string",
	"tinytext":   "string",
	"mediumtext": "string",
	"text":       "string",
	"longtext":   "string",
	"blob":       "string",
	"tinyblob":   "string",
	"mediumblob": "string",
	"longblob":   "string",
	"date":       "time.Time",
	"datetime":   "time.Time",
	"timestamp":  "time.Time",
	"time":       "time.Time",
	"float":      "float64",
	"double":     "float64",
}
