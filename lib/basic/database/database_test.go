package database

import (
	"database/sql"
	"fmt"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

// 测试数据库基本操作

const dbUrl = "homework:homework@tcp(10.112.36.52:6060)/hxx_school_qa"

func TestMysql(t *testing.T) {
	db, err := sql.Open("mysql", dbUrl)
	if err != nil {
		fmt.Println("test sql gain error, ", err)
		return
	}
	defer db.Close()

	fmt.Println(db.Ping())

	var name string
	row := db.QueryRow("select app_name as name from tblApp limit 1;")
	fmt.Println(row.Scan(&name))
	fmt.Println(name)
}
