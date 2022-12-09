# model 数据库表生成model层

## conf 数据类型映射关系

### 基础数据类型映射  

Db数据类型映射到Golang数据类型 ```conf/config.yaml / db_type_map```

### 数据对象空

参考文档: 

1. [SqlNull处理](https://iamdual.com/en/posts/handle-sql-null-golang/)
2. [golang 官方wiki](https://github.com/golang/go/wiki/SQLInterface)
3. [sql文档](https://pkg.go.dev/database/sql#ColumnType.Nullable)

config中 ```go_nullable_map``` 定义空类型映射关系

### sql 基础字段约定
1. tbl{content} tbl标识业务字段表
2. ORM常用 GROM 实现关系映射, GORM规范 gorm.Model 定义基础类型

```sql
CREATE TABLE `tblModelTable` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键id',

  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '最后更新时间',
  `deleted_at` datetime DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='数据表模型';
```

### comment 注释识别类型
通过 ```gorm.io/datatypes``` 扩展数据类型。 json 定义类型

1. ```datatypes.Date``` 日期类型，使用Sql NullTime 原始类型。
2. ```datatypes.RawMessage``` comment 包含json, 转换该类型
3. ```datatypes.Time``` comment 包含timestamp 转换
4. ```datatypes.URL``` 非空转换类型 